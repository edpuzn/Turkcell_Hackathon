<template>
  <div class="digital-services">
    <div class="page-header">
      <h1>Dijital İŞ Servisleri</h1>
      <p>İşinizi dijitalleştirin, verimliliğinizi artırın</p>
    </div>

    <div class="services-grid">
      <div class="service-card">
        <h3>Kesintisiz Erişim</h3>
        <p>7/24 kesintisiz internet ve ağ hizmetleri</p>
      </div>

      <div class="service-card">
        <h3>Bulut Çözümleri</h3>
        <p>Güvenli ve ölçeklenebilir bulut altyapısı</p>
      </div>

      <div class="service-card">
        <h3>Siber Güvenlik</h3>
        <p>Kapsamlı siber güvenlik koruması</p>
      </div>

      <div class="service-card">
        <h3>IoT Çözümleri</h3>
        <p>Nesnelerin interneti teknolojileri</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.digital-services {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.page-header h1 {
  color: var(--color-text);
  margin-bottom: 1rem;
}

.page-header p {
  color: var(--color-text-light);
  font-size: 1.2rem;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.service-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
}

.service-card:hover {
  border-color: var(--color-primary);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.service-card h3 {
  color: var(--color-primary);
  margin-bottom: 1rem;
}

.service-card p {
  color: var(--color-text-light);
  margin: 0;
}
</style>
