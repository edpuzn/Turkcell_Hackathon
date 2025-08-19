<template>
  <div class="api-test-view">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <h1 class="mb-4">
            <i class="bi bi-gear-wide-connected"></i>
            API Test Sayfası
          </h1>
          <p class="lead">
            Bu sayfa backend API'sinin frontend ile entegrasyonunu test eder.
          </p>
        </div>
      </div>

      <!-- API Status Component -->
      <div class="row mb-4">
        <div class="col-12">
          <ApiStatus />
        </div>
      </div>

      <!-- API Test Sections -->
      <div class="row">
        <!-- Health Check -->
        <div class="col-md-6 mb-4">
          <div class="card">
            <div class="card-header">
              <h5 class="card-title mb-0">
                <i class="bi bi-heart-pulse"></i>
                Health Check
              </h5>
            </div>
            <div class="card-body">
              <button 
                @click="testHealth" 
                :disabled="healthLoading"
                class="btn btn-primary"
              >
                <span v-if="healthLoading" class="spinner-border spinner-border-sm me-2"></span>
                {{ healthLoading ? 'Kontrol Ediliyor...' : 'Health Check' }}
              </button>
              
              <div v-if="healthResult" class="mt-3">
                <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(healthResult, null, 2) }}</code></pre>
              </div>
            </div>
          </div>
        </div>

        <!-- Catalog Summary -->
        <div class="col-md-6 mb-4">
          <div class="card">
            <div class="card-header">
              <h5 class="card-title mb-0">
                <i class="bi bi-collection"></i>
                Katalog Özeti
              </h5>
            </div>
            <div class="card-body">
              <button 
                @click="testCatalog" 
                :disabled="catalogLoading"
                class="btn btn-success"
              >
                <span v-if="catalogLoading" class="spinner-border spinner-border-sm me-2"></span>
                {{ catalogLoading ? 'Yükleniyor...' : 'Katalog Getir' }}
              </button>
              
              <div v-if="catalogResult" class="mt-3">
                <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(catalogResult, null, 2) }}</code></pre>
              </div>
            </div>
          </div>
        </div>

        <!-- Coverage Test -->
        <div class="col-md-6 mb-4">
          <div class="card">
            <div class="card-header">
              <h5 class="card-title mb-0">
                <i class="bi bi-geo-alt"></i>
                Kapsama Testi
              </h5>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <label for="addressId" class="form-label">Adres ID:</label>
                <input 
                  v-model="addressId" 
                  type="text" 
                  class="form-control" 
                  id="addressId"
                  placeholder="Örnek: 12345"
                >
              </div>
              
              <button 
                @click="testCoverage" 
                :disabled="coverageLoading || !addressId"
                class="btn btn-info"
              >
                <span v-if="coverageLoading" class="spinner-border spinner-border-sm me-2"></span>
                {{ coverageLoading ? 'Kontrol Ediliyor...' : 'Kapsama Kontrol Et' }}
              </button>
              
              <div v-if="coverageResult" class="mt-3">
                <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(coverageResult, null, 2) }}</code></pre>
              </div>
            </div>
          </div>
        </div>

        <!-- Recommendation Test -->
        <div class="col-md-6 mb-4">
          <div class="card">
            <div class="card-header">
              <h5 class="card-title mb-0">
                <i class="bi bi-lightbulb"></i>
                Öneri Testi
              </h5>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <label for="recommendationData" class="form-label">Öneri Verileri:</label>
                <textarea 
                  v-model="recommendationData" 
                  class="form-control" 
                  id="recommendationData"
                  rows="3"
                  placeholder='{"user_id": "123", "preferences": {"internet": true, "mobile": true}}'
                ></textarea>
              </div>
              
              <button 
                @click="testRecommendation" 
                :disabled="recommendationLoading || !recommendationData"
                class="btn btn-warning"
              >
                <span v-if="recommendationLoading" class="spinner-border spinner-border-sm me-2"></span>
                {{ recommendationLoading ? 'İşleniyor...' : 'Öneri Al' }}
              </button>
              
              <div v-if="recommendationResult" class="mt-3">
                <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(recommendationResult, null, 2) }}</code></pre>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Error Display -->
      <div v-if="error" class="row">
        <div class="col-12">
          <div class="alert alert-danger" role="alert">
            <h6 class="alert-heading">
              <i class="bi bi-exclamation-triangle"></i>
              Hata
            </h6>
            <p class="mb-0">{{ error }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import ApiStatus from '@/components/ApiStatus.vue'
import { 
  healthCheck, 
  getCatalogSummary, 
  getCoverageByAddress, 
  postRecommendation 
} from '@/utils/api'

export default {
  name: 'ApiTestView',
  components: {
    ApiStatus
  },
  setup() {
    const healthLoading = ref(false)
    const catalogLoading = ref(false)
    const coverageLoading = ref(false)
    const recommendationLoading = ref(false)
    
    const healthResult = ref(null)
    const catalogResult = ref(null)
    const coverageResult = ref(null)
    const recommendationResult = ref(null)
    
    const addressId = ref('12345')
    const recommendationData = ref('{"user_id": "123", "preferences": {"internet": true, "mobile": true}}')
    const error = ref('')

    const clearError = () => {
      error.value = ''
    }

    const testHealth = async () => {
      clearError()
      healthLoading.value = true
      healthResult.value = null
      
      try {
        const result = await healthCheck()
        healthResult.value = result
      } catch (err) {
        error.value = `Health check hatası: ${err.message}`
      } finally {
        healthLoading.value = false
      }
    }

    const testCatalog = async () => {
      clearError()
      catalogLoading.value = true
      catalogResult.value = null
      
      try {
        const result = await getCatalogSummary()
        catalogResult.value = result
      } catch (err) {
        error.value = `Katalog hatası: ${err.message}`
      } finally {
        catalogLoading.value = false
      }
    }

    const testCoverage = async () => {
      clearError()
      coverageLoading.value = true
      coverageResult.value = null
      
      try {
        const result = await getCoverageByAddress(addressId.value)
        coverageResult.value = result
      } catch (err) {
        error.value = `Kapsama hatası: ${err.message}`
      } finally {
        coverageLoading.value = false
      }
    }

    const testRecommendation = async () => {
      clearError()
      recommendationLoading.value = true
      recommendationResult.value = null
      
      try {
        const data = JSON.parse(recommendationData.value)
        const result = await postRecommendation(data)
        recommendationResult.value = result
      } catch (err) {
        if (err instanceof SyntaxError) {
          error.value = 'JSON formatı geçersiz'
        } else {
          error.value = `Öneri hatası: ${err.message}`
        }
      } finally {
        recommendationLoading.value = false
      }
    }

    return {
      healthLoading,
      catalogLoading,
      coverageLoading,
      recommendationLoading,
      healthResult,
      catalogResult,
      coverageResult,
      recommendationResult,
      addressId,
      recommendationData,
      error,
      testHealth,
      testCatalog,
      testCoverage,
      testRecommendation
    }
  }
}
</script>

<style scoped>
.api-test-view {
  padding: 2rem 0;
}

.card-header {
  background-color: #f8f9fa;
  border-bottom: 1px solid #dee2e6;
}

.card-title {
  color: #495057;
  font-weight: 600;
}

.card-title i {
  margin-right: 0.5rem;
}

pre {
  font-size: 0.875rem;
  max-height: 300px;
  overflow-y: auto;
}

.form-label {
  font-weight: 500;
  color: #495057;
}

.btn {
  font-weight: 500;
}

.btn i {
  margin-right: 0.5rem;
}
</style>
