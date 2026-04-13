import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { pomodoroAPI } from '@/api/pomodoro'
import type { PomodoroSession, PomodoroSettings, PomodoroStats } from '@/types/models'

export const usePomodoroStore = defineStore('pomodoro', () => {
  const currentSession = ref<PomodoroSession | null>(null)
  const settings = ref<PomodoroSettings>({
    user_id: 0,
    work_duration: 1500,
    short_break: 300,
    long_break: 900,
    long_break_interval: 4,
    auto_start_breaks: false,
    auto_start_pomodoro: false
  })
  const stats = ref<PomodoroStats>({
    today_completed: 0,
    today_total_time: 0,
    weekly_completed: 0
  })
  const sessions = ref<PomodoroSession[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Timer state
  const elapsedSeconds = ref(0)
  const isOnBreak = ref(false)
  const breakSecondsLeft = ref(0)
  let timerInterval: ReturnType<typeof setInterval> | null = null
  let breakInterval: ReturnType<typeof setInterval> | null = null

  const isRunning = computed(() => currentSession.value?.status === 'running')

  const timeRemaining = computed(() => {
    if (!currentSession.value) return 0
    return Math.max(0, currentSession.value.duration - elapsedSeconds.value)
  })

  const formattedTimeRemaining = computed(() => {
    const total = isOnBreak.value ? breakSecondsLeft.value : timeRemaining.value
    const minutes = Math.floor(total / 60)
    const seconds = total % 60
    return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`
  })

  const progress = computed(() => {
    if (!currentSession.value) return 0
    return (elapsedSeconds.value / currentSession.value.duration) * 100
  })

  const fetchSettings = async () => {
    try {
      const response = await pomodoroAPI.getSettings()
      settings.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch settings'
    }
  }

  const updateSettings = async (newSettings: Partial<PomodoroSettings>) => {
    loading.value = true
    try {
      const response = await pomodoroAPI.updateSettings(newSettings)
      settings.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to update settings'
    } finally {
      loading.value = false
    }
  }

  const fetchStats = async () => {
    try {
      const response = await pomodoroAPI.getStats()
      stats.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch stats'
    }
  }

  const fetchSessions = async (params?: { task_id?: number; date?: string }) => {
    loading.value = true
    try {
      const response = await pomodoroAPI.getSessions(params)
      sessions.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch sessions'
    } finally {
      loading.value = false
    }
  }

  const startPomodoro = async (taskId: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await pomodoroAPI.start(taskId, settings.value.work_duration)
      currentSession.value = response.data
      elapsedSeconds.value = 0
      isOnBreak.value = false

      if (timerInterval) clearInterval(timerInterval)
      timerInterval = setInterval(() => {
        elapsedSeconds.value++
        if (elapsedSeconds.value >= (currentSession.value?.duration || 0)) {
          completePomodoro()
        }
      }, 1000)

      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to start pomodoro'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const completePomodoro = async () => {
    if (!currentSession.value) return

    if (timerInterval) {
      clearInterval(timerInterval)
      timerInterval = null
    }

    try {
      await pomodoroAPI.complete(currentSession.value.id)
      stats.value.today_completed++
      stats.value.today_total_time += currentSession.value.duration

      // Start break
      const breakDuration = (stats.value.today_completed % settings.value.long_break_interval === 0)
        ? settings.value.long_break
        : settings.value.short_break

      isOnBreak.value = true
      breakSecondsLeft.value = breakDuration

      if (breakInterval) clearInterval(breakInterval)
      breakInterval = setInterval(() => {
        breakSecondsLeft.value--
        if (breakSecondsLeft.value <= 0) {
          endBreak()
        }
      }, 1000)

      currentSession.value = null
      elapsedSeconds.value = 0
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to complete pomodoro'
    }
  }

  const cancelPomodoro = async () => {
    if (!currentSession.value) return

    if (timerInterval) {
      clearInterval(timerInterval)
      timerInterval = null
    }

    try {
      await pomodoroAPI.cancel(currentSession.value.id)
      currentSession.value = null
      elapsedSeconds.value = 0
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to cancel pomodoro'
    }
  }

  const endBreak = () => {
    if (breakInterval) {
      clearInterval(breakInterval)
      breakInterval = null
    }
    isOnBreak.value = false
    breakSecondsLeft.value = 0
  }

  const skipBreak = () => {
    endBreak()
  }

  return {
    currentSession,
    settings,
    stats,
    sessions,
    loading,
    error,
    elapsedSeconds,
    isOnBreak,
    breakSecondsLeft,
    isRunning,
    timeRemaining,
    formattedTimeRemaining,
    progress,
    fetchSettings,
    updateSettings,
    fetchStats,
    fetchSessions,
    startPomodoro,
    completePomodoro,
    cancelPomodoro,
    skipBreak
  }
})
