<template>
  <div class="api-status">
    <div class="card">
      <div class="card-header">
        <h5 class="card-title mb-0">
          <i class="bi bi-activity"></i>
          API Durumu
        </h5>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-6">
            <div class="status-item">
              <span class="status-label">Backend:</span>
              <span 
                :class="['status-badge', backendStatus ? 'status-ok' : 'status-error']"
              >
                {{ backendStatus ? 'Çalışıyor' : 'Bağlantı Yok' }}
              </span>
            </div>
          </div>
          <div class="col-md-6">
            <div class="status-item">
              <span class="status-label">Katalog:</span>
              <span 
                :class="['status-badge', catalogStatus ? 'status-ok' : 'status-error']"
              >
                {{ catalogStatus ? 'Yüklendi' : 'Yüklenemedi' }}
              </span>
            </div>
          </div>
        </div>
        
        <div v-if="backendInfo" class="mt-3">
          <small class="text-muted">
            <strong>Backend Bilgileri:</strong><br>
            Servis: {{ backendInfo.service }}<br>
            Zaman: {{ formatTime(backendInfo.timestamp) }}
          </small>
        </div>
        
        <div v-if="catalogInfo" class="mt-3">
          <small class="text-muted">
            <strong>Katalog Özeti:</strong><br>
            Mobil Planlar: {{ catalogInfo.mobile_plans || 0 }}<br>
            Ev Planları: {{ catalogInfo.home_plans || 0 }}<br>
            TV Planları: {{ catalogInfo.tv_plans || 0 }}
          </small>
        </div>
        
        <div class="mt-3">
          <button 
            @click="checkStatus" 
            :disabled="loading"
            class="btn btn-primary btn-sm"
          >
            <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
            {{ loading ? 'Kontrol Ediliyor...' : 'Durumu Kontrol Et' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { healthCheck, getCatalogSummary } from '@/utils/api'

export default {
  name: 'ApiStatus',
  setup() {
    const backendStatus = ref(false)
    const catalogStatus = ref(false)
    const backendInfo = ref(null)
    const catalogInfo = ref(null)
    const loading = ref(false)

    const checkStatus = async () => {
      loading.value = true
      
      try {
        // Backend health check
        const health = await healthCheck()
        backendStatus.value = true
        backendInfo.value = health
        
        // Catalog summary
        const catalog = await getCatalogSummary()
        catalogStatus.value = true
        catalogInfo.value = catalog
        
      } catch (error) {
        console.error('API status check failed:', error)
        backendStatus.value = false
        catalogStatus.value = false
        backendInfo.value = null
        catalogInfo.value = null
      } finally {
        loading.value = false
      }
    }

    const formatTime = (timestamp) => {
      if (!timestamp) return 'N/A'
      return new Date(timestamp).toLocaleString('tr-TR')
    }

    onMounted(() => {
      checkStatus()
    })

    return {
      backendStatus,
      catalogStatus,
      backendInfo,
      catalogInfo,
      loading,
      checkStatus,
      formatTime
    }
  }
}
</script>

<style scoped>
.api-status {
  margin: 1rem 0;
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.status-label {
  font-weight: 500;
}

.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
  font-weight: 500;
}

.status-ok {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.status-error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.card-header {
  background-color: #f8f9fa;
  border-bottom: 1px solid #dee2e6;
}

.card-title {
  color: #495057;
}

.card-title i {
  margin-right: 0.5rem;
}
</style>
