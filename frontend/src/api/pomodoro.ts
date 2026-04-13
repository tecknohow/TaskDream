import client from './client'
import type { PomodoroSession, PomodoroSettings, PomodoroStats } from '@/types/models'

export const pomodoroAPI = {
  getSessions: (params?: { task_id?: number; date?: string }) => {
    return client.get<PomodoroSession[]>('/pomodoro/sessions', { params })
  },

  start: (taskId: number, duration?: number) => {
    return client.post<PomodoroSession>('/pomodoro/start', {
      task_id: taskId,
      duration: duration || 0
    })
  },

  complete: (id: number) => {
    return client.post<PomodoroSession>(`/pomodoro/${id}/complete`)
  },

  cancel: (id: number) => {
    return client.post(`/pomodoro/${id}/cancel`)
  },

  getSettings: () => {
    return client.get<PomodoroSettings>('/pomodoro/settings')
  },

  updateSettings: (settings: Partial<PomodoroSettings>) => {
    return client.put<PomodoroSettings>('/pomodoro/settings', settings)
  },

  getStats: (date?: string) => {
    return client.get<PomodoroStats>('/pomodoro/stats', { params: { date } })
  }
}
