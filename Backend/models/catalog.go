package models

// Coverage (Kapsama) - Bir adres için kapsama bilgilerini temsil eder
// Adres bazında hangi teknolojinin (fiber, vdsl, fwa) mevcut olduğunu ve
// indirme hızını bilgilerini tutar
type Coverage struct {
	AddressID string  `json:"address_id"` // Adres kimliği
	Tech      string  `json:"tech"`       // Teknoloji türü (fiber, vdsl, fwa)
	DownMbps  float64 `json:"down_mbps"`  // İndirme hızı (Mbps cinsinden)
}

// MobilePlan (Mobil Plan) - Mobil telefon planını temsil eder
// Her planın GB kotası, dakika kotası, aylık ücreti ve
// kota aşımı ücretlerini içerir
type MobilePlan struct {
	ID          string  `json:"id"`           // Plan kimliği
	Name        string  `json:"name"`         // Plan adı
	QuotaGB     int     `json:"quota_gb"`     // GB kotası
	QuotaMin    int     `json:"quota_min"`    // Dakika kotası
	MonthlyPrice float64 `json:"monthly_price"` // Aylık ücret
	OverageGB   float64 `json:"overage_gb"`   // GB aşımı ücreti
	OverageMin  float64 `json:"overage_min"`  // Dakika aşımı ücreti
}

// HomePlan (Ev İnternet Planı) - Ev internet planını temsil eder
// Fiber, VDSL veya FWA teknolojilerinden birini kullanarak
// ev internet hizmeti sunan planları tanımlar
type HomePlan struct {
	ID          string  `json:"id"`           // Plan kimliği
	Name        string  `json:"name"`         // Plan adı
	Tech        string  `json:"tech"`         // Teknoloji (fiber, vdsl, fwa)
	DownMbps    float64 `json:"down_mbps"`    // İndirme hızı (Mbps)
	MonthlyPrice float64 `json:"monthly_price"` // Aylık ücret
	InstallFee  float64 `json:"install_fee"`  // Kurulum ücreti
}

// TVPlan (TV Planı) - Dijital TV planını temsil eder
// HD kanal saatleri ve aylık ücret bilgilerini içerir
type TVPlan struct {
	ID            string  `json:"id"`             // Plan kimliği
	Name          string  `json:"name"`           // Plan adı
	HDHoursIncl   int     `json:"hd_hours_incl"`  // Dahil HD kanal saatleri
	MonthlyPrice  float64 `json:"monthly_price"`  // Aylık ücret
}

// BundlingRule (Paketleme Kuralı) - Paketleme indirim kurallarını temsil eder
// Birden fazla hizmet alındığında uygulanacak indirim kurallarını tanımlar
// Örnek: Mobil + Ev internet + TV alındığında %20 indirim
type BundlingRule struct {
	ID             string  `json:"id"`              // Kural kimliği
	RuleType       string  `json:"rule_type"`       // Kural türü
	AppliesTo      string  `json:"applies_to"`      // Hangi hizmetlere uygulanır
	DiscountPercent float64 `json:"discount_percent"` // İndirim yüzdesi
	Description    string  `json:"description"`     // Kural açıklaması
}
