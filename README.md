# Turkcell Hackathon - Paket Öneri Sistemi

Bu proje, Turkcell Hackathon için geliştirilmiş bir paket öneri sistemidir. Go (Fiber) backend ve Vue 3 frontend'den oluşur.

## 🚀 Özellikler

- **Backend**: Go (Fiber) ile RESTful API
- **Frontend**: Vue 3 + TypeScript + Pinia + Bootstrap
- **Veri Kaynağı**: CSV dosyaları
- **Öneri Motoru**: Akıllı paket önerileri
- **Kapsama Kontrolü**: Adres bazlı teknoloji kontrolü

## 📁 Proje Yapısı

```
TurkcellHackathon/
├── Backend/                 # Go backend
│   ├── cmd/
│   │   └── server/         # Ana uygulama
│   ├── internal/
│   │   ├── csvload/        # CSV veri yükleme
│   │   ├── http/           # HTTP handlers
│   │   ├── models/         # Veri modelleri
│   │   ├── service/        # İş mantığı
│   │   └── utils/          # Yardımcı fonksiyonlar
│   ├── data/               # CSV veri dosyaları
│   └── go.mod
├── frontend/               # Vue 3 frontend
│   ├── src/
│   │   ├── components/     # Vue bileşenleri
│   │   ├── stores/         # Pinia store'ları
│   │   ├── views/          # Sayfa görünümleri
│   │   └── utils/          # Yardımcı fonksiyonlar
│   └── package.json
├── docker-compose.yml      # Production Docker
├── docker-compose.dev.yml  # Development Docker
├── Makefile               # Geliştirme komutları
└── README.md
```

## 🛠️ Kurulum

### Gereksinimler

- Go 1.23+
- Node.js 20+
- Yarn veya npm

### Hızlı Başlangıç

1. **Repository'yi klonlayın:**
```bash
git clone https://github.com/edpuzn/Turkcell_Hackathon.git
cd Turkcell_Hackathon
```

2. **Backend'i başlatın:**
```bash
cd Backend
go mod tidy
go run cmd/server/main.go
```

3. **Frontend'i başlatın:**
```bash
cd frontend
yarn install
yarn dev
```

4. **Tarayıcıda açın:**
- Frontend: http://localhost:5173
- Backend API: http://localhost:8000

## 📋 API Endpoint'leri

### Health Check
- `GET /health` - Sistem durumu

### Katalog
- `GET /api/catalog/summary` - Katalog özeti

### Kapsama
- `GET /api/coverage/:address_id` - Adres bazlı kapsama bilgisi

### Öneriler
- `POST /api/recommendation` - Paket önerileri

## 🐳 Docker ile Çalıştırma

### Development
```bash
docker compose -f docker-compose.dev.yml up --build
```

### Production
```bash
docker compose up --build
```

## 🔧 Geliştirme Komutları

```bash
# Her iki servisi de başlat
make dev

# Sadece backend
make dev-backend

# Sadece frontend
make dev-frontend

# Test
make test

# Build
make build

# Temizlik
make clean
```

## 📊 Veri Yapısı

### CSV Dosyaları
- `coverage.csv` - Kapsama bilgileri
- `mobile_plans.csv` - Mobil planlar
- `home_plans.csv` - Ev internet planları
- `tv_plans.csv` - TV planları
- `bundling_rules.csv` - Paketleme kuralları
- `install_slots.csv` - Kurulum slotları
- `users.csv` - Kullanıcı bilgileri

## 🎯 Özellikler

### Frontend
- **Onboarding**: Kullanıcı bilgileri girişi
- **Wizard**: Adım adım paket seçimi
- **Kapsama Kontrolü**: Adres bazlı teknoloji kontrolü
- **Öneri Sistemi**: Akıllı paket önerileri
- **Responsive Design**: Mobil uyumlu tasarım

### Backend
- **RESTful API**: Modern API tasarımı
- **CORS Desteği**: Cross-origin istekler
- **Logging**: Detaylı log sistemi
- **Error Handling**: Hata yönetimi
- **Validation**: Veri doğrulama

## 🤝 Katkıda Bulunma

1. Fork edin
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit edin (`git commit -m 'Add amazing feature'`)
4. Push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## 📄 Lisans

Bu proje MIT lisansı altında lisanslanmıştır.

## 👥 Takım

- **Backend**: Go (Fiber)
- **Frontend**: Vue 3 + TypeScript
- **Veritabanı**: CSV (PostgreSQL hazır)
- **Deployment**: Docker + Docker Compose

---

**Turkcell Hackathon 2024** - Paket Öneri Sistemi
