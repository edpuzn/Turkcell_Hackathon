# Turkcell Telekom Paket Sihirbazı

Bu proje, Turkcell kurumsal müşterileri için geliştirilmiş modern bir telekom paket seçim sistemidir. Vue.js 3, Pinia ve modern web teknolojileri kullanılarak geliştirilmiştir.

## 🚀 Özellikler

### 1. Girdi Sihirbazı (Tek Ekran)
- **Adres Seçimi**: Şehir → İlçe → Adres ID (hierarchical dropdown)
- **Hane Hatları**: Dinamik satır ekleme/çıkarma
- **GB ve Dakika İhtiyacı**: Her hat için ayrı ayrı belirleme
- **TV HD Saat**: Opsiyonel TV paketi seçimi

### 2. Kapsama Rozetleri
- **Teknoloji Gösterimi**: Fiber/VDSL/FWA kapsama durumu
- **GET /catalog** entegrasyonu
- **Görsel rozetler** ile kolay anlaşılır arayüz

### 3. Top-3 Öneri Kartları
- **POST /recommendation** entegrasyonu
- **Combo paketler**: Mobil+Ev / Mobil+Ev+TV
- **Fiyat ve tasarruf bilgileri**
- **Rozetler**: Kapsama, Bundle indirimi, Ek hat indirimi
- **Alternatif öneri** etiketleri

### 4. Detay Modal ve Seçim
- **Detaylı paket bilgileri**
- **Fiyat dökümü** ve indirim açıklamaları
- **Randevu slot seçimi**
- **POST /checkout** entegrasyonu

### 5. Tema ve Erişilebilirlik
- **Karanlık/Aydınlık tema** toggle
- **Responsive tasarım**
- **ARIA labels** ve klavye navigasyonu
- **Yüksek kontrast** desteği

## 🏗️ Mimari

### SOLID Prensipleri
- **Single Responsibility**: Her component tek bir sorumluluğa sahip
- **Open/Closed**: Yeni özellikler için genişletilebilir
- **Liskov Substitution**: Interface'ler tutarlı
- **Interface Segregation**: Küçük, odaklanmış interface'ler
- **Dependency Inversion**: Store pattern ile loose coupling

### Modüler Yapı
```
src/
├── components/          # Yeniden kullanılabilir bileşenler
│   ├── ThemeToggle.vue
│   ├── RecommendationModal.vue
│   └── SelectionModal.vue
├── views/              # Sayfa bileşenleri
│   ├── HomeView.vue    # Ana bundle wizard
│   └── [Other Views]   # Diğer sayfalar
├── stores/             # Pinia state management
│   └── bundle.js       # Bundle veri store'u
├── router/             # Vue Router konfigürasyonu
└── assets/             # CSS ve görsel dosyalar
```

## 🛠️ Teknolojiler

- **Vue.js 3** - Composition API
- **Pinia** - State management
- **Vue Router 4** - Client-side routing
- **CSS Variables** - Tema sistemi
- **Grid/Flexbox** - Modern layout
- **ES6+** - Modern JavaScript

## 📱 Responsive Tasarım

- **Mobile-first** yaklaşım
- **Breakpoints**: 480px, 768px, 1024px
- **Touch-friendly** butonlar (min 44px)
- **Flexible grids** ve responsive cards

## ♿ Erişilebilirlik

- **ARIA labels** ve roles
- **Klavye navigasyonu** desteği
- **Focus management** ve visible focus
- **Screen reader** uyumluluğu
- **Yüksek kontrast** mod desteği
- **Reduced motion** desteği

## 🚀 Kurulum

```bash
# Bağımlılıkları yükle
npm install

# Geliştirme sunucusunu başlat
npm run dev

# Production build
npm run build

# Lint ve format
npm run lint
npm run format
```

## 🔧 API Entegrasyonu

### Endpoints
- `GET /api/catalog?address_id={id}` - Kapsama bilgisi
- `POST /api/recommendation` - Paket önerileri
- `GET /api/install-slots` - Randevu saatleri
- `POST /api/checkout` - Sipariş onayı

### Mock Data
Şu anda mock data kullanılmaktadır. Gerçek API entegrasyonu için store dosyalarındaki mock fonksiyonları gerçek API çağrıları ile değiştirin.

## 🎨 Tema Sistemi

### CSS Variables
```css
:root {
  --color-primary: #0066cc;
  --color-secondary: #00a651;
  --color-success: #00a651;
  --color-warning: #ff9900;
  --color-danger: #dc3545;
  /* ... */
}
```

### Dark Theme
```css
body.dark {
  --color-background: #0d1117;
  --color-text: #f0f6fc;
  /* ... */
}
```

## 📊 State Management

### Pinia Store (bundle.js)
- **Cities, Districts, Addresses** - Lokasyon verileri
- **Coverage Data** - Teknoloji kapsama bilgileri
- **Recommendations** - Paket önerileri
- **Install Slots** - Randevu saatleri

### Actions
- `initializeCoverage()` - Başlangıç verilerini yükle
- `fetchCoverage(addressId)` - Kapsama bilgisi al
- `fetchRecommendations(addressId, household)` - Önerileri al
- `fetchInstallSlots(addressId, techType)` - Randevu saatleri al

## 🔄 Workflow

1. **Adres Seçimi** → Şehir/İlçe/Adres dropdown'ları
2. **Hane Hatları** → GB, dakika, TV ihtiyaçları
3. **Kapsama Kontrolü** → Fiber/VDSL/FWA rozetleri
4. **Öneri Hesaplama** → Top-3 paket önerisi
5. **Detay Görüntüleme** → Modal ile detaylı bilgi
6. **Randevu Seçimi** → Kurulum tarihi ve saati
7. **Sipariş Onayı** → Checkout ve order_id

## 🧪 Test

```bash
# Unit testler (gelecekte eklenecek)
npm run test:unit

# E2E testler (gelecekte eklenecek)
npm run test:e2e
```

## 📈 Performance

- **Lazy loading** - Route-based code splitting
- **Computed properties** - Reactive data optimization
- **Debounced inputs** - API çağrı optimizasyonu
- **Virtual scrolling** - Büyük listeler için (gelecekte)

## 🔒 Güvenlik

- **Input validation** - Form veri doğrulama
- **XSS protection** - Vue.js built-in protection
- **CSRF tokens** - API güvenliği (gelecekte)
- **Rate limiting** - API abuse koruması (gelecekte)

## 🌐 Browser Support

- **Chrome** 90+
- **Firefox** 88+
- **Safari** 14+
- **Edge** 90+

## 📝 Lisans

Bu proje Turkcell için geliştirilmiştir. Tüm hakları saklıdır.

## 🤝 Katkıda Bulunma

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit yapın (`git commit -m 'Add amazing feature'`)
4. Push yapın (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## 📞 Destek

Teknik destek için:
- **Email**: tech-support@turkcell.com.tr
- **Telefon**: 0850 222 0 444
- **Dokümantasyon**: [docs.turkcell.com.tr](https://docs.turkcell.com.tr)
