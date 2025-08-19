<template>
  <div class="onboarding-container">
    <!-- Header Section -->
    <div class="onboarding-header">
      <div class="header-content">
        <div class="logo-section">
          <div class="logo-placeholder">TURKCELL</div>
          <h1 class="main-title">Ev+Mobil Paket Danışmanı</h1>
          <p class="subtitle">Size en uygun paketi bulmak için 5 dakika ayırın</p>
        </div>
      </div>
    </div>

    <!-- API Status Section -->
    <div class="api-status-section">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <ApiStatus />
          </div>
        </div>
      </div>
    </div>

    <!-- Steps Overview -->
    <div class="steps-container">
      <div class="steps-header">
        <h2 class="steps-title">Bu Sihirbazda Neler Olacak?</h2>
        <p class="steps-description">Adım adım size en uygun paketi bulalım</p>
      </div>

      <div class="steps-grid">
        <!-- Step 1 -->
        <div class="step-card" :class="{ active: currentStep === 1 }">
          <div class="step-number">1</div>
          <div class="step-content">
            <h3 class="step-title">Hane & Adres Girişi</h3>
            <p class="step-description">
              Şehir, ilçe seçimi ve hane bilgilerinizi girin. Kaç mobil hat, internet ihtiyacı ve TV
              izleme sürenizi belirtin.
            </p>
            <div class="step-features">
              <span class="feature-tag">📍 Konum</span>
              <span class="feature-tag">👥 Hane</span>
              <span class="feature-tag">📱 Hat</span>
            </div>
          </div>
        </div>

        <!-- Step 2 -->
        <div class="step-card" :class="{ active: currentStep === 2 }">
          <div class="step-number">2</div>
          <div class="step-content">
            <h3 class="step-title">Kapsama Kontrolü</h3>
            <p class="step-description">
              Adresinizde hangi internet teknolojileri mevcut? Fiber, VDSL veya FWA uygunluğunu
              kontrol edelim.
            </p>
            <div class="step-features">
              <span class="feature-tag">🌐 Fiber</span>
              <span class="feature-tag">📡 VDSL</span>
              <span class="feature-tag">📶 FWA</span>
            </div>
          </div>
        </div>

        <!-- Step 3 -->
        <div class="step-card" :class="{ active: currentStep === 3 }">
          <div class="step-number">3</div>
          <div class="step-content">
            <h3 class="step-title">Öneri Kartları</h3>
            <p class="step-description">
              Size en uygun 3 paket kombinasyonunu gösterelim. Mobil + Ev internet + TV paketleri,
              indirimler ve tasarruflar.
            </p>
            <div class="step-features">
              <span class="feature-tag">💳 Paket</span>
              <span class="feature-tag">💰 Tasarruf</span>
              <span class="feature-tag">📊 Detay</span>
            </div>
          </div>
        </div>

        <!-- Step 4 -->
        <div class="step-card" :class="{ active: currentStep === 4 }">
          <div class="step-number">4</div>
          <div class="step-content">
            <h3 class="step-title">Özet & Randevu</h3>
            <p class="step-description">
              Seçtiğiniz paketlerin özeti ve kurulum için uygun randevu slotları. Tarih ve saat
              seçimi.
            </p>
            <div class="step-features">
              <span class="feature-tag">📋 Özet</span>
              <span class="feature-tag">📅 Randevu</span>
              <span class="feature-tag">✅ Onay</span>
            </div>
          </div>
        </div>

        <!-- Step 5 -->
        <div class="step-card" :class="{ active: currentStep === 5 }">
          <div class="step-number">5</div>
          <div class="step-content">
            <h3 class="step-title">Sipariş Onayı</h3>
            <p class="step-description">
              Siparişiniz alındı! Randevu bilgileri ve kurulum detayları ile birlikte onay mesajı.
            </p>
            <div class="step-features">
              <span class="feature-tag">🎉 Başarı</span>
              <span class="feature-tag">📞 Kurulum</span>
              <span class="feature-tag">🏠 Ana Sayfa</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Section -->
    <div class="action-section">
      <div class="action-content">
        <div class="info-box">
          <div class="info-icon">
            <i class="bi bi-clock"></i>
          </div>
          <div class="info-text">
            <h4>Tahmini Süre: 5-7 dakika</h4>
            <p>Basit sorularla size en uygun paketi bulalım</p>
          </div>
        </div>

        <div class="cta-buttons">
          <button class="btn-secondary" @click="goToHome">
            <i class="bi bi-house me-2"></i>
            Ana Sayfa
          </button>
          <button class="btn-primary" @click="startWizard">
            <i class="bi bi-play-circle me-2"></i>
            Sihirbazı Başlat
          </button>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <div class="onboarding-footer">
      <p>Turkcell Ev+Mobil Paket Danışmanı ile size en uygun paketi bulun</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const currentStep = ref(1)

// Auto-advance steps for demo
onMounted(() => {
  const interval = setInterval(() => {
    currentStep.value = currentStep.value < 5 ? currentStep.value + 1 : 1
  }, 2000)

  return () => clearInterval(interval)
})

const startWizard = () => {
  console.log('startWizard function called')
  console.log('Router object:', router)
  console.log('Current route:', router.currentRoute.value)

  try {
    router.push('/wizard')
    console.log('Router.push called successfully')
  } catch (error) {
    console.error('Router.push error:', error)
  }
}

const goToHome = () => {
  console.log('goToHome function called')
  try {
    router.push('/wizard')
    console.log('Router.push called successfully')
  } catch (error) {
    console.error('Router.push error:', error)
  }
}
</script>

<style scoped>
.onboarding-container {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--color-background-soft), var(--color-background));
  padding: 2rem 1rem;
}

.onboarding-header {
  text-align: center;
  margin-bottom: 3rem;
}

.header-content {
  max-width: 800px;
  margin: 0 auto;
}

.logo-section {
  margin-bottom: 2rem;
}

.logo-placeholder {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-primary);
  margin-bottom: 1rem;
  letter-spacing: 2px;
}

.main-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 1rem;
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 1.2rem;
  color: var(--color-text-mute);
  font-weight: 500;
}

.steps-container {
  max-width: 1200px;
  margin: 0 auto 3rem;
}

.steps-header {
  text-align: center;
  margin-bottom: 3rem;
}

.steps-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 1rem;
}

.steps-description {
  font-size: 1.1rem;
  color: var(--color-text-mute);
}

.steps-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
}

.step-card {
  background: var(--color-background-card);
  border: 2px solid var(--color-border);
  border-radius: 20px;
  padding: 2rem;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.step-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--color-primary), var(--color-primary-light));
  transform: scaleX(0);
  transition: transform 0.3s ease;
}

.step-card.active::before {
  transform: scaleX(1);
}

.step-card:hover {
  transform: translateY(-5px);
  box-shadow: var(--shadow-lg);
  border-color: var(--color-primary-light);
}

.step-card.active {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-xl);
}

.step-number {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  color: var(--color-text-inverse);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.step-title {
  font-size: 1.3rem;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 1rem;
}

.step-description {
  color: var(--color-text-mute);
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.step-features {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.feature-tag {
  background: var(--color-primary-lighter);
  color: var(--color-primary-dark);
  padding: 0.4rem 0.8rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.action-section {
  max-width: 800px;
  margin: 0 auto 3rem;
}

.action-content {
  background: var(--color-background-card);
  border: 2px solid var(--color-border);
  border-radius: 20px;
  padding: 2rem;
}

.info-box {
  display: flex;
  align-items: center;
  gap: 1rem;
  background: var(--color-primary-lighter);
  padding: 1.5rem;
  border-radius: 16px;
  margin-bottom: 2rem;
}

.info-icon {
  width: 50px;
  height: 50px;
  background: var(--color-primary);
  color: var(--color-text-inverse);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
}

.info-text h4 {
  color: var(--color-primary-dark);
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.info-text p {
  color: var(--color-text-mute);
  margin: 0;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn-primary,
.btn-secondary {
  padding: 1rem 2rem;
  font-size: 1.1rem;
  font-weight: 600;
  border-radius: 16px;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-primary {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  color: var(--color-text-inverse);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-lg);
}

.btn-secondary {
  background: var(--color-background-mute);
  color: var(--color-text);
  border: 2px solid var(--color-border);
}

.btn-secondary:hover {
  background: var(--color-border);
  transform: translateY(-2px);
}

.onboarding-footer {
  text-align: center;
  color: var(--color-text-mute);
  font-size: 0.9rem;
  padding: 2rem 0;
}

/* Responsive */
@media (max-width: 768px) {
  .onboarding-container {
    padding: 1rem 0.5rem;
  }

  .main-title {
    font-size: 2rem;
  }

  .steps-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .step-card {
    padding: 1.5rem;
  }

  .cta-buttons {
    flex-direction: column;
  }

  .btn-primary,
  .btn-secondary {
    width: 100%;
    justify-content: center;
  }
}
</style>
