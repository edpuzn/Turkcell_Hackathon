<template>
  <div class="min-h-screen bg-light py-4">
    <div class="container">
      <!-- Header -->
      <div class="text-center mb-5">
        <p class="text-muted fs-5">Adres Kapsama Durumu</p>
      </div>

      <!-- Coverage Content -->
      <div class="row justify-content-center">
        <div class="col-lg-8">
          <div class="card shadow-sm">
            <div class="card-header py-3">
              <h4 class="mb-0">
                <i class="bi bi-wifi me-2"></i>
                Adres Kapsama Durumu
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

              <!-- Coverage Badge -->
              <div v-if="addressId" class="text-center">
                <h5 class="mb-3">Seçilen Adres: {{ selectedCity }} - {{ selectedDistrict }}</h5>
                <CoverageBadge :addressId="addressId" />
              </div>

              <!-- Navigation Buttons -->
              <div class="d-flex justify-content-between align-items-center mt-4 pt-3 border-top">
                <RouterLink to="/onboarding" class="btn btn-outline-secondary">
                  <i class="bi bi-house me-2"></i>
                  Ana Sayfa
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
</template>

<script setup>
import { ref, computed } from 'vue'
import CoverageBadge from '@/components/CoverageBadge.vue'

// Local state for direct access
const selectedCity = ref('')
const selectedDistrict = ref('')

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

// Methods
const onCityChange = () => {
  selectedDistrict.value = '' // Reset district when city changes
}

const onDistrictChange = () => {
  // District selected, coverage will be shown automatically
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
