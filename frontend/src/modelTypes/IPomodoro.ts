import type {IAbstract} from './IAbstract'
import type {ITask} from './ITask'
import type {IUser} from './IUser'

export interface IPomodoroSession extends IAbstract {
	id: number
	taskId: ITask['id']
	userId: IUser['id']
	duration: number
	status: 'running' | 'completed' | 'cancelled'
	startedAt: Date | null
	completedAt: Date | null
	created: Date
	updated: Date
}

export interface IPomodoroSettings extends IAbstract {
	userId: IUser['id']
	workDuration: number
	shortBreak: number
	longBreak: number
	longBreakInterval: number
	autoStartBreaks: boolean
	autoStartPomodoro: boolean
}

export interface IPomodoroStats {
	todayCompleted: number
	todayTotalTime: number
	weeklyCompleted: number
}
