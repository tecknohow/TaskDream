import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { tasksAPI, type TaskFilters } from '@/api/tasks'
import type { Task } from '@/types/models'

export const useTasksStore = defineStore('tasks', () => {
  const tasks = ref<Task[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const selectedTask = ref<Task | null>(null)

  const tasksByProject = computed(() => {
    return (projectId: number) => tasks.value.filter(t => t.project_id === projectId)
  })

  const tasksByBucket = computed(() => {
    return (bucketId: number) => tasks.value.filter(t => t.bucket_id === bucketId)
  })

  const overdueTasks = computed(() => {
    return tasks.value.filter(t => t.due_date && new Date(t.due_date) < new Date() && !t.done)
  })

  const todayTasks = computed(() => {
    const today = new Date()
    today.setHours(0, 0, 0, 0)
    const tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)
    return tasks.value.filter(t => {
      if (!t.due_date) return false
      const d = new Date(t.due_date)
      return d >= today && d < tomorrow
    })
  })

  const doneTasks = computed(() => tasks.value.filter(t => t.done))
  const openTasks = computed(() => tasks.value.filter(t => !t.done))

  // Eisenhower matrix quadrants
  const eisenhowerQuadrants = computed(() => ({
    urgentImportant: tasks.value.filter(t => !t.done && t.urgency === 1 && t.importance === 1),
    notUrgentImportant: tasks.value.filter(t => !t.done && t.urgency === 0 && t.importance === 1),
    urgentNotImportant: tasks.value.filter(t => !t.done && t.urgency === 1 && t.importance === 0),
    notUrgentNotImportant: tasks.value.filter(t => !t.done && t.urgency === 0 && t.importance === 0),
  }))

  const fetchAll = async (filters?: TaskFilters) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.getAll(filters)
      tasks.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch tasks'
    } finally {
      loading.value = false
    }
  }

  const fetchById = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await tasksAPI.getById(id)
      const taskData = response.data.task
      const index = tasks.value.findIndex(t => t.id === id)
      if (index > -1) {
        tasks.value[index] = taskData
      } else {
        tasks.value.push(taskData)
      }
      selectedTask.value = taskData
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch task'
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
      error.value = err.response?.data?.error || 'Failed to create task'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const update = async (id: number, task: Partial<Task>) => {
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
      error.value = err.response?.data?.error || 'Failed to update task'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const toggleDone = async (id: number) => {
    const task = tasks.value.find(t => t.id === id)
    if (task) {
      return update(id, { done: !task.done })
    }
  }

  const deleteTask = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      await tasksAPI.delete(id)
      tasks.value = tasks.value.filter(t => t.id !== id)
      if (selectedTask.value?.id === id) {
        selectedTask.value = null
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to delete task'
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
    tasksByBucket,
    overdueTasks,
    todayTasks,
    doneTasks,
    openTasks,
    eisenhowerQuadrants,
    fetchAll,
    fetchById,
    create,
    update,
    toggleDone,
    deleteTask,
    selectTask
  }
})
