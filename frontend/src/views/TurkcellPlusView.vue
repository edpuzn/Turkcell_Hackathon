<template>
  <div class="turkcell-plus">
    <div class="page-header">
      <h1>İŞTurkcell Plus</h1>
      <p>Şirketinize özel avantajlar ve çözümler</p>
    </div>

    <div class="plus-grid">
      <div class="plus-card">
        <h3>Şirketinize Özel</h3>
        <p>Kurumsal ihtiyaçlarınıza özel paketler</p>
      </div>

      <div class="plus-card">
        <h3>Çalışanlara Özel</h3>
        <p>Çalışanlarınız için özel indirimler</p>
      </div>

      <div class="plus-card">
        <h3>Öncelikli Destek</h3>
        <p>7/24 öncelikli müşteri desteği</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.turkcell-plus {
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

.plus-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.plus-card {
  background: var(--color-background-soft);
  border: 1px solid var(--color-border);
  border-radius: 12px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
}

.plus-card:hover {
  border-color: var(--color-primary);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.plus-card h3 {
  color: var(--color-primary);
  margin-bottom: 1rem;
}

.plus-card p {
  color: var(--color-text-light);
  margin: 0;
}
</style>
