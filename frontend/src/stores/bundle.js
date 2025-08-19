import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useBundleStore = defineStore('bundle', () => {
  // State
  const coverageData = ref([])
  const cities = ref([])
  const districts = ref({})
  const addresses = ref({})
  const recommendations = ref([])
  const isLoading = ref(false)
  const error = ref(null)

  // Mock data for MVP
  const mockCities = ['İstanbul', 'Ankara', 'İzmir', 'Bursa', 'Antalya']
  const mockDistricts = {
    'İstanbul': ['Kadıköy', 'Beşiktaş', 'Şişli', 'Beyoğlu', 'Fatih'],
    'Ankara': ['Çankaya', 'Keçiören', 'Mamak', 'Yenimahalle', 'Etimesgut'],
    'İzmir': ['Konak', 'Bornova', 'Karşıyaka', 'Buca', 'Çiğli'],
    'Bursa': ['Nilüfer', 'Osmangazi', 'Yıldırım', 'Mudanya', 'Gemlik'],
    'Antalya': ['Muratpaşa', 'Kepez', 'Konyaaltı', 'Aksu', 'Döşemealtı']
  }
  const mockAddresses = {
    'İstanbul': {
      'Kadıköy': [
        { id: 'IST_KAD_001', street: 'Bağdat Caddesi', building: 'No:123' },
        { id: 'IST_KAD_002', street: 'Fenerbahçe Mahallesi', building: 'No:45' },
        { id: 'IST_KAD_003', street: 'Sahrayıcedit Mahallesi', building: 'No:67' }
      ],
      'Beşiktaş': [
        { id: 'IST_BES_001', street: 'Levent Mahallesi', building: 'No:89' },
        { id: 'IST_BES_002', street: 'Etiler Mahallesi', building: 'No:12' }
      ]
    },
    'Ankara': {
      'Çankaya': [
        { id: 'ANK_CAN_001', street: 'Kızılay Meydanı', building: 'No:34' },
        { id: 'ANK_CAN_002', street: 'Bahçelievler Mahallesi', building: 'No:56' }
      ]
    }
  }

  // Coverage data for different address types
  const mockCoverageData = {
    'IST_KAD_001': [
      { type: 'fiber', name: 'Fiber Optik', available: true, speed: '1000 Mbps' },
      { type: 'vdsl', name: 'VDSL', available: true, speed: '100 Mbps' },
      { type: 'fwa', name: '4.5G Sabit İnternet', available: true, speed: '150 Mbps' }
    ],
    'IST_KAD_002': [
      { type: 'fiber', name: 'Fiber Optik', available: false, speed: null },
      { type: 'vdsl', name: 'VDSL', available: true, speed: '50 Mbps' },
      { type: 'fwa', name: '4.5G Sabit İnternet', available: true, speed: '100 Mbps' }
    ],
    'IST_BES_001': [
      { type: 'fiber', name: 'Fiber Optik', available: true, speed: '2000 Mbps' },
      { type: 'vdsl', name: 'VDSL', available: true, speed: '200 Mbps' },
      { type: 'fwa', name: '4.5G Sabit İnternet', available: true, speed: '200 Mbps' }
    ]
  }

  // Catalog data for plans and combinations
  const mockCatalog = {
    mobile: [
      { id: 'mob_1', name: '5 GB + 100 dk', gb: 5, minutes: 100, price: 49.90 },
      { id: 'mob_2', name: '10 GB + 500 dk', gb: 10, minutes: 500, price: 79.90 },
      { id: 'mob_3', name: '20 GB + 1000 dk', gb: 20, minutes: 1000, price: 119.90 },
      { id: 'mob_4', name: '50 GB + 1000 dk', gb: 50, minutes: 1000, price: 159.90 },
      { id: 'mob_5', name: '100 GB + 1000 dk', gb: 100, minutes: 1000, price: 199.90 },
      { id: 'mob_6', name: 'Sınırsız + Sınırsız', gb: 'unlimited', minutes: 'unlimited', price: 299.90 }
    ],
    home: [
      { id: 'home_1', name: '50 Mbps', speed: 50, price: 89.90, technology: ['fiber', 'vdsl'] },
      { id: 'home_2', name: '100 Mbps', speed: 100, price: 119.90, technology: ['fiber', 'vdsl'] },
      { id: 'home_3', name: '200 Mbps', speed: 200, price: 149.90, technology: ['fiber'] },
      { id: 'home_4', name: '500 Mbps', speed: 500, price: 199.90, technology: ['fiber'] },
      { id: 'home_5', name: '1000 Mbps', speed: 1000, price: 299.90, technology: ['fiber'] }
    ],
    tv: [
      { id: 'tv_1', name: 'TV Plus 4 Saat', hours: 4, price: 29.90 },
      { id: 'tv_2', name: 'TV Plus 8 Saat', hours: 8, price: 39.90 },
      { id: 'tv_3', name: 'TV Plus 12 Saat', hours: 12, price: 49.90 },
      { id: 'tv_4', name: 'TV Plus 24 Saat', hours: 24, price: 69.90 }
    ]
  }

  // Getters
  const getCities = computed(() => cities.value)
  const getDistricts = computed(() => (city) => districts.value[city] || [])
  const getAddresses = computed(() => (city, district) => addresses.value[city]?.[district] || [])
  const getCoverage = computed(() => coverageData.value)
  const getRecommendations = computed(() => recommendations.value)
  const getLoading = computed(() => isLoading.value)
  const getError = computed(() => error.value)

  // Actions
  const initializeCoverage = async () => {
    try {
      isLoading.value = true
      
      // Simulate API call delay
      await new Promise(resolve => setTimeout(resolve, 500))
      
      // Initialize mock data
      cities.value = mockCities
      districts.value = mockDistricts
      addresses.value = mockAddresses
      
      isLoading.value = false
    } catch (err) {
      error.value = err.message
      isLoading.value = false
    }
  }

  const fetchCoverage = async (addressId) => {
    try {
      isLoading.value = true
      
      // Simulate API call delay
      await new Promise(resolve => setTimeout(resolve, 800))
      
      // Get coverage data for address
      const coverage = mockCoverageData[addressId] || [
        { type: 'fiber', name: 'Fiber Optik', available: false, speed: null },
        { type: 'vdsl', name: 'VDSL', available: true, speed: '50 Mbps' },
        { type: 'fwa', name: '4.5G Sabit İnternet', available: true, speed: '100 Mbps' }
      ]
      
      coverageData.value = coverage
      isLoading.value = false
      
      return coverage
    } catch (err) {
      error.value = err.message
      isLoading.value = false
      throw err
    }
  }

  const fetchRecommendations = async (addressId, household) => {
    try {
      isLoading.value = true
      
      // Simulate API call delay
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      // Get coverage for address
      const coverage = mockCoverageData[addressId] || [
        { type: 'fiber', name: 'Fiber Optik', available: false, speed: null },
        { type: 'vdsl', name: 'VDSL', available: true, speed: '50 Mbps' },
        { type: 'fwa', name: '4.5G Sabit İnternet', available: true, speed: '100 Mbps' }
      ]
      
      // Find best available technology
      const bestTech = coverage.find(tech => tech.available) || coverage[0]
      
      // Generate recommendations based on household needs
      const recommendations = generateRecommendations(household, bestTech, coverage)
      
      recommendations.value = recommendations
      isLoading.value = false
      
      return recommendations
    } catch (err) {
      error.value = err.message
      isLoading.value = false
      throw err
    }
  }

  const fetchInstallSlots = async (addressId, techType) => {
    try {
      isLoading.value = true
      
      // Simulate API call delay
      await new Promise(resolve => setTimeout(resolve, 600))
      
      // Generate mock install slots
      const slots = generateInstallSlots(techType)
      
      isLoading.value = false
      return slots
    } catch (err) {
      error.value = err.message
      isLoading.value = false
      throw err
    }
  }

  const clearError = () => {
    error.value = null
  }

  // Helper functions for MVP features
  const generateRecommendations = (household, bestTech, coverage) => {
    const recommendations = []
    
    // Calculate total GB and minutes needed
    const totalGB = household.reduce((sum, line) => {
      if (line.gb === 'unlimited') return sum + 100 // Assume unlimited = 100 GB for calculation
      return sum + parseInt(line.gb)
    }, 0)
    
    const totalMinutes = household.reduce((sum, line) => {
      if (line.minutes === 'unlimited') return sum + 1000 // Assume unlimited = 1000 min for calculation
      return sum + parseInt(line.minutes)
    }, 0)
    
    const totalTVHours = household.reduce((sum, line) => sum + parseInt(line.tvHours), 0)
    
    // Find best mobile plans for each line
    const mobilePlans = findBestMobilePlans(household)
    
    // Find best home internet plan
    const homePlan = findBestHomePlan(totalGB, bestTech)
    
    // Find best TV plan
    const tvPlan = findBestTVPlan(totalTVHours)
    
    // Generate combination recommendations
    if (mobilePlans.length > 0 && homePlan) {
      // Mobile + Home combination
      const mobileHomeTotal = mobilePlans.reduce((sum, plan) => sum + plan.price, 0) + homePlan.price
      const bundleDiscount = 0.10 // 10% for dual bundle
      const discountedTotal = mobileHomeTotal * (1 - bundleDiscount)
      
      recommendations.push({
        id: 'rec_1',
        name: 'Mobil + Ev İnternet Paketi',
        summary: `${mobilePlans.length} hat, ${homePlan.speed} Mbps internet`,
        coverageType: bestTech.type,
        monthlyTotal: Math.round(discountedTotal),
        savings: Math.round(mobileHomeTotal - discountedTotal),
        bundleDiscount: 10,
        extraLineDiscount: mobilePlans.length > 1 ? (mobilePlans.length === 2 ? 5 : 10) : 0,
        isAlternative: !bestTech.available,
        mobilePlans,
        homePlan,
        tvPlan: null,
        totalOriginal: mobileHomeTotal
      })
    }
    
    if (mobilePlans.length > 0 && homePlan && tvPlan) {
      // Mobile + Home + TV combination
      const mobileHomeTVTotal = mobilePlans.reduce((sum, plan) => sum + plan.price, 0) + homePlan.price + tvPlan.price
      const bundleDiscount = 0.15 // 15% for triple bundle
      const discountedTotal = mobileHomeTVTotal * (1 - bundleDiscount)
      
      recommendations.push({
        id: 'rec_2',
        name: 'Mobil + Ev İnternet + TV Paketi',
        summary: `${mobilePlans.length} hat, ${homePlan.speed} Mbps internet, ${tvPlan.hours} saat TV`,
        coverageType: bestTech.type,
        monthlyTotal: Math.round(discountedTotal),
        savings: Math.round(mobileHomeTVTotal - discountedTotal),
        bundleDiscount: 15,
        extraLineDiscount: mobilePlans.length > 1 ? (mobilePlans.length === 2 ? 5 : 10) : 0,
        isAlternative: !bestTech.available,
        mobilePlans,
        homePlan,
        tvPlan,
        totalOriginal: mobileHomeTVTotal
      })
    }
    
    // Add alternative technology recommendation if best tech not available
    if (!bestTech.available && coverage.length > 1) {
      const alternativeTech = coverage.find(tech => tech.available && tech.type !== bestTech.type)
      if (alternativeTech) {
        const altHomePlan = findBestHomePlan(totalGB, alternativeTech)
        if (altHomePlan) {
          const altTotal = mobilePlans.reduce((sum, plan) => sum + plan.price, 0) + altHomePlan.price
          const altDiscountedTotal = altTotal * 0.9 // 10% discount
          
          recommendations.push({
            id: 'rec_3',
            name: `Alternatif ${alternativeTech.name} Paketi`,
            summary: `${mobilePlans.length} hat, ${altHomePlan.speed} Mbps ${alternativeTech.name}`,
            coverageType: alternativeTech.type,
            monthlyTotal: Math.round(altDiscountedTotal),
            savings: Math.round(altTotal - altDiscountedTotal),
            bundleDiscount: 10,
            extraLineDiscount: mobilePlans.length > 1 ? (mobilePlans.length === 2 ? 5 : 10) : 0,
            isAlternative: true,
            mobilePlans,
            homePlan: altHomePlan,
            tvPlan: null,
            totalOriginal: altTotal
          })
        }
      }
    }
    
    // Sort by monthly total (lowest first)
    return recommendations.sort((a, b) => a.monthlyTotal - b.monthlyTotal).slice(0, 3)
  }

  const findBestMobilePlans = (household) => {
    return household.map(line => {
      const gb = line.gb === 'unlimited' ? 100 : parseInt(line.gb)
      const minutes = line.minutes === 'unlimited' ? 1000 : parseInt(line.minutes)
      
      // Find best matching plan
      let bestPlan = mockCatalog.mobile[0]
      for (const plan of mockCatalog.mobile) {
        if (plan.gb >= gb && plan.minutes >= minutes) {
          bestPlan = plan
          break
        }
      }
      
      return {
        ...bestPlan,
        requiredGB: gb,
        requiredMinutes: minutes,
        overageGB: Math.max(0, gb - (bestPlan.gb === 'unlimited' ? 100 : bestPlan.gb)),
        overageMinutes: Math.max(0, minutes - (bestPlan.minutes === 'unlimited' ? 1000 : bestPlan.minutes))
      }
    })
  }

  const findBestHomePlan = (totalGB, tech) => {
    // Calculate required speed based on GB usage
    let requiredSpeed = 50
    if (totalGB > 100) requiredSpeed = 100
    if (totalGB > 200) requiredSpeed = 200
    if (totalGB > 500) requiredSpeed = 500
    if (totalGB > 1000) requiredSpeed = 1000
    
    // Find best plan for technology
    const availablePlans = mockCatalog.home.filter(plan => 
      plan.technology.includes(tech.type) && plan.speed >= requiredSpeed
    )
    
    return availablePlans.length > 0 ? availablePlans[0] : null
  }

  const findBestTVPlan = (totalHours) => {
    if (totalHours === 0) return null
    
    // Find best TV plan
    for (const plan of mockCatalog.tv) {
      if (plan.hours >= totalHours) {
        return plan
      }
    }
    
    return mockCatalog.tv[mockCatalog.tv.length - 1] // Return highest plan
  }

  const generateInstallSlots = (techType) => {
    const slots = []
    const today = new Date()
    
    // Generate slots for next 7 days
    for (let i = 1; i <= 7; i++) {
      const date = new Date(today)
      date.setDate(today.getDate() + i)
      
      // Morning slots (9:00 - 12:00)
      for (let hour = 9; hour < 12; hour++) {
        slots.push({
          id: `slot_${i}_${hour}`,
          date: date.toISOString().split('T')[0],
          time: `${hour.toString().padStart(2, '0')}:00`,
          available: Math.random() > 0.3, // 70% chance of availability
          technology: techType
        })
      }
      
      // Afternoon slots (13:00 - 17:00)
      for (let hour = 13; hour < 17; hour++) {
        slots.push({
          id: `slot_${i}_${hour}`,
          date: date.toISOString().split('T')[0],
          time: `${hour.toString().padStart(2, '0')}:00`,
          available: Math.random() > 0.3,
          technology: techType
        })
      }
    }
    
    return slots.filter(slot => slot.available)
  }

  return {
    // State
    coverageData,
    cities,
    districts,
    addresses,
    recommendations,
    isLoading,
    error,
    
    // Getters
    getCities,
    getDistricts,
    getAddresses,
    getCoverage,
    getRecommendations,
    getLoading,
    getError,
    
    // Actions
    initializeCoverage,
    fetchCoverage,
    fetchRecommendations,
    fetchInstallSlots,
    clearError
  }
}) 