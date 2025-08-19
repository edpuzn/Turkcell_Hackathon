<template>
  <div class="modal-overlay" @click="$emit('close')">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>Seçim Özeti ve Kurulum Randevusu</h2>
        <button @click="$emit('close')" class="btn btn-close" aria-label="Kapat">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M18 6L6 18M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <!-- Selected Package Summary -->
        <div class="summary-section">
          <h3>Seçilen Paket</h3>
          <div class="package-summary card">
            <div class="summary-header">
              <h4>{{ recommendation.name }}</h4>
              <div class="package-badges">
                <span class="badge" :class="getCoverageBadgeClass(recommendation.coverageType)">
                  {{ recommendation.coverageType.toUpperCase() }}
                </span>
                <span v-if="recommendation.bundleDiscount" class="badge badge-success">
                  %{{ recommendation.bundleDiscount }} İndirim
                </span>
              </div>
            </div>

            <div class="summary-content">
              <p>{{ recommendation.summary }}</p>
              <div class="price-highlight">
                <span class="monthly-total">{{ recommendation.monthlyTotal }}₺</span>
                <span class="period">/ay</span>
                <span v-if="recommendation.savings" class="savings">
                  {{ recommendation.savings }}₺ tasarruf
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Package Breakdown -->
        <div class="breakdown-section">
          <h3>Paket Detayları</h3>
          <div class="breakdown-grid">
            <!-- Mobile Lines -->
            <div class="breakdown-item card">
              <h4>Mobil Hatlar</h4>
              <div class="item-details">
                <div
                  v-for="(plan, index) in recommendation.mobilePlans"
                  :key="plan.id"
                  class="detail-row"
                >
                  <span class="label">Hat {{ index + 1 }}:</span>
                  <span class="value">{{ plan.name }} - {{ plan.price }}₺</span>
                </div>
              </div>
            </div>

            <!-- Home Internet -->
            <div class="breakdown-item card">
              <h4>Ev İnternet</h4>
              <div class="item-details">
                <div class="detail-row">
                  <span class="label">Plan:</span>
                  <span class="value">{{ recommendation.homePlan.name }}</span>
                </div>
                <div class="detail-row">
                  <span class="label">Hız:</span>
                  <span class="value">{{ recommendation.homePlan.speed }} Mbps</span>
                </div>
                <div class="detail-row">
                  <span class="label">Fiyat:</span>
                  <span class="value">{{ recommendation.homePlan.price }}₺</span>
                </div>
              </div>
            </div>

            <!-- TV Package -->
            <div v-if="recommendation.tvPlan" class="breakdown-item card">
              <h4>TV Paketi</h4>
              <div class="item-details">
                <div class="detail-row">
                  <span class="label">Plan:</span>
                  <span class="value">{{ recommendation.tvPlan.name }}</span>
                </div>
                <div class="detail-row">
                  <span class="label">HD Saat:</span>
                  <span class="value">{{ recommendation.tvPlan.hours }} saat</span>
                </div>
                <div class="detail-row">
                  <span class="label">Fiyat:</span>
                  <span class="value">{{ recommendation.tvPlan.price }}₺</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Install Slot Selection -->
        <div class="install-section">
          <h3>Kurulum Randevusu</h3>
          <p>Lütfen uygun bir kurulum tarihi ve saati seçin:</p>

          <div class="slot-selection">
            <div class="date-groups">
              <div v-for="dateGroup in groupedSlots" :key="dateGroup.date" class="date-group">
                <h4 class="date-header">{{ formatDate(dateGroup.date) }}</h4>
                <div class="time-slots">
                  <button
                    v-for="slot in dateGroup.slots"
                    :key="slot.id"
                    @click="selectSlot(slot)"
                    class="time-slot"
                    :class="{
                      selected: selectedSlot.id === slot.id,
                      available: slot.available,
                    }"
                    :disabled="!slot.available"
                    type="button"
                  >
                    {{ slot.time }}
                  </button>
                </div>
              </div>
            </div>

            <div v-if="selectedSlot.date" class="selected-slot-info">
              <div class="selected-info card">
                <h4>Seçilen Randevu</h4>
                <div class="slot-details">
                  <div class="detail-row">
                    <span class="label">Tarih:</span>
                    <span class="value">{{ formatDate(selectedSlot.date) }}</span>
                  </div>
                  <div class="detail-row">
                    <span class="label">Saat:</span>
                    <span class="value">{{ selectedSlot.time }}</span>
                  </div>
                  <div class="detail-row">
                    <span class="label">Teknoloji:</span>
                    <span class="value">{{ selectedSlot.technology.toUpperCase() }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Additional Services -->
        <div class="services-section">
          <h3>Ek Hizmetler</h3>
          <div class="services-grid">
            <div class="service-item card">
              <div class="service-header">
                <h4>Hat Taşıma</h4>
                <input
                  type="checkbox"
                  id="lineTransfer"
                  v-model="additionalServices.lineTransfer"
                  class="service-checkbox"
                />
              </div>
              <p>Mevcut hatlarınızı yeni pakete taşıyın</p>
              <span class="service-price">Ücretsiz</span>
            </div>

            <div class="service-item card">
              <div class="service-header">
                <h4>Modem Teslimatı</h4>
                <input
                  type="checkbox"
                  id="modemDelivery"
                  v-model="additionalServices.modemDelivery"
                  class="service-checkbox"
                />
              </div>
              <p>Yeni modem cihazı teslim edilsin</p>
              <span class="service-price">Ücretsiz</span>
            </div>

            <div class="service-item card">
              <div class="service-header">
                <h4>Teknik Destek</h4>
                <input
                  type="checkbox"
                  id="techSupport"
                  v-model="additionalServices.techSupport"
                  class="service-checkbox"
                />
              </div>
              <p>7/24 teknik destek hizmeti</p>
              <span class="service-price">Ücretsiz</span>
            </div>
          </div>
        </div>

        <!-- Final Summary -->
        <div class="final-summary-section">
          <h3>Sipariş Özeti</h3>
          <div class="final-summary card">
            <div class="summary-rows">
              <div class="summary-row">
                <span class="label">Paket Tutarı:</span>
                <span class="value">{{ recommendation.monthlyTotal }}₺</span>
              </div>
              <div class="summary-row">
                <span class="label">Kurulum Ücreti:</span>
                <span class="value">0₺</span>
              </div>
              <div class="summary-row">
                <span class="label">Ek Hizmetler:</span>
                <span class="value">0₺</span>
              </div>
              <div class="summary-row total">
                <span class="label">İlk Ay Toplam:</span>
                <span class="value">{{ recommendation.monthlyTotal }}₺</span>
              </div>
            </div>

            <div class="summary-note">
              <p>
                💡 İlk ay kurulum ve ek hizmetler ücretsizdir. Sonraki aylarda sadece paket tutarı
                ödenecektir.
              </p>
            </div>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button @click="$emit('close')" class="btn btn-secondary">İptal</button>
        <button
          @click="confirmSelection"
          class="btn btn-primary btn-lg"
          :disabled="!selectedSlot.date"
          type="button"
        >
          Siparişi Onayla
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  recommendation: {
    type: Object,
    required: true,
  },
  installSlots: {
    type: Array,
    required: true,
  },
})

const emit = defineEmits(['close', 'confirm'])

// Selected slot
const selectedSlot = ref({})

// Additional services
const additionalServices = ref({
  lineTransfer: true,
  modemDelivery: true,
  techSupport: false,
})

// Computed properties
const groupedSlots = computed(() => {
  const groups = {}
  props.installSlots.forEach((slot) => {
    if (!groups[slot.date]) {
      groups[slot.date] = []
    }
    groups[slot.date].push(slot)
  })

  // Sort dates and times
  return Object.keys(groups)
    .sort()
    .reduce((result, date) => {
      result[date] = groups[date].sort((a, b) => a.time.localeCompare(b.time))
      return result
    }, {})
})

// Methods
const selectSlot = (slot) => {
  selectedSlot.value = slot
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  const options = {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }
  return date.toLocaleDateString('tr-TR', options)
}

const getCoverageBadgeClass = (type) => {
  const classes = {
    fiber: 'badge-primary',
    vdsl: 'badge-secondary',
    fwa: 'badge-tertiary',
  }
  return classes[type] || 'badge-info'
}

const confirmSelection = () => {
  if (!selectedSlot.value.date) {
    alert('Lütfen bir kurulum randevusu seçin')
    return
  }

  const checkoutData = {
    recommendation: props.recommendation,
    installSlot: selectedSlot.value,
    additionalServices: additionalServices.value,
  }

  emit('confirm', checkoutData)
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: var(--spacing-md);
}

.modal-content {
  background: var(--color-background);
  border-radius: var(--radius-lg);
  max-width: 900px;
  width: 100%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: var(--shadow-lg);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-lg);
  border-bottom: 1px solid var(--color-border);
  position: sticky;
  top: 0;
  background: var(--color-background);
  z-index: 10;
}

.modal-header h2 {
  margin: 0;
  color: var(--color-primary);
}

.btn-close {
  background: none;
  border: none;
  color: var(--color-text-light);
  cursor: pointer;
  padding: var(--spacing-sm);
  border-radius: var(--radius-sm);
  transition: all var(--transition-fast);
  min-width: auto;
  min-height: auto;
}

.btn-close:hover {
  background: var(--color-background-mute);
  color: var(--color-text);
}

.btn-close svg {
  width: 20px;
  height: 20px;
}

.modal-body {
  padding: var(--spacing-lg);
}

.modal-footer {
  padding: var(--spacing-lg);
  border-top: 1px solid var(--color-border);
  display: flex;
  justify-content: space-between;
  gap: var(--spacing-md);
}

/* Sections */
.summary-section,
.breakdown-section,
.install-section,
.services-section,
.final-summary-section {
  margin-bottom: var(--spacing-xl);
}

.summary-section h3,
.breakdown-section h3,
.install-section h3,
.services-section h3,
.final-summary-section h3 {
  margin-bottom: var(--spacing-lg);
  color: var(--color-primary);
  border-bottom: 2px solid var(--color-primary);
  padding-bottom: var(--spacing-sm);
}

/* Package Summary */
.package-summary {
  border: 2px solid var(--color-primary);
}

.summary-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--spacing-md);
  flex-wrap: wrap;
  gap: var(--spacing-md);
}

.summary-header h4 {
  margin: 0;
  color: var(--color-primary);
  font-size: 1.3rem;
}

.package-badges {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
}

.summary-content p {
  margin-bottom: var(--spacing-md);
  font-size: 1.1rem;
}

.price-highlight {
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

.period {
  color: var(--color-text-light);
  font-size: 1.2rem;
}

.savings {
  background: var(--color-success);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 600;
}

/* Breakdown Grid */
.breakdown-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-lg);
}

.breakdown-item {
  border: 1px solid var(--color-border);
}

.breakdown-item h4 {
  margin: 0 0 var(--spacing-md) 0;
  color: var(--color-primary);
  border-bottom: 1px solid var(--color-border);
  padding-bottom: var(--spacing-sm);
}

.item-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.detail-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.detail-row .label {
  font-weight: 600;
  color: var(--color-text-light);
}

.detail-row .value {
  font-weight: 600;
  color: var(--color-text);
}

/* Install Section */
.install-section p {
  margin-bottom: var(--spacing-lg);
  color: var(--color-text-light);
}

.slot-selection {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: var(--spacing-xl);
}

.date-groups {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.date-group {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  overflow: hidden;
}

.date-header {
  background: var(--color-primary);
  color: white;
  margin: 0;
  padding: var(--spacing-md);
  font-size: 1.1rem;
}

.time-slots {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(80px, 1fr));
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
}

.time-slot {
  padding: var(--spacing-sm);
  border: 2px solid var(--color-border);
  border-radius: var(--radius-sm);
  background: var(--color-background);
  color: var(--color-text);
  cursor: pointer;
  transition: all var(--transition-fast);
  font-weight: 600;
  min-height: 40px;
}

.time-slot:hover:not(:disabled) {
  border-color: var(--color-primary);
  background: var(--color-primary-light);
}

.time-slot.selected {
  border-color: var(--color-primary);
  background: var(--color-primary);
  color: white;
}

.time-slot:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.selected-slot-info {
  position: sticky;
  top: 20px;
}

.selected-info h4 {
  margin: 0 0 var(--spacing-md) 0;
  color: var(--color-primary);
}

.slot-details {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

/* Services Grid */
.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-lg);
}

.service-item {
  border: 1px solid var(--color-border);
  transition: all var(--transition-normal);
}

.service-item:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-md);
}

.service-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
}

.service-header h4 {
  margin: 0;
  color: var(--color-primary);
}

.service-checkbox {
  width: 20px;
  height: 20px;
  accent-color: var(--color-primary);
}

.service-item p {
  margin: 0 0 var(--spacing-sm) 0;
  color: var(--color-text-light);
  font-size: 0.9rem;
}

.service-price {
  color: var(--color-success);
  font-weight: 700;
  font-size: 1.1rem;
}

/* Final Summary */
.final-summary {
  background: var(--color-background-soft);
  border: 2px solid var(--color-primary);
}

.summary-rows {
  margin-bottom: var(--spacing-lg);
}

.summary-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) 0;
  border-bottom: 1px solid var(--color-border-light);
}

.summary-row:last-child {
  border-bottom: none;
}

.summary-row.total {
  border-top: 2px solid var(--color-primary);
  border-bottom: 2px solid var(--color-primary);
  padding: var(--spacing-md) 0;
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--color-primary);
}

.summary-row .label {
  font-weight: 600;
}

.summary-row .value {
  font-weight: 700;
}

.summary-note {
  background: var(--color-info-light);
  border: 1px solid var(--color-info);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
}

.summary-note p {
  margin: 0;
  color: var(--color-text);
  font-size: 0.9rem;
  line-height: 1.5;
}

/* Responsive */
@media (max-width: 768px) {
  .modal-content {
    margin: var(--spacing-sm);
    max-height: 95vh;
  }

  .modal-header,
  .modal-body,
  .modal-footer {
    padding: var(--spacing-md);
  }

  .slot-selection {
    grid-template-columns: 1fr;
    gap: var(--spacing-lg);
  }

  .selected-slot-info {
    position: static;
  }

  .breakdown-grid {
    grid-template-columns: 1fr;
  }

  .services-grid {
    grid-template-columns: 1fr;
  }

  .summary-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .modal-footer {
    flex-direction: column;
  }

  .btn {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .modal-overlay {
    padding: var(--spacing-sm);
  }

  .modal-content {
    margin: 0;
  }

  .modal-header h2 {
    font-size: 1.25rem;
  }

  .time-slots {
    grid-template-columns: repeat(2, 1fr);
  }

  .monthly-total {
    font-size: 1.5rem;
  }
}
</style>
