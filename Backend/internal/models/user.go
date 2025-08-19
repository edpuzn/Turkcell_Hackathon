package models

// HouseholdLine (Hane Hattı) - Bir hanedeki telefon hattını temsil eder
// Her hattın beklenen GB kullanımı, dakika kullanımı ve TV HD saat bilgilerini içerir
type HouseholdLine struct {
	UserID       int    `json:"user_id" csv:"user_id"`             // Kullanıcı kimliği
	LineID       string `json:"line_id" csv:"line_id"`             // Hat kimliği
	ExpectedGB   int    `json:"expected_gb" csv:"expected_gb"`     // Beklenen GB kullanımı
	ExpectedMin  int    `json:"expected_min" csv:"expected_min"`   // Beklenen dakika kullanımı
	TVHDHours    int    `json:"tv_hd_hours" csv:"tv_hd_hours"`     // TV HD kanal saatleri
}

// User (Kullanıcı) - Sisteme kayıtlı kullanıcıyı temsil eder
// Kullanıcının temel bilgileri ve mevcut paket bilgisini içerir
type User struct {
	ID                 int    `json:"id" csv:"id"`                           // Kullanıcı kimliği
	Name               string `json:"name" csv:"name"`                       // Kullanıcı adı
	AddressID          string `json:"address_id" csv:"address_id"`           // Adres kimliği
	CurrentBundleLabel string `json:"current_bundle_label" csv:"current_bundle_label"` // Mevcut paket etiketi
}

// CurrentServices (Mevcut Hizmetler) - Kullanıcının mevcut hizmetlerini temsil eder
// Kullanıcının hangi hizmetlere sahip olduğunu ve detaylarını gösterir
type CurrentServices struct {
	UserID        int    `json:"user_id" csv:"user_id"`               // Kullanıcı kimliği
	HasHome       bool   `json:"has_home" csv:"has_home"`             // Ev interneti var mı?
	HomeTech      string `json:"home_tech" csv:"home_tech"`           // Ev internet teknolojisi
	HomeSpeed     int    `json:"home_speed" csv:"home_speed"`         // Ev internet hızı
	HasTV         bool   `json:"has_tv" csv:"has_tv"`                 // TV hizmeti var mı?
	MobilePlanIDs string `json:"mobile_plan_ids" csv:"mobile_plan_ids"` // Mevcut mobil plan kimlikleri (virgülle ayrılmış)
}
