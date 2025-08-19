<template>
  <div class="devices">
    <div class="page-header">
      <h1>Cihazlar</h1>
      <p>En son teknoloji cihazlar ve çözümler</p>
    </div>

    <div class="devices-grid">
      <div class="device-card">
        <h3>Akıllı Telefonlar</h3>
        <p>En son model akıllı telefonlar</p>
      </div>

      <div class="device-card">
        <h3>Tablet & Bilgisayar</h3>
        <p>Tablet ve bilgisayar çözümleri</p>
      </div>

      <div class="device-card">
        <h3>Modem & Router</h3>
        <p>Yüksek performanslı ağ cihazları</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.devices {
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

.devices-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.device-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
}

.device-card:hover {
  border-color: var(--color-primary);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.device-card h3 {
  color: var(--color-primary);
  margin-bottom: 1rem;
}

.device-card p {
  color: var(--color-text-light);
  margin: 0;
}
</style>
