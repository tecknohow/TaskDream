import {defineStore, acceptHMRUpdate} from 'pinia'
import {ref, computed} from 'vue'

import PomodoroService from '@/services/pomodoro'
import type {IPomodoroSession, IPomodoroSettings, IPomodoroStats} from '@/modelTypes/IPomodoro'

export const usePomodoroStore = defineStore('pomodoro', () => {
	const pomodoroService = new PomodoroService()

	const currentSession = ref<IPomodoroSession | null>(null)
	const settings = ref<IPomodoroSettings>({
		userId: 0,
		workDuration: 1500,
		shortBreak: 300,
		longBreak: 900,
		longBreakInterval: 4,
		autoStartBreaks: false,
		autoStartPomodoro: false,
		maxPermission: null,
	})
	const stats = ref<IPomodoroStats>({
		todayCompleted: 0,
		todayTotalTime: 0,
		weeklyCompleted: 0,
	})
	const sessions = ref<IPomodoroSession[]>([])
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

	async function fetchSettings() {
		try {
			const result = await pomodoroService.getSettings()
			settings.value = result
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to fetch settings'
			error.value = message
		}
	}

	async function updateSettings(newSettings: Partial<IPomodoroSettings>) {
		loading.value = true
		try {
			const result = await pomodoroService.updateSettings(newSettings)
			settings.value = result
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to update settings'
			error.value = message
		} finally {
			loading.value = false
		}
	}

	async function fetchStats() {
		try {
			const result = await pomodoroService.getStats()
			stats.value = result
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to fetch stats'
			error.value = message
		}
	}

	async function fetchSessions(params?: {taskId?: number; date?: string}) {
		loading.value = true
		try {
			const result = await pomodoroService.getAll(undefined, params || {})
			sessions.value = result
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to fetch sessions'
			error.value = message
		} finally {
			loading.value = false
		}
	}

	async function startPomodoro(taskId: number) {
		loading.value = true
		error.value = null

		try {
			const result = await pomodoroService.start(taskId, settings.value.workDuration)
			currentSession.value = result
			elapsedSeconds.value = 0
			isOnBreak.value = false

			if (timerInterval) clearInterval(timerInterval)
			timerInterval = setInterval(() => {
				elapsedSeconds.value++
				if (elapsedSeconds.value >= (currentSession.value?.duration || 0)) {
					completePomodoro()
				}
			}, 1000)

			return result
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to start pomodoro'
			error.value = message
			throw new Error(message)
		} finally {
			loading.value = false
		}
	}

	async function completePomodoro() {
		if (!currentSession.value) return

		if (timerInterval) {
			clearInterval(timerInterval)
			timerInterval = null
		}

		try {
			await pomodoroService.complete(currentSession.value.id)
			stats.value.todayCompleted++
			stats.value.todayTotalTime += currentSession.value.duration

			// Start break
			const breakDuration = (stats.value.todayCompleted % settings.value.longBreakInterval === 0)
				? settings.value.longBreak
				: settings.value.shortBreak

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
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to complete pomodoro'
			error.value = message
		}
	}

	async function cancelPomodoro() {
		if (!currentSession.value) return

		if (timerInterval) {
			clearInterval(timerInterval)
			timerInterval = null
		}

		try {
			await pomodoroService.cancel(currentSession.value.id)
			currentSession.value = null
			elapsedSeconds.value = 0
		} catch (err: unknown) {
			const message = err instanceof Error ? err.message : 'Failed to cancel pomodoro'
			error.value = message
		}
	}

	function endBreak() {
		if (breakInterval) {
			clearInterval(breakInterval)
			breakInterval = null
		}
		isOnBreak.value = false
		breakSecondsLeft.value = 0
	}

	function skipBreak() {
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
		skipBreak,
	}
})

// support hot reloading
if (import.meta.hot) {
	import.meta.hot.accept(acceptHMRUpdate(usePomodoroStore, import.meta.hot))
}
