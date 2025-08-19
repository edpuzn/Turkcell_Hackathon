package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"turkcell-hackathon-backend/internal/csvload"
	"turkcell-hackathon-backend/internal/models"
	"turkcell-hackathon-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Test için örnek katalog oluştur
func createTestCatalog() *csvload.Catalog {
	catalog := &csvload.Catalog{
		Coverages: []models.Coverage{
			{AddressID: "ADR-1001", Tech: "fiber", DownMbps: 1000},
			{AddressID: "ADR-1001", Tech: "vdsl", DownMbps: 100},
		},
		MobilePlans: []models.MobilePlan{
			{ID: 1, Name: "Test Plan", QuotaGB: 10, QuotaMin: 500, MonthlyPrice: 29.99},
		},
		HomePlans: []models.HomePlan{
			{ID: 1, Name: "Test Home", Tech: "fiber", DownMbps: 100, MonthlyPrice: 89.99},
		},
		TVPlans: []models.TVPlan{
			{ID: 1, Name: "Test TV", HDHoursIncl: 100, MonthlyPrice: 29.99},
		},
		BundlingRules: []models.BundlingRule{
			{ID: 1, RuleType: "test", AppliesTo: "test", DiscountPercent: 10.0},
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
	
	// Index'leri oluştur
	catalog.BuildIndexes()
	return catalog
}

func TestHealthCheck(t *testing.T) {
	// Test katalog oluştur
	catalog := createTestCatalog()
	
	// Service ve handler oluştur
	catalogService := service.NewCatalogService(catalog)
	handler := NewHandler(catalogService)

	// Fiber app oluştur
	app := fiber.New()
	app.Get("/health", handler.HealthCheck)

	// Test request oluştur
	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Response body'yi kontrol et
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestGetCatalogSummary(t *testing.T) {
	// Test katalog oluştur
	catalog := createTestCatalog()
	
	// Service ve handler oluştur
	catalogService := service.NewCatalogService(catalog)
	handler := NewHandler(catalogService)

	// Fiber app oluştur
	app := fiber.New()
	app.Get("/summary", handler.GetCatalogSummary)

	// Test request oluştur
	req := httptest.NewRequest("GET", "/summary", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Response body'yi kontrol et
	var response map[string]int
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	
	// Gerekli alanların varlığını kontrol et
	assert.Contains(t, response, "coverages")
	assert.Contains(t, response, "mobile_plans")
	assert.Contains(t, response, "home_plans")
	assert.Contains(t, response, "tv_plans")
	assert.Contains(t, response, "bundling_rules")
	assert.Contains(t, response, "install_slots")
	assert.Contains(t, response, "users")
	assert.Contains(t, response, "household_lines")
	assert.Contains(t, response, "current_services")
	
	// Değerleri kontrol et
	assert.Equal(t, 2, response["coverages"])
	assert.Equal(t, 1, response["mobile_plans"])
	assert.Equal(t, 1, response["home_plans"])
}

func TestGetCoverageByAddress(t *testing.T) {
	// Test katalog oluştur
	catalog := createTestCatalog()
	
	// Service ve handler oluştur
	catalogService := service.NewCatalogService(catalog)
	handler := NewHandler(catalogService)

	// Fiber app oluştur
	app := fiber.New()
	app.Get("/coverage/:address_id", handler.GetCoverageByAddress)

	// Test request oluştur - mevcut adres
	req := httptest.NewRequest("GET", "/coverage/ADR-1001", nil)
	resp, err := app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Response body'yi kontrol et
	var coverages []models.Coverage
	err = json.NewDecoder(resp.Body).Decode(&coverages)
	assert.NoError(t, err)
	assert.Len(t, coverages, 2)
	assert.Equal(t, "ADR-1001", coverages[0].AddressID)
	assert.Equal(t, "fiber", coverages[0].Tech)
	assert.Equal(t, "vdsl", coverages[1].Tech)

	// Test request oluştur - olmayan adres
	req = httptest.NewRequest("GET", "/coverage/ADR-9999", nil)
	resp, err = app.Test(req)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Response body'yi kontrol et - boş array olmalı
	err = json.NewDecoder(resp.Body).Decode(&coverages)
	assert.NoError(t, err)
	assert.Len(t, coverages, 0)
}
