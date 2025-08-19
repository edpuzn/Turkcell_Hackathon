# Turkcell Hackathon Backend

Production-ready Go backend skeleton for Turkcell Hackathon case using Fiber and CSV loader.

## 🚀 Hızlı Başlangıç

### Gereksinimler
- Go 1.21+
- Git

### Kurulum

1. **Environment dosyasını kopyala:**
```bash
cp env.example .env
```

2. **Bağımlılıkları yükle:**
```bash
make tidy
```

3. **Server'ı çalıştır:**
```bash
make run
```

Server varsayılan olarak `http://localhost:8000` adresinde çalışacaktır.

## 📁 Proje Yapısı

```
backend/
├── cmd/server/main.go          # Ana server dosyası
├── internal/
│   ├── http/                   # HTTP katmanı
│   │   ├── router.go           # Route yapılandırması
│   │   ├── handlers.go         # HTTP handler'ları
│   │   └── middleware.go       # Middleware'ler
│   ├── csvload/                # CSV yükleme sistemi
│   │   ├── loader.go           # CSV loader
│   │   └── files.go            # Dosya tanımları
│   ├── models/                 # Veri modelleri
│   │   ├── catalog.go          # Katalog modelleri
│   │   └── user.go             # Kullanıcı modelleri
│   ├── service/                # Service katmanı
│   │   └── catalog.go          # Katalog servisi
│   └── utils/                  # Yardımcı fonksiyonlar
│       ├── env.go              # Environment yardımcıları
│       └── log.go              # Logging yardımcıları
├── data/                       # CSV veri dosyaları
│   ├── coverage.csv
│   ├── mobile_plans.csv
│   ├── home_plans.csv
│   ├── tv_plans.csv
│   ├── bundling_rules.csv
│   ├── install_slots.csv
│   ├── users.csv
│   ├── household.csv
│   └── current_services.csv
├── Makefile                    # Build komutları
├── go.mod                      # Go modül tanımı
└── README.md                   # Bu dosya
```

## 🔧 Komutlar

```bash
make tidy   # Go modüllerini düzenle
make run    # Server'ı çalıştır
make test   # Testleri çalıştır
make build  # Binary dosyasını oluştur
make lint   # Lint (placeholder)
```

## 🔧 Konfigürasyon

### Environment Değişkenleri

```bash
# Server Configuration
PORT=8000                    # Server portu (varsayılan: 8000)
CSV_DIR=./data              # CSV dosyalarının bulunduğu dizin

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000  # İzin verilen origin'ler
CORS_CREDENTIALS=false      # Credentials desteği (varsayılan: false)
```

### CORS Test

CORS yapılandırmasını test etmek için:

```bash
# Preflight OPTIONS isteği
curl -s -H "Origin: http://localhost:5173" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -X OPTIONS http://localhost:8000/api/recommendation -v

# POST isteği ile CORS header'larını kontrol et
curl -s -X POST http://localhost:8000/api/recommendation \
  -H 'Content-Type: application/json' \
  -H 'Origin: http://localhost:5173' \
  -d '{"user_id":1,"address_id":"ADR-1001","household":[{"line_id":"L1","expected_gb":20,"expected_min":400,"tv_hd_hours":6}]}' -v
```

## 🌐 API Endpoint'leri

### Health Check
```bash
curl http://localhost:8000/health
```
**Yanıt:** `{"status":"ok"}`

### Katalog Özeti
```bash
curl http://localhost:8000/api/catalog/summary
```
**Yanıt:** Tüm veri setlerinin sayıları

### Adres Kapsaması
```bash
curl http://localhost:8000/api/coverage/ADR-1001
```
**Yanıt:** Belirtilen adres için kapsama bilgileri

### Öneri Motoru
```bash
curl -s -X POST http://localhost:8000/api/recommendation \
  -H 'Content-Type: application/json' \
  -d '{
    "user_id": 1,
    "address_id": "ADR-1001",
    "household": [
      {
        "line_id": "L1",
        "expected_gb": 20,
        "expected_min": 400,
        "tv_hd_hours": 6
      },
      {
        "line_id": "L2",
        "expected_gb": 10,
        "expected_min": 200,
        "tv_hd_hours": 3
      }
    ],
    "prefer_tech": ["fiber", "vdsl"]
  }' | jq .
```
**Yanıt:** En uygun 3 hizmet kombinasyonu

## 📊 Veri Modelleri

### Katalog Modelleri
- **Coverage**: Adres bazında teknoloji kapsaması
- **MobilePlan**: Mobil plan detayları
- **HomePlan**: Ev internet planları
- **TVPlan**: TV planları
- **BundlingRule**: Paketleme indirim kuralları
- **InstallSlot**: Kurulum zaman dilimleri

### Kullanıcı Modelleri
- **User**: Kullanıcı bilgileri
- **HouseholdLine**: Hane halkı hatları
- **CurrentServices**: Mevcut hizmetler

## 🔧 Konfigürasyon

Environment değişkenleri:
- `PORT`: Server port (varsayılan: 8000)
- `CSV_DIR`: CSV dosyalarının dizini (varsayılan: ./data)

## 🧪 Test

```bash
make test
```

## 📝 Özellikler

- ✅ Go 1.21+ ve Fiber v2
- ✅ In-memory CSV loader
- ✅ Clean architecture
- ✅ Middleware (RequestID, Logger, Recover, CORS)
- ✅ Graceful shutdown
- ✅ JSON logging
- ✅ Environment konfigürasyonu
- ✅ Makefile ile kolay kullanım
- ✅ Recommendation engine (pricing rules, beam search)
- ✅ Comprehensive test coverage
