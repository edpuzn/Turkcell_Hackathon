<template>
  <div class="partnership">
    <div class="page-header">
      <h1>İŞ Ortaklığı</h1>
      <p>Birlikte büyüyelim, başarıya ulaşalım</p>
    </div>

    <div class="partnership-grid">
      <div class="partnership-card">
        <h3>Ortaklık Programları</h3>
        <p>Çeşitli ortaklık programları ve avantajlar</p>
      </div>

      <div class="partnership-card">
        <h3>Teknik Destek</h3>
        <p>Ortaklarımıza özel teknik destek</p>
      </div>

      <div class="partnership-card">
        <h3>Eğitim Programları</h3>
        <p>Ortaklarımız için eğitim ve sertifikasyon</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.partnership {
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

.partnership-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.partnership-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
}

.partnership-card:hover {
  border-color: var(--color-primary);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.partnership-card h3 {
  color: var(--color-primary);
  margin-bottom: 1rem;
}

.partnership-card p {
  color: var(--color-text-light);
  margin: 0;
}
</style>
