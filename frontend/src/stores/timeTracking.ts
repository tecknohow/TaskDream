import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { timeTrackingAPI } from '@/api/timeTracking'
import type { TimeTrackingEntry } from '@/types/models'

export const useTimeTrackingStore = defineStore('timeTracking', () => {
  const entries = ref<TimeTrackingEntry[]>([])
  const currentEntry = ref<TimeTrackingEntry | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const elapsedSeconds = ref(0)

  let timerInterval: NodeJS.Timeout | null = null

  const isTracking = computed(() => !!currentEntry.value)

  const formattedTime = computed(() => {
    const seconds = elapsedSeconds.value % 60
    const minutes = Math.floor((elapsedSeconds.value / 60) % 60)
    const hours = Math.floor(elapsedSeconds.value / 3600)

    const pad = (num: number) => String(num).padStart(2, '0')
    return `${pad(hours)}:${pad(minutes)}:${pad(seconds)}`
  })

  const totalTimeTrackedForTask = computed(() => {
    return (taskId: string) => {
      return entries.value
        .filter(e => e.taskId === taskId)
        .reduce((sum, e) => sum + e.duration, 0)
    }
  })

  const startTracking = async (taskId: string) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.startTracking(taskId)
      currentEntry.value = response.data
      elapsedSeconds.value = 0

      if (timerInterval) clearInterval(timerInterval)
      timerInterval = setInterval(() => {
        elapsedSeconds.value++
      }, 1000)

      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to start tracking'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const stopTracking = async (entryId: string) => {
    loading.value = true
    error.value = null

    if (timerInterval) {
      clearInterval(timerInterval)
      timerInterval = null
    }

    try {
      const response = await timeTrackingAPI.stopTracking(entryId)
      const index = entries.value.findIndex(e => e.id === entryId)
      if (index > -1) {
        entries.value[index] = response.data
      } else {
        entries.value.push(response.data)
      }
      currentEntry.value = null
      elapsedSeconds.value = 0
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to stop tracking'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const fetchEntries = async (filters?: { taskId?: string; userId?: string }) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.getEntries(filters)
      entries.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch entries'
    } finally {
      loading.value = false
    }
  }

  const createEntry = async (entry: Partial<TimeTrackingEntry>) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.createEntry(entry)
      entries.value.push(response.data)
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to create entry'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const updateEntry = async (id: string, entry: Partial<TimeTrackingEntry>) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.updateEntry(id, entry)
      const index = entries.value.findIndex(e => e.id === id)
      if (index > -1) {
        entries.value[index] = response.data
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to update entry'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const deleteEntry = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      await timeTrackingAPI.deleteEntry(id)
      entries.value = entries.value.filter(e => e.id !== id)
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to delete entry'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const getCurrentTimer = async () => {
    try {
      const response = await timeTrackingAPI.getCurrentTimer()
      if (response.data) {
        currentEntry.value = response.data
        elapsedSeconds.value = 0

        if (timerInterval) clearInterval(timerInterval)
        timerInterval = setInterval(() => {
          elapsedSeconds.value++
        }, 1000)
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to get current timer'
    }
  }

  return {
    entries,
    currentEntry,
    loading,
    error,
    elapsedSeconds,
    isTracking,
    formattedTime,
    totalTimeTrackedForTask,
    startTracking,
    stopTracking,
    fetchEntries,
    createEntry,
    updateEntry,
    deleteEntry,
    getCurrentTimer
  }
})
