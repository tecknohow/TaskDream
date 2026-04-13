import client from './client'
import type { Task } from '@/types/models'

export const tasksAPI = {
  getAll: (filters?: { projectId?: string; status?: string }) => {
    return client.get<Task[]>('/tasks', { params: filters })
  },

  getById: (id: string) => {
    return client.get<Task>(`/tasks/${id}`)
  },

  create: (task: Partial<Task>) => {
    return client.post<Task>('/tasks', task)
  },

  update: (id: string, task: Partial<Task>) => {
    return client.put<Task>(`/tasks/${id}`, task)
  },

  delete: (id: string) => {
    return client.delete(`/tasks/${id}`)
  },

  getByProject: (projectId: string) => {
    return client.get<Task[]>(`/projects/${projectId}/tasks`)
  },

  updateStatus: (id: string, status: string) => {
    return client.patch<Task>(`/tasks/${id}/status`, { status })
  },

  updateOrder: (id: string, order: number, bucketId?: string) => {
    return client.patch<Task>(`/tasks/${id}/order`, { order, bucketId })
  },

  bulkUpdate: (ids: string[], updates: Partial<Task>) => {
    return client.post<Task[]>('/tasks/bulk', { ids, updates })
  },

  getTodayTasks: () => {
    return client.get<Task[]>('/tasks/today')
  },

  getOverdueTasks: () => {
    return client.get<Task[]>('/tasks/overdue')
  }
}
