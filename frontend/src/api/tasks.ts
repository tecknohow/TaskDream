import client from './client'
import type { Task, TaskWithSubtasks } from '@/types/models'

export interface TaskFilters {
  project_id?: number
  status?: 'done' | 'undone'
  priority?: number
  sort?: 'priority' | 'due_date' | 'created' | 'position'
  order?: 'asc' | 'desc'
  filter?: 'overdue' | 'today' | 'upcoming' | 'no_date'
}

export const tasksAPI = {
  getAll: (filters?: TaskFilters) => {
    return client.get<Task[]>('/tasks', { params: filters })
  },

  getById: (id: number) => {
    return client.get<TaskWithSubtasks>(`/tasks/${id}`)
  },

  create: (task: Partial<Task>) => {
    return client.post<Task>('/tasks', task)
  },

  update: (id: number, task: Partial<Task>) => {
    return client.put<Task>(`/tasks/${id}`, task)
  },

  delete: (id: number) => {
    return client.delete(`/tasks/${id}`)
  },

  // Subtasks
  getSubtasks: (taskId: number) => {
    return client.get<Task[]>(`/tasks/${taskId}/subtasks`)
  },

  createSubtask: (taskId: number, subtask: Partial<Task>) => {
    return client.post<Task>(`/tasks/${taskId}/subtasks`, subtask)
  },

  // Time tracking for task
  getTimeTracking: (taskId: number) => {
    return client.get(`/tasks/${taskId}/time-tracking`)
  },

  createTimeTracking: (taskId: number, entry: any) => {
    return client.post(`/tasks/${taskId}/time-tracking`, entry)
  }
}
