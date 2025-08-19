package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"turkcell-hackathon-backend/internal/csvload"
	"turkcell-hackathon-backend/internal/http"
	"turkcell-hackathon-backend/internal/service"
	"turkcell-hackathon-backend/internal/utils"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// .env dosyasını yükle (opsiyonel)
	if err := godotenv.Load(); err != nil {
		log.Println(".env dosyası bulunamadı, environment değişkenleri kullanılıyor")
	}

	// Logger'ı başlat
	logger := utils.ConfigureLogger()
	logger.Info("Turkcell Hackathon Backend başlatılıyor")

	// Environment değişkenlerini yükle
	port := utils.GetEnv("PORT", "8000")

	logger.WithFields(logrus.Fields{
		"port": port,
	}).Info("Konfigürasyon yüklendi")

	// CSV modu
	csvDir := utils.GetEnv("CSV_DIR", "./data")
	logger.WithField("csvDir", csvDir).Info("CSV modu aktif")

	// CSV verilerini yükle
	ctx := context.Background()
	catalog, err := csvload.GlobalLoader.LoadAll(ctx, csvDir)
	if err != nil {
		logger.WithError(err).Fatal("CSV verileri yüklenemedi")
	}

	// Service katmanını başlat
	catalogService := service.NewCatalogService(catalog)

	logger.WithFields(logrus.Fields{
		"coverages":        len(catalog.Coverages),
		"mobile_plans":     len(catalog.MobilePlans),
		"home_plans":       len(catalog.HomePlans),
		"tv_plans":         len(catalog.TVPlans),
		"bundling_rules":   len(catalog.BundlingRules),
		"install_slots":    len(catalog.InstallSlots),
		"users":            len(catalog.Users),
		"household_lines":  len(catalog.HouseholdLines),
		"current_services": len(catalog.CurrentServices),
	}).Info("CSV verileri başarıyla yüklendi")

	// HTTP router
	router := http.NewRouter(catalogService, logger)
	app := router.App()

	// Server'ı goroutine'de başlat
	go func() {
		addr := fmt.Sprintf(":%s", port)
		logger.WithField("address", addr).Info("HTTP server başlatılıyor")
		if err := app.Listen(addr); err != nil {
			logger.WithError(err).Fatal("Server başlatılamadı")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Server kapatılıyor...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.WithError(err).Error("Server zorla kapatıldı")
	} else {
		logger.Info("Server düzgün şekilde kapatıldı")
	}
}
