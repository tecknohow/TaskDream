import client from './client'
import type { TimeTrackingEntry } from '@/types/models'

export const timeTrackingAPI = {
  getEntries: (taskId: number) => {
    return client.get<TimeTrackingEntry[]>(`/tasks/${taskId}/time-tracking`)
  },

  createEntry: (taskId: number, entry: Partial<TimeTrackingEntry>) => {
    return client.post<TimeTrackingEntry>(`/tasks/${taskId}/time-tracking`, entry)
  },

  updateEntry: (id: number, entry: Partial<TimeTrackingEntry>) => {
    return client.put<TimeTrackingEntry>(`/time-tracking/${id}`, entry)
  },

  deleteEntry: (id: number) => {
    return client.delete(`/time-tracking/${id}`)
  },

  startTracking: (taskId: number) => {
    const now = new Date().toISOString()
    return client.post<TimeTrackingEntry>(`/tasks/${taskId}/time-tracking`, {
      start: now,
      duration: 0
    })
  },

  stopTracking: (id: number, duration: number) => {
    const now = new Date().toISOString()
    return client.put<TimeTrackingEntry>(`/time-tracking/${id}`, {
      end: now,
      duration: duration
    })
  }
}
