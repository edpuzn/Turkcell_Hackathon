package csvload

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"turkcell-hackathon-backend/internal/models"
)

// Catalog - Tüm veri modellerini içeren ana yapı
type Catalog struct {
	Coverages       []models.Coverage
	MobilePlans     []models.MobilePlan
	HomePlans       []models.HomePlan
	TVPlans         []models.TVPlan
	BundlingRules   []models.BundlingRule
	InstallSlots    []models.InstallSlot
	Users           []models.User
	HouseholdLines  []models.HouseholdLine
	CurrentServices []models.CurrentServices

	// Indexes for faster lookups
	coverageByAddress map[string][]models.Coverage
	householdByUser   map[int][]models.HouseholdLine
	servicesByUser    map[int]*models.CurrentServices

	mu sync.RWMutex
}

// Loader - CSV yükleyici singleton
type Loader struct {
	once    sync.Once
	catalog *Catalog
}

// Global loader instance
var GlobalLoader = &Loader{}

// LoadAll - Tüm CSV dosyalarını yükler ve Catalog oluşturur
func (l *Loader) LoadAll(ctx context.Context, dir string) (*Catalog, error) {
	var err error
	l.once.Do(func() {
		l.catalog, err = l.loadAllFiles(ctx, dir)
	})
	return l.catalog, err
}

// loadAllFiles - Tüm CSV dosyalarını sırayla yükler
func (l *Loader) loadAllFiles(ctx context.Context, dir string) (*Catalog, error) {
	catalog := &Catalog{
		coverageByAddress: make(map[string][]models.Coverage),
		householdByUser:   make(map[int][]models.HouseholdLine),
		servicesByUser:    make(map[int]*models.CurrentServices),
	}

	fileConfigs := GetAllFileConfigs()

	for _, config := range fileConfigs {
		filepath := fmt.Sprintf("%s/%s", dir, config.Filename)

		if err := l.loadFile(ctx, filepath, config, catalog); err != nil {
			return nil, fmt.Errorf("error loading %s: %w", config.Filename, err)
		}
	}

	// Build indexes
	catalog.BuildIndexes()

	return catalog, nil
}

// loadFile - Tek bir CSV dosyasını yükler
func (l *Loader) loadFile(ctx context.Context, filepath string, config FileConfig, catalog *Catalog) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Skip UTF-8 BOM if present
	reader := io.Reader(file)
	if bom := make([]byte, 3); len(bom) == 3 {
		if _, err := io.ReadFull(file, bom); err == nil && bom[0] == 0xEF && bom[1] == 0xBB && bom[2] == 0xBF {
			reader = file
		} else {
			file.Seek(0, 0)
			reader = file
		}
	}

	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read header
	header, err := csvReader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Validate header
	if err := l.validateHeader(header, config.Headers, config.Filename); err != nil {
		return err
	}

	// Read data rows
	lineNum := 2 // Start from line 2 (after header)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading line %d: %w", lineNum, err)
		}

		// Skip blank lines
		if len(record) == 0 || (len(record) == 1 && strings.TrimSpace(record[0]) == "") {
			lineNum++
			continue
		}

		// Parse row based on file type
		if err := l.parseRow(record, config.Filename, catalog); err != nil {
			return fmt.Errorf("error parsing line %d in %s: %w", lineNum, config.Filename, err)
		}

		lineNum++
	}

	return nil
}

// validateHeader - Header'ı doğrular
func (l *Loader) validateHeader(header, expected []string, filename string) error {
	if len(header) != len(expected) {
		return fmt.Errorf("header length mismatch in %s: expected %d, got %d", filename, len(expected), len(header))
	}

	for i, h := range header {
		if strings.TrimSpace(h) != expected[i] {
			return fmt.Errorf("header mismatch in %s at column %d: expected '%s', got '%s'", filename, i+1, expected[i], h)
		}
	}

	return nil
}

// parseRow - Satırı dosya tipine göre parse eder
func (l *Loader) parseRow(record []string, filename string, catalog *Catalog) error {
	switch filename {
	case CoverageFile:
		return l.parseCoverage(record, catalog)
	case MobilePlansFile:
		return l.parseMobilePlan(record, catalog)
	case HomePlansFile:
		return l.parseHomePlan(record, catalog)
	case TVPlansFile:
		return l.parseTVPlan(record, catalog)
	case BundlingRulesFile:
		return l.parseBundlingRule(record, catalog)
	case InstallSlotsFile:
		return l.parseInstallSlot(record, catalog)
	case UsersFile:
		return l.parseUser(record, catalog)
	case HouseholdFile:
		return l.parseHouseholdLine(record, catalog)
	case CurrentServicesFile:
		return l.parseCurrentServices(record, catalog)
	default:
		return fmt.Errorf("unknown file type: %s", filename)
	}
}

// Parse functions for each model type
// parseCoverage - Yeni CSV formatına göre coverage parse eder
// Yeni format: address_id,city,district,fiber,vdsl,fwa
func (l *Loader) parseCoverage(record []string, catalog *Catalog) error {
	if len(record) < 6 {
		return fmt.Errorf("insufficient fields for coverage: %d", len(record))
	}

	addressID := strings.TrimSpace(record[0])

	// Her teknoloji için ayrı coverage kaydı oluştur
	techs := []string{"fiber", "vdsl", "fwa"}
	techValues := []string{record[3], record[4], record[5]}

	for i, tech := range techs {
		techValue := strings.TrimSpace(techValues[i])
		if techValue == "1" {
			// Varsayılan hız değerleri (gerçek uygulamada CSV'den alınmalı)
			var downMbps int
			switch tech {
			case "fiber":
				downMbps = 1000
			case "vdsl":
				downMbps = 100
			case "fwa":
				downMbps = 50
			default:
				downMbps = 50
			}

			coverage := models.Coverage{
				AddressID: addressID,
				Tech:      tech,
				DownMbps:  downMbps,
			}
			catalog.Coverages = append(catalog.Coverages, coverage)
		}
	}

	return nil
}

func (l *Loader) parseMobilePlan(record []string, catalog *Catalog) error {
	if len(record) < 7 {
		return fmt.Errorf("insufficient fields for mobile plan: %d", len(record))
	}

	id, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid plan_id: %s", record[0])
	}

	quotaGB, err := strconv.Atoi(strings.TrimSpace(record[2]))
	if err != nil {
		return fmt.Errorf("invalid quota_gb: %s", record[2])
	}

	quotaMin, err := strconv.Atoi(strings.TrimSpace(record[3]))
	if err != nil {
		return fmt.Errorf("invalid quota_min: %s", record[3])
	}

	monthlyPrice, err := strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
	if err != nil {
		return fmt.Errorf("invalid monthly_price: %s", record[4])
	}

	overageGB, err := strconv.ParseFloat(strings.TrimSpace(record[5]), 64)
	if err != nil {
		return fmt.Errorf("invalid overage_gb: %s", record[5])
	}

	overageMin, err := strconv.ParseFloat(strings.TrimSpace(record[6]), 64)
	if err != nil {
		return fmt.Errorf("invalid overage_min: %s", record[6])
	}

	plan := models.MobilePlan{
		ID:           id,
		Name:         strings.TrimSpace(record[1]),
		QuotaGB:      quotaGB,
		QuotaMin:     quotaMin,
		MonthlyPrice: monthlyPrice,
		OverageGB:    overageGB,
		OverageMin:   overageMin,
	}

	catalog.MobilePlans = append(catalog.MobilePlans, plan)
	return nil
}

func (l *Loader) parseHomePlan(record []string, catalog *Catalog) error {
	if len(record) < 6 {
		return fmt.Errorf("insufficient fields for home plan: %d", len(record))
	}

	id, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid home_id: %s", record[0])
	}

	downMbps, err := strconv.Atoi(strings.TrimSpace(record[3]))
	if err != nil {
		return fmt.Errorf("invalid down_mbps: %s", record[3])
	}

	monthlyPrice, err := strconv.ParseFloat(strings.TrimSpace(record[4]), 64)
	if err != nil {
		return fmt.Errorf("invalid monthly_price: %s", record[4])
	}

	installFee, err := strconv.ParseFloat(strings.TrimSpace(record[5]), 64)
	if err != nil {
		return fmt.Errorf("invalid install_fee: %s", record[5])
	}

	plan := models.HomePlan{
		ID:           id,
		Name:         strings.TrimSpace(record[1]),
		Tech:         strings.TrimSpace(record[2]),
		DownMbps:     downMbps,
		MonthlyPrice: monthlyPrice,
		InstallFee:   installFee,
	}

	catalog.HomePlans = append(catalog.HomePlans, plan)
	return nil
}

func (l *Loader) parseTVPlan(record []string, catalog *Catalog) error {
	if len(record) < 4 {
		return fmt.Errorf("insufficient fields for TV plan: %d", len(record))
	}

	id, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid tv_id: %s", record[0])
	}

	hdHoursIncl, err := strconv.Atoi(strings.TrimSpace(record[2]))
	if err != nil {
		return fmt.Errorf("invalid hd_hours_included: %s", record[2])
	}

	monthlyPrice, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
	if err != nil {
		return fmt.Errorf("invalid monthly_price: %s", record[3])
	}

	plan := models.TVPlan{
		ID:           id,
		Name:         strings.TrimSpace(record[1]),
		HDHoursIncl:  hdHoursIncl,
		MonthlyPrice: monthlyPrice,
	}

	catalog.TVPlans = append(catalog.TVPlans, plan)
	return nil
}

func (l *Loader) parseBundlingRule(record []string, catalog *Catalog) error {
	if len(record) < 5 {
		return fmt.Errorf("insufficient fields for bundling rule: %d", len(record))
	}

	id, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid rule_id: %s", record[0])
	}

	discountPercent, err := strconv.ParseFloat(strings.TrimSpace(record[3]), 64)
	if err != nil {
		return fmt.Errorf("invalid discount_percent: %s", record[3])
	}

	rule := models.BundlingRule{
		ID:              id,
		RuleType:        strings.TrimSpace(record[1]),
		Description:     strings.TrimSpace(record[2]),
		DiscountPercent: discountPercent,
		AppliesTo:       strings.TrimSpace(record[4]),
	}

	catalog.BundlingRules = append(catalog.BundlingRules, rule)
	return nil
}

func (l *Loader) parseInstallSlot(record []string, catalog *Catalog) error {
	if len(record) < 5 {
		return fmt.Errorf("insufficient fields for install slot: %d", len(record))
	}

	slot := models.InstallSlot{
		SlotID:    strings.TrimSpace(record[0]),
		AddressID: strings.TrimSpace(record[1]),
		StartISO:  strings.TrimSpace(record[2]),
		EndISO:    strings.TrimSpace(record[3]),
		Tech:      strings.TrimSpace(record[4]),
	}

	catalog.InstallSlots = append(catalog.InstallSlots, slot)
	return nil
}

func (l *Loader) parseUser(record []string, catalog *Catalog) error {
	if len(record) < 4 {
		return fmt.Errorf("insufficient fields for user: %d", len(record))
	}

	id, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid user_id: %s", record[0])
	}

	user := models.User{
		ID:                 id,
		Name:               strings.TrimSpace(record[1]),
		AddressID:          strings.TrimSpace(record[2]),
		CurrentBundleLabel: strings.TrimSpace(record[3]),
	}

	catalog.Users = append(catalog.Users, user)
	return nil
}

func (l *Loader) parseHouseholdLine(record []string, catalog *Catalog) error {
	if len(record) < 5 {
		return fmt.Errorf("insufficient fields for household line: %d", len(record))
	}

	userID, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid user_id: %s", record[0])
	}

	expectedGB, err := strconv.Atoi(strings.TrimSpace(record[2]))
	if err != nil {
		return fmt.Errorf("invalid expected_gb: %s", record[2])
	}

	expectedMin, err := strconv.Atoi(strings.TrimSpace(record[3]))
	if err != nil {
		return fmt.Errorf("invalid expected_min: %s", record[3])
	}

	tvHDHours, err := strconv.Atoi(strings.TrimSpace(record[4]))
	if err != nil {
		return fmt.Errorf("invalid tv_hd_hours: %s", record[4])
	}

	line := models.HouseholdLine{
		UserID:      userID,
		LineID:      strings.TrimSpace(record[1]),
		ExpectedGB:  expectedGB,
		ExpectedMin: expectedMin,
		TVHDHours:   tvHDHours,
	}

	catalog.HouseholdLines = append(catalog.HouseholdLines, line)
	return nil
}

func (l *Loader) parseCurrentServices(record []string, catalog *Catalog) error {
	if len(record) < 6 {
		return fmt.Errorf("insufficient fields for current services: %d", len(record))
	}

	userID, err := strconv.Atoi(strings.TrimSpace(record[0]))
	if err != nil {
		return fmt.Errorf("invalid user_id: %s", record[0])
	}

	hasHome, err := strconv.ParseBool(strings.TrimSpace(record[1]))
	if err != nil {
		return fmt.Errorf("invalid has_home: %s", record[1])
	}

	homeSpeed, err := strconv.Atoi(strings.TrimSpace(record[3]))
	if err != nil {
		return fmt.Errorf("invalid home_speed: %s", record[3])
	}

	hasTV, err := strconv.ParseBool(strings.TrimSpace(record[4]))
	if err != nil {
		return fmt.Errorf("invalid has_tv: %s", record[4])
	}

	services := models.CurrentServices{
		UserID:        userID,
		HasHome:       hasHome,
		HomeTech:      strings.TrimSpace(record[2]),
		HomeSpeed:     homeSpeed,
		HasTV:         hasTV,
		MobilePlanIDs: strings.TrimSpace(record[5]),
	}

	catalog.CurrentServices = append(catalog.CurrentServices, services)
	return nil
}

// BuildIndexes - Hızlı arama için indexleri oluşturur
func (c *Catalog) BuildIndexes() {
	// Map'leri initialize et
	if c.coverageByAddress == nil {
		c.coverageByAddress = make(map[string][]models.Coverage)
	}
	if c.householdByUser == nil {
		c.householdByUser = make(map[int][]models.HouseholdLine)
	}
	if c.servicesByUser == nil {
		c.servicesByUser = make(map[int]*models.CurrentServices)
	}

	// Coverage by address
	for _, coverage := range c.Coverages {
		c.coverageByAddress[coverage.AddressID] = append(c.coverageByAddress[coverage.AddressID], coverage)
	}

	// Household lines by user
	for _, line := range c.HouseholdLines {
		c.householdByUser[line.UserID] = append(c.householdByUser[line.UserID], line)
	}

	// Current services by user
	for i := range c.CurrentServices {
		c.servicesByUser[c.CurrentServices[i].UserID] = &c.CurrentServices[i]
	}
}

// Safe getters with read-only access
func (c *Catalog) ByAddressCoverage(addressID string) []models.Coverage {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if coverages, exists := c.coverageByAddress[addressID]; exists {
		result := make([]models.Coverage, len(coverages))
		copy(result, coverages)
		return result
	}
	return []models.Coverage{}
}

func (c *Catalog) Counts() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return map[string]int{
		"coverages":        len(c.Coverages),
		"mobile_plans":     len(c.MobilePlans),
		"home_plans":       len(c.HomePlans),
		"tv_plans":         len(c.TVPlans),
		"bundling_rules":   len(c.BundlingRules),
		"install_slots":    len(c.InstallSlots),
		"users":            len(c.Users),
		"household_lines":  len(c.HouseholdLines),
		"current_services": len(c.CurrentServices),
	}
}
