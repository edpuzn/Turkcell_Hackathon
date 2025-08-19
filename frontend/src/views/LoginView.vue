<template>
  <div class="login-page">
    <div class="login-container">
      <!-- Left Side - Branding -->
      <div class="branding-section">
        <div class="branding-content">
          <div class="logo-placeholder">TURKCELL</div>
          <h1 class="brand-title">Ev+Mobil Paket Danışmanı</h1>
          <p class="brand-subtitle">
            Size en uygun paketi bulmak için giriş yapın ve kişiselleştirilmiş 
            önerilerimizi keşfedin.
          </p>
          <div class="features-list">
            <div class="feature-item">
              <i class="bi bi-check-circle-fill"></i>
              <span>Kişiselleştirilmiş paket önerileri</span>
            </div>
            <div class="feature-item">
              <i class="bi bi-check-circle-fill"></i>
              <span>Hızlı ve kolay sipariş süreci</span>
            </div>
            <div class="feature-item">
              <i class="bi bi-check-circle-fill"></i>
              <span>7/24 müşteri desteği</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Side - Login Form -->
      <div class="form-section">
        <div class="form-container">
          <div class="form-header">
            <h2 class="form-title">Giriş Yap</h2>
            <p class="form-subtitle">Hesabınıza giriş yaparak devam edin</p>
          </div>

          <form class="login-form" @submit.prevent="handleLogin">
            <div class="form-group">
              <label class="form-label">E-posta veya Telefon</label>
              <div class="input-wrapper">
                <i class="bi bi-envelope input-icon"></i>
                <input 
                  type="email" 
                  class="form-control" 
                  placeholder="ornek@email.com"
                  v-model="formData.email"
                  required
                />
              </div>
            </div>

            <div class="form-group">
              <label class="form-label">Şifre</label>
              <div class="input-wrapper">
                <i class="bi bi-lock input-icon"></i>
                <input 
                  :type="showPassword ? 'text' : 'password'" 
                  class="form-control" 
                  placeholder="Şifrenizi girin"
                  v-model="formData.password"
                  required
                />
                <button 
                  type="button" 
                  class="password-toggle"
                  @click="togglePassword"
                >
                  <i :class="showPassword ? 'bi bi-eye-slash' : 'bi bi-eye'"></i>
                </button>
              </div>
            </div>

            <div class="form-options">
              <label class="checkbox-label">
                <input type="checkbox" v-model="formData.rememberMe" />
                <span>Beni hatırla</span>
              </label>
              <RouterLink to="/forgot-password" class="forgot-password">
                Şifremi unuttum
              </RouterLink>
            </div>

            <button type="submit" class="btn-login" :disabled="isLoading">
              <span v-if="!isLoading">Giriş Yap</span>
              <div v-else class="loading-spinner"></div>
            </button>
          </form>

          <div class="login-divider">
            <span>veya</span>
          </div>

          <div class="social-login">
            <button type="button" class="btn-social google">
              <i class="bi bi-google"></i>
              Google ile Giriş
            </button>
            <button type="button" class="btn-social apple">
              <i class="bi bi-apple"></i>
              Apple ile Giriş
            </button>
          </div>

          <div class="register-section">
            <p class="register-text">
              Hesabınız yok mu? 
              <RouterLink to="/register" class="register-link">
                Hemen üye olun
              </RouterLink>
            </p>
          </div>

          <div class="back-section">
            <RouterLink to="/onboarding" class="back-link">
              <i class="bi bi-arrow-left me-2"></i>
              Geri Dön
            </RouterLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isLoading = ref(false)
const showPassword = ref(false)

const formData = reactive({
  email: '',
  password: '',
  rememberMe: false
})

const togglePassword = () => {
  showPassword.value = !showPassword.value
}

const handleLogin = async () => {
  isLoading.value = true
  
  // Simulate API call
  setTimeout(() => {
    isLoading.value = false
    // For now, just redirect to wizard
    router.push('/wizard')
  }, 1500)
}
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--color-background-soft), var(--color-background));
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem 1rem;
}

.login-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  max-width: 1200px;
  width: 100%;
  background: var(--color-background-card);
  border-radius: 24px;
  overflow: hidden;
  box-shadow: var(--shadow-2xl);
  border: 1px solid var(--color-border);
}

/* Branding Section */
.branding-section {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  color: var(--color-text-inverse);
  padding: 3rem 2rem;
  display: flex;
  align-items: center;
  position: relative;
  overflow: hidden;
}

.branding-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grain" width="100" height="100" patternUnits="userSpaceOnUse"><circle cx="25" cy="25" r="1" fill="rgba(255,255,255,0.1)"/><circle cx="75" cy="75" r="1" fill="rgba(255,255,255,0.1)"/><circle cx="50" cy="10" r="0.5" fill="rgba(255,255,255,0.1)"/></pattern></defs><rect width="100" height="100" fill="url(%23grain)"/></svg>');
  opacity: 0.3;
}

.branding-content {
  position: relative;
  z-index: 1;
  text-align: center;
}

.logo-placeholder {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 2rem;
  letter-spacing: 3px;
  opacity: 0.9;
}

.brand-title {
  font-size: 2.2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.brand-subtitle {
  font-size: 1.1rem;
  line-height: 1.6;
  margin-bottom: 2.5rem;
  opacity: 0.9;
}

.features-list {
  text-align: left;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
  font-size: 1rem;
}

.feature-item i {
  color: var(--color-yellow);
  font-size: 1.2rem;
}

/* Form Section */
.form-section {
  padding: 3rem 2rem;
  display: flex;
  align-items: center;
}

.form-container {
  width: 100%;
  max-width: 400px;
  margin: 0 auto;
}

.form-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.form-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-text);
  margin-bottom: 0.5rem;
}

.form-subtitle {
  color: var(--color-text-mute);
  font-size: 1rem;
}

.login-form {
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: block;
  color: var(--color-text);
  font-weight: 600;
  font-size: 0.95rem;
  margin-bottom: 0.75rem;
}

.input-wrapper {
  position: relative;
}

.input-icon {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: var(--color-text-mute);
  font-size: 1.1rem;
  z-index: 2;
}

.form-control {
  width: 100%;
  padding: 1rem 1rem 1rem 3rem;
  border: 2px solid var(--color-border);
  border-radius: 16px;
  background: var(--color-background-input);
  color: var(--color-text);
  font-size: 1rem;
  transition: all 0.3s ease;
  box-sizing: border-box;
}

.form-control:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.15);
  background: var(--color-background-card);
}

.form-control::placeholder {
  color: var(--color-text-mute);
  opacity: 0.7;
}

.password-toggle {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--color-text-mute);
  cursor: pointer;
  font-size: 1.1rem;
  padding: 0.25rem;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.password-toggle:hover {
  color: var(--color-primary);
  background: var(--color-primary-lighter);
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  font-size: 0.9rem;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  color: var(--color-text);
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  width: 18px;
  height: 18px;
  accent-color: var(--color-primary);
}

.forgot-password {
  color: var(--color-primary);
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}

.forgot-password:hover {
  color: var(--color-primary-dark);
}

.btn-login {
  width: 100%;
  padding: 1rem;
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  color: var(--color-text-inverse);
  border: none;
  border-radius: 16px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.btn-login:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: var(--shadow-lg);
}

.btn-login:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.login-divider {
  text-align: center;
  margin: 2rem 0;
  position: relative;
}

.login-divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: var(--color-border);
}

.login-divider span {
  background: var(--color-background-card);
  padding: 0 1rem;
  color: var(--color-text-mute);
  font-size: 0.9rem;
}

.social-login {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.btn-social {
  width: 100%;
  padding: 1rem;
  border: 2px solid var(--color-border);
  border-radius: 16px;
  background: var(--color-background-card);
  color: var(--color-text);
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
}

.btn-social:hover {
  border-color: var(--color-primary);
  background: var(--color-primary-lighter);
  transform: translateY(-2px);
}

.btn-social.google:hover {
  border-color: #ea4335;
  background: rgba(234, 67, 53, 0.1);
}

.btn-social.apple:hover {
  border-color: #000;
  background: rgba(0, 0, 0, 0.1);
}

.register-section {
  text-align: center;
  margin-bottom: 2rem;
}

.register-text {
  color: var(--color-text-mute);
  margin: 0;
}

.register-link {
  color: var(--color-primary);
  text-decoration: none;
  font-weight: 600;
  margin-left: 0.25rem;
}

.register-link:hover {
  text-decoration: underline;
}

.back-section {
  text-align: center;
}

.back-link {
  color: var(--color-text-mute);
  text-decoration: none;
  font-size: 0.9rem;
  transition: color 0.2s ease;
  display: inline-flex;
  align-items: center;
}

.back-link:hover {
  color: var(--color-primary);
}

/* Responsive */
@media (max-width: 768px) {
  .login-container {
    grid-template-columns: 1fr;
    max-width: 500px;
  }
  
  .branding-section {
    padding: 2rem 1.5rem;
  }
  
  .brand-title {
    font-size: 1.8rem;
  }
  
  .form-section {
    padding: 2rem 1.5rem;
  }
  
  .form-title {
    font-size: 1.6rem;
  }
}

@media (max-width: 480px) {
  .login-page {
    padding: 1rem 0.5rem;
  }
  
  .branding-section,
  .form-section {
    padding: 1.5rem 1rem;
  }
  
  .form-options {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }
}
</style>
