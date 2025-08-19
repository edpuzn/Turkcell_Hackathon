.PHONY: dev build clean install-backend install-frontend test-backend test-frontend

# Varsayılan hedef
all: dev

# Development ortamını başlat
dev:
	@echo "🚀 Development ortamı başlatılıyor..."
	@echo "📦 Backend bağımlılıkları yükleniyor..."
	@cd Backend && go mod download
	@echo "📦 Frontend bağımlılıkları yükleniyor..."
	@cd frontend && npm install
	@echo "🔧 Backend ve Frontend paralel olarak başlatılıyor..."
	@cd Backend && go run cmd/server/main.go & \
	cd frontend && npm run dev & \
	wait

# Backend'i başlat
dev-backend:
	@echo "🔧 Backend başlatılıyor..."
	@cd Backend && go run cmd/server/main.go

# Frontend'i başlat
dev-frontend:
	@echo "🎨 Frontend başlatılıyor..."
	@cd frontend && npm run dev

# Production build
build:
	@echo "🏗️ Production build başlatılıyor..."
	@echo "🔧 Backend build ediliyor..."
	@cd Backend && go build -o bin/server cmd/server/main.go
	@echo "🎨 Frontend build ediliyor..."
	@cd frontend && npm run build

# Backend bağımlılıklarını yükle
install-backend:
	@echo "📦 Backend bağımlılıkları yükleniyor..."
	@cd Backend && go mod download

# Frontend bağımlılıklarını yükle
install-frontend:
	@echo "📦 Frontend bağımlılıkları yükleniyor..."
	@cd frontend && npm install

# Backend testlerini çalıştır
test-backend:
	@echo "🧪 Backend testleri çalıştırılıyor..."
	@cd Backend && go test ./...

# Frontend testlerini çalıştır
test-frontend:
	@echo "🧪 Frontend testleri çalıştırılıyor..."
	@cd frontend && npm run test

# Temizlik
clean:
	@echo "🧹 Temizlik yapılıyor..."
	@rm -rf Backend/bin
	@rm -rf frontend/dist
	@rm -rf frontend/node_modules
	@echo "✅ Temizlik tamamlandı"

# Docker ile çalıştır
docker-dev:
	@echo "🐳 Docker development ortamı başlatılıyor..."
	@docker-compose -f docker-compose.dev.yml up --build

# Docker production
docker-prod:
	@echo "🐳 Docker production ortamı başlatılıyor..."
	@docker-compose up --build

# Docker durdur
docker-stop:
	@echo "🛑 Docker container'ları durduruluyor..."
	@docker-compose down

# Health check
health:
	@echo "🏥 Health check yapılıyor..."
	@curl -f http://localhost:8000/health || echo "❌ Backend sağlıksız"
	@curl -f http://localhost:5173 || echo "❌ Frontend sağlıksız"
	@echo "✅ Health check tamamlandı"

# Integration test
test-integration:
	@echo "🧪 Entegrasyon testi başlatılıyor..."
	@./test-integration.sh

migrate:
	@echo "🔄 DB migrasyonları uygulanıyor..."
	@cd Backend && go run cmd/migrate/main.go

seed:
	@echo "🌱 DB seed ediliyor (CSV'den)..."
	@cd Backend && DB_SSLMODE=disable go run cmd/seed/main.go

test-api:
	@echo "🧪 API testleri çalıştırılıyor..."
	@echo "Backend health check..."
	@curl -f http://localhost:8000/health || (echo "❌ Backend çalışmıyor" && exit 1)
	@echo "✅ Backend çalışıyor"
	@echo "Katalog özeti testi..."
	@curl -f http://localhost:8000/api/catalog/summary || (echo "❌ Katalog API çalışmıyor" && exit 1)
	@echo "✅ Katalog API çalışıyor"
	@echo "Kapsama testi..."
	@curl -f http://localhost:8000/api/coverage/A1001 || (echo "❌ Kapsama API çalışmıyor" && exit 1)
	@echo "✅ Kapsama API çalışıyor"
	@echo "🎉 Tüm API testleri başarılı!"

# Yardım
help:
	@echo "📖 Kullanılabilir komutlar:"
	@echo "  dev          - Development ortamını başlat (backend + frontend)"
	@echo "  dev-backend  - Sadece backend'i başlat"
	@echo "  dev-frontend - Sadece frontend'i başlat"
	@echo "  build        - Production build"
	@echo "  install-backend  - Backend bağımlılıklarını yükle"
	@echo "  install-frontend - Frontend bağımlılıklarını yükle"
	@echo "  test-backend     - Backend testlerini çalıştır"
	@echo "  test-frontend    - Frontend testlerini çalıştır"
	@echo "  clean        - Temizlik"
	@echo "  docker-dev   - Docker development"
	@echo "  docker-prod  - Docker production"
	@echo "  docker-stop  - Docker durdur"
	@echo "  health       - Health check"
	@echo "  test-integration - Entegrasyon testi"
	@echo "  help         - Bu yardım mesajını göster"
