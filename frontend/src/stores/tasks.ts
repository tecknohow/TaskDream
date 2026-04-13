import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { tasksAPI } from '@/api/tasks'
import type { Task, TaskStatus } from '@/types/models'

export const useTasksStore = defineStore('tasks', () => {
  const tasks = ref<Task[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const selectedTask = ref<Task | null>(null)

  const tasksByProject = computed(() => {
    return (projectId: string) => tasks.value.filter(t => t.projectId === projectId)
  })

  const tasksByStatus = computed(() => {
    return (status: TaskStatus) => tasks.value.filter(t => t.status === status)
  })

  const tasksByBucket = computed(() => {
    return (bucketId: string) => tasks.value.filter(t => t.bucketId === bucketId)
  })

  const overdueTasks = computed(() => {
    return tasks.value.filter(t => t.dueDate && new Date(t.dueDate) < new Date() && t.status !== 'done')
  })

  const todayTasks = computed(() => {
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    return tasks.value.filter(t => {
      if (!t.dueDate) return false
      const dueDate = new Date(t.dueDate)
      dueDate.setHours(0, 0, 0, 0)
      return dueDate.getTime() === today.getTime()
    })
  })

  const fetchAll = async (filters?: { projectId?: string; status?: string }) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.getAll(filters)
      tasks.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch tasks'
    } finally {
      loading.value = false
    }
  }

  const fetchById = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.getById(id)
      const index = tasks.value.findIndex(t => t.id === id)
      if (index > -1) {
        tasks.value[index] = response.data
      } else {
        tasks.value.push(response.data)
      }
      selectedTask.value = response.data
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch task'
    } finally {
      loading.value = false
    }
  }

  const create = async (task: Partial<Task>) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.create(task)
      tasks.value.push(response.data)
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to create task'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const update = async (id: string, task: Partial<Task>) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.update(id, task)
      const index = tasks.value.findIndex(t => t.id === id)
      if (index > -1) {
        tasks.value[index] = response.data
      }
      if (selectedTask.value?.id === id) {
        selectedTask.value = response.data
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to update task'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const updateStatus = async (id: string, status: string) => {
    return update(id, { status: status as TaskStatus })
  }

  const updateOrder = async (id: string, order: number, bucketId?: string) => {
    try {
      const response = await tasksAPI.updateOrder(id, order, bucketId)
      const index = tasks.value.findIndex(t => t.id === id)
      if (index > -1) {
        tasks.value[index] = response.data
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to update task order'
      throw error.value
    }
  }

  const deleteTask = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      await tasksAPI.delete(id)
      tasks.value = tasks.value.filter(t => t.id !== id)
      if (selectedTask.value?.id === id) {
        selectedTask.value = null
      }
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to delete task'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const selectTask = (task: Task | null) => {
    selectedTask.value = task
  }

  return {
    tasks,
    selectedTask,
    loading,
    error,
    tasksByProject,
    tasksByStatus,
    tasksByBucket,
    overdueTasks,
    todayTasks,
    fetchAll,
    fetchById,
    create,
    update,
    updateStatus,
    updateOrder,
    deleteTask,
    selectTask
  }
})
