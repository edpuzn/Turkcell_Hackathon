package service

import (
	"turkcell-hackathon-backend/internal/csvload"
	"turkcell-hackathon-backend/internal/models"
)

// CatalogService - Katalog servis katmanı
type CatalogService struct {
	catalog *csvload.Catalog
}

// NewCatalogService - Yeni katalog servisi oluşturur
func NewCatalogService(catalog *csvload.Catalog) *CatalogService {
	return &CatalogService{
		catalog: catalog,
	}
}

// Summary - Tüm veri setlerinin sayılarını döndürür
func (s *CatalogService) Summary() map[string]int {
	return s.catalog.Counts()
}

// CoverageByAddress - Belirli bir adres için kapsama bilgilerini döndürür
func (s *CatalogService) CoverageByAddress(addressID string) []models.Coverage {
	return s.catalog.ByAddressCoverage(addressID)
}

// GetCatalog - Catalog'u döndürür
func (s *CatalogService) GetCatalog() *csvload.Catalog {
	return s.catalog
}
