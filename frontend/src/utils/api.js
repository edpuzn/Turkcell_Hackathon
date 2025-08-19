/**
 * API utility functions for frontend
 */

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '';

/**
 * API isteği gönderir
 * @param {string} endpoint - API endpoint'i
 * @param {Object} options - Fetch options
 * @returns {Promise<Object>} API response
 */
export async function apiRequest(endpoint, options = {}) {
  const url = `${API_BASE_URL}${endpoint}`;
  
  const defaultOptions = {
    headers: {
      'Content-Type': 'application/json',
    },
    ...options,
  };

  try {
    const response = await fetch(url, defaultOptions);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('API request failed:', error);
    throw error;
  }
}

/**
 * GET isteği gönderir
 * @param {string} endpoint - API endpoint'i
 * @returns {Promise<Object>} API response
 */
export async function apiGet(endpoint) {
  return apiRequest(endpoint, { method: 'GET' });
}

/**
 * POST isteği gönderir
 * @param {string} endpoint - API endpoint'i
 * @param {Object} data - Request body
 * @returns {Promise<Object>} API response
 */
export async function apiPost(endpoint, data) {
  return apiRequest(endpoint, {
    method: 'POST',
    body: JSON.stringify(data),
  });
}

/**
 * PUT isteği gönderir
 * @param {string} endpoint - API endpoint'i
 * @param {Object} data - Request body
 * @returns {Promise<Object>} API response
 */
export async function apiPut(endpoint, data) {
  return apiRequest(endpoint, {
    method: 'PUT',
    body: JSON.stringify(data),
  });
}

/**
 * DELETE isteği gönderir
 * @param {string} endpoint - API endpoint'i
 * @returns {Promise<Object>} API response
 */
export async function apiDelete(endpoint) {
  return apiRequest(endpoint, { method: 'DELETE' });
}

/**
 * Health check yapar
 * @returns {Promise<Object>} Health status
 */
export async function healthCheck() {
  return apiGet('/health');
}

/**
 * Katalog özetini getirir
 * @returns {Promise<Object>} Catalog summary
 */
export async function getCatalogSummary() {
  return apiGet('/api/catalog/summary');
}

/**
 * Adres bazlı kapsama bilgisini getirir
 * @param {string} addressId - Adres ID'si
 * @returns {Promise<Object>} Coverage information
 */
export async function getCoverageByAddress(addressId) {
  return apiGet(`/api/coverage/${addressId}`);
}

/**
 * Öneri sistemi isteği gönderir
 * @param {Object} recommendationData - Öneri verileri
 * @returns {Promise<Object>} Recommendation response
 */
export async function postRecommendation(recommendationData) {
  return apiPost('/api/recommendation', recommendationData);
}

// API endpoints constants
export const API_ENDPOINTS = {
  HEALTH: '/health',
  CATALOG_SUMMARY: '/api/catalog/summary',
  COVERAGE: '/api/coverage',
  RECOMMENDATION: '/api/recommendation',
};
