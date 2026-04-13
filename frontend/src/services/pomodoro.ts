import AbstractService from './abstractService'
import type {IPomodoroSession, IPomodoroSettings, IPomodoroStats} from '@/modelTypes/IPomodoro'
import {AuthenticatedHTTPFactory} from '@/helpers/fetcher'

export default class PomodoroService extends AbstractService<IPomodoroSession> {
	constructor() {
		super({
			create: '/pomodoro/start',
			getAll: '/pomodoro/sessions',
			get: '/pomodoro/sessions/{id}',
			update: '/pomodoro/sessions/{id}',
			delete: '/pomodoro/sessions/{id}',
		})
	}

	modelFactory(data: Partial<IPomodoroSession>) {
		return {
			...data,
			startedAt: data.startedAt ? new Date(data.startedAt) : null,
			completedAt: data.completedAt ? new Date(data.completedAt) : null,
			created: data.created ? new Date(data.created) : new Date(),
			updated: data.updated ? new Date(data.updated) : new Date(),
		} as IPomodoroSession
	}

	async start(taskId: number, duration: number): Promise<IPomodoroSession> {
		const cancel = this.setLoading()
		try {
			const response = await this.http.put('/pomodoro/start', {
				task_id: taskId,
				duration,
			})
			return this.modelFactory(response.data)
		} finally {
			cancel()
		}
	}

	async complete(sessionId: number): Promise<IPomodoroSession> {
		const cancel = this.setLoading()
		try {
			const response = await this.http.post(`/pomodoro/${sessionId}/complete`, {})
			return this.modelFactory(response.data)
		} finally {
			cancel()
		}
	}

	async cancel(sessionId: number): Promise<void> {
		const cancel = this.setLoading()
		try {
			await this.http.post(`/pomodoro/${sessionId}/cancel`, {})
		} finally {
			cancel()
		}
	}

	async getSettings(): Promise<IPomodoroSettings> {
		const cancel = this.setLoading()
		try {
			const response = await this.http.get('/pomodoro/settings')
			return response.data as IPomodoroSettings
		} finally {
			cancel()
		}
	}

	async updateSettings(settings: Partial<IPomodoroSettings>): Promise<IPomodoroSettings> {
		const cancel = this.setLoading()
		try {
			const response = await this.http.post('/pomodoro/settings', settings)
			return response.data as IPomodoroSettings
		} finally {
			cancel()
		}
	}

	async getStats(date?: string): Promise<IPomodoroStats> {
		const cancel = this.setLoading()
		try {
			const response = await this.http.get('/pomodoro/stats', {
				params: date ? {date} : undefined,
			})
			return response.data as IPomodoroStats
		} finally {
			cancel()
		}
	}
}
