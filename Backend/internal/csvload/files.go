package csvload

// CSV dosya adları ve header tanımları
const (
	CoverageFile        = "coverage.csv"
	MobilePlansFile     = "mobile_plans.csv"
	HomePlansFile       = "home_plans.csv"
	TVPlansFile         = "tv_plans.csv"
	BundlingRulesFile   = "bundling_rules.csv"
	InstallSlotsFile    = "install_slots.csv"
	UsersFile           = "users.csv"
	HouseholdFile       = "household.csv"
	CurrentServicesFile = "current_services.csv"
)

// CSVHeader tanımları - tam eşleşme için case-sensitive
var (
	CoverageHeaders        = []string{"address_id", "city", "district", "fiber", "vdsl", "fwa"}
	MobilePlansHeaders     = []string{"plan_id", "plan_name", "quota_gb", "quota_min", "monthly_price", "overage_gb", "overage_min"}
	HomePlansHeaders       = []string{"home_id", "name", "tech", "down_mbps", "monthly_price", "install_fee"}
	TVPlansHeaders         = []string{"tv_id", "name", "hd_hours_included", "monthly_price"}
	BundlingRulesHeaders   = []string{"rule_id", "rule_type", "description", "discount_percent", "applies_to"}
	InstallSlotsHeaders    = []string{"slot_id", "address_id", "slot_start", "slot_end", "tech"}
	UsersHeaders           = []string{"user_id", "name", "address_id", "current_bundle_label"}
	HouseholdHeaders       = []string{"user_id", "line_id", "expected_gb", "expected_min", "tv_hd_hours"}
	CurrentServicesHeaders = []string{"user_id", "has_home", "home_tech", "home_speed", "has_tv", "mobile_plan_ids"}
)

// FileConfig - CSV dosya konfigürasyonu
type FileConfig struct {
	Filename string
	Headers  []string
}

// GetAllFileConfigs - Tüm CSV dosya konfigürasyonlarını döndürür
func GetAllFileConfigs() []FileConfig {
	return []FileConfig{
		{Filename: CoverageFile, Headers: CoverageHeaders},
		{Filename: MobilePlansFile, Headers: MobilePlansHeaders},
		{Filename: HomePlansFile, Headers: HomePlansHeaders},
		{Filename: TVPlansFile, Headers: TVPlansHeaders},
		{Filename: BundlingRulesFile, Headers: BundlingRulesHeaders},
		{Filename: InstallSlotsFile, Headers: InstallSlotsHeaders},
		{Filename: UsersFile, Headers: UsersHeaders},
		{Filename: HouseholdFile, Headers: HouseholdHeaders},
		{Filename: CurrentServicesFile, Headers: CurrentServicesHeaders},
	}
}
