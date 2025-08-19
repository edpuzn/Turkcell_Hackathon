<template>
  <div class="recommendation-cards">
    <!-- Header -->
    <div class="text-center mb-4">
      <h4 class="text-primary mb-2">Size Özel Paket Önerileri</h4>
      <p class="text-muted small mb-0">İhtiyaçlarınıza göre en uygun 3 kombinasyon</p>
    </div>

    <!-- Recommendation Cards -->
    <div class="row g-3">
      <div
        v-for="(recommendation, index) in recommendations"
        :key="index"
        class="col-lg-4 col-md-6"
      >
        <div class="card h-100 border-0 shadow-sm recommendation-card">
          <!-- Card Header -->
          <div class="card-header bg-primary-light text-primary py-3">
            <div class="d-flex justify-content-between align-items-center">
              <h6 class="mb-0 fw-bold">
                <i class="bi bi-star-fill me-2"></i>
                {{ recommendation.bundleName }}
              </h6>
              <span class="badge bg-primary">{{ index + 1 }}. Öneri</span>
            </div>
          </div>

          <!-- Card Body -->
          <div class="card-body p-3">
            <!-- Package Summary -->
            <div class="mb-3">
              <div class="d-flex align-items-center mb-2">
                <i class="bi bi-phone text-primary me-2"></i>
                <span class="small fw-bold">{{ recommendation.mobileSummary }}</span>
              </div>
              <div class="d-flex align-items-center mb-2">
                <i class="bi bi-wifi text-primary me-2"></i>
                <span class="small fw-bold">{{ recommendation.homeSummary }}</span>
              </div>
              <div v-if="recommendation.tvSummary" class="d-flex align-items-center">
                <i class="bi bi-tv text-primary me-2"></i>
                <span class="small fw-bold">{{ recommendation.tvSummary }}</span>
              </div>
            </div>

            <!-- Price Info -->
            <div class="border-top pt-3">
              <div class="d-flex justify-content-between align-items-center mb-2">
                <span class="text-muted small">Aylık Toplam:</span>
                <span class="h5 text-primary mb-0">{{ recommendation.totalPrice }} ₺</span>
              </div>
              <div class="d-flex justify-content-between align-items-center">
                <span class="text-muted small">Aylık Tasarruf:</span>
                <span class="text-success fw-bold">{{ recommendation.monthlySavings }} ₺</span>
              </div>
            </div>

            <!-- Discount Badge -->
            <div v-if="recommendation.discountPercent > 0" class="text-center mt-3">
              <span class="badge bg-success text-white">
                <i class="bi bi-percent me-1"></i>
                {{ recommendation.discountPercent }}% İndirim
              </span>
            </div>
          </div>

          <!-- Card Footer -->
          <div class="card-footer bg-light border-0 p-3">
            <div class="d-grid gap-2">
              <button @click="showDetails(recommendation)" class="btn btn-outline-primary btn-sm">
                <i class="bi bi-info-circle me-2"></i>
                Detayları Gör
              </button>
              <button @click="selectRecommendation(recommendation)" class="btn btn-primary btn-sm">
                <i class="bi bi-check-circle me-2"></i>
                Seç ve Devam Et
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- No Recommendations -->
    <div v-if="recommendations.length === 0" class="text-center py-5">
      <i class="bi bi-search text-muted" style="font-size: 3rem"></i>
      <p class="text-muted mt-3">Henüz öneri bulunamadı</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useMockDataStore } from '@/stores/mockData'

const props = defineProps({
  householdData: {
    type: Object,
    required: true,
  },
  addressId: {
    type: String,
    required: true,
  },
})

const emit = defineEmits(['select-recommendation', 'show-details'])

const mockDataStore = useMockDataStore()

// Generate recommendations based on household data and coverage
const recommendations = computed(() => {
  if (!props.householdData || !props.addressId) return []

  const coverage = mockDataStore.getCoverageByAddress(props.addressId)
  if (!coverage) return []

  const recommendations = []

  // Calculate total needs
  const totalGb = props.householdData.household.reduce((sum, line) => sum + line.expected_gb, 0)
  const totalMinutes = props.householdData.household.reduce(
    (sum, line) => sum + line.expected_min,
    0,
  )
  const totalTVHours = props.householdData.household.reduce(
    (sum, line) => sum + line.tv_hd_hours,
    0,
  )
  const lineCount = props.householdData.household.length

  // Recommendation 1: Basic Bundle (Mobile + Home)
  if (coverage.fiber || coverage.vdsl || coverage.fwa) {
    const mobilePlan =
      mockDataStore.getMobilePlansByQuota(totalGb, totalGb + 10)[0] || mockDataStore.mobilePlans[0]
    const homePlan = mockDataStore.getHomePlansByTech(
      coverage.fiber ? 'fiber' : coverage.vdsl ? 'vdsl' : 'fwa',
    )[0]

    if (mobilePlan && homePlan) {
      const basePrice = mobilePlan.monthly_price * lineCount + homePlan.monthly_price
      const discount = mockDataStore.calculateBundleDiscount('Mobil+Ev')
      const discountedPrice = basePrice * (1 - discount / 100)

      recommendations.push({
        bundleName: 'Mobil + Ev İnternet',
        mobileSummary: `${lineCount} Hat × ${mobilePlan.plan_name}`,
        homeSummary: `${homePlan.name} (${homePlan.down_mbps} Mbps)`,
        tvSummary: null,
        totalPrice: Math.round(discountedPrice),
        monthlySavings: Math.round(basePrice - discountedPrice),
        discountPercent: discount,
        type: 'mobile_home',
        mobilePlan,
        homePlan,
        tvPlan: null,
      })
    }
  }

  // Recommendation 2: Premium Bundle (Mobile + Home + TV)
  if (totalTVHours > 0 && (coverage.fiber || coverage.vdsl || coverage.fwa)) {
    const mobilePlan =
      mockDataStore.getMobilePlansByQuota(totalGb, totalGb + 20)[1] || mockDataStore.mobilePlans[1]
    const homePlan =
      mockDataStore.getHomePlansByTech(
        coverage.fiber ? 'fiber' : coverage.vdsl ? 'vdsl' : 'fwa',
      )[1] ||
      mockDataStore.getHomePlansByTech(coverage.fiber ? 'fiber' : coverage.vdsl ? 'vdsl' : 'fwa')[0]
    const tvPlan = mockDataStore.getTVPlansByHours(totalTVHours)[0]

    if (mobilePlan && homePlan && tvPlan) {
      const basePrice =
        mobilePlan.monthly_price * lineCount + homePlan.monthly_price + tvPlan.monthly_price
      const discount = mockDataStore.calculateBundleDiscount('Mobil+Ev+TV')
      const discountedPrice = basePrice * (1 - discount / 100)

      recommendations.push({
        bundleName: 'Mobil + Ev + TV Plus',
        mobileSummary: `${lineCount} Hat × ${mobilePlan.plan_name}`,
        homeSummary: `${homePlan.name} (${homePlan.down_mbps} Mbps)`,
        tvSummary: `${tvPlan.name} (${tvPlan.hd_hours_included} saat)`,
        totalPrice: Math.round(discountedPrice),
        monthlySavings: Math.round(basePrice - discountedPrice),
        discountPercent: discount,
        type: 'mobile_home_tv',
        mobilePlan,
        homePlan,
        tvPlan,
      })
    }
  }

  // Recommendation 3: Value Bundle (Optimized for cost)
  if (recommendations.length < 3) {
    const mobilePlan = mockDataStore.mobilePlans.find((p) => p.quota_gb >= totalGb * 0.8)
    const homePlan = mockDataStore
      .getHomePlansByTech(coverage.fiber ? 'fiber' : coverage.vdsl ? 'vdsl' : 'fwa')
      .find((p) => p.down_mbps >= 50)

    if (mobilePlan && homePlan) {
      const basePrice = mobilePlan.monthly_price * lineCount + homePlan.monthly_price
      const discount = mockDataStore.calculateBundleDiscount('Mobil+Ev')
      const discountedPrice = basePrice * (1 - discount / 100)

      recommendations.push({
        bundleName: 'Ekonomik Paket',
        mobileSummary: `${lineCount} Hat × ${mobilePlan.plan_name}`,
        homeSummary: `${homePlan.name} (${homePlan.down_mbps} Mbps)`,
        tvSummary: null,
        totalPrice: Math.round(discountedPrice),
        monthlySavings: Math.round(basePrice - discountedPrice),
        discountPercent: discount,
        type: 'economic',
        mobilePlan,
        homePlan,
        tvPlan: null,
      })
    }
  }

  // Sort by total price (lowest first)
  return recommendations.sort((a, b) => a.totalPrice - b.totalPrice)
})

const showDetails = (recommendation) => {
  emit('show-details', recommendation)
}

const selectRecommendation = (recommendation) => {
  emit('select-recommendation', recommendation)
}
</script>

<style scoped>
/* Component-specific styles - most styles are now in main.css */
.recommendation-cards {
  color: var(--color-text);
}

/* Custom hover effects for recommendation cards */
.recommendation-card:hover .card-header {
  background: linear-gradient(135deg, var(--color-primary-light), var(--color-primary)) !important;
  color: var(--color-text-inverse) !important;
}

/* Price highlighting */
.text-primary {
  color: var(--color-primary) !important;
}

.text-success {
  color: var(--color-success) !important;
}

.text-muted {
  color: var(--color-text-light) !important;
}

.bg-primary {
  background-color: var(--color-primary) !important;
}

.bg-success {
  background-color: var(--color-success) !important;
}

.bg-primary-light {
  background-color: var(--color-primary-lighter) !important;
}
</style>
