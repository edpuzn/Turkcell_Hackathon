package http

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// RequestID - Her istek için benzersiz ID oluşturur
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			// 16 byte rastgele ID oluştur
			bytes := make([]byte, 16)
			rand.Read(bytes)
			requestID = hex.EncodeToString(bytes)
		}
		
		c.Set("X-Request-ID", requestID)
		c.Locals("request_id", requestID)
		
		return c.Next()
	}
}

// Logger - İstek loglama middleware'i
func Logger(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		requestID := c.Locals("request_id").(string)
		
		// İstek başlangıcını logla
		reqLogger := logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Method(),
			"path":       c.Path(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		})
		
		reqLogger.Info("İstek başladı")
		
		// İsteği işle
		err := c.Next()
		
		// Yanıtı logla
		duration := time.Since(start)
		status := c.Response().StatusCode()
		
		reqLogger.WithFields(logrus.Fields{
			"status_code": status,
			"duration_ms": duration.Milliseconds(),
			"bytes":       len(c.Response().Body()),
		}).Info("İstek tamamlandı")
		
		return err
	}
}

// Recover - Panic kurtarma middleware'i
func Recover(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				requestID := c.Locals("request_id").(string)
				
				logger.WithFields(logrus.Fields{
					"request_id": requestID,
					"method":     c.Method(),
					"path":       c.Path(),
					"panic":      r,
				}).Error("Panic kurtarıldı")
				
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Sunucu iç hatası",
				})
			}
		}()
		
		return c.Next()
	}
}

// CORS - Manuel CORS middleware'i
func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// CORS header'larını her zaman set et
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-ID")
		c.Set("Access-Control-Expose-Headers", "Content-Length")
		c.Set("Access-Control-Max-Age", "86400")
		
		// OPTIONS istekleri için erken dön
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		
		return c.Next()
	}
}

// RegisterMiddlewares - Tüm middleware'leri kaydeder
func RegisterMiddlewares(app *fiber.App, logger *logrus.Logger) {
	// 1. CORS - EN BAŞTA OLMALI (tüm route'lardan önce)
	app.Use(CORS())
	
	// 2. Diğer middleware'ler (CORS'tan sonra)
	app.Use(RequestID())
	app.Use(Logger(logger))
	app.Use(Recover(logger))
	
	// CORS yapılandırmasını logla
	logger.Info("CORS middleware yapılandırıldı - AllowOrigins: *")
}
