package service

import (
	"context"
	"fmt"
	"testing"

	"turkcell-hackathon-backend/internal/csvload"
	"turkcell-hackathon-backend/internal/models"

	"github.com/stretchr/testify/assert"
)

// Test için örnek katalog oluştur
func createTestCatalogForRecommender() *csvload.Catalog {
	catalog := &csvload.Catalog{
		Coverages: []models.Coverage{
			{AddressID: "ADR-1001", Tech: "fiber", DownMbps: 1000},
			{AddressID: "ADR-1001", Tech: "vdsl", DownMbps: 100},
		},
		MobilePlans: []models.MobilePlan{
			{ID: 1, Name: "Temel Plan", QuotaGB: 5, QuotaMin: 500, MonthlyPrice: 29.99, OverageGB: 0.50, OverageMin: 0.10},
			{ID: 2, Name: "Orta Plan", QuotaGB: 15, QuotaMin: 1000, MonthlyPrice: 49.99, OverageGB: 0.30, OverageMin: 0.08},
			{ID: 3, Name: "İleri Plan", QuotaGB: 30, QuotaMin: 2000, MonthlyPrice: 79.99, OverageGB: 0.20, OverageMin: 0.05},
		},
		HomePlans: []models.HomePlan{
			{ID: 1, Name: "Fiber 100M", Tech: "fiber", DownMbps: 100, MonthlyPrice: 89.99, InstallFee: 0.00},
			{ID: 2, Name: "Fiber 500M", Tech: "fiber", DownMbps: 500, MonthlyPrice: 149.99, InstallFee: 0.00},
			{ID: 3, Name: "VDSL 50M", Tech: "vdsl", DownMbps: 50, MonthlyPrice: 69.99, InstallFee: 50.00},
		},
		TVPlans: []models.TVPlan{
			{ID: 1, Name: "Temel TV", HDHoursIncl: 100, MonthlyPrice: 29.99},
			{ID: 2, Name: "Orta TV", HDHoursIncl: 200, MonthlyPrice: 49.99},
			{ID: 3, Name: "İleri TV", HDHoursIncl: 500, MonthlyPrice: 79.99},
		},
		BundlingRules: []models.BundlingRule{
			{ID: 1, RuleType: "mobile_home", AppliesTo: "Mobile + Home", DiscountPercent: 10.0, Description: "Mobil ve ev internet paketleme indirimi"},
		},
		InstallSlots: []models.InstallSlot{
			{SlotID: "SLOT-001", AddressID: "ADR-1001", Tech: "fiber"},
		},
		Users: []models.User{
			{ID: 1, Name: "Test User", AddressID: "ADR-1001"},
		},
		HouseholdLines: []models.HouseholdLine{
			{UserID: 1, LineID: "LINE-001", ExpectedGB: 10, ExpectedMin: 500},
		},
		CurrentServices: []models.CurrentServices{
			{UserID: 1, HasHome: true, HomeTech: "fiber", HomeSpeed: 100, HasTV: true},
		},
	}
	
	catalog.BuildIndexes()
	return catalog
}

func TestRecommender_SingleLineExactFit(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-1001",
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0},
		},
	}

	response, err := recommender.Recommend(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Top3)
	assert.Len(t, response.Top3, 3) // En iyi 3 kombinasyon

	combo := response.Top3[0]
	assert.NotEmpty(t, combo.ComboLabel)
	assert.NotEmpty(t, combo.Reasoning)
	assert.True(t, combo.MonthlyTotal > 0)
	assert.Equal(t, 0.0, combo.Savings) // Mevcut hizmet yok
}

func TestRecommender_TwoLinesWithDiscount(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-1001",
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0},
			{LineID: "L2", ExpectedGB: 5, ExpectedMin: 200, TVHDHours: 0},
		},
	}

	response, err := recommender.Recommend(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Top3)

	combo := response.Top3[0]
	assert.Len(t, combo.Items.Mobile, 2) // 2 hat olmalı
	assert.Contains(t, combo.Reasoning, "2 hat için %5 indirim") // %5 hat indirimi
}

func TestRecommender_ThreeLinesWithBundleDiscount(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-1001",
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 50},
			{LineID: "L2", ExpectedGB: 5, ExpectedMin: 200, TVHDHours: 30},
			{LineID: "L3", ExpectedGB: 15, ExpectedMin: 800, TVHDHours: 100},
		},
	}

	response, err := recommender.Recommend(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Top3)

	combo := response.Top3[0]
	assert.Len(t, combo.Items.Mobile, 3) // 3 hat olmalı
	assert.Contains(t, combo.Reasoning, "3 hat için %10 indirim") // %10 hat indirimi
	assert.Contains(t, combo.Reasoning, "paketleme %15 indirim") // %15 paketleme indirimi
	assert.NotNil(t, combo.Items.Home) // Ev internet olmalı
	assert.NotNil(t, combo.Items.TV)   // TV olmalı
}

func TestRecommender_NoTVvsWithTV(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-1001",
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0}, // TV yok
		},
	}

	response, err := recommender.Recommend(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Top3)

	// TV olmayan kombinasyonlar olmalı
	foundNoTV := false
	for _, combo := range response.Top3 {
		if combo.Items.TV == nil {
			foundNoTV = true
			break
		}
	}
	assert.True(t, foundNoTV, "Should have at least one combination without TV")
}

func TestRecommender_PreferTechFallback(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:     1,
		AddressID:  "ADR-1001",
		PreferTech: []string{"vdsl", "fiber"}, // VDSL öncelikli
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0},
		},
	}

	response, err := recommender.Recommend(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.Top3)

	// VDSL planı seçilmeli
	combo := response.Top3[0]
	if combo.Items.Home != nil {
		assert.Equal(t, "vdsl", combo.Items.Home.Tech)
	}
}

func TestRecommender_InvalidAddress(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-9999", // Olmayan adres
		Household: []models.HouseholdLine{
			{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0},
		},
	}

	_, err := recommender.Recommend(context.Background(), input)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "has no coverage")
}

func TestRecommender_EmptyHousehold(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	input := models.RecommendationInput{
		UserID:    1,
		AddressID: "ADR-1001",
		Household: []models.HouseholdLine{}, // Boş household
	}

	_, err := recommender.Recommend(context.Background(), input)

	assert.Error(t, err)
}

func TestRecommender_LineDiscountCalculation(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	testCases := []struct {
		name     string
		lines    int
		expected float64
	}{
		{"1 line", 1, 0.0},
		{"2 lines", 2, 0.05},
		{"3 lines", 3, 0.10},
		{"4 lines", 4, 0.10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			household := make([]models.HouseholdLine, tc.lines)
			for i := 0; i < tc.lines; i++ {
				household[i] = models.HouseholdLine{
					LineID:      fmt.Sprintf("L%d", i+1),
					ExpectedGB:  10,
					ExpectedMin: 500,
					TVHDHours:   0,
				}
			}

			input := models.RecommendationInput{
				UserID:    1,
				AddressID: "ADR-1001",
				Household: household,
			}

			response, err := recommender.Recommend(context.Background(), input)
			assert.NoError(t, err)
			assert.NotEmpty(t, response.Top3)

			combo := response.Top3[0]
			assert.Len(t, combo.Items.Mobile, tc.lines)
		})
	}
}

func TestRecommender_BundleDiscountCalculation(t *testing.T) {
	catalog := createTestCatalogForRecommender()
	recommender := NewRecommender(catalog)

	testCases := []struct {
		name     string
		hasHome  bool
		hasTV    bool
		expected float64
	}{
		{"mobile only", false, false, 0.0},
		{"mobile + home", true, false, 0.10},
		{"mobile + home + tv", true, true, 0.15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			household := []models.HouseholdLine{
				{LineID: "L1", ExpectedGB: 10, ExpectedMin: 500, TVHDHours: 0},
			}

			if tc.hasTV {
				household[0].TVHDHours = 100
			}

			input := models.RecommendationInput{
				UserID:    1,
				AddressID: "ADR-1001",
				Household: household,
			}

			response, err := recommender.Recommend(context.Background(), input)
			assert.NoError(t, err)
			assert.NotEmpty(t, response.Top3)

			combo := response.Top3[0]
			if tc.hasHome {
				assert.NotNil(t, combo.Items.Home)
			}
			if tc.hasTV {
				assert.NotNil(t, combo.Items.TV)
			}
		})
	}
}
