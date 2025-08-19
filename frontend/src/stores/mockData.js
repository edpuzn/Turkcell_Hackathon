import { defineStore } from 'pinia'
import { ref } from 'vue'

// Mock data based on CSV files
export const useMockDataStore = defineStore('mockData', () => {
  // 1. Users CSV
  const users = ref([
    {
      user_id: 1001,
      name: 'Ahmet Yılmaz',
      address_id: 'A1001',
      current_bundle_label: 'Mobil(101,102)+Ev100',
    },
    {
      user_id: 1002,
      name: 'Fatma Demir',
      address_id: 'A1002',
      current_bundle_label: 'Mobil(103)+Ev201',
    },
    {
      user_id: 1003,
      name: 'Mehmet Kaya',
      address_id: 'A1003',
      current_bundle_label: 'Mobil(102)+Ev203',
    },
  ])

  // 2. Coverage CSV
  const coverage = ref([
    { address_id: 'A1001', city: 'Istanbul', district: 'Kadikoy', fiber: 1, vdsl: 1, fwa: 1 },
    { address_id: 'A1002', city: 'Ankara', district: 'Cankaya', fiber: 0, vdsl: 1, fwa: 1 },
    { address_id: 'A1003', city: 'Izmir', district: 'Bornova', fiber: 1, vdsl: 0, fwa: 1 },
    { address_id: 'A1004', city: 'Istanbul', district: 'Besiktas', fiber: 1, vdsl: 1, fwa: 0 },
    { address_id: 'A1005', city: 'Ankara', district: 'Kecıoren', fiber: 0, vdsl: 0, fwa: 1 },
  ])

  // 3. Mobile Plans CSV
  const mobilePlans = ref([
    {
      plan_id: 101,
      plan_name: 'GNÇ 10GB',
      quota_gb: 10,
      quota_min: 500,
      monthly_price: 230,
      overage_gb: 25,
      overage_min: 0.5,
    },
    {
      plan_id: 102,
      plan_name: 'Bireysel 20GB',
      quota_gb: 20,
      quota_min: 1000,
      monthly_price: 350,
      overage_gb: 22,
      overage_min: 0.4,
    },
    {
      plan_id: 103,
      plan_name: 'Platinum 30GB',
      quota_gb: 30,
      quota_min: 2000,
      monthly_price: 520,
      overage_gb: 20,
      overage_min: 0.3,
    },
    {
      plan_id: 104,
      plan_name: 'Premium 50GB',
      quota_gb: 50,
      quota_min: 5000,
      monthly_price: 750,
      overage_gb: 18,
      overage_min: 0.2,
    },
  ])

  // 4. Home Plans CSV
  const homePlans = ref([
    {
      home_id: 201,
      name: 'Fiber 100',
      tech: 'fiber',
      down_mbps: 100,
      monthly_price: 399,
      install_fee: 0,
    },
    {
      home_id: 202,
      name: 'Fiber 200',
      tech: 'fiber',
      down_mbps: 200,
      monthly_price: 499,
      install_fee: 0,
    },
    {
      home_id: 203,
      name: 'VDSL 35',
      tech: 'vdsl',
      down_mbps: 35,
      monthly_price: 329,
      install_fee: 0,
    },
    {
      home_id: 204,
      name: 'FWA 30',
      tech: 'fwa',
      down_mbps: 30,
      monthly_price: 349,
      install_fee: 0,
    },
    {
      home_id: 205,
      name: 'Fiber 500',
      tech: 'fiber',
      down_mbps: 500,
      monthly_price: 699,
      install_fee: 0,
    },
  ])

  // 5. TV Plans CSV
  const tvPlans = ref([
    { tv_id: 301, name: 'TV Plus Temel', hd_hours_included: 20, monthly_price: 69 },
    { tv_id: 302, name: 'TV Plus Aile', hd_hours_included: 60, monthly_price: 109 },
    { tv_id: 303, name: 'TV Plus Sinema+', hd_hours_included: 120, monthly_price: 149 },
    { tv_id: 304, name: 'TV Plus Premium', hd_hours_included: 200, monthly_price: 199 },
  ])

  // 6. Bundling Rules CSV
  const bundlingRules = ref([
    {
      rule_id: 401,
      rule_type: 'bundle',
      description: 'Mobil+Ev',
      discount_percent: 10,
      applies_to: 'all',
    },
    {
      rule_id: 402,
      rule_type: 'bundle',
      description: 'Mobil+Ev+TV',
      discount_percent: 15,
      applies_to: 'all',
    },
    {
      rule_id: 403,
      rule_type: 'extra_line',
      description: '2. Hat -%5 (mobil bileşen)',
      discount_percent: 5,
      applies_to: 'mobile',
    },
    {
      rule_id: 404,
      rule_type: 'extra_line',
      description: '3+ Hat -%10 (mobil bileşen)',
      discount_percent: 10,
      applies_to: 'mobile',
    },
  ])

  // 7. Household CSV
  const household = ref([
    { user_id: 1001, line_id: 'L-1', expected_gb: 14, expected_min: 400, tv_hd_hours: 40 },
    { user_id: 1001, line_id: 'L-2', expected_gb: 7, expected_min: 200, tv_hd_hours: 0 },
    { user_id: 1001, line_id: 'L-3', expected_gb: 25, expected_min: 800, tv_hd_hours: 60 },
    { user_id: 1002, line_id: 'L-1', expected_gb: 20, expected_min: 600, tv_hd_hours: 30 },
    { user_id: 1003, line_id: 'L-1', expected_gb: 15, expected_min: 300, tv_hd_hours: 20 },
  ])

  // 8. Current Services CSV
  const currentServices = ref([
    {
      user_id: 1001,
      has_home: 1,
      home_tech: 'fiber',
      home_speed: 100,
      has_tv: 0,
      mobile_plan_ids: '101;102',
    },
    {
      user_id: 1002,
      has_home: 1,
      home_tech: 'fiber',
      home_speed: 200,
      has_tv: 1,
      mobile_plan_ids: '103',
    },
    {
      user_id: 1003,
      has_home: 1,
      home_tech: 'vdsl',
      home_speed: 35,
      has_tv: 0,
      mobile_plan_ids: '102',
    },
  ])

  // 9. Install Slots CSV
  const installSlots = ref([
    {
      slot_id: 'S1',
      address_id: 'A1001',
      slot_start: '2025-08-20T10:00',
      slot_end: '2025-08-20T12:00',
      tech: 'fiber',
    },
    {
      slot_id: 'S2',
      address_id: 'A1001',
      slot_start: '2025-08-20T14:00',
      slot_end: '2025-08-20T16:00',
      tech: 'fiber',
    },
    {
      slot_id: 'S3',
      address_id: 'A1002',
      slot_start: '2025-08-21T10:00',
      slot_end: '2025-08-21T12:00',
      tech: 'vdsl',
    },
    {
      slot_id: 'S4',
      address_id: 'A1002',
      slot_start: '2025-08-21T14:00',
      slot_end: '2025-08-21T16:00',
      tech: 'vdsl',
    },
    {
      slot_id: 'S5',
      address_id: 'A1003',
      slot_start: '2025-08-22T10:00',
      slot_end: '2025-08-22T12:00',
      tech: 'fiber',
    },
  ])

  // Helper functions
  const getCoverageByAddress = (addressId) => {
    return coverage.value.find((c) => c.address_id === addressId)
  }

  const getHomePlansByTech = (tech) => {
    return homePlans.value.filter((h) => h.tech === tech)
  }

  const getMobilePlansByQuota = (minGb, maxGb) => {
    return mobilePlans.value.filter((m) => m.quota_gb >= minGb && m.quota_gb <= maxGb)
  }

  const getTVPlansByHours = (hours) => {
    return tvPlans.value.filter((t) => t.hd_hours_included >= hours)
  }

  const calculateBundleDiscount = (bundleType) => {
    const rule = bundlingRules.value.find((r) => r.description.includes(bundleType))
    return rule ? rule.discount_percent : 0
  }

  const calculateExtraLineDiscount = (lineCount) => {
    if (lineCount >= 3) return 10
    if (lineCount === 2) return 5
    return 0
  }

  return {
    users,
    coverage,
    mobilePlans,
    homePlans,
    tvPlans,
    bundlingRules,
    household,
    currentServices,
    installSlots,
    getCoverageByAddress,
    getHomePlansByTech,
    getMobilePlansByQuota,
    getTVPlansByHours,
    calculateBundleDiscount,
    calculateExtraLineDiscount,
  }
})
