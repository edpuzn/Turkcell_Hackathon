package service

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"turkcell-hackathon-backend/internal/csvload"
	"turkcell-hackathon-backend/internal/models"
)

// Recommender - Öneri motoru arayüzü
type Recommender interface {
	Recommend(ctx context.Context, in models.RecommendationInput) (models.RecommendationResponse, error)
}

// recommender - Öneri motoru implementasyonu
type recommender struct {
	catalog *csvload.Catalog
}

// NewRecommender - Yeni öneri motoru oluşturur
func NewRecommender(catalog *csvload.Catalog) Recommender {
	return &recommender{
		catalog: catalog,
	}
}

// Recommend - Ana öneri fonksiyonu
func (r *recommender) Recommend(ctx context.Context, in models.RecommendationInput) (models.RecommendationResponse, error) {
	// Household kontrolü
	if len(in.Household) == 0 {
		return models.RecommendationResponse{}, fmt.Errorf("household cannot be empty")
	}

	// Kapsama kontrolü
	coverages := r.catalog.ByAddressCoverage(in.AddressID)
	if len(coverages) == 0 {
		return models.RecommendationResponse{}, fmt.Errorf("address %s has no coverage", in.AddressID)
	}

	// Mevcut hizmetler (tasarruf hesaplaması için)
	currentServices := r.getCurrentServices(in.UserID)
	currentTotal := r.calculateCurrentTotal(currentServices)

	// Mobil plan adayları
	mobileCandidates := r.generateMobileCandidates(in.Household)

	// Ev internet plan adayları
	homeCandidates := r.chooseHomePlans(in.AddressID, in.PreferTech, r.catalog.HomePlans, coverages)

	// TV plan adayları
	totalHDHours := r.calculateTotalHDHours(in.Household)
	tvCandidates := r.chooseTVPlan(totalHDHours, r.catalog.TVPlans)

	// Kombinasyonları oluştur ve skorla
	combinations := r.generateCombinations(mobileCandidates, homeCandidates, tvCandidates)

	// En iyi 3'ü seç
	top3 := r.selectTop3(combinations, currentTotal)

	return models.RecommendationResponse{Top3: top3}, nil
}

// generateMobileCandidates - Her hat için mobil plan adayları oluşturur
func (r *recommender) generateMobileCandidates(household []models.HouseholdLine) [][]models.MobilePlan {
	var candidates [][]models.MobilePlan

	for _, line := range household {
		lineCandidates := r.candidateMobilePlansForLine(line, r.catalog.MobilePlans)
		candidates = append(candidates, lineCandidates)
	}

	return candidates
}

// candidateMobilePlansForLine - Bir hat için mobil plan adayları
func (r *recommender) candidateMobilePlansForLine(line models.HouseholdLine, allPlans []models.MobilePlan) []models.MobilePlan {
	var candidates []models.MobilePlan

	// Planları GB kotasına göre sırala
	sortedPlans := make([]models.MobilePlan, len(allPlans))
	copy(sortedPlans, allPlans)
	sort.Slice(sortedPlans, func(i, j int) bool {
		return sortedPlans[i].QuotaGB < sortedPlans[j].QuotaGB
	})

	// Fit plan (ihtiyacı karşılayan ilk plan)
	fitPlan := r.findFitPlan(line, sortedPlans)
	if fitPlan != nil {
		candidates = append(candidates, *fitPlan)
	}

	// Down plan (bir seviye aşağı, eğer varsa)
	downPlan := r.findDownPlan(line, sortedPlans)
	if downPlan != nil {
		candidates = append(candidates, *downPlan)
	}

	// Up plan (bir seviye yukarı, eğer varsa)
	upPlan := r.findUpPlan(line, sortedPlans)
	if upPlan != nil {
		candidates = append(candidates, *upPlan)
	}

	// En fazla 2 aday döndür
	if len(candidates) > 2 {
		candidates = candidates[:2]
	}

	return candidates
}

// findFitPlan - İhtiyacı karşılayan ilk planı bulur
func (r *recommender) findFitPlan(line models.HouseholdLine, plans []models.MobilePlan) *models.MobilePlan {
	for _, plan := range plans {
		if plan.QuotaGB >= line.ExpectedGB && plan.QuotaMin >= line.ExpectedMin {
			return &plan
		}
	}
	return nil
}

// findDownPlan - Bir seviye aşağı planı bulur
func (r *recommender) findDownPlan(line models.HouseholdLine, plans []models.MobilePlan) *models.MobilePlan {
	for i := len(plans) - 1; i >= 0; i-- {
		plan := plans[i]
		if plan.QuotaGB < line.ExpectedGB || plan.QuotaMin < line.ExpectedMin {
			return &plan
		}
	}
	return nil
}

// findUpPlan - Bir seviye yukarı planı bulur
func (r *recommender) findUpPlan(line models.HouseholdLine, plans []models.MobilePlan) *models.MobilePlan {
	fitPlan := r.findFitPlan(line, plans)
	if fitPlan == nil {
		return nil
	}

	// Fit plan'dan sonraki planı bul
	for i, plan := range plans {
		if plan.ID == fitPlan.ID && i+1 < len(plans) {
			return &plans[i+1]
		}
	}
	return nil
}

// chooseHomePlans - Ev internet plan adaylarını seçer
func (r *recommender) chooseHomePlans(addressID string, preferTech []string, allHome []models.HomePlan, coverages []models.Coverage) []*models.HomePlan {
	// Kapsama teknolojilerini al
	availableTechs := make(map[string]bool)
	for _, coverage := range coverages {
		availableTechs[coverage.Tech] = true
	}

	// Tercih edilen teknolojileri sırala
	techPriority := r.buildTechPriority(preferTech, availableTechs)

	var candidates []*models.HomePlan

	// Her teknoloji için en uygun planları bul
	for _, tech := range techPriority {
		techPlans := r.filterPlansByTech(allHome, tech)
		if len(techPlans) == 0 {
			continue
		}

		// En ucuz plan
		cheapest := r.findCheapestPlan(techPlans)
		if cheapest != nil {
			candidates = append(candidates, cheapest)
		}

		// Bir seviye yukarı plan (eğer varsa)
		upPlan := r.findNextTierPlan(techPlans, cheapest)
		if upPlan != nil {
			candidates = append(candidates, upPlan)
		}

		// En fazla 2 aday per teknoloji
		if len(candidates) >= 2 {
			break
		}
	}

	return candidates
}

// buildTechPriority - Teknoloji öncelik sırasını oluşturur
func (r *recommender) buildTechPriority(preferTech []string, availableTechs map[string]bool) []string {
	priority := []string{"fiber", "vdsl", "fwa"} // Varsayılan sıralama

	// Tercih edilen teknolojileri öne al
	if len(preferTech) > 0 {
		var customPriority []string
		for _, tech := range preferTech {
			if availableTechs[tech] {
				customPriority = append(customPriority, tech)
			}
		}
		// Kalan teknolojileri ekle
		for _, tech := range priority {
			if !contains(customPriority, tech) && availableTechs[tech] {
				customPriority = append(customPriority, tech)
			}
		}
		return customPriority
	}

	// Sadece mevcut teknolojileri döndür
	var result []string
	for _, tech := range priority {
		if availableTechs[tech] {
			result = append(result, tech)
		}
	}
	return result
}

// filterPlansByTech - Teknolojiye göre planları filtreler
func (r *recommender) filterPlansByTech(plans []models.HomePlan, tech string) []models.HomePlan {
	var filtered []models.HomePlan
	for _, plan := range plans {
		if strings.EqualFold(plan.Tech, tech) {
			filtered = append(filtered, plan)
		}
	}
	return filtered
}

// findCheapestPlan - En ucuz planı bulur
func (r *recommender) findCheapestPlan(plans []models.HomePlan) *models.HomePlan {
	if len(plans) == 0 {
		return nil
	}

	cheapest := plans[0]
	for _, plan := range plans {
		if plan.MonthlyPrice < cheapest.MonthlyPrice {
			cheapest = plan
		}
	}
	return &cheapest
}

// findNextTierPlan - Bir seviye yukarı planı bulur
func (r *recommender) findNextTierPlan(plans []models.HomePlan, current *models.HomePlan) *models.HomePlan {
	if current == nil || len(plans) <= 1 {
		return nil
	}

	// Hızına göre sırala
	sortedPlans := make([]models.HomePlan, len(plans))
	copy(sortedPlans, plans)
	sort.Slice(sortedPlans, func(i, j int) bool {
		return sortedPlans[i].DownMbps < sortedPlans[j].DownMbps
	})

	// Current plan'dan sonraki planı bul
	for i, plan := range sortedPlans {
		if plan.ID == current.ID && i+1 < len(sortedPlans) {
			return &sortedPlans[i+1]
		}
	}
	return nil
}

// chooseTVPlan - TV plan adaylarını seçer
func (r *recommender) chooseTVPlan(totalHDHours int, allTV []models.TVPlan) []*models.TVPlan {
	var candidates []*models.TVPlan

	if totalHDHours == 0 {
		// TV yok seçeneği
		candidates = append(candidates, nil)

		// En ucuz plan (küçük saatler için)
		cheapest := r.findCheapestTVPlan(allTV)
		if cheapest != nil {
			candidates = append(candidates, cheapest)
		}
	} else {
		// İhtiyacı karşılayan en ucuz plan
		fitPlan := r.findFitTVPlan(totalHDHours, allTV)
		if fitPlan != nil {
			candidates = append(candidates, fitPlan)
		}
	}

	return candidates
}

// findCheapestTVPlan - En ucuz TV planını bulur
func (r *recommender) findCheapestTVPlan(plans []models.TVPlan) *models.TVPlan {
	if len(plans) == 0 {
		return nil
	}

	cheapest := plans[0]
	for _, plan := range plans {
		if plan.MonthlyPrice < cheapest.MonthlyPrice {
			cheapest = plan
		}
	}
	return &cheapest
}

// findFitTVPlan - İhtiyacı karşılayan en ucuz TV planını bulur
func (r *recommender) findFitTVPlan(totalHDHours int, plans []models.TVPlan) *models.TVPlan {
	var fitPlans []models.TVPlan
	for _, plan := range plans {
		if plan.HDHoursIncl >= totalHDHours {
			fitPlans = append(fitPlans, plan)
		}
	}

	if len(fitPlans) == 0 {
		return nil
	}

	// En ucuz olanı seç
	cheapest := fitPlans[0]
	for _, plan := range fitPlans {
		if plan.MonthlyPrice < cheapest.MonthlyPrice {
			cheapest = plan
		}
	}
	return &cheapest
}

// calculateTotalHDHours - Toplam HD saatlerini hesaplar
func (r *recommender) calculateTotalHDHours(household []models.HouseholdLine) int {
	total := 0
	for _, line := range household {
		total += line.TVHDHours
	}
	return total
}

// generateCombinations - Tüm kombinasyonları oluşturur
func (r *recommender) generateCombinations(mobileCandidates [][]models.MobilePlan, homeCandidates []*models.HomePlan, tvCandidates []*models.TVPlan) []models.ComboResult {
	var combinations []models.ComboResult

	// Mobil kombinasyonları oluştur (beam search)
	mobileCombinations := r.generateMobileCombinations(mobileCandidates)

	// Her mobil kombinasyon için ev internet ve TV kombinasyonları
	for _, mobileCombo := range mobileCombinations {
		for _, homePlan := range homeCandidates {
			for _, tvPlan := range tvCandidates {
				combo := r.buildCombo(mobileCombo, homePlan, tvPlan)
				if combo != nil {
					combinations = append(combinations, *combo)
				}
			}
		}
	}

	return combinations
}

// generateMobileCombinations - Mobil kombinasyonları oluşturur (beam search)
func (r *recommender) generateMobileCombinations(candidates [][]models.MobilePlan) [][]models.MobilePlan {
	if len(candidates) == 0 {
		return [][]models.MobilePlan{}
	}

	// İlk hattın adayları
	result := make([][]models.MobilePlan, 0)
	for _, plan := range candidates[0] {
		result = append(result, []models.MobilePlan{plan})
	}

	// Her hattı ekle
	for i := 1; i < len(candidates); i++ {
		var newResult [][]models.MobilePlan

		for _, existing := range result {
			for _, plan := range candidates[i] {
				newCombo := make([]models.MobilePlan, len(existing)+1)
				copy(newCombo, existing)
				newCombo[len(existing)] = plan
				newResult = append(newResult, newCombo)
			}
		}

		// Beam search: en fazla 2 aday tut
		if len(newResult) > 2 {
			// En ucuz 2'sini seç
			sort.Slice(newResult, func(i, j int) bool {
				return r.calculateMobileTotal(newResult[i]) < r.calculateMobileTotal(newResult[j])
			})
			newResult = newResult[:2]
		}

		result = newResult
	}

	return result
}

// buildCombo - Kombinasyon oluşturur
func (r *recommender) buildCombo(mobilePlans []models.MobilePlan, homePlan *models.HomePlan, tvPlan *models.TVPlan) *models.ComboResult {
	// Mobil maliyetleri hesapla
	mobileItems := r.calculateMobileItems(mobilePlans)
	mobileTotal := r.calculateMobileTotal(mobilePlans)

	// Hat indirimi uygula
	lineDiscount := r.applyLineDiscount(mobileTotal, len(mobilePlans))
	mobileTotalAfterDiscount := mobileTotal * (1 - lineDiscount)

	// Ev internet maliyeti
	homeCost := 0.0
	if homePlan != nil {
		homeCost = homePlan.MonthlyPrice
	}

	// TV maliyeti
	tvCost := 0.0
	if tvPlan != nil {
		tvCost = tvPlan.MonthlyPrice
	}

	// Paketleme indirimi
	bundleDiscount := r.bundleDiscount(homePlan != nil, tvPlan != nil)
	subtotal := mobileTotalAfterDiscount + homeCost + tvCost
	grandTotal := subtotal * (1 - bundleDiscount)

	// Kombinasyon etiketi oluştur
	comboLabel := r.buildComboLabel(mobilePlans, homePlan, tvPlan)

	// Gerekçe oluştur
	reasoning := r.buildReasoning(mobilePlans, homePlan, tvPlan, lineDiscount, bundleDiscount)

	// Tasarruf hesapla (şimdilik 0)
	savings := 0.0

	return &models.ComboResult{
		ComboLabel: comboLabel,
		Items: models.ComboItems{
			Mobile: mobileItems,
			Home:   homePlan,
			TV:     tvPlan,
		},
		MonthlyTotal: grandTotal,
		Savings:      savings,
		Reasoning:    reasoning,
	}
}

// calculateMobileItems - Mobil hat öğelerini hesaplar
func (r *recommender) calculateMobileItems(plans []models.MobilePlan) []models.MobileLineItem {
	var items []models.MobileLineItem

	// Bu basit implementasyonda her plan için sabit bir line_id kullanıyoruz
	// Gerçek uygulamada household bilgisi kullanılmalı
	for i, plan := range plans {
		cost := plan.MonthlyPrice // Basit hesaplama
		items = append(items, models.MobileLineItem{
			LineID: fmt.Sprintf("L%d", i+1),
			Plan:   plan,
			Cost:   cost,
		})
	}

	return items
}

// calculateMobileTotal - Mobil toplam maliyeti hesaplar
func (r *recommender) calculateMobileTotal(plans []models.MobilePlan) float64 {
	total := 0.0
	for _, plan := range plans {
		total += plan.MonthlyPrice
	}
	return total
}

// applyLineDiscount - Hat indirimini uygular
func (r *recommender) applyLineDiscount(mobileTotal float64, n int) float64 {
	switch n {
	case 1:
		return 0.0
	case 2:
		return 0.05 // 5%
	default:
		return 0.10 // 10%
	}
}

// bundleDiscount - Paketleme indirimini hesaplar
func (r *recommender) bundleDiscount(hasHome, hasTV bool) float64 {
	if hasHome && hasTV {
		return 0.15 // 15%
	} else if hasHome {
		return 0.10 // 10%
	}
	return 0.0
}

// buildComboLabel - Kombinasyon etiketi oluşturur
func (r *recommender) buildComboLabel(mobilePlans []models.MobilePlan, homePlan *models.HomePlan, tvPlan *models.TVPlan) string {
	var parts []string

	// Ev internet
	if homePlan != nil {
		parts = append(parts, fmt.Sprintf("%s %dM", strings.Title(homePlan.Tech), homePlan.DownMbps))
	}

	// TV
	if tvPlan != nil {
		parts = append(parts, fmt.Sprintf("TV %s", tvPlan.Name))
	}

	// Mobil planlar
	if len(mobilePlans) > 0 {
		var mobileParts []string
		for _, plan := range mobilePlans {
			mobileParts = append(mobileParts, fmt.Sprintf("%dGB", plan.QuotaGB))
		}
		parts = append(parts, fmt.Sprintf("(%s)", strings.Join(mobileParts, ",")))
	}

	return strings.Join(parts, " + ")
}

// buildReasoning - Gerekçe oluşturur
func (r *recommender) buildReasoning(mobilePlans []models.MobilePlan, homePlan *models.HomePlan, tvPlan *models.TVPlan, lineDiscount, bundleDiscount float64) string {
	var reasons []string

	// Hat indirimi
	if lineDiscount > 0 {
		reasons = append(reasons, fmt.Sprintf("%d hat için %%%.0f indirim", len(mobilePlans), lineDiscount*100))
	}

	// Paketleme indirimi
	if bundleDiscount > 0 {
		reasons = append(reasons, fmt.Sprintf("paketleme %%%.0f indirim", bundleDiscount*100))
	}

	// Teknoloji seçimi
	if homePlan != nil {
		reasons = append(reasons, fmt.Sprintf("%s teknolojisi", homePlan.Tech))
	}

	if len(reasons) == 0 {
		return "En uygun fiyatlı kombinasyon"
	}

	return strings.Join(reasons, ", ")
}

// selectTop3 - En iyi 3 kombinasyonu seçer
func (r *recommender) selectTop3(combinations []models.ComboResult, currentTotal float64) []models.ComboResult {
	// Tasarruf hesapla
	for i := range combinations {
		if currentTotal > 0 {
			combinations[i].Savings = currentTotal - combinations[i].MonthlyTotal
		}
	}

	// Fiyata göre sırala
	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].MonthlyTotal < combinations[j].MonthlyTotal
	})

	// Deduplication (basit implementasyon)
	seen := make(map[string]bool)
	var unique []models.ComboResult

	for _, combo := range combinations {
		key := fmt.Sprintf("%s-%.2f", combo.ComboLabel, combo.MonthlyTotal)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, combo)
		}
	}

	// En iyi 3'ü döndür
	if len(unique) > 3 {
		return unique[:3]
	}
	return unique
}

// getCurrentServices - Mevcut hizmetleri alır
func (r *recommender) getCurrentServices(userID int) *models.CurrentServices {
	// Bu basit implementasyonda nil döndürüyoruz
	// Gerçek uygulamada catalog'dan alınmalı
	return nil
}

// calculateCurrentTotal - Mevcut toplam maliyeti hesaplar
func (r *recommender) calculateCurrentTotal(services *models.CurrentServices) float64 {
	if services == nil {
		return 0.0
	}
	// Basit hesaplama - gerçek uygulamada detaylı hesaplama yapılmalı
	return 0.0
}

// contains - Slice'ta eleman var mı kontrol eder
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
