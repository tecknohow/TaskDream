import client from './client'
import type { TimeTrackingEntry } from '@/types/models'

export const timeTrackingAPI = {
  getEntries: (filters?: { taskId?: string; userId?: string }) => {
    return client.get<TimeTrackingEntry[]>('/time-tracking', { params: filters })
  },

  getEntry: (id: string) => {
    return client.get<TimeTrackingEntry>(`/time-tracking/${id}`)
  },

  startTracking: (taskId: string) => {
    return client.post<TimeTrackingEntry>('/time-tracking/start', { taskId })
  },

  stopTracking: (entryId: string) => {
    return client.post<TimeTrackingEntry>(`/time-tracking/${entryId}/stop`)
  },

  createEntry: (entry: Partial<TimeTrackingEntry>) => {
    return client.post<TimeTrackingEntry>('/time-tracking', entry)
  },

  updateEntry: (id: string, entry: Partial<TimeTrackingEntry>) => {
    return client.put<TimeTrackingEntry>(`/time-tracking/${id}`, entry)
  },

  deleteEntry: (id: string) => {
    return client.delete(`/time-tracking/${id}`)
  },

  getTaskTimeTracked: (taskId: string) => {
    return client.get<{ totalDuration: number }>(`/tasks/${taskId}/time-tracked`)
  },

  getCurrentTimer: () => {
    return client.get<TimeTrackingEntry | null>('/time-tracking/current')
  }
}
