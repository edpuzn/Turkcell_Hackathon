package models

// Coverage (Kapsama) - Bir adres için kapsama bilgilerini temsil eder
// Adres bazında hangi teknolojinin (fiber, vdsl, fwa) mevcut olduğunu ve
// indirme hızını bilgilerini tutar
type Coverage struct {
	AddressID string `json:"address_id" csv:"address_id"` // Adres kimliği
	Tech      string `json:"tech" csv:"tech"`             // Teknoloji türü (fiber, vdsl, fwa)
	DownMbps  int    `json:"down_mbps" csv:"down_mbps"`   // İndirme hızı (Mbps cinsinden)
}

// MobilePlan (Mobil Plan) - Mobil telefon planını temsil eder
// Her planın GB kotası, dakika kotası, aylık ücreti ve
// kota aşımı ücretlerini içerir
type MobilePlan struct {
	ID           int     `json:"id" csv:"id"`                     // Plan kimliği
	Name         string  `json:"name" csv:"name"`                 // Plan adı
	QuotaGB      int     `json:"quota_gb" csv:"quota_gb"`         // GB kotası
	QuotaMin     int     `json:"quota_min" csv:"quota_min"`       // Dakika kotası
	MonthlyPrice float64 `json:"monthly_price" csv:"monthly_price"` // Aylık ücret
	OverageGB    float64 `json:"overage_gb" csv:"overage_gb"`     // GB aşımı ücreti
	OverageMin   float64 `json:"overage_min" csv:"overage_min"`   // Dakika aşımı ücreti
}

// HomePlan (Ev İnternet Planı) - Ev internet planını temsil eder
// Fiber, VDSL veya FWA teknolojilerinden birini kullanarak
// ev internet hizmeti sunan planları tanımlar
type HomePlan struct {
	ID           int     `json:"id" csv:"id"`                     // Plan kimliği
	Name         string  `json:"name" csv:"name"`                 // Plan adı
	Tech         string  `json:"tech" csv:"tech"`                 // Teknoloji (fiber, vdsl, fwa)
	DownMbps     int     `json:"down_mbps" csv:"down_mbps"`       // İndirme hızı (Mbps)
	MonthlyPrice float64 `json:"monthly_price" csv:"monthly_price"` // Aylık ücret
	InstallFee   float64 `json:"install_fee" csv:"install_fee"`   // Kurulum ücreti
}

// TVPlan (TV Planı) - Dijital TV planını temsil eder
// HD kanal saatleri ve aylık ücret bilgilerini içerir
type TVPlan struct {
	ID           int     `json:"id" csv:"id"`                       // Plan kimliği
	Name         string  `json:"name" csv:"name"`                   // Plan adı
	HDHoursIncl  int     `json:"hd_hours_incl" csv:"hd_hours_incl"` // Dahil HD kanal saatleri
	MonthlyPrice float64 `json:"monthly_price" csv:"monthly_price"` // Aylık ücret
}

// BundlingRule (Paketleme Kuralı) - Paketleme indirim kurallarını temsil eder
// Birden fazla hizmet alındığında uygulanacak indirim kurallarını tanımlar
// Örnek: Mobil + Ev internet + TV alındığında %20 indirim
type BundlingRule struct {
	ID              int     `json:"id" csv:"id"`                           // Kural kimliği
	RuleType        string  `json:"rule_type" csv:"rule_type"`             // Kural türü
	AppliesTo       string  `json:"applies_to" csv:"applies_to"`           // Hangi hizmetlere uygulanır
	DiscountPercent float64 `json:"discount_percent" csv:"discount_percent"` // İndirim yüzdesi
	Description     string  `json:"description" csv:"description"`         // Kural açıklaması
}

// InstallSlot (Kurulum Zaman Dilimi) - Kurulum için müsait zaman dilimini temsil eder
// Teknisyenlerin hangi adreste, hangi teknoloji için ne zaman müsait olduğunu gösterir
type InstallSlot struct {
	SlotID    string `json:"slot_id" csv:"slot_id"`       // Zaman dilimi kimliği
	AddressID string `json:"address_id" csv:"address_id"` // Adres kimliği
	Tech      string `json:"tech" csv:"tech"`             // Teknoloji türü
	StartISO  string `json:"start_iso" csv:"start_iso"`   // Başlangıç zamanı
	EndISO    string `json:"end_iso" csv:"end_iso"`       // Bitiş zamanı
}
