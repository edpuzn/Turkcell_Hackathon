package http

import (
	"fmt"
	"turkcell-hackathon-backend/internal/models"
	"turkcell-hackathon-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

// PostRecommendation - Öneri endpoint'i
func (h *Handler) PostRecommendation(c *fiber.Ctx) error {
	// JSON'ı parse et
	var input models.RecommendationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	// Validasyon
	if err := h.validateRecommendationInput(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Öneri motorunu çağır
	recommender := service.NewRecommender(h.catalogService.GetCatalog())
	response, err := recommender.Recommend(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// validateRecommendationInput - Öneri girişini doğrular
func (h *Handler) validateRecommendationInput(input models.RecommendationInput) error {
	// Address ID kontrolü
	if input.AddressID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "address_id is required")
	}

	// Household kontrolü
	if len(input.Household) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "household cannot be empty")
	}

	// Her household line'ı kontrol et
	for i, line := range input.Household {
		if line.LineID == "" {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("household[%d].line_id is required", i))
		}
		if line.ExpectedGB < 0 {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("household[%d].expected_gb must be non-negative", i))
		}
		if line.ExpectedMin < 0 {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("household[%d].expected_min must be non-negative", i))
		}
		if line.TVHDHours < 0 {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("household[%d].tv_hd_hours must be non-negative", i))
		}
	}

	// PreferTech kontrolü (opsiyonel)
	if len(input.PreferTech) > 0 {
		validTechs := map[string]bool{"fiber": true, "vdsl": true, "fwa": true}
		for i, tech := range input.PreferTech {
			if !validTechs[tech] {
				return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("prefer_tech[%d] must be one of: fiber, vdsl, fwa", i))
			}
		}
	}

	return nil
}
