<template>
  <div class="coverage-badges">
    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center p-4">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="ml-2 text-gray-600">Kapsama bilgileri yükleniyor...</span>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4">
      <div class="flex items-center">
        <svg
          class="w-5 h-5 text-red-400 mr-2"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg>
        <span class="text-red-800">{{ error }}</span>
      </div>
    </div>

    <!-- Coverage Badges -->
    <div v-else-if="coverageData" class="space-y-3">
      <!-- Header -->
      <div class="text-center mb-4">
        <h5 class="text-primary mb-2">Adres Kapsama Durumu</h5>
        <p class="text-muted small mb-0">Seçilen adres için mevcut teknoloji seçenekleri</p>
      </div>

      <!-- Technology Badges -->
      <div class="row g-3">
        <!-- Fiber Badge -->
        <div class="col-md-4">
          <div class="text-center">
            <div class="coverage-badge" :class="getBadgeClass('fiber')">
              <i class="bi bi-lightning me-2"></i>
              {{ getBadgeText('fiber') }}
            </div>
          </div>
        </div>

        <!-- VDSL Badge -->
        <div class="col-md-4">
          <div class="text-center">
            <div class="coverage-badge" :class="getBadgeClass('vdsl')">
              <i class="bi bi-check-circle me-2"></i>
              {{ getBadgeText('vdsl') }}
            </div>
          </div>
        </div>

        <!-- FWA Badge -->
        <div class="col-md-4">
          <div class="text-center">
            <div class="coverage-badge" :class="getBadgeClass('fwa')">
              <i class="bi bi-wifi me-2"></i>
              {{ getBadgeText('fwa') }}
            </div>
          </div>
        </div>
      </div>

      <!-- Alternative Technology Notice -->
      <div
        v-if="showAlternativeNotice"
        class="mt-4 p-3 bg-warning-light border border-warning rounded"
      >
        <div class="d-flex align-items-center">
          <i class="bi bi-exclamation-triangle text-warning me-2"></i>
          <div>
            <h6 class="text-warning mb-1">Alternatif Teknoloji Önerisi</h6>
            <p class="text-warning small mb-0">
              Fiber internet mevcut değil, ancak {{ getAlternativeTechnologies() }} teknolojileri
              kullanılabilir.
            </p>
          </div>
        </div>
      </div>

      <!-- Coverage Summary -->
      <div class="mt-4 p-3 bg-light rounded">
        <h6 class="text-primary mb-2">Kapsama Özeti</h6>
        <div class="small text-muted">
          <p class="mb-1"><strong>Adres ID:</strong> {{ addressId }}</p>
          <p class="mb-1"><strong>Mevcut Teknolojiler:</strong> {{ getAvailableTechnologies() }}</p>
          <p class="mb-0"><strong>Önerilen Teknoloji:</strong> {{ getRecommendedTechnology() }}</p>
        </div>
      </div>
    </div>

    <!-- No Data State -->
    <div v-else class="text-center p-8">
      <svg
        class="w-16 h-16 text-gray-300 mx-auto mb-4"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
        ></path>
      </svg>
      <p class="text-gray-500">Kapsama bilgisi bulunamadı</p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { getCoverageByAddress } from '@/utils/api'

// Props
const props = defineProps({
  addressId: {
    type: String,
    required: true,
  },
})

// Reactive state
const loading = ref(true)
const error = ref(null)
const coverageData = ref(null)

// Computed properties
const showAlternativeNotice = computed(() => {
  if (!coverageData.value) return false
  return (
    coverageData.value.fiber === 0 &&
    (coverageData.value.vdsl === 1 || coverageData.value.fwa === 1)
  )
})

// Methods
const loadCoverageData = async () => {
  try {
    loading.value = true
    error.value = null

    // Backend'den kapsama verisini al
    const apiData = await getCoverageByAddress(props.addressId)
    // apiData: [{ address_id, tech: 'fiber'|'vdsl'|'fwa', down_mbps }]
    if (Array.isArray(apiData) && apiData.length > 0) {
      const flags = { fiber: 0, vdsl: 0, fwa: 0 }
      apiData.forEach((c) => {
        const t = String(c.tech || '').toLowerCase()
        if (t === 'fiber' || t === 'vdsl' || t === 'fwa') flags[t] = 1
      })
      coverageData.value = { address_id: props.addressId, ...flags }
    } else {
      error.value = `Adres ID "${props.addressId}" için kapsama bilgisi bulunamadı`
    }
  } catch (err) {
    error.value = 'Kapsama verileri yüklenirken hata oluştu: ' + (err?.message || err)
  } finally {
    loading.value = false
  }
}

const getBadgeClass = (technology) => {
  if (!coverageData.value) return 'badge-unavailable'

  const isAvailable = coverageData.value[technology] === 1

  switch (technology) {
    case 'fiber':
      return isAvailable ? 'badge-fiber' : 'badge-unavailable'
    case 'vdsl':
      return isAvailable ? 'badge-vdsl' : 'badge-unavailable'
    case 'fwa':
      return isAvailable ? 'badge-fwa' : 'badge-unavailable'
    default:
      return 'badge-unavailable'
  }
}

const getBadgeText = (technology) => {
  if (!coverageData.value) return 'Bilinmiyor'

  const isAvailable = coverageData.value[technology] === 1

  switch (technology) {
    case 'fiber':
      return isAvailable ? 'Fiber Uygun' : 'Fiber Yok'
    case 'vdsl':
      return isAvailable ? 'VDSL Uygun' : 'VDSL Yok'
    case 'fwa':
      return isAvailable ? 'FWA Uygun' : 'FWA Yok'
    default:
      return 'Bilinmiyor'
  }
}

const getAlternativeTechnologies = () => {
  if (!coverageData.value) return ''

  const alternatives = []
  if (coverageData.value.vdsl === 1) alternatives.push('VDSL')
  if (coverageData.value.fwa === 1) alternatives.push('FWA')

  return alternatives.join(' / ')
}

const getAvailableTechnologies = () => {
  if (!coverageData.value) return 'Bilinmiyor'

  const available = []
  if (coverageData.value.fiber === 1) available.push('Fiber')
  if (coverageData.value.vdsl === 1) available.push('VDSL')
  if (coverageData.value.fwa === 1) available.push('FWA')

  return available.length > 0 ? available.join(', ') : 'Hiçbiri'
}

const getRecommendedTechnology = () => {
  if (!coverageData.value) return 'Bilinmiyor'

  // Priority: Fiber > VDSL > FWA
  if (coverageData.value.fiber === 1) return 'Fiber (En Hızlı)'
  if (coverageData.value.vdsl === 1) return 'VDSL (Orta Hız)'
  if (coverageData.value.fwa === 1) return 'FWA (Sabit Kablosuz)'

  return 'Uygun Teknoloji Yok'
}

// Lifecycle
onMounted(() => {
  loadCoverageData()
})

// Watch for addressId changes
watch(
  () => props.addressId,
  (newAddressId) => {
    if (newAddressId) {
      loadCoverageData()
    }
  },
)
</script>

<style scoped>
.coverage-badges {
  color: var(--color-text);
}

.bg-warning-light {
  background-color: var(--color-warning-lightest) !important;
  border-color: var(--color-warning) !important;
}

.text-warning {
  color: var(--color-warning-dark) !important;
}

.bg-light {
  background-color: var(--color-background-soft) !important;
}

/* Remove old badge styles since they're now in main.css */
</style>
