<template>
  <div class="min-h-screen bg-light py-4">
    <div class="container">
      <!-- Header -->

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

              <!-- Household Members Count -->
              <div class="row mb-4">
                <div class="col-md-6 mb-3">
                  <label for="householdMembers" class="form-label fw-bold">
                    <i class="bi bi-people me-2"></i>
                    Hane Üyeleri Sayısı *
                  </label>
                  <input
                    id="householdMembers"
                    type="number"
                    v-model.number="wizardStore.householdMembers"
                    @input="wizardStore.setHouseholdMembers($event.target.value)"
                    :min="wizardStore.activeLines"
                    max="10"
                    class="form-control form-control-sm"
                    placeholder="Örn: 4"
                    required
                  />
                  <div class="form-text">
                    Hane içinde yaşayan toplam kişi sayısı (en az hat sayısı kadar olmalı)
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <label for="activeLines" class="form-label fw-bold">
                    <i class="bi bi-phone me-2"></i>
                    Aktif Mobil Hat Sayısı *
                  </label>
                  <input
                    id="activeLines"
                    type="number"
                    v-model.number="wizardStore.activeLines"
                    @input="wizardStore.setActiveLines($event.target.value)"
                    min="1"
                    :max="wizardStore.householdMembers"
                    class="form-control form-control-sm"
                    placeholder="Örn: 3"
                    required
                  />
                  <div class="form-text">
                    Şu anda aktif olan mobil hat sayısı (hane üyeleri sayısından fazla olamaz)
                  </div>
                </div>
              </div>

              <!-- Household Lines -->
              <div class="mb-4">
                <h5 class="mb-3">
                  <i class="bi bi-phone me-2"></i>
                  Mobil Hat Detayları
                </h5>
                <p class="text-muted small mb-3">
                  Her hat için aylık GB, dakika ve TV izleme sürelerini belirtin
                </p>

                <div
                  class="row g-3"
                  v-for="(line, index) in wizardStore.household"
                  :key="line.line_id"
                >
                  <div class="col-12">
                    <div class="card border-primary-light bg-light">
                      <div class="card-body py-3">
                        <div class="d-flex justify-content-between align-items-center mb-3">
                          <h6 class="mb-0 text-primary">{{ line.line_id }}</h6>
                          <button
                            v-if="wizardStore.household.length > 1"
                            @click="wizardStore.removeHouseholdLine(index)"
                            type="button"
                            class="btn btn-outline-danger btn-sm"
                            title="Bu hattı kaldır"
                          >
                            <i class="bi bi-trash"></i>
                          </button>
                        </div>

                        <div class="row g-2">
                          <div class="col-md-4">
                            <label :for="`gb-${index}`" class="form-label small">GB İhtiyacı</label>
                            <input
                              :id="`gb-${index}`"
                              type="number"
                              v-model.number="line.expected_gb"
                              @input="
                                wizardStore.updateHouseholdLine(
                                  index,
                                  'expected_gb',
                                  $event.target.value,
                                )
                              "
                              min="1"
                              max="1000"
                              class="form-control form-control-sm"
                              placeholder="10"
                            />
                          </div>

                          <div class="col-md-4">
                            <label :for="`min-${index}`" class="form-label small"
                              >Dakika İhtiyacı</label
                            >
                            <input
                              :id="`min-${index}`"
                              type="number"
                              v-model.number="line.expected_min"
                              @input="
                                wizardStore.updateHouseholdLine(
                                  index,
                                  'expected_min',
                                  $event.target.value,
                                )
                              "
                              min="0"
                              max="10000"
                              class="form-control form-control-sm"
                              placeholder="500"
                            />
                          </div>

                          <div class="col-md-4">
                            <label :for="`tv-${index}`" class="form-label small"
                              >Ortalama TV İzleme Saati</label
                            >
                            <input
                              :id="`tv-${index}`"
                              type="number"
                              v-model.number="line.tv_hd_hours"
                              @input="
                                wizardStore.updateHouseholdLine(
                                  index,
                                  'tv_hd_hours',
                                  $event.target.value,
                                )
                              "
                              min="0"
                              max="24"
                              class="form-control form-control-sm"
                              placeholder="4"
                            />
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="text-center mt-3">
                  <button
                    @click="wizardStore.addHouseholdLine"
                    type="button"
                    class="btn btn-outline-primary btn-sm"
                  >
                    <i class="bi bi-plus-circle me-2"></i>
                    Yeni Hat Ekle
                  </button>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="d-flex justify-content-between align-items-center pt-3 border-top">
                <div class="d-flex gap-2">
                  <RouterLink to="/onboarding" class="btn btn-outline-secondary btn-sm">
                    <i class="bi bi-house me-2"></i>
                    Ana Menüye Dön
                  </RouterLink>

                  <button
                    @click="wizardStore.resetWizard"
                    type="button"
                    class="btn btn-outline-warning btn-sm"
                  >
                    <i class="bi bi-arrow-clockwise me-2"></i>
                    Sıfırla
                  </button>
                </div>

                <button
                  @click="handleContinue"
                  :disabled="!wizardStore.canContinue"
                  type="button"
                  class="btn btn-primary"
                >
                  Devam Et
                  <i class="bi bi-arrow-right ms-2"></i>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Step 2: Recommendation Cards -->
      <div v-if="showRecommendations && !showCoverage" class="row justify-content-center">
        <div class="col-lg-10">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <div class="d-flex align-items-center justify-content-between flex-wrap gap-2">
                <div>
                  <h4 class="mb-0">
                    <i class="bi bi-lightbulb me-2"></i>
                    Size Özel Paket Önerileri
                  </h4>
                  <p class="text-muted mb-0 small">İhtiyaçlarınıza göre en uygun 3 kombinasyon</p>
                </div>
                <div class="btn-group btn-group-sm" role="group" aria-label="Currency">
                  <button
                    type="button"
                    class="btn"
                    :class="currency === 'TRY' ? 'btn-primary' : 'btn-outline-primary'"
                    @click="setCurrency('TRY')"
                  >
                    ₺ TRY
                  </button>
                  <button
                    type="button"
                    class="btn"
                    :class="currency === 'USD' ? 'btn-primary' : 'btn-outline-primary'"
                    @click="setCurrency('USD')"
                  >
                    $ USD
                  </button>
                </div>
              </div>
            </div>
            <div class="card-body py-4">
              <!-- Quick Edit Panel: District + Add Line -->
              <div class="row g-3 align-items-end mb-4">
                <div class="col-md-12 col-lg-6">
                  <label for="district-step2" class="form-label fw-bold">İlçe</label>
                  <select
                    id="district-step2"
                    v-model="wizardStore.district"
                    @change="wizardStore.setDistrict($event.target.value)"
                    class="form-select form-select-sm"
                    :disabled="!wizardStore.city"
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

              <!-- API driven recommendations -->
              <div v-if="wizardStore.loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Yükleniyor...</span>
                </div>
                <p class="mt-3">En uygun paketler hesaplanıyor...</p>
              </div>

              <div v-else-if="wizardStore.recommendations && wizardStore.recommendations.top3 && wizardStore.recommendations.top3.length" class="row">
                <div
                  v-for="(rec, idx) in wizardStore.recommendations.top3"
                  :key="idx"
                  class="col-md-4 mb-4"
                >
                  <div
                    class="card h-100 border-primary position-relative recommendation-card"
                    :class="{ 'selected-card': isSelected(rec) }"
                    @click="handleSelectRecommendation(rec)"
                    style="cursor: pointer;"
                  >
                    <div class="card-header text-white text-center" :class="isSelected(rec) ? 'bg-success' : 'bg-primary'">
                      <h5 class="mb-0">{{ idx + 1 }}. Öneri</h5>
                    </div>

                    <!-- Seçildi Rozeti -->
                    <div v-if="isSelected(rec)" class="selected-ribbon">
                      <i class="bi bi-check2-circle me-1"></i>
                      Seçildi
                    </div>

                    <div class="card-body">
                      <h6 class="card-title">{{ rec.combo_label }}</h6>
                      <p class="card-text">{{ rec.reasoning }}</p>

                      <div class="mb-2">
                        <strong>Mobil Planlar</strong>
                        <ul class="small mb-2 ms-3">
                          <li v-for="m in rec.items.mobile" :key="m.line_id">
                            {{ m.line_id }}: {{ m.plan.name }} ({{ m.plan.quota_gb }}GB)
                          </li>
                        </ul>
                      </div>

                      <div v-if="rec.items.home" class="mb-2 small">
                        <strong>Ev İnternet:</strong>
                        {{ rec.items.home.name }} ({{ rec.items.home.down_mbps }} Mbps)
                      </div>

                      <div v-if="rec.items.tv" class="mb-2 small">
                        <strong>TV:</strong>
                        {{ rec.items.tv.name }}
                      </div>

                      <div class="text-center mt-3">
                        <h4 class="text-primary mb-2">{{ formatPrice(rec.monthly_total) }}<span class="ms-1 text-muted small">{{ currency === 'TRY' ? '/ay' : '/mo' }}</span></h4>
                        <button
                          class="btn btn-sm"
                          :class="isSelected(rec) ? 'btn-success' : 'btn-outline-primary'"
                          @click.stop="handleSelectRecommendation(rec)"
                        >
                          <i class="bi" :class="isSelected(rec) ? 'bi-check2-circle me-2' : 'bi-cursor me-2'"></i>
                          {{ isSelected(rec) ? 'Seçildi' : 'Seç' }}
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div v-else class="text-center text-muted py-5">
                Henüz öneri bulunamadı
              </div>
            </div>
          </div>

          <!-- Navigation Buttons -->
          <div class="d-flex justify-content-between align-items-center mt-4">
            <button @click="goBack" type="button" class="btn btn-outline-secondary">
              <i class="bi bi-arrow-left me-2"></i>
              Geri Dön
            </button>

            <div class="d-flex gap-2">
              <button
                @click="continueToCoverage"
                :disabled="!selectedRecommendation"
                type="button"
                class="btn btn-primary"
              >
                Kapsama Kontrolü
                <i class="bi bi-wifi ms-2"></i>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Step 3: Coverage Badges -->
      <div v-if="showCoverage" class="row justify-content-center">
        <div class="col-lg-8">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-wifi me-2"></i>
                Adres Kapsama Durumu
              </h4>
            </div>
            <div class="card-body py-4">
              <CoverageBadge :addressId="wizardStore.getMockAddressId()" />
            </div>
          </div>

          <!-- Navigation Buttons -->
          <div class="d-flex justify-content-between align-items-center mt-4">
            <button @click="goBack" type="button" class="btn btn-outline-secondary">
              <i class="bi bi-arrow-left me-2"></i>
              Geri Dön
            </button>

            <button @click="continueToAppointment" type="button" class="btn btn-primary">
              Randevu Seç
              <i class="bi bi-calendar-check ms-2"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Success Popup -->
  <div v-if="showSuccessPopup" class="success-popup-overlay">
    <div class="success-popup-content">
      <div class="success-icon">
        <i class="bi bi-check-circle"></i>
      </div>
      <h3 class="success-title">✅ Randevunuz Oluşturuldu!</h3>
      <div class="success-details">
        <div class="detail-item">
          <i class="bi bi-calendar-event"></i>
          <span class="detail-label">Tarih:</span>
          <span class="detail-value">{{ successAppointment.date }}</span>
        </div>
        <div class="detail-item">
          <i class="bi bi-clock"></i>
          <span class="detail-label">Saat:</span>
          <span class="detail-value">{{ successAppointment.time }}</span>
        </div>
        <div class="detail-item">
          <i class="bi bi-geo-alt"></i>
          <span class="detail-label">Adres:</span>
          <span class="detail-value">{{ successAppointment.address }}</span>
        </div>
      </div>
      <p class="success-message">Kurulum ekibi belirtilen tarih ve saatte adresinizde olacaktır.</p>
      <div class="success-progress">
        <div class="progress-bar">
          <div class="progress-fill" :style="{ width: progressWidth + '%' }"></div>
        </div>
        <span class="progress-text">{{ progressText }}</span>
      </div>
    </div>
  </div>

  <!-- Appointment Modal -->
  <div v-if="showAppointmentModal" class="appointment-modal-overlay" @click="closeAppointmentModal">
    <div class="appointment-modal-content" @click.stop>
      <div class="modal-header">
        <div class="header-content">
          <div class="header-icon">
            <i class="bi bi-calendar-check"></i>
          </div>
          <div class="header-text">
            <h4 class="modal-title">Randevu Seçimi</h4>
            <p class="modal-subtitle">Kurulum için uygun tarih ve saati seçin</p>
          </div>
        </div>
        <button class="modal-close" @click="closeAppointmentModal">
          <i class="bi bi-x-lg"></i>
        </button>
      </div>

      <div class="modal-body">
        <!-- Date Selection -->
        <div class="date-section">
          <h5 class="section-title">
            <i class="bi bi-calendar-week me-2"></i>
            Tarih Seçimi
          </h5>
          <div class="date-grid">
            <div
              v-for="date in availableDates"
              :key="date.value"
              class="date-option"
              :class="{ active: selectedDate === date.value }"
              @click="selectDate(date.value)"
            >
              <div class="date-day">{{ date.day }}</div>
              <div class="date-number">{{ date.number }}</div>
              <div class="date-month">{{ date.month }}</div>
            </div>
          </div>
        </div>

        <!-- Time Selection -->
        <div class="time-section">
          <h5 class="section-title">
            <i class="bi bi-clock me-2"></i>
            Saat Seçimi
          </h5>
          <div class="time-grid">
            <div
              v-for="time in availableTimes"
              :key="time.value"
              class="time-option"
              :class="{ active: selectedTime === time.value }"
              @click="selectTime(time.value)"
            >
              {{ time.label }}
            </div>
          </div>
        </div>

        <!-- Appointment Summary -->
        <div class="appointment-summary">
          <div class="summary-header">
            <i class="bi bi-info-circle me-2"></i>
            <h6>Randevu Özeti</h6>
          </div>
          <div class="summary-content">
            <div class="summary-item">
              <span class="summary-label">Seçilen Tarih:</span>
              <span class="summary-value">{{ getSelectedDateText() }}</span>
            </div>
            <div class="summary-item">
              <span class="summary-label">Seçilen Saat:</span>
              <span class="summary-value">{{ getSelectedTimeText() }}</span>
            </div>
            <div class="summary-item">
              <span class="summary-label">Kurulum Süresi:</span>
              <span class="summary-value">2-4 saat</span>
            </div>
          </div>
        </div>

        <!-- Additional Notes -->
        <div class="notes-section">
          <label for="appointment-notes" class="form-label">
            <i class="bi bi-chat-text me-2"></i>
            Ek Notlar (Opsiyonel)
          </label>
          <textarea
            id="appointment-notes"
            v-model="appointmentNotes"
            class="form-control"
            rows="3"
            placeholder="Kurulum ekibi için özel notlarınızı buraya yazabilirsiniz..."
          ></textarea>
        </div>
      </div>

      <div class="modal-footer">
        <button class="btn btn-outline-secondary" @click="closeAppointmentModal">
          <i class="bi bi-x me-2"></i>
          İptal
        </button>
        <button
          class="btn btn-primary"
          :disabled="!selectedDate || !selectedTime"
          @click="confirmAppointment"
        >
          <i class="bi bi-check-lg me-2"></i>
          Randevuyu Onayla
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useWizardStore } from '@/stores/useWizardStore'
import CoverageBadge from '@/components/CoverageBadge.vue'
import RecommendationCards from '@/components/RecommendationCards.vue'

// Props
const props = defineProps({
  initialStep: {
    type: String,
    default: 'form',
    validator: (value) => ['form', 'recommendations', 'coverage'].includes(value),
  },
})

const router = useRouter()
const route = useRoute()
const wizardStore = useWizardStore()

// State
const showCoverage = ref(false)
const showRecommendations = ref(false)
const selectedRecommendation = ref(null)
const currency = ref('TRY') // 'TRY' | 'USD'
const USD_RATE = 38.7
const showAppointmentModal = ref(false)
const selectedDate = ref('')
const selectedTime = ref('')
const appointmentNotes = ref('')
const showSuccessPopup = ref(false)
const successAppointment = ref({})
const progressWidth = ref(0)
const progressText = ref('')

// Initialize step based on props or route
onMounted(() => {
  const step = props.initialStep || route.name?.replace('wizard-step', '') || 'form'
  setStep(step)
})

// Set step and update URL
const setStep = (step) => {
  showRecommendations.value = false
  showCoverage.value = false

  switch (step) {
    case '1':
    case 'form':
      // Stay on form
      break
    case '2':
    case 'recommendations':
      showRecommendations.value = true
      if (route.path !== '/wizard/step2') router.push('/wizard/step2')
      break
    case '3':
    case 'coverage':
      showRecommendations.value = true
      showCoverage.value = true
      if (route.path !== '/wizard/step3') router.push('/wizard/step3')
      break
  }
}

// Watch for route changes
watch(
  () => route.name,
  (newRoute) => {
    if (newRoute && newRoute.startsWith('wizard-step')) {
      const step = newRoute.replace('wizard-step', '')
      setStep(step)
    }
  },
)

const handleContinue = async () => {
  if (wizardStore.canContinue) {
    console.log('Form Data:', {
      city: wizardStore.city,
      district: wizardStore.district,
      householdMembers: wizardStore.householdMembers,
      activeLines: wizardStore.activeLines,
      household: wizardStore.household,
    })
    // Önce kapsama kontrolü, sonra öneriler
    await wizardStore.checkCoverage()
    await wizardStore.getRecommendations()
    showRecommendations.value = true
    router.push('/wizard/step2')
  }
}

const continueToCoverage = () => {
  showRecommendations.value = false
  showCoverage.value = true
  router.push('/wizard/step3')
}

const skipToCoverage = () => {
  // Öneri seçmeden doğrudan kapsama kontrolüne geç
  showRecommendations.value = false
  showCoverage.value = true
  router.push('/wizard/step3')
  console.log('Öneri adımı atlandı, doğrudan kapsama kontrolüne geçiliyor')
}

const goBack = () => {
  if (showCoverage.value) {
    showCoverage.value = false
    showRecommendations.value = true
    router.push('/wizard/step2')
  } else if (showRecommendations.value) {
    showRecommendations.value = false
    router.push('/wizard/step1')
  } else {
    router.push('/onboarding')
  }
}

const handleSelectRecommendation = (recommendation) => {
  if (
    selectedRecommendation.value &&
    recommendation &&
    recommendation.combo_label === selectedRecommendation.value.combo_label &&
    Number(recommendation.monthly_total).toFixed(2) === Number(selectedRecommendation.value.monthly_total).toFixed(2)
  ) {
    // Aynı karta tekrar tıklandı: seçimi kaldır
    selectedRecommendation.value = null
  } else {
    // Farklı kart: seçimi ata
    selectedRecommendation.value = recommendation
  }
}

const isSelected = (rec) => {
  if (!selectedRecommendation.value) return false
  // Benzersiz anahtar olarak etiket + toplamı kullan
  return (
    rec.combo_label === selectedRecommendation.value.combo_label &&
    Number(rec.monthly_total).toFixed(2) === Number(selectedRecommendation.value.monthly_total).toFixed(2)
  )
}

const setCurrency = (c) => {
  currency.value = c
}

const formatPrice = (amount) => {
  const n = Number(amount || 0)
  if (currency.value === 'USD') {
    return `$${(n / USD_RATE).toFixed(2)}`
  }
  return `₺${n.toFixed(2)}`
}

const handleShowDetails = (recommendation) => {
  console.log('Show details for:', recommendation)
  // TODO: Show detailed modal
}

// Available dates for the next 2 weeks
const availableDates = ref([
  { value: '2024-01-15', day: 'Pzt', number: '15', month: 'Oca' },
  { value: '2024-01-16', day: 'Sal', number: '16', month: 'Oca' },
  { value: '2024-01-17', day: 'Çar', number: '17', month: 'Oca' },
  { value: '2024-01-18', day: 'Per', number: '18', month: 'Oca' },
  { value: '2024-01-19', day: 'Cum', number: '19', month: 'Oca' },
  { value: '2024-01-22', day: 'Pzt', number: '22', month: 'Oca' },
  { value: '2024-01-23', day: 'Sal', number: '23', month: 'Oca' },
  { value: '2024-01-24', day: 'Çar', number: '24', month: 'Oca' },
  { value: '2024-01-25', day: 'Per', number: '25', month: 'Oca' },
  { value: '2024-01-26', day: 'Cum', number: '26', month: 'Oca' },
])

// Available time slots
const availableTimes = ref([
  { value: '09:00', label: '09:00 - 11:00' },
  { value: '11:00', label: '11:00 - 13:00' },
  { value: '13:00', label: '13:00 - 15:00' },
  { value: '15:00', label: '15:00 - 17:00' },
  { value: '17:00', label: '17:00 - 19:00' },
])

const continueToAppointment = () => {
  showAppointmentModal.value = true
}

const closeAppointmentModal = () => {
  showAppointmentModal.value = false
  selectedDate.value = ''
  selectedTime.value = ''
  appointmentNotes.value = ''
}

const selectDate = (date) => {
  selectedDate.value = date
}

const selectTime = (time) => {
  selectedTime.value = time
}

const getSelectedDateText = () => {
  if (!selectedDate.value) return 'Seçilmedi'
  const date = availableDates.value.find((d) => d.value === selectedDate.value)
  return date ? `${date.number} ${date.month} ${date.day}` : 'Seçilmedi'
}

const getSelectedTimeText = () => {
  if (!selectedTime.value) return 'Seçilmedi'
  const time = availableTimes.value.find((t) => t.value === selectedTime.value)
  return time ? time.label : 'Seçilmedi'
}

const confirmAppointment = () => {
  if (selectedDate.value && selectedTime.value) {
    // Set success appointment data BEFORE closing modal
    const selectedDateText = getSelectedDateText()
    const selectedTimeText = getSelectedTimeText()
    const addressText =
      wizardStore.city && wizardStore.district
        ? `${wizardStore.city} / ${wizardStore.district}`
        : 'Adres bilgisi girilmemiş'

    console.log('Selected Date:', selectedDate.value)
    console.log('Selected Time:', selectedTime.value)
    console.log('Date Text:', selectedDateText)
    console.log('Time Text:', selectedTimeText)
    console.log('City from store:', wizardStore.city)
    console.log('District from store:', wizardStore.district)
    console.log('Final Address Text:', addressText)

    successAppointment.value = {
      date: selectedDateText,
      time: selectedTimeText,
      address: addressText,
    }

    // Close modal and reset
    closeAppointmentModal()

    // Show success popup
    showSuccessPopup.value = true

    // Start progress animation
    startProgressAnimation()
  }
}

const startProgressAnimation = () => {
  progressWidth.value = 0
  progressText.value = 'Ana sayfaya yönlendiriliyor...'

  const duration = 5000 // 5 seconds (increased from 2)
  const interval = 50 // Update every 50ms
  const steps = duration / interval
  const increment = 100 / steps

  const timer = setInterval(() => {
    progressWidth.value += increment

    if (progressWidth.value >= 100) {
      clearInterval(timer)
      progressWidth.value = 100
      progressText.value = 'Yönlendiriliyor...'

      // Navigate to onboarding after animation
      setTimeout(() => {
        router.push('/onboarding')
      }, 1000) // Increased from 500ms
    }
  }, interval)
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

/* Success Popup Styles */
.success-popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(12px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10000;
  padding: 1rem;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.success-popup-content {
  background: var(--color-background-card);
  border-radius: 24px;
  max-width: 500px;
  width: 100%;
  padding: 3rem 2rem;
  text-align: center;
  box-shadow: 0 25px 80px rgba(0, 0, 0, 0.4);
  border: 2px solid var(--color-success-light);
  animation: popupSlideIn 0.4s ease-out;
  position: relative;
  overflow: hidden;
}

@keyframes popupSlideIn {
  from {
    opacity: 0;
    transform: translateY(-30px) scale(0.9);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.success-popup-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--color-success), var(--color-success-light));
  border-radius: 24px 24px 0 0;
}

.success-icon {
  width: 100px;
  height: 100px;
  background: linear-gradient(135deg, var(--color-success), var(--color-success-dark));
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 2rem;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
  animation: iconBounce 0.6s ease-out 0.3s both;
}

@keyframes iconBounce {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}

.success-icon i {
  font-size: 3rem;
  color: white;
}

.success-title {
  color: var(--color-success-dark);
  margin-bottom: 2rem;
  font-size: 1.8rem;
  font-weight: 700;
  animation: titleSlideIn 0.5s ease-out 0.4s both;
}

@keyframes titleSlideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.success-details {
  margin-bottom: 2rem;
  animation: detailsSlideIn 0.5s ease-out 0.5s both;
}

@keyframes detailsSlideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.detail-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
  padding: 1rem;
  background: var(--color-background-soft);
  border-radius: 12px;
  border: 1px solid var(--color-border);
  flex-wrap: wrap;
  text-align: center;
}

.detail-item i {
  color: var(--color-primary);
  font-size: 1.2rem;
  width: 20px;
}

.detail-label {
  color: var(--color-text-mute);
  font-weight: 500;
  min-width: 60px;
}

.detail-value {
  color: var(--color-text);
  font-weight: 600;
  font-size: 1.1rem;
  word-break: break-word;
  max-width: 200px;
  text-align: center;
}

.success-message {
  color: var(--color-text-mute);
  margin-bottom: 2rem;
  line-height: 1.6;
  animation: messageSlideIn 0.5s ease-out 0.6s both;
}

@keyframes messageSlideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.success-progress {
  animation: progressSlideIn 0.5s ease-out 0.7s both;
}

@keyframes progressSlideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Responsive Success Popup */
@media (max-width: 768px) {
  .success-popup-content {
    max-width: 90%;
    padding: 2rem 1.5rem;
  }

  .success-title {
    font-size: 1.5rem;
  }

  .detail-item {
    flex-direction: column;
    gap: 0.5rem;
    padding: 0.75rem;
  }

  .detail-label {
    min-width: auto;
  }

  .detail-value {
    max-width: 100%;
    font-size: 1rem;
  }
}

/* Responsive Navigation Buttons */
@media (max-width: 768px) {
  .d-flex.justify-content-between.align-items-center {
    flex-direction: column;
    gap: 1rem;
  }

  .d-flex.gap-2 {
    gap: 0.5rem !important;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.5rem 0.75rem;
  }
}

/* Skip Button Styles */
.btn-outline-warning {
  border-color: var(--color-warning);
  color: var(--color-warning);
  transition: all var(--transition-fast);
}

.btn-outline-warning:hover {
  background-color: var(--color-warning);
  color: var(--color-text-inverse);
  transform: translateY(-1px);
  box-shadow: var(--shadow-sm);
}

.btn-outline-warning:focus {
  box-shadow: 0 0 0 3px rgba(255, 193, 7, 0.25);
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: var(--color-background-soft);
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 1rem;
  border: 1px solid var(--color-border);
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--color-success), var(--color-success-light));
  border-radius: 4px;
  transition: width 0.05s linear;
  position: relative;
  overflow: hidden;
}

.progress-fill::after {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: shimmer 1.5s infinite;
}

@keyframes shimmer {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

.progress-text {
  color: var(--color-success);
  font-size: 0.9rem;
  font-weight: 500;
}

/* Appointment Modal Styles */
.appointment-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  padding: 1rem;
}

.appointment-modal-content {
  background: var(--color-background-card);
  border-radius: 24px;
  max-width: 600px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  border: 2px solid var(--color-border);
  animation: modalSlideIn 0.3s ease-out;
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.modal-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 2rem 2rem 1.5rem;
  border-bottom: 2px solid var(--color-border);
  background: linear-gradient(135deg, var(--color-primary-lighter), var(--color-primary-light));
  border-radius: 24px 24px 0 0;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  flex: 1;
}

.header-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.header-icon i {
  font-size: 1.8rem;
  color: white;
}

.header-text {
  flex: 1;
}

.modal-title {
  color: var(--color-primary-dark);
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  font-weight: 700;
}

.modal-subtitle {
  color: var(--color-text-mute);
  margin: 0;
  font-size: 0.95rem;
}

.modal-close {
  background: none;
  border: none;
  color: var(--color-text-mute);
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 50%;
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.modal-close:hover {
  background: var(--color-background-soft);
  color: var(--color-text);
}

.modal-body {
  padding: 2rem;
}

.section-title {
  color: var(--color-primary);
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  font-weight: 600;
  display: flex;
  align-items: center;
}

.date-section,
.time-section {
  margin-bottom: 2rem;
}

.date-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(80px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.date-option {
  background: var(--color-background-soft);
  border: 2px solid var(--color-border);
  border-radius: 16px;
  padding: 1rem 0.5rem;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.date-option:hover {
  border-color: var(--color-primary-light);
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.date-option.active {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  border-color: var(--color-primary);
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
}

.date-day {
  font-size: 0.8rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
  opacity: 0.8;
}

.date-number {
  font-size: 1.2rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
}

.date-month {
  font-size: 0.75rem;
  opacity: 0.8;
}

.time-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
  gap: 1rem;
}

.time-option {
  background: var(--color-background-soft);
  border: 2px solid var(--color-border);
  border-radius: 12px;
  padding: 1rem;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
}

.time-option:hover {
  border-color: var(--color-primary-light);
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.time-option.active {
  background: linear-gradient(135deg, var(--color-primary), var(--color-primary-dark));
  border-color: var(--color-primary);
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
}

.appointment-summary {
  background: var(--color-background-soft);
  border: 2px solid var(--color-border);
  border-radius: 16px;
  padding: 1.5rem;
  margin-bottom: 2rem;
}

.summary-header {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.summary-header h6 {
  color: var(--color-primary);
  margin: 0;
  font-weight: 600;
}

.summary-content {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: var(--color-background-card);
  border-radius: 12px;
  border: 1px solid var(--color-border);
}

.summary-label {
  color: var(--color-text-mute);
  font-weight: 500;
}

.summary-value {
  color: var(--color-text);
  font-weight: 600;
}

.notes-section {
  margin-bottom: 2rem;
}

.notes-section .form-label {
  color: var(--color-text);
  font-weight: 600;
  margin-bottom: 0.75rem;
  display: flex;
  align-items: center;
}

.notes-section .form-control {
  border: 2px solid var(--color-border);
  border-radius: 12px;
  background: var(--color-background-input);
  color: var(--color-text);
  transition: all 0.3s ease;
}

.notes-section .form-control:focus {
  border-color: var(--color-primary);
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.15);
  background: var(--color-background-card);
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  padding: 1.5rem 2rem 2rem;
  border-top: 2px solid var(--color-border);
  background: var(--color-background-soft);
  border-radius: 0 0 24px 24px;
}

.modal-footer .btn {
  padding: 0.875rem 1.5rem;
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.modal-footer .btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
}

/* Responsive */
@media (max-width: 768px) {
  .appointment-modal-content {
    max-width: 95vw;
    margin: 1rem;
  }

  .modal-header {
    padding: 1.5rem 1.5rem 1rem;
  }

  .modal-body {
    padding: 1.5rem;
  }

  .modal-footer {
    padding: 1rem 1.5rem 1.5rem;
    flex-direction: column;
  }

  .header-content {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }

  .date-grid {
    grid-template-columns: repeat(auto-fit, minmax(70px, 1fr));
  }

  .time-grid {
    grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  }
}

/* Seçili kart görselleştirme */
.recommendation-card {
  transition: transform 0.12s ease, box-shadow 0.12s ease, border-color 0.12s ease;
}
.recommendation-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.08);
}
.selected-card {
  border-width: 2px !important;
  border-color: #28a745 !important; /* success */
  box-shadow: 0 0 0 0.25rem rgba(40, 167, 69, 0.15);
}
.selected-ribbon {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: #28a745;
  color: #fff;
  padding: 4px 10px;
  border-radius: 999px;
  font-size: 0.8rem;
  display: inline-flex;
  align-items: center;
}
</style>
