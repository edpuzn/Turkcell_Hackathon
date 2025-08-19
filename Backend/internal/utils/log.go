package utils

import (
	"github.com/sirupsen/logrus"
)

// ConfigureLogger - Logrus logger'ını yapılandırır
func ConfigureLogger() *logrus.Logger {
	logger := logrus.New()
	
	// JSON formatter kullan
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	
	// Log seviyesini ayarla
	logger.SetLevel(logrus.InfoLevel)
	
	return logger
}

// GetRequestLogger - Request-scoped logger oluşturur
func GetRequestLogger(baseLogger *logrus.Logger, requestID string) *logrus.Entry {
	return baseLogger.WithField("request_id", requestID)
}
