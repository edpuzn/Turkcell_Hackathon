package http

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCORSHeaders(t *testing.T) {
	// Test app oluştur
	app := fiber.New()
	
	// Test logger oluştur
	logger := &logrus.Logger{}
	
	// Middleware'leri kaydet
	RegisterMiddlewares(app, logger)
	
	// Test route ekle
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	
	// OPTIONS isteği test et
	t.Run("OPTIONS preflight request", func(t *testing.T) {
		req := httptest.NewRequest("OPTIONS", "/test", nil)
		req.Header.Set("Origin", "http://localhost:5173")
		req.Header.Set("Access-Control-Request-Method", "GET")
		req.Header.Set("Access-Control-Request-Headers", "Content-Type")
		
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 204, resp.StatusCode)
		
		// CORS header'larını kontrol et
		assert.Equal(t, "http://localhost:5173", resp.Header.Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "GET,POST,PUT,DELETE,OPTIONS", resp.Header.Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "Origin,Content-Type,Accept,Authorization,X-Request-ID", resp.Header.Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "86400", resp.Header.Get("Access-Control-Max-Age"))
	})
	
	// GET isteği test et
	t.Run("GET request with CORS headers", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Origin", "http://localhost:5173")
		
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		// CORS header'larını kontrol et
		assert.Equal(t, "http://localhost:5173", resp.Header.Get("Access-Control-Allow-Origin"))
	})
}

func TestCORSMultipleOrigins(t *testing.T) {
	// Environment değişkenini geçici olarak ayarla
	originalEnv := os.Getenv("ALLOWED_ORIGINS")
	defer os.Setenv("ALLOWED_ORIGINS", originalEnv)
	
	os.Setenv("ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:3000")
	
	// Test app oluştur
	app := fiber.New()
	logger := &logrus.Logger{}
	
	// Middleware'leri kaydet
	RegisterMiddlewares(app, logger)
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	
	// İlk origin test et
	t.Run("First allowed origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Origin", "http://localhost:5173")
		
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "http://localhost:5173", resp.Header.Get("Access-Control-Allow-Origin"))
	})
	
	// İkinci origin test et
	t.Run("Second allowed origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "http://localhost:3000", resp.Header.Get("Access-Control-Allow-Origin"))
	})
	
	// İzin verilmeyen origin test et
	t.Run("Disallowed origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Origin", "http://malicious-site.com")
		
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "", resp.Header.Get("Access-Control-Allow-Origin"))
	})
}
