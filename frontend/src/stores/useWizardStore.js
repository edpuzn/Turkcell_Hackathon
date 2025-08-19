import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getCoverageByAddress, postRecommendation } from '@/utils/api'

export const useWizardStore = defineStore('wizard', () => {
  // State
  const city = ref('')
  const district = ref('')
  const householdMembers = ref(1)
  const activeLines = ref(1)
  const household = ref([
    {
      line_id: 'Ana Hat',
      expected_gb: 10,
      expected_min: 500,
      tv_hd_hours: 4,
    },
  ])

  // API State
  const coverage = ref(null)
  const recommendations = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Cities and Districts data
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
    if (!city.value) return []
    return districtsByCity.value[city.value] || []
  })

  const canContinue = computed(() => {
    return city.value && district.value && householdMembers.value > 0 && activeLines.value > 0
  })

  // Address mapping for demo
  const getAddressId = () => {
    const cityMap = {
      'İstanbul': 'A1001',
      'Ankara': 'A2001', 
      'İzmir': 'A3001'
    }
    return cityMap[city.value] || 'A1001'
  }
  
  // Backward-compat alias for legacy usage in components
  const getMockAddressId = () => getAddressId()

  // Actions
  const setCity = (newCity) => {
    city.value = newCity
    district.value = '' // Reset district when city changes
  }

  const setDistrict = (newDistrict) => {
    district.value = newDistrict
  }

  const setHouseholdMembers = (count) => {
    householdMembers.value = Math.max(1, Math.min(10, parseInt(count) || 1))
  }

  const setActiveLines = (count) => {
    const newCount = Math.max(1, Math.min(5, parseInt(count) || 1))
    activeLines.value = newCount

    // Auto-adjust household lines based on active lines
    if (newCount > household.value.length) {
      // Add more lines
      for (let i = household.value.length; i < newCount; i++) {
        addHouseholdLine()
      }
    } else if (newCount < household.value.length) {
      // Remove excess lines
      household.value.splice(newCount)
      // Rename remaining lines
      const lineNames = ['Ana Hat', 'İkinci Hat', 'Üçüncü Hat', 'Dördüncü Hat', 'Beşinci Hat']
      household.value.forEach((line, idx) => {
        line.line_id = lineNames[idx] || `${idx + 1}. Hat`
      })
    }
  }

  const addHouseholdLine = () => {
    if (household.value.length < 5) {
      const newLineNumber = household.value.length + 1
      const lineNames = ['Ana Hat', 'İkinci Hat', 'Üçüncü Hat', 'Dördüncü Hat', 'Beşinci Hat']
      const lineName = lineNames[newLineNumber - 1] || `${newLineNumber}. Hat`
      household.value.push({
        line_id: lineName,
        expected_gb: 10,
        expected_min: 500,
        tv_hd_hours: 4,
      })
    }
  }

  const updateHouseholdLine = (index, field, value) => {
    if (household.value[index]) {
      household.value[index][field] = parseInt(value) || 0
    }
  }

  const removeHouseholdLine = (index) => {
    if (household.value.length > 1) {
      household.value.splice(index, 1)
      activeLines.value = household.value.length
    }
  }

  // API Actions
  const checkCoverage = async () => {
    if (!canContinue.value) return
    
    loading.value = true
    error.value = null
    
    try {
      const addressId = getAddressId()
      coverage.value = await getCoverageByAddress(addressId)
    } catch (err) {
      error.value = err.message
      console.error('Coverage check failed:', err)
    } finally {
      loading.value = false
    }
  }

  const getRecommendations = async () => {
    if (!canContinue.value) return
    
    loading.value = true
    error.value = null
    
    try {
      const addressId = getAddressId()
      const requestData = {
        user_id: 1001, // Demo user
        address_id: addressId,
        household: household.value.map(line => ({
          user_id: 1001,
          line_id: line.line_id,
          expected_gb: line.expected_gb,
          expected_min: line.expected_min,
          tv_hd_hours: line.tv_hd_hours
        })),
        prefer_tech: coverage.value?.map(c => c.tech) || ['fiber', 'vdsl']
      }
      
      recommendations.value = await postRecommendation(requestData)
    } catch (err) {
      error.value = err.message
      console.error('Recommendation failed:', err)
    } finally {
      loading.value = false
    }
  }

  const reset = () => {
    city.value = ''
    district.value = ''
    householdMembers.value = 1
    activeLines.value = 1
    household.value = [{
      line_id: 'Ana Hat',
      expected_gb: 10,
      expected_min: 500,
      tv_hd_hours: 4,
    }]
    coverage.value = null
    recommendations.value = null
    error.value = null
  }

  return {
    // State
    city,
    district,
    householdMembers,
    activeLines,
    household,
    coverage,
    recommendations,
    loading,
    error,
    
    // Data
    cities,
    districtsByCity,
    
    // Computed
    availableDistricts,
    canContinue,
    
    // Actions
    setCity,
    setDistrict,
    setHouseholdMembers,
    setActiveLines,
    addHouseholdLine,
    updateHouseholdLine,
    removeHouseholdLine,
    
    // API Actions
    checkCoverage,
    getRecommendations,
    reset,
    
    // Helpers
    getAddressId,
    getMockAddressId,
  }
})
