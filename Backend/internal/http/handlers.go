package http

import (
	"turkcell-hackathon-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Handler - HTTP handler'ları
type Handler struct {
	catalogService *service.CatalogService
}

// NewHandler - Yeni handler oluşturur
func NewHandler(catalogService *service.CatalogService) *Handler {
	return &Handler{
		catalogService: catalogService,
	}
}

// HealthCheck - Sağlık kontrolü endpoint'i
func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

// GetCatalogSummary - Katalog özeti endpoint'i
func (h *Handler) GetCatalogSummary(c *fiber.Ctx) error {
	summary := h.catalogService.Summary()

	return c.Status(fiber.StatusOK).JSON(summary)
}

// GetCoverageByAddress - Adres bazında kapsama bilgileri endpoint'i
func (h *Handler) GetCoverageByAddress(c *fiber.Ctx) error {
	addressID := c.Params("address_id")
	if addressID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "address_id parameter is required",
		})
	}

	coverages := h.catalogService.CoverageByAddress(addressID)

	return c.Status(fiber.StatusOK).JSON(coverages)
}
