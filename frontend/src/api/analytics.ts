import client from './client'
import type { DashboardStats, ProductivityTrend, SearchResults } from '@/types/models'

export const analyticsAPI = {
  getDashboardStats: () => {
    return client.get<DashboardStats>('/analytics/dashboard')
  },

  getProductivityTrend: (days?: number) => {
    return client.get<ProductivityTrend[]>('/analytics/trend', {
      params: days ? { days: `${days * 24}h` } : undefined
    })
  }
}

export const searchAPI = {
  search: (query: string) => {
    return client.get<SearchResults>('/search', { params: { q: query } })
  }
}
