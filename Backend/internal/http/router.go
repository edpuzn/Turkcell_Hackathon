package http

import (
	"encoding/json"
	"turkcell-hackathon-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Router - HTTP router yapılandırması
type Router struct {
	app     *fiber.App
	handler *Handler
	logger  *logrus.Logger
}

// NewRouter - Yeni router oluşturur
func NewRouter(catalogService *service.CatalogService, logger *logrus.Logger) *Router {
	app := fiber.New(fiber.Config{
		AppName:      "Turkcell Hackathon Backend",
		ServerHeader: "Turkcell-Hackathon",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	handler := NewHandler(catalogService)

	router := &Router{
		app:     app,
		handler: handler,
		logger:  logger,
	}

	router.setupMiddleware()
	router.setupRoutes()

	return router
}

// setupMiddleware - Middleware'leri yapılandırır
func (r *Router) setupMiddleware() {
	// Tüm middleware'leri kaydet (CORS dahil)
	RegisterMiddlewares(r.app, r.logger)
}

// setupRoutes - Route'ları yapılandırır
func (r *Router) setupRoutes() {
	// Health check
	r.app.Get("/health", r.handler.HealthCheck)

	// API routes
	api := r.app.Group("/api")

	// Catalog routes
	catalog := api.Group("/catalog")
	catalog.Get("/summary", r.handler.GetCatalogSummary)

	// Coverage routes
	api.Get("/coverage/:address_id", r.handler.GetCoverageByAddress)

	// Recommendation routes
	api.Post("/recommendation", r.handler.PostRecommendation)
}

// App - Fiber app'ini döndürür
func (r *Router) App() *fiber.App {
	return r.app
}
