<script setup>
import { ref, computed, onMounted } from 'vue'
import RecommendationModal from '@/components/RecommendationModal.vue'
import SelectionModal from '@/components/SelectionModal.vue'
import { useBundleStore } from '@/stores/bundle'

const bundleStore = useBundleStore()

// Step management
const currentStep = ref(1)

// Household profile
const householdProfile = ref({
  size: '',
  activeLines: '',
})

// Form data
const selectedCity = ref('')
const selectedDistrict = ref('')
const selectedAddressId = ref('')
const householdLines = ref([{ id: 1, gb: '10', minutes: '500', tvHours: '0' }])

// Coverage and recommendations
const coverageData = ref([])
const recommendations = ref([])
const installSlots = ref([])

// Modal states
const showModal = ref(false)
const showSelectionModal = ref(false)
const selectedRecommendation = ref(null)

// Computed properties
const cities = computed(() => bundleStore.getCities())
const filteredDistricts = computed(() =>
  selectedCity.value ? bundleStore.getDistricts(selectedCity.value) : [],
)
const filteredAddresses = computed(() =>
  selectedDistrict.value
    ? bundleStore.getAddresses(selectedCity.value, selectedDistrict.value)
    : [],
)
const canContinue = computed(
  () =>
    householdProfile.value.size &&
    householdProfile.value.activeLines &&
    selectedCity.value &&
    selectedDistrict.value &&
    selectedAddressId.value &&
    householdLines.value.length > 0,
)

// Methods
const updateHouseholdLines = () => {
  const targetLines = parseInt(householdProfile.value.activeLines) || 1
  const currentLines = householdLines.value.length

  if (targetLines > currentLines) {
    // Add lines
    for (let i = currentLines; i < targetLines; i++) {
      addLine()
    }
  } else if (targetLines < currentLines) {
    // Remove lines
    while (householdLines.value.length > targetLines) {
      removeLine(householdLines.value.length - 1)
    }
  }
}

const onCityChange = () => {
  selectedDistrict.value = ''
  selectedAddressId.value = ''
}

const onDistrictChange = () => {
  selectedAddressId.value = ''
}

const addLine = () => {
  const newId = Math.max(...householdLines.value.map((l) => l.id)) + 1
  householdLines.value.push({
    id: newId,
    gb: '10',
    minutes: '500',
    tvHours: '0',
  })
}

const removeLine = (index) => {
  if (householdLines.value.length > 1) {
    householdLines.value.splice(index, 1)
  }
}

const continueToRecommendations = async () => {
  try {
    // Get coverage data
    coverageData.value = await bundleStore.fetchCoverage(selectedAddressId.value)

    currentStep.value = 2
  } catch (error) {
    console.error('Error fetching coverage:', error)
  }
}

const calculateRecommendations = async () => {
  try {
    // Get recommendations
    recommendations.value = await bundleStore.fetchRecommendations(
      selectedAddressId.value,
      householdLines.value,
    )

    currentStep.value = 3
  } catch (error) {
    console.error('Error fetching recommendations:', error)
  }
}

const getTechIcon = (type) => {
  const icons = {
    fiber: '/icons/fiber.svg',
    vdsl: '/icons/vdsl.svg',
    fwa: '/icons/fwa.svg',
  }
  return icons[type] || '/icons/default.svg'
}

const getCoverageBadgeClass = (type) => {
  const classes = {
    fiber: 'badge-primary',
    vdsl: 'badge-secondary',
    fwa: 'badge-tertiary',
  }
  return classes[type] || 'badge-info'
}

const showDetails = (recommendation) => {
  selectedRecommendation.value = recommendation
  showModal.value = true
}

const selectRecommendation = async (recommendation) => {
  selectedRecommendation.value = recommendation

  try {
    // Get install slots
    installSlots.value = await bundleStore.fetchInstallSlots(
      selectedAddressId.value,
      recommendation.coverageType,
    )

    showSelectionModal.value = true
  } catch (error) {
    console.error('Error fetching install slots:', error)
  }
}

const confirmSelection = async (installSlot) => {
  try {
    const checkoutData = {
      user_id: 'mock_user_id',
      address_id: selectedAddressId.value,
      household: householdLines.value,
      recommendation_id: selectedRecommendation.value.id,
      install_slot: installSlot,
    }

    const response = await fetch('/api/checkout', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(checkoutData),
    })

    const result = await response.json()

    // Show success message
    alert(`Sipariş başarıyla oluşturuldu! Sipariş ID: ${result.order_id}`)

    // Reset wizard
    currentStep.value = 1
    showSelectionModal.value = false
    selectedRecommendation.value = null
  } catch (error) {
    console.error('Error during checkout:', error)
    alert('Sipariş oluşturulurken bir hata oluştu.')
  }
}

onMounted(() => {
  // Initialize store with coverage data
  bundleStore.initializeCoverage()
})
</script>

<template>
  <div class="bundle-wizard container">
    <div class="wizard-header">
      <h1>Telekom Paket Sihirbazı</h1>
      <p>İhtiyaçlarınıza uygun en iyi paketi bulun</p>
    </div>

    <!-- Step 1: Input Wizard -->
    <div class="wizard-step" v-if="currentStep === 1">
      <div class="step-content">
        <h2>Hane Profili ve Adres Bilgileri</h2>

        <!-- Household Profile -->
        <div class="form-section">
          <h3>Hane Profili</h3>
          <div class="grid grid-2">
            <div class="form-group">
              <label for="householdSize">Hane Üyeleri Sayısı *</label>
              <select
                id="householdSize"
                v-model="householdProfile.size"
                @change="updateHouseholdLines"
                required
                class="form-select"
              >
                <option value="">Seçiniz</option>
                <option value="1">1 Kişi</option>
                <option value="2">2 Kişi</option>
                <option value="3">3 Kişi</option>
                <option value="4">4 Kişi</option>
                <option value="5">5+ Kişi</option>
              </select>
            </div>

            <div class="form-group">
              <label for="activeLines">Aktif Mobil Hat Sayısı *</label>
              <select
                id="activeLines"
                v-model="householdProfile.activeLines"
                @change="updateHouseholdLines"
                required
                class="form-select"
              >
                <option value="">Seçiniz</option>
                <option value="1">1 Hat</option>
                <option value="2">2 Hat</option>
                <option value="3">3 Hat</option>
                <option value="4">4 Hat</option>
                <option value="5">5+ Hat</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Address Selection -->
        <div class="form-section">
          <h3>Adres Bilgileri</h3>
          <div class="grid grid-3">
            <div class="form-group">
              <label for="city">Şehir *</label>
              <select
                id="city"
                v-model="selectedCity"
                @change="onCityChange"
                required
                class="form-select"
              >
                <option value="">Şehir seçiniz</option>
                <option v-for="city in cities" :key="city" :value="city">
                  {{ city }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label for="district">İlçe *</label>
              <select
                id="district"
                v-model="selectedDistrict"
                @change="onDistrictChange"
                required
                class="form-select"
                :disabled="!selectedCity"
              >
                <option value="">İlçe seçiniz</option>
                <option v-for="district in filteredDistricts" :key="district" :value="district">
                  {{ district }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label for="address">Adres ID *</label>
              <select
                id="address"
                v-model="selectedAddressId"
                required
                class="form-select"
                :disabled="!selectedDistrict"
              >
                <option value="">Adres seçiniz</option>
                <option v-for="address in filteredAddresses" :key="address.id" :value="address.id">
                  {{ address.street }} {{ address.building }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <!-- Household Lines -->
        <div class="form-section">
          <h3>Hane Hatları Detayları</h3>
          <div class="household-lines">
            <div v-for="(line, index) in householdLines" :key="line.id" class="line-item card">
              <div class="line-header">
                <h4>Hat {{ index + 1 }}</h4>
                <button
                  v-if="householdLines.length > 1"
                  @click="removeLine(index)"
                  class="btn btn-danger"
                  type="button"
                >
                  Kaldır
                </button>
              </div>

              <div class="grid grid-3">
                <div class="form-group">
                  <label>GB İhtiyacı</label>
                  <select v-model="line.gb" class="form-select">
                    <option value="5">5 GB</option>
                    <option value="10">10 GB</option>
                    <option value="20">20 GB</option>
                    <option value="50">50 GB</option>
                    <option value="100">100 GB</option>
                    <option value="unlimited">Sınırsız</option>
                  </select>
                </div>

                <div class="form-group">
                  <label>Dakika İhtiyacı</label>
                  <select v-model="line.minutes" class="form-select">
                    <option value="100">100 dk</option>
                    <option value="500">500 dk</option>
                    <option value="1000">1000 dk</option>
                    <option value="unlimited">Sınırsız</option>
                  </select>
                </div>

                <div class="form-group">
                  <label>TV HD Saat</label>
                  <select v-model="line.tvHours" class="form-select">
                    <option value="0">TV Yok</option>
                    <option value="4">4 Saat</option>
                    <option value="8">8 Saat</option>
                    <option value="12">12 Saat</option>
                    <option value="24">24 Saat</option>
                  </select>
                </div>
              </div>
            </div>

            <button @click="addLine" class="btn btn-primary" type="button">+ Yeni Hat Ekle</button>
          </div>
        </div>

        <!-- Continue Button -->
        <div class="form-actions">
          <button
            @click="continueToRecommendations"
            class="btn btn-primary btn-lg"
            :disabled="!canContinue"
            type="button"
          >
            Devam Et
          </button>
        </div>
      </div>
    </div>

    <!-- Step 2: Coverage Badges -->
    <div class="wizard-step" v-if="currentStep === 2">
      <div class="step-content">
        <h2>Kapsama Durumu</h2>
        <p>Seçilen adres için mevcut teknolojiler:</p>

        <div class="coverage-badges grid grid-3">
          <div
            v-for="tech in coverageData"
            :key="tech.type"
            class="coverage-badge card"
            :class="tech.available ? 'available' : 'unavailable'"
          >
            <div class="badge-icon">
              <img :src="getTechIcon(tech.type)" :alt="tech.name" />
            </div>
            <div class="badge-content">
              <h4>{{ tech.name }}</h4>
              <p>{{ tech.available ? 'Mevcut' : 'Mevcut Değil' }}</p>
              <span v-if="tech.speed" class="speed">{{ tech.speed }}</span>
            </div>
          </div>
        </div>

        <div class="form-actions">
          <button @click="currentStep = 1" class="btn btn-secondary">Geri</button>
          <button @click="calculateRecommendations" class="btn btn-primary">
            Önerileri Hesapla
          </button>
        </div>
      </div>
    </div>

    <!-- Step 3: Recommendations -->
    <div class="wizard-step" v-if="currentStep === 3">
      <div class="step-content">
        <h2>Size Özel Öneriler</h2>
        <p>İhtiyaçlarınıza göre en uygun paketler:</p>

        <div class="recommendations grid grid-1">
          <div
            v-for="(recommendation, index) in recommendations"
            :key="recommendation.id"
            class="recommendation-card card"
            :class="{ alternative: recommendation.isAlternative }"
          >
            <div class="card-header">
              <h3>{{ recommendation.name }}</h3>
              <div class="card-badges">
                <span
                  v-if="recommendation.isAlternative"
                  class="badge badge-warning"
                  title="Bu özellik mevcut değil, alternatif önerildi"
                >
                  Alternatif Öneri
                </span>
                <span class="badge" :class="getCoverageBadgeClass(recommendation.coverageType)">
                  {{ recommendation.coverageType.toUpperCase() }}
                </span>
                <span v-if="recommendation.bundleDiscount" class="badge badge-success">
                  %{{ recommendation.bundleDiscount }} İndirim
                </span>
                <span v-if="recommendation.extraLineDiscount" class="badge badge-info">
                  Ek Hat -%{{ recommendation.extraLineDiscount }}
                </span>
              </div>
            </div>

            <div class="card-content">
              <div class="summary">
                <p>{{ recommendation.summary }}</p>
                <div class="price-info">
                  <span class="monthly-total">{{ recommendation.monthlyTotal }}₺</span>
                  <span class="savings">/ay</span>
                  <span v-if="recommendation.savings" class="savings-badge">
                    {{ recommendation.savings }}₺ tasarruf
                  </span>
                </div>
              </div>
            </div>

            <div class="card-actions">
              <button @click="showDetails(recommendation)" class="btn btn-secondary" type="button">
                Detayları Gör
              </button>
              <button
                @click="selectRecommendation(recommendation)"
                class="btn btn-primary"
                type="button"
              >
                Seç ve Devam Et
              </button>
            </div>
          </div>
        </div>

        <div class="form-actions">
          <button @click="currentStep = 2" class="btn btn-secondary">Geri</button>
        </div>
      </div>
    </div>

    <!-- Recommendation Details Modal -->
    <RecommendationModal
      v-if="showModal"
      :recommendation="selectedRecommendation"
      @close="showModal = false"
    />

    <!-- Selection Summary Modal -->
    <SelectionModal
      v-if="showSelectionModal"
      :recommendation="selectedRecommendation"
      :installSlots="installSlots"
      @close="showSelectionModal = false"
      @confirm="confirmSelection"
    />
  </div>
</template>

<style scoped>
.bundle-wizard {
  padding: var(--spacing-xl) 0;
}

.wizard-header {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
}

.wizard-header h1 {
  margin-bottom: var(--spacing-md);
}

.wizard-header p {
  font-size: 1.2rem;
}

.wizard-step {
  background: var(--color-background-soft);
  border-radius: var(--radius-xl);
  padding: var(--spacing-xl);
  margin-bottom: var(--spacing-xl);
  box-shadow: var(--shadow-md);
}

.step-content h2 {
  margin-bottom: var(--spacing-lg);
  text-align: center;
}

.form-section {
  margin-bottom: var(--spacing-xl);
}

.form-section h3 {
  margin-bottom: var(--spacing-lg);
  border-bottom: 2px solid var(--color-primary);
  padding-bottom: var(--spacing-sm);
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: var(--spacing-sm);
  font-weight: 600;
  color: var(--color-text);
}

.household-lines {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.line-item {
  border: 2px solid var(--color-border);
}

.line-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.line-header h4 {
  margin: 0;
  color: var(--color-primary);
}

.form-actions {
  display: flex;
  justify-content: center;
  gap: var(--spacing-md);
  margin-top: var(--spacing-xl);
}

.btn-lg {
  padding: var(--spacing-md) var(--spacing-xl);
  font-size: 1.1rem;
  min-height: 56px;
}

.coverage-badges {
  margin: var(--spacing-xl) 0;
}

.coverage-badge {
  text-align: center;
  border: 2px solid var(--color-border);
  transition: all var(--transition-normal);
}

.coverage-badge.available {
  border-color: var(--color-success);
  background: var(--color-success-light);
}

.coverage-badge.unavailable {
  border-color: var(--color-danger);
  background: var(--color-danger-light);
  opacity: 0.6;
}

.badge-icon img {
  width: 48px;
  height: 48px;
  margin-bottom: var(--spacing-md);
}

.badge-content h4 {
  margin-bottom: var(--spacing-sm);
}

.badge-content p {
  margin-bottom: var(--spacing-sm);
}

.speed {
  background: var(--color-primary);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: 0.8rem;
  font-weight: 600;
}

.recommendations {
  margin: var(--spacing-xl) 0;
}

.recommendation-card {
  border: 2px solid var(--color-border);
  transition: all var(--transition-normal);
}

.recommendation-card:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-lg);
  transform: translateY(-2px);
}

.recommendation-card.alternative {
  border-color: var(--color-warning);
  background: var(--color-warning-light);
}

.card-header {
  margin-bottom: var(--spacing-lg);
}

.card-header h3 {
  margin-bottom: var(--spacing-md);
}

.card-badges {
  display: flex;
  flex-wrap: wrap;
  gap: var(--spacing-sm);
}

.card-content {
  margin-bottom: var(--spacing-lg);
}

.summary p {
  margin-bottom: var(--spacing-md);
}

.price-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.monthly-total {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-primary);
}

.savings {
  color: var(--color-text-light);
  font-size: 1.2rem;
}

.savings-badge {
  background: var(--color-success);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 600;
}

.card-actions {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .wizard-step {
    padding: var(--spacing-lg);
    margin: var(--spacing-lg);
  }

  .grid-3 {
    grid-template-columns: 1fr;
  }

  .card-actions {
    flex-direction: column;
  }

  .btn {
    width: 100%;
  }

  .form-actions {
    flex-direction: column;
    align-items: center;
  }

  .btn-lg {
    width: 100%;
    max-width: 300px;
  }
}

@media (max-width: 480px) {
  .wizard-step {
    padding: var(--spacing-md);
    margin: var(--spacing-md);
  }

  .wizard-header {
    margin-bottom: var(--spacing-xl);
  }

  .coverage-badges {
    grid-template-columns: 1fr;
  }
}
</style>
