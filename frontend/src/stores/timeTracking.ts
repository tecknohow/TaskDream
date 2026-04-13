import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { timeTrackingAPI } from '@/api/timeTracking'
import type { TimeTrackingEntry } from '@/types/models'

export const useTimeTrackingStore = defineStore('timeTracking', () => {
  const entries = ref<TimeTrackingEntry[]>([])
  const currentEntry = ref<TimeTrackingEntry | null>(null)
  const currentTaskId = ref<number | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const elapsedSeconds = ref(0)

  let timerInterval: ReturnType<typeof setInterval> | null = null

  const isTracking = computed(() => !!currentEntry.value)

  const formattedTime = computed(() => {
    const seconds = elapsedSeconds.value % 60
    const minutes = Math.floor((elapsedSeconds.value / 60) % 60)
    const hours = Math.floor(elapsedSeconds.value / 3600)

    const pad = (num: number) => String(num).padStart(2, '0')
    return `${pad(hours)}:${pad(minutes)}:${pad(seconds)}`
  })

  const startTracking = async (taskId: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.startTracking(taskId)
      currentEntry.value = response.data
      currentTaskId.value = taskId
      elapsedSeconds.value = 0

      if (timerInterval) clearInterval(timerInterval)
      timerInterval = setInterval(() => {
        elapsedSeconds.value++
      }, 1000)

      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to start tracking'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const stopTracking = async () => {
    if (!currentEntry.value) return

    loading.value = true
    error.value = null

    if (timerInterval) {
      clearInterval(timerInterval)
      timerInterval = null
    }

    try {
      const response = await timeTrackingAPI.stopTracking(currentEntry.value.id, elapsedSeconds.value)
      entries.value.unshift(response.data)
      currentEntry.value = null
      currentTaskId.value = null
      elapsedSeconds.value = 0
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to stop tracking'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const fetchEntries = async (taskId: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await timeTrackingAPI.getEntries(taskId)
      entries.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch entries'
    } finally {
      loading.value = false
    }
  }

  const deleteEntry = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      await timeTrackingAPI.deleteEntry(id)
      entries.value = entries.value.filter(e => e.id !== id)
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to delete entry'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  return {
    entries,
    currentEntry,
    currentTaskId,
    loading,
    error,
    elapsedSeconds,
    isTracking,
    formattedTime,
    startTracking,
    stopTracking,
    fetchEntries,
    deleteEntry
  }
})
