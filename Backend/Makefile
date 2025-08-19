.PHONY: tidy run test build lint

# Go modüllerini düzenle
tidy:
	go mod tidy

# Server'ı çalıştır
run:
	go run cmd/server/main.go

# Testleri çalıştır
test:
	go test ./...

# Binary dosyasını oluştur
build:
	go build -o bin/server cmd/server/main.go

# Lint (placeholder)
lint:
	echo "Lint işlemi henüz yapılandırılmadı"
