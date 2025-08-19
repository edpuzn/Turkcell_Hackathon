<template>
  <div class="modal-overlay" @click="$emit('close')">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>{{ recommendation.name }}</h2>
        <button @click="$emit('close')" class="btn btn-close" aria-label="Kapat">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M18 6L6 18M6 6l12 12" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <!-- Package Summary -->
        <div class="summary-section">
          <h3>Paket Özeti</h3>
          <div class="summary-grid">
            <div class="summary-item">
              <span class="label">Teknoloji:</span>
              <span class="value">{{ recommendation.coverageType.toUpperCase() }}</span>
            </div>
            <div class="summary-item">
              <span class="label">Toplam Hat:</span>
              <span class="value">{{ recommendation.mobilePlans.length }} hat</span>
            </div>
            <div class="summary-item">
              <span class="label">Ev İnternet:</span>
              <span class="value">{{ recommendation.homePlan.speed }} Mbps</span>
            </div>
            <div class="summary-item" v-if="recommendation.tvPlan">
              <span class="label">TV Paketi:</span>
              <span class="value">{{ recommendation.tvPlan.hours }} saat</span>
            </div>
          </div>
        </div>

        <!-- Mobile Plans -->
        <div class="plans-section">
          <h3>Mobil Hatlar</h3>
          <div class="mobile-plans">
            <div
              v-for="(plan, index) in recommendation.mobilePlans"
              :key="plan.id"
              class="mobile-plan card"
            >
              <div class="plan-header">
                <h4>Hat {{ index + 1 }}</h4>
                <span class="plan-name">{{ plan.name }}</span>
              </div>

              <div class="plan-details grid grid-2">
                <div class="detail-item">
                  <span class="label">GB:</span>
                  <span class="value">{{
                    plan.gb === 'unlimited' ? 'Sınırsız' : plan.gb + ' GB'
                  }}</span>
                </div>
                <div class="detail-item">
                  <span class="label">Dakika:</span>
                  <span class="value">{{
                    plan.minutes === 'unlimited' ? 'Sınırsız' : plan.minutes + ' dk'
                  }}</span>
                </div>
                <div class="detail-item">
                  <span class="label">İhtiyaç GB:</span>
                  <span class="value">{{ plan.requiredGB }} GB</span>
                </div>
                <div class="detail-item">
                  <span class="label">İhtiyaç Dk:</span>
                  <span class="value">{{ plan.requiredMinutes }} dk</span>
                </div>
              </div>

              <div class="plan-price">
                <span class="price">{{ plan.price }}₺</span>
                <span class="period">/ay</span>
              </div>

              <!-- Overage warnings -->
              <div v-if="plan.overageGB > 0 || plan.overageMinutes > 0" class="overage-warning">
                <span class="warning-icon">⚠️</span>
                <span class="warning-text">
                  {{ plan.overageGB > 0 ? `${plan.overageGB} GB aşım` : '' }}
                  {{ plan.overageGB > 0 && plan.overageMinutes > 0 ? ' + ' : '' }}
                  {{ plan.overageMinutes > 0 ? `${plan.overageMinutes} dk aşım` : '' }}
                  ek ücretlendirilecektir
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Home Internet -->
        <div class="plans-section">
          <h3>Ev İnternet</h3>
          <div class="home-plan card">
            <div class="plan-header">
              <h4>{{ recommendation.homePlan.name }}</h4>
              <span class="technology-badge">{{ recommendation.coverageType.toUpperCase() }}</span>
            </div>

            <div class="plan-details grid grid-2">
              <div class="detail-item">
                <span class="label">Hız:</span>
                <span class="value">{{ recommendation.homePlan.speed }} Mbps</span>
              </div>
              <div class="detail-item">
                <span class="label">Teknoloji:</span>
                <span class="value">{{
                  recommendation.coverageType === 'fiber'
                    ? 'Fiber Optik'
                    : recommendation.coverageType === 'vdsl'
                      ? 'VDSL'
                      : '4.5G Sabit İnternet'
                }}</span>
              </div>
            </div>

            <div class="plan-price">
              <span class="price">{{ recommendation.homePlan.price }}₺</span>
              <span class="period">/ay</span>
            </div>
          </div>
        </div>

        <!-- TV Package -->
        <div v-if="recommendation.tvPlan" class="plans-section">
          <h3>TV Paketi</h3>
          <div class="tv-plan card">
            <div class="plan-header">
              <h4>{{ recommendation.tvPlan.name }}</h4>
            </div>

            <div class="plan-details grid grid-2">
              <div class="detail-item">
                <span class="label">HD Saat:</span>
                <span class="value">{{ recommendation.tvPlan.hours }} saat</span>
              </div>
              <div class="detail-item">
                <span class="label">Paket:</span>
                <span class="value">TV Plus</span>
              </div>
            </div>

            <div class="plan-price">
              <span class="price">{{ recommendation.tvPlan.price }}₺</span>
              <span class="period">/ay</span>
            </div>
          </div>
        </div>

        <!-- Discounts Applied -->
        <div class="discounts-section">
          <h3>Uygulanan İndirimler</h3>
          <div class="discounts-list">
            <div v-if="recommendation.bundleDiscount" class="discount-item">
              <span class="discount-label">Paket İndirimi:</span>
              <span class="discount-value">%{{ recommendation.bundleDiscount }}</span>
              <span class="discount-desc">
                {{ recommendation.bundleDiscount === 15 ? 'Mobil + Ev + TV' : 'Mobil + Ev' }} paketi
              </span>
            </div>

            <div v-if="recommendation.extraLineDiscount" class="discount-item">
              <span class="discount-label">Ek Hat İndirimi:</span>
              <span class="discount-value">%{{ recommendation.extraLineDiscount }}</span>
              <span class="discount-desc">
                {{ recommendation.extraLineDiscount === 5 ? '2. hat' : '3. ve üzeri hat' }} için
              </span>
            </div>
          </div>
        </div>

        <!-- Price Breakdown -->
        <div class="price-breakdown-section">
          <h3>Fiyat Dökümü</h3>
          <div class="price-breakdown card">
            <div class="breakdown-row">
              <span class="label">Mobil Hatlar:</span>
              <span class="value">{{ getMobileTotal() }}₺</span>
            </div>

            <div class="breakdown-row">
              <span class="label">Ev İnternet:</span>
              <span class="value">{{ recommendation.homePlan.price }}₺</span>
            </div>

            <div v-if="recommendation.tvPlan" class="breakdown-row">
              <span class="label">TV Paketi:</span>
              <span class="value">{{ recommendation.tvPlan.price }}₺</span>
            </div>

            <div class="breakdown-row subtotal">
              <span class="label">Ara Toplam:</span>
              <span class="value">{{ recommendation.totalOriginal }}₺</span>
            </div>

            <div class="breakdown-row discount">
              <span class="label">Toplam İndirim:</span>
              <span class="value">-{{ recommendation.savings }}₺</span>
            </div>

            <div class="breakdown-row total">
              <span class="label">Aylık Toplam:</span>
              <span class="value">{{ recommendation.monthlyTotal }}₺</span>
            </div>
          </div>
        </div>

        <!-- Alternative Technology Notice -->
        <div v-if="recommendation.isAlternative" class="alternative-notice">
          <div class="notice-icon">ℹ️</div>
          <div class="notice-content">
            <h4>Alternatif Teknoloji Önerisi</h4>
            <p>
              İstediğiniz
              {{
                recommendation.coverageType === 'fiber'
                  ? 'Fiber Optik'
                  : recommendation.coverageType === 'vdsl'
                    ? 'VDSL'
                    : '4.5G Sabit İnternet'
              }}
              teknolojisi bu adreste mevcut değil. En iyi alternatif teknoloji önerilmiştir.
            </p>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button @click="$emit('close')" class="btn btn-secondary">Kapat</button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  recommendation: {
    type: Object,
    required: true,
  },
})

defineEmits(['close'])

const getMobileTotal = () => {
  return this.recommendation.mobilePlans.reduce((sum, plan) => sum + plan.price, 0)
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
  max-width: 800px;
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
  justify-content: flex-end;
  gap: var(--spacing-md);
}

/* Sections */
.summary-section,
.plans-section,
.discounts-section,
.price-breakdown-section {
  margin-bottom: var(--spacing-xl);
}

.summary-section h3,
.plans-section h3,
.discounts-section h3,
.price-breakdown-section h3 {
  margin-bottom: var(--spacing-lg);
  color: var(--color-primary);
  border-bottom: 2px solid var(--color-primary);
  padding-bottom: var(--spacing-sm);
}

/* Summary Grid */
.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  background: var(--color-background-soft);
  border-radius: var(--radius-md);
}

.summary-item .label {
  font-weight: 600;
  color: var(--color-text-light);
}

.summary-item .value {
  font-weight: 700;
  color: var(--color-primary);
}

/* Mobile Plans */
.mobile-plans {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.mobile-plan {
  border: 2px solid var(--color-border);
  transition: all var(--transition-normal);
}

.mobile-plan:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-md);
}

.plan-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-md);
}

.plan-header h4 {
  margin: 0;
  color: var(--color-primary);
}

.plan-name {
  background: var(--color-primary);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: 0.9rem;
  font-weight: 600;
}

.plan-details {
  margin-bottom: var(--spacing-md);
}

.detail-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) 0;
}

.detail-item .label {
  font-weight: 600;
  color: var(--color-text-light);
}

.detail-item .value {
  font-weight: 600;
  color: var(--color-text);
}

.plan-price {
  text-align: right;
  margin-bottom: var(--spacing-md);
}

.plan-price .price {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-primary);
}

.plan-price .period {
  color: var(--color-text-light);
  font-size: 1rem;
}

/* Technology Badge */
.technology-badge {
  background: var(--color-secondary);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
}

/* Overage Warning */
.overage-warning {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm);
  background: var(--color-warning-light);
  border: 1px solid var(--color-warning);
  border-radius: var(--radius-sm);
  margin-top: var(--spacing-sm);
}

.warning-icon {
  font-size: 1.2rem;
}

.warning-text {
  color: var(--color-warning-dark);
  font-size: 0.9rem;
  font-weight: 600;
}

/* Discounts */
.discounts-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.discount-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  background: var(--color-success-light);
  border: 1px solid var(--color-success);
  border-radius: var(--radius-md);
}

.discount-label {
  font-weight: 600;
  color: var(--color-text);
  min-width: 120px;
}

.discount-value {
  background: var(--color-success);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-weight: 700;
  min-width: 60px;
  text-align: center;
}

.discount-desc {
  color: var(--color-text-light);
  font-size: 0.9rem;
}

/* Price Breakdown */
.price-breakdown {
  background: var(--color-background-soft);
}

.breakdown-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) 0;
  border-bottom: 1px solid var(--color-border-light);
}

.breakdown-row:last-child {
  border-bottom: none;
}

.breakdown-row.subtotal {
  border-top: 2px solid var(--color-border);
  border-bottom: 2px solid var(--color-border);
  padding: var(--spacing-md) 0;
  font-weight: 600;
}

.breakdown-row.discount {
  color: var(--color-success);
  font-weight: 600;
}

.breakdown-row.total {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--color-primary);
  border-top: 2px solid var(--color-primary);
  padding: var(--spacing-md) 0;
}

.breakdown-row .label {
  font-weight: 600;
}

.breakdown-row .value {
  font-weight: 700;
}

/* Alternative Notice */
.alternative-notice {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background: var(--color-info-light);
  border: 1px solid var(--color-info);
  border-radius: var(--radius-md);
  margin-top: var(--spacing-lg);
}

.notice-icon {
  font-size: 2rem;
  flex-shrink: 0;
}

.notice-content h4 {
  margin: 0 0 var(--spacing-sm) 0;
  color: var(--color-info-dark);
}

.notice-content p {
  margin: 0;
  color: var(--color-text);
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

  .summary-grid {
    grid-template-columns: 1fr;
  }

  .plan-details {
    grid-template-columns: 1fr;
  }

  .plan-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-sm);
  }

  .discount-item {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-sm);
  }

  .discount-label {
    min-width: auto;
  }

  .alternative-notice {
    flex-direction: column;
    text-align: center;
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

  .plan-price .price {
    font-size: 1.25rem;
  }
}
</style>
