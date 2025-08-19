<template>
  <div class="min-h-screen bg-light py-4">
    <div class="container">
      <!-- Header -->
      <div class="text-center mb-5">
        <h1 class="text-primary mb-2">Turkcell Ev+Mobil Paket Danışmanı</h1>
        <p class="text-muted fs-5">
          Size en uygun paketi bulmak için birkaç bilgiye ihtiyacımız var
        </p>
      </div>

      <!-- Progress Steps -->
      <div class="row justify-content-center mb-5">
        <div class="col-md-8">
          <div class="d-flex align-items-center justify-content-between">
            <div class="d-flex align-items-center">
              <div class="progress-step active">1</div>
              <span class="ms-2 text-primary fw-bold">Hane & Adres</span>
            </div>
            <div
              class="progress-line"
              :class="{ active: showRecommendations || showCoverage }"
            ></div>
            <div class="d-flex align-items-center">
              <div
                class="progress-step"
                :class="{ active: showRecommendations, completed: showCoverage }"
              >
                2
              </div>
              <span
                class="ms-2"
                :class="{
                  'text-primary fw-bold': showRecommendations || showCoverage,
                  'text-muted': !showRecommendations && !showCoverage,
                }"
                >Öneriler</span
              >
            </div>
            <div class="progress-line" :class="{ active: showCoverage }"></div>
            <div class="d-flex align-items-center">
              <div class="progress-step" :class="{ active: showCoverage, completed: showCoverage }">
                3
              </div>
              <span
                class="ms-2"
                :class="{ 'text-primary fw-bold': showCoverage, 'text-muted': !showCoverage }"
                >Kapsama</span
              >
            </div>
          </div>
        </div>
      </div>

      <!-- Step 1: Form (Hidden when showing recommendations or coverage) -->
      <div v-if="!showRecommendations && !showCoverage" class="row justify-content-center">
        <div class="col-lg-8">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-house-door me-2"></i>
                Hane Profili & Adres Bilgileri
              </h4>
            </div>
            <div class="card-body py-4">
              <!-- Address Selection -->
              <div class="row mb-4">
                <div class="col-md-6 mb-3">
                  <label for="city" class="form-label fw-bold">Şehir *</label>
                  <select
                    id="city"
                    v-model="wizardStore.city"
                    @change="wizardStore.setCity($event.target.value)"
                    class="form-select form-select-sm"
                    required
                  >
                    <option value="">Şehir seçiniz</option>
                    <option v-for="city in wizardStore.cities" :key="city" :value="city">
                      {{ city }}
                    </option>
                  </select>
                </div>

                <div class="col-md-6 mb-3">
                  <label for="district" class="form-label fw-bold">İlçe *</label>
                  <select
                    id="district"
                    v-model="wizardStore.district"
                    @change="wizardStore.setDistrict($event.target.value)"
                    class="form-select form-select-sm"
                    :disabled="!wizardStore.city"
                    required
                  >
                    <option value="">İlçe seçiniz</option>
                    <option
                      v-for="district in wizardStore.availableDistricts"
                      :key="district"
                      :value="district"
                    >
                      {{ district }}
                    </option>
                  </select>
                </div>
              </div>

              <!-- Household Information -->
              <div class="row mb-4">
                <div class="col-md-6 mb-3">
                  <label for="householdMembers" class="form-label fw-bold">Hane Halkı Sayısı *</label>
                  <input
                    id="householdMembers"
                    v-model.number="wizardStore.householdMembers"
                    @input="wizardStore.setHouseholdMembers($event.target.value)"
                    type="number"
                    class="form-control form-control-sm"
                    min="1"
                    max="10"
                    required
                  />
                </div>

                <div class="col-md-6 mb-3">
                  <label for="activeLines" class="form-label fw-bold">Aktif Mobil Hat Sayısı *</label>
                  <input
                    id="activeLines"
                    v-model.number="wizardStore.activeLines"
                    @input="wizardStore.setActiveLines($event.target.value)"
                    type="number"
                    class="form-control form-control-sm"
                    min="1"
                    max="5"
                    required
                  />
                </div>
              </div>

              <!-- Household Lines -->
              <div class="mb-4">
                <h5 class="mb-3">
                  <i class="bi bi-phone me-2"></i>
                  Mobil Hat Detayları
                </h5>
                <div
                  v-for="(line, index) in wizardStore.household"
                  :key="index"
                  class="row mb-3 p-3 border rounded"
                >
                  <div class="col-12 mb-2">
                    <h6 class="mb-0">{{ line.line_id }}</h6>
                  </div>
                  <div class="col-md-4 mb-2">
                    <label class="form-label small">GB İhtiyacı</label>
                    <input
                      v-model.number="line.expected_gb"
                      @input="wizardStore.updateHouseholdLine(index, 'expected_gb', $event.target.value)"
                      type="number"
                      class="form-control form-control-sm"
                      min="1"
                      max="100"
                    />
                  </div>
                  <div class="col-md-4 mb-2">
                    <label class="form-label small">Dakika İhtiyacı</label>
                    <input
                      v-model.number="line.expected_min"
                      @input="wizardStore.updateHouseholdLine(index, 'expected_min', $event.target.value)"
                      type="number"
                      class="form-control form-control-sm"
                      min="0"
                      max="2000"
                    />
                  </div>
                  <div class="col-md-4 mb-2">
                    <label class="form-label small">TV HD Saat</label>
                    <input
                      v-model.number="line.tv_hd_hours"
                      @input="wizardStore.updateHouseholdLine(index, 'tv_hd_hours', $event.target.value)"
                      type="number"
                      class="form-control form-control-sm"
                      min="0"
                      max="24"
                    />
                  </div>
                </div>
              </div>

              <!-- Continue Button -->
              <div class="text-center">
                <button
                  @click="handleContinue"
                  :disabled="!wizardStore.canContinue || wizardStore.loading"
                  class="btn btn-primary btn-lg px-5"
                >
                  <span v-if="wizardStore.loading" class="spinner-border spinner-border-sm me-2"></span>
                  {{ wizardStore.loading ? 'Kontrol Ediliyor...' : 'Devam Et' }}
                </button>
              </div>

              <!-- Error Message -->
              <div v-if="wizardStore.error" class="alert alert-danger mt-3">
                <i class="bi bi-exclamation-triangle me-2"></i>
                {{ wizardStore.error }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Step 2: Recommendations -->
      <div v-if="showRecommendations" class="row justify-content-center">
        <div class="col-lg-10">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-lightbulb me-2"></i>
                Size Özel Paket Önerileri
              </h4>
            </div>
            <div class="card-body py-4">
              <div v-if="wizardStore.loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Yükleniyor...</span>
                </div>
                <p class="mt-3">En uygun paketler hesaplanıyor...</p>
              </div>

              <div v-else-if="wizardStore.recommendations?.top3" class="row">
                <div
                  v-for="(recommendation, index) in wizardStore.recommendations.top3"
                  :key="index"
                  class="col-md-4 mb-4"
                >
                  <div class="card h-100 border-primary">
                    <div class="card-header bg-primary text-white text-center">
                      <h5 class="mb-0">{{ index + 1 }}. Öneri</h5>
                    </div>
                    <div class="card-body">
                      <h6 class="card-title">{{ recommendation.combo_label }}</h6>
                      <p class="card-text">{{ recommendation.reasoning }}</p>
                      
                      <div class="mb-3">
                        <strong>Mobil Planlar:</strong>
                        <ul class="list-unstyled ms-3">
                          <li v-for="mobile in recommendation.items.mobile" :key="mobile.line_id">
                            {{ mobile.line_id }}: {{ mobile.plan.name }} ({{ mobile.plan.quota_gb }}GB)
                          </li>
                        </ul>
                      </div>

                      <div v-if="recommendation.items.home" class="mb-3">
                        <strong>Ev İnternet:</strong>
                        <p class="mb-1">{{ recommendation.items.home.name }} ({{ recommendation.items.home.down_mbps }}Mbps)</p>
                      </div>

                      <div v-if="recommendation.items.tv" class="mb-3">
                        <strong>TV:</strong>
                        <p class="mb-1">{{ recommendation.items.tv.name }}</p>
                      </div>

                      <div class="text-center">
                        <h4 class="text-primary mb-2">₺{{ recommendation.monthly_total.toFixed(2) }}/ay</h4>
                        <button class="btn btn-outline-primary btn-sm">Seç</button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="text-center mt-4">
                <button @click="showRecommendations = false" class="btn btn-secondary me-2">
                  <i class="bi bi-arrow-left me-2"></i>
                  Geri
                </button>
                <button @click="showCoverage = true" class="btn btn-primary">
                  Kapsama Kontrolü
                  <i class="bi bi-arrow-right ms-2"></i>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Step 3: Coverage -->
      <div v-if="showCoverage" class="row justify-content-center">
        <div class="col-lg-8">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-geo-alt me-2"></i>
                Adres Kapsama Bilgileri
              </h4>
            </div>
            <div class="card-body py-4">
              <div v-if="wizardStore.coverage" class="row">
                <div
                  v-for="tech in wizardStore.coverage"
                  :key="tech.tech"
                  class="col-md-4 mb-3"
                >
                  <div class="card text-center">
                    <div class="card-body">
                      <h5 class="card-title">{{ tech.tech.toUpperCase() }}</h5>
                      <p class="card-text">{{ tech.down_mbps }} Mbps</p>
                      <span class="badge bg-success">Mevcut</span>
                    </div>
                  </div>
                </div>
              </div>

              <div class="text-center mt-4">
                <button @click="showCoverage = false" class="btn btn-secondary me-2">
                  <i class="bi bi-arrow-left me-2"></i>
                  Geri
                </button>
                <button @click="resetWizard" class="btn btn-success">
                  <i class="bi bi-check-circle me-2"></i>
                  Tamamla
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useWizardStore } from '@/stores/useWizardStore'

const wizardStore = useWizardStore()
const showCoverage = ref(false)
const showRecommendations = ref(false)

const handleContinue = async () => {
  if (wizardStore.canContinue) {
    console.log('Form Data:', {
      city: wizardStore.city,
      district: wizardStore.district,
      householdMembers: wizardStore.householdMembers,
      activeLines: wizardStore.activeLines,
      household: wizardStore.household,
    })
    
    // Önce kapsama kontrolü yap
    await wizardStore.checkCoverage()
    
    // Sonra önerileri al
    await wizardStore.getRecommendations()
    
    showRecommendations.value = true
  }
}

const resetWizard = () => {
  wizardStore.reset()
  showRecommendations.value = false
  showCoverage.value = false
}
</script>

<style scoped>
.bg-light {
  background-color: var(--color-background-soft) !important;
}

.border-primary-light {
  border-color: var(--color-primary-lighter) !important;
}

.text-primary {
  color: var(--color-primary) !important;
}

.text-muted {
  color: var(--color-text-light) !important;
}

.bg-primary-light {
  background-color: var(--color-primary-lighter) !important;
}
</style>
