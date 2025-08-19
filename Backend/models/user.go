package models

import "time"

// InstallSlot (Kurulum Zaman Dilimi) - Kurulum için müsait zaman dilimini temsil eder
// Teknisyenlerin hangi adreste, hangi teknoloji için ne zaman müsait olduğunu gösterir
type InstallSlot struct {
	SlotID    string    `json:"slot_id"`    // Zaman dilimi kimliği
	AddressID string    `json:"address_id"` // Adres kimliği
	Tech      string    `json:"tech"`       // Teknoloji türü
	StartISO  time.Time `json:"start_iso"`  // Başlangıç zamanı
	EndISO    time.Time `json:"end_iso"`    // Bitiş zamanı
}

// HouseholdLine (Hane Hattı) - Bir hanedeki telefon hattını temsil eder
// Her hattın beklenen GB kullanımı, dakika kullanımı ve TV HD saat bilgilerini içerir
type HouseholdLine struct {
	LineID        string `json:"line_id"`        // Hat kimliği
	ExpectedGB    int    `json:"expected_gb"`    // Beklenen GB kullanımı
	ExpectedMin   int    `json:"expected_min"`   // Beklenen dakika kullanımı
	TVHDHours     int    `json:"tv_hd_hours"`    // TV HD kanal saatleri
}

// User (Kullanıcı) - Sisteme kayıtlı kullanıcıyı temsil eder
// Kullanıcının temel bilgileri ve mevcut paket bilgisini içerir
type User struct {
	ID                  string `json:"id"`                   // Kullanıcı kimliği
	Name                string `json:"name"`                 // Kullanıcı adı
	AddressID           string `json:"address_id"`           // Adres kimliği
	CurrentBundleLabel  string `json:"current_bundle_label"` // Mevcut paket etiketi
}

// CurrentServices (Mevcut Hizmetler) - Kullanıcının mevcut hizmetlerini temsil eder
// Kullanıcının hangi hizmetlere sahip olduğunu ve detaylarını gösterir
type CurrentServices struct {
	UserID        string   `json:"user_id"`        // Kullanıcı kimliği
	HasHome       bool     `json:"has_home"`       // Ev interneti var mı?
	HomeTech      string   `json:"home_tech"`      // Ev internet teknolojisi
	HomeSpeed     float64  `json:"home_speed"`     // Ev internet hızı
	HasTV         bool     `json:"has_tv"`         // TV hizmeti var mı?
	MobilePlanIDs []string `json:"mobile_plan_ids"` // Mevcut mobil plan kimlikleri
}
