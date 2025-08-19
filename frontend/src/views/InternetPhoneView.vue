<template>
  <div class="internet-phone">
    <div class="page-header">
      <h1>İnternet & Telefon</h1>
      <p>Hızlı internet ve kaliteli telefon hizmetleri</p>
    </div>

    <div class="services-grid">
      <div class="service-card">
        <h3>Fiber İnternet</h3>
        <p>1000 Mbps'e kadar hızlı fiber internet</p>
      </div>

      <div class="service-card">
        <h3>VDSL İnternet</h3>
        <p>100 Mbps'e kadar VDSL internet</p>
      </div>

      <div class="service-card">
        <h3>Sabit Telefon</h3>
        <p>Kaliteli sabit telefon hizmeti</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.internet-phone {
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
