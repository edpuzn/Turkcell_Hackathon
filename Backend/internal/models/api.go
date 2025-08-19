package models

// RecommendationInput (Öneri Girişi) - Öneri API'si için giriş verilerini temsil eder
// Kullanıcının hane bilgileri ve tercih ettiği teknolojileri içerir
type RecommendationInput struct {
	UserID    int              `json:"user_id"`     // Kullanıcı kimliği
	AddressID string           `json:"address_id"`  // Adres kimliği
	Household []HouseholdLine  `json:"household"`   // Hane halkı hatları
	PreferTech []string        `json:"prefer_tech,omitempty"` // Tercih edilen teknolojiler (fiber, vdsl, fwa)
}

// ComboResult (Kombinasyon Sonucu) - Önerilen hizmet kombinasyonunu temsil eder
// En uygun mobil, ev internet ve TV planlarının birleşimini gösterir
type ComboResult struct {
	ComboLabel   string                 `json:"combo_label"`   // Kombinasyon etiketi
	Items        ComboItems             `json:"items"`         // Kombinasyon içindeki hizmetler
	MonthlyTotal float64                `json:"monthly_total"` // Aylık toplam ücret
	Savings      float64                `json:"savings"`       // Tasarruf miktarı
	Reasoning    string                 `json:"reasoning"`     // Öneri gerekçesi
}

// ComboItems (Kombinasyon Öğeleri) - Kombinasyon içindeki hizmetleri temsil eder
// Mobil hatlar, ev internet ve TV planlarını içerir
type ComboItems struct {
	Mobile []MobileLineItem `json:"mobile"` // Mobil hat planları
	Home   *HomePlan        `json:"home,omitempty"` // Ev internet planı (opsiyonel)
	TV     *TVPlan          `json:"tv,omitempty"`   // TV planı (opsiyonel)
}

// MobileLineItem (Mobil Hat Öğesi) - Bir mobil hattın plan ve maliyet bilgilerini temsil eder
// Her hat için seçilen plan ve o hattın aylık maliyetini gösterir
type MobileLineItem struct {
	LineID string     `json:"line_id"` // Hat kimliği
	Plan   MobilePlan `json:"plan"`    // Seçilen mobil plan
	Cost   float64    `json:"cost"`    // Aylık maliyet
}

// RecommendationResponse (Öneri Yanıtı) - Öneri API'sinden dönen yanıtı temsil eder
// En uygun 3 kombinasyonu içerir
type RecommendationResponse struct {
	Top3 []ComboResult `json:"top3"` // En iyi 3 kombinasyon
}
