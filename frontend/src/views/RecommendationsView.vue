<template>
  <div class="min-h-screen bg-light py-4">
    <div class="container">
      <!-- Header -->
      <div class="text-center mb-5">
        <h1 class="text-primary mb-2">Turkcell Ev+Mobil Paket Danışmanı</h1>
        <p class="text-muted fs-5">Paket Önerileri</p>
      </div>

      <!-- Recommendations Content -->
      <div class="row justify-content-center">
        <div class="col-lg-10">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-lightbulb me-2"></i>
                Paket Önerileri
              </h4>
            </div>
            <div class="card-body py-4">
              <!-- Address Input for Direct Access -->
              <div class="row mb-4">
                <div class="col-md-6 mb-3">
                  <label for="city" class="form-label fw-bold">Şehir *</label>
                  <select
                    id="city"
                    v-model="selectedCity"
                    @change="onCityChange"
                    class="form-select form-select-sm"
                    required
                  >
                    <option value="">Şehir seçiniz</option>
                    <option v-for="city in cities" :key="city" :value="city">
                      {{ city }}
                    </option>
                  </select>
                </div>

                <div class="col-md-6 mb-3">
                  <label for="district" class="form-label fw-bold">İlçe *</label>
                  <select
                    id="district"
                    v-model="selectedDistrict"
                    @change="onDistrictChange"
                    class="form-select form-select-sm"
                    :disabled="!selectedCity"
                    required
                  >
                    <option value="">İlçe seçiniz</option>
                    <option
                      v-for="district in availableDistricts"
                      :key="district"
                      :value="district"
                    >
                      {{ district }}
                    </option>
                  </select>
                </div>
              </div>

              <!-- Household Inputs -->
              <div class="row mb-4">
                <div class="col-md-4 mb-3">
                  <label for="householdMembers" class="form-label fw-bold">
                    <i class="bi bi-people me-2"></i>
                    Hane Üyeleri Sayısı *
                  </label>
                  <input
                    id="householdMembers"
                    type="number"
                    v-model.number="householdMembers"
                    min="1"
                    max="10"
                    class="form-control form-control-sm"
                    placeholder="Örn: 4"
                    required
                  />
                </div>

                <div class="col-md-4 mb-3">
                  <label for="activeLines" class="form-label fw-bold">
                    <i class="bi bi-phone me-2"></i>
                    Aktif Mobil Hat Sayısı *
                  </label>
                  <input
                    id="activeLines"
                    type="number"
                    v-model.number="activeLines"
                    min="1"
                    max="5"
                    class="form-control form-control-sm"
                    placeholder="Örn: 3"
                    required
                  />
                </div>

                <div class="col-md-4 mb-3">
                  <label for="avgGB" class="form-label fw-bold">
                    <i class="bi bi-wifi me-2"></i>
                    Ortalama GB İhtiyacı *
                  </label>
                  <input
                    id="avgGB"
                    type="number"
                    v-model.number="avgGB"
                    min="1"
                    max="1000"
                    class="form-control form-control-sm"
                    placeholder="Örn: 50"
                    required
                  />
                </div>
              </div>

              <!-- Recommendations -->
              <div v-if="canShowRecommendations" class="text-center">
                <h5 class="mb-3">{{ selectedCity }} - {{ selectedDistrict }} için Öneriler</h5>
                <RecommendationCards
                  :householdData="{
                    city: selectedCity,
                    district: selectedDistrict,
                    household: generateMockHousehold(),
                  }"
                  :addressId="addressId"
                  @select-recommendation="handleSelectRecommendation"
                  @show-details="handleShowDetails"
                />
              </div>

              <!-- Navigation Buttons -->
              <div class="d-flex justify-content-between align-items-center mt-4 pt-3 border-top">
                <RouterLink to="/onboarding" class="btn btn-outline-secondary">
                  <i class="bi bi-house me-2"></i>
                  Ana Sayfa
                </RouterLink>

                <div class="d-flex gap-2">
                  <RouterLink to="/coverage" class="btn btn-outline-primary">
                    <i class="bi bi-wifi me-2"></i>
                    Kapsama Kontrolü
                  </RouterLink>

                  <RouterLink to="/wizard" class="btn btn-primary">
                    <i class="bi bi-arrow-right me-2"></i>
                    Sihirbazı Başlat
                  </RouterLink>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import RecommendationCards from '@/components/RecommendationCards.vue'

// Local state for direct access
const selectedCity = ref('')
const selectedDistrict = ref('')
const householdMembers = ref(1)
const activeLines = ref(1)
const avgGB = ref(50)

// Cities and Districts data (same as wizard store)
const cities = ref([
  'İstanbul',
  'Ankara',
  'İzmir',
  'Bursa',
  'Antalya',
  'Adana',
  'Konya',
  'Gaziantep',
  'Mersin',
  'Diyarbakır',
])

const districtsByCity = ref({
  İstanbul: ['Kadıköy', 'Beşiktaş', 'Şişli', 'Beyoğlu', 'Üsküdar', 'Maltepe', 'Ataşehir'],
  Ankara: ['Çankaya', 'Keçiören', 'Mamak', 'Yenimahalle', 'Etimesgut', 'Sincan'],
  İzmir: ['Konak', 'Bornova', 'Karşıyaka', 'Buca', 'Çiğli', 'Bayraklı'],
  Bursa: ['Nilüfer', 'Osmangazi', 'Yıldırım', 'Mudanya', 'Gürsu'],
  Antalya: ['Muratpaşa', 'Kepez', 'Döşemealtı', 'Aksu', 'Konyaaltı'],
  Adana: ['Seyhan', 'Çukurova', 'Sarıçam', 'Karaisalı'],
  Konya: ['Selçuklu', 'Meram', 'Karatay', 'Cumra', 'Ereğli'],
  Gaziantep: ['Şahinbey', 'Şehitkamil', 'Oğuzeli', 'Nizip'],
  Mersin: ['Yenişehir', 'Akdeniz', 'Toroslar', 'Mezitli'],
  Diyarbakır: ['Bağlar', 'Kayapınar', 'Sur', 'Yenişehir'],
})

// Computed
const availableDistricts = computed(() => {
  if (!selectedCity.value) return []
  return districtsByCity.value[selectedCity.value] || []
})

const addressId = computed(() => {
  if (!selectedCity.value || !selectedDistrict.value) return null
  const cityCode = selectedCity.value.substring(0, 3).toUpperCase()
  const districtCode = selectedDistrict.value.substring(0, 3).toUpperCase()
  return `${cityCode}_${districtCode}_001`
})

const canShowRecommendations = computed(() => {
  return (
    selectedCity.value &&
    selectedDistrict.value &&
    householdMembers.value > 0 &&
    activeLines.value > 0 &&
    avgGB.value > 0
  )
})

// Methods
const onCityChange = () => {
  selectedDistrict.value = '' // Reset district when city changes
}

const onDistrictChange = () => {
  // District selected, recommendations will be shown automatically
}

const generateMockHousehold = () => {
  const household = []
  for (let i = 0; i < activeLines.value; i++) {
    const lineNames = ['Ana Hat', 'İkinci Hat', 'Üçüncü Hat', 'Dördüncü Hat', 'Beşinci Hat']
    household.push({
      line_id: lineNames[i] || `${i + 1}. Hat`,
      expected_gb: Math.floor(avgGB.value / activeLines.value),
      expected_min: 500,
      tv_hd_hours: 4,
    })
  }
  return household
}

const handleSelectRecommendation = (recommendation) => {
  console.log('Selected recommendation:', recommendation)
  // TODO: Handle recommendation selection
}

const handleShowDetails = (recommendation) => {
  console.log('Show details for:', recommendation)
  // TODO: Show detailed modal
}
</script>

<style scoped>
.bg-light {
  background-color: var(--color-background-soft) !important;
}

.text-primary {
  color: var(--color-primary) !important;
}

.text-muted {
  color: var(--color-text-light) !important;
}
</style>
