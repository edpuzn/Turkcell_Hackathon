<script setup>
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import ThemeToggle from '@/components/ThemeToggle.vue'

const isDark = ref(false)
const showLoginModal = ref(false)

const toggleTheme = () => {
  isDark.value = !isDark.value
  document.body.classList.toggle('dark', isDark.value)
}
</script>

<template>
  <div id="app" :class="{ dark: isDark }">
    <header class="header">
      <div class="header-container">
        <div class="logo-section">
          <div class="logo-placeholder">TURKCELL</div>
          <span class="logo-text">Kurumsal</span>
        </div>

        <nav class="main-nav">
          <div class="nav-dropdown">
            <button class="nav-dropdown-btn">
              Hizmetler
              <i class="bi bi-chevron-down ms-1"></i>
            </button>
            <div class="nav-dropdown-content">
              <RouterLink to="/paketler" class="dropdown-link">Paketler</RouterLink>
              <RouterLink to="/mobil" class="dropdown-link">Mobil</RouterLink>
              <RouterLink to="/ev-internet" class="dropdown-link">Ev İnternet</RouterLink>
              <RouterLink to="/tv" class="dropdown-link">TV+</RouterLink>
            </div>
          </div>
          <RouterLink to="/kurumsal" class="nav-link">Kurumsal</RouterLink>
          <RouterLink to="/destek" class="nav-link">Destek</RouterLink>
        </nav>

        <div class="header-actions">
          <div class="search-section">
            <div class="search-box">
              <i class="bi bi-search search-icon"></i>
              <input type="text" placeholder="Ara..." class="search-input" />
            </div>
          </div>

          <div class="login-section">
            <RouterLink to="/login" class="login-btn">
              <i class="bi bi-person-circle me-2"></i>
              Giriş Yap
            </RouterLink>
          </div>

          <ThemeToggle :isDark="isDark" @toggle="toggleTheme" />
        </div>
      </div>
    </header>

    <main class="main-content">
      <RouterView />
    </main>

    <footer class="footer">
      <div class="footer-container">
        <div class="footer-section">
          <h5 class="footer-title">Kurumsal</h5>
          <ul class="footer-links">
            <li><a href="#" class="footer-link">Hakkımızda</a></li>
            <li><a href="#" class="footer-link">Kariyer</a></li>
            <li><a href="#" class="footer-link">Basın</a></li>
            <li><a href="#" class="footer-link">İletişim</a></li>
          </ul>
        </div>

        <div class="footer-section">
          <h5 class="footer-title">Hizmetler</h5>
          <ul class="footer-links">
            <li><a href="#" class="footer-link">Mobil</a></li>
            <li><a href="#" class="footer-link">Ev İnternet</a></li>
            <li><a href="#" class="footer-link">TV+</a></li>
            <li><a href="#" class="footer-link">Kurumsal</a></li>
          </ul>
        </div>

        <div class="footer-section">
          <h5 class="footer-title">Destek</h5>
          <ul class="footer-links">
            <li><a href="#" class="footer-link">Müşteri Hizmetleri</a></li>
            <li><a href="#" class="footer-link">Teknik Destek</a></li>
            <li><a href="#" class="footer-link">Sık Sorulan Sorular</a></li>
            <li><a href="#" class="footer-link">Şube Bul</a></li>
          </ul>
        </div>

        <div class="footer-section">
          <h5 class="footer-title">İletişim</h5>
          <ul class="footer-links">
            <li><a href="#" class="footer-link">444 0 532</a></li>
            <li><a href="#" class="footer-link">info@turkcell.com.tr</a></li>
            <li><a href="#" class="footer-link">Sosyal Medya</a></li>
          </ul>
        </div>
      </div>

      <div class="footer-bottom">
        <div class="footer-bottom-container">
          <p class="footer-copyright">© 2024 Turkcell. Tüm hakları saklıdır.</p>
          <div class="footer-legal">
            <a href="#" class="footer-legal-link">Gizlilik Politikası</a>
            <a href="#" class="footer-legal-link">Kullanım Şartları</a>
            <a href="#" class="footer-legal-link">KVKK</a>
          </div>
        </div>
      </div>
    </footer>

    <!-- Login Modal -->
    <div v-if="showLoginModal" class="modal-overlay" @click="showLoginModal = false">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h4 class="modal-title">
            <i class="bi bi-person-circle me-2"></i>
            Giriş Yap
          </h4>
          <button class="modal-close" @click="showLoginModal = false">
            <i class="bi bi-x-lg"></i>
          </button>
        </div>

        <div class="modal-body">
          <form class="login-form">
            <div class="form-group">
              <label for="email" class="form-label">E-posta / Telefon</label>
              <input
                type="text"
                id="email"
                class="form-control"
                placeholder="ornek@email.com veya 05XX XXX XX XX"
                required
              />
            </div>

            <div class="form-group">
              <label for="password" class="form-label">Şifre</label>
              <input
                type="password"
                id="password"
                class="form-control"
                placeholder="Şifrenizi giriniz"
                required
              />
            </div>

            <div class="form-options">
              <label class="checkbox-label">
                <input type="checkbox" class="form-check-input" />
                <span class="checkmark"></span>
                Beni hatırla
              </label>
              <a href="#" class="forgot-password">Şifremi unuttum</a>
            </div>

            <button type="submit" class="btn btn-primary btn-login">
              <i class="bi bi-box-arrow-in-right me-2"></i>
              Giriş Yap
            </button>
          </form>

          <div class="login-divider">
            <span>veya</span>
          </div>

          <div class="social-login">
            <button class="btn btn-outline-primary btn-social">
              <i class="bi bi-google me-2"></i>
              Google ile Giriş
            </button>
            <button class="btn btn-outline-primary btn-social">
              <i class="bi bi-apple me-2"></i>
              Apple ile Giriş
            </button>
          </div>

          <div class="register-section">
            <p class="register-text">
              Hesabınız yok mu?
              <a href="#" class="register-link">Hemen üye olun</a>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.header {
  background: var(--color-background);
  border-bottom: 1px solid var(--color-border);
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.header-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 80px;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.logo-placeholder {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  color: white;
  padding: 0.75rem 1.25rem;
  border-radius: 8px;
  font-weight: bold;
  font-size: 1.2rem;
  letter-spacing: 1px;
  box-shadow: var(--shadow-md);
}

.logo-text {
  color: var(--color-text-light);
  font-size: 0.9rem;
  font-weight: 500;
}

.main-nav {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.nav-link {
  color: var(--color-text);
  text-decoration: none;
  padding: 0.75rem 1.25rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  font-weight: 500;
  position: relative;
}

.nav-link:hover {
  background: var(--color-primary-lighter);
  color: var(--color-primary);
  transform: translateY(-1px);
}

.nav-link.router-link-active {
  background: var(--color-primary);
  color: white;
  box-shadow: var(--shadow-sm);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.btn-primary {
  background: var(--color-primary);
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
}

.btn-primary:hover {
  background: var(--color-primary-dark);
}

.main-content {
  min-height: calc(100vh - 160px);
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.footer {
  background: var(--color-background-soft);
  border-top: 1px solid var(--color-border);
  margin-top: 4rem;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 3rem 2rem 2rem;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.footer-section h4 {
  color: var(--color-text);
  margin-bottom: 1rem;
  font-size: 1.1rem;
}

.footer-section ul {
  list-style: none;
  padding: 0;
}

.footer-section ul li {
  margin-bottom: 0.5rem;
}

.footer-section ul li a {
  color: var(--color-text-light);
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-section ul li a:hover {
  color: var(--color-primary);
}

.footer-bottom {
  border-top: 1px solid var(--color-border);
  padding: 1.5rem 2rem;
  text-align: center;
  color: var(--color-text-light);
}

@media (max-width: 768px) {
  .header-container {
    flex-direction: column;
    height: auto;
    padding: 1rem;
    gap: 1rem;
  }

  .main-nav {
    flex-wrap: wrap;
    justify-content: center;
    gap: 1rem;
  }

  .nav-link {
    font-size: 0.9rem;
    padding: 0.4rem 0.8rem;
  }

  .main-content {
    padding: 1rem;
  }
}
</style>
