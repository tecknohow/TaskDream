<template>
	<div class="pomodoro-timer" :class="{ 'is-running': pomodoroStore.isRunning, 'is-break': pomodoroStore.isOnBreak }">
		<div class="timer-circle">
			<svg class="progress-ring" viewBox="0 0 120 120">
				<circle class="progress-ring-bg" cx="60" cy="60" r="54" />
				<circle
					class="progress-ring-fill"
					cx="60" cy="60" r="54"
					:stroke-dasharray="circumference"
					:stroke-dashoffset="strokeDashoffset"
				/>
			</svg>
			<div class="timer-display">
				<span class="timer-time">{{ pomodoroStore.formattedTimeRemaining }}</span>
				<span class="timer-label">
					{{ pomodoroStore.isOnBreak ? 'Break' : pomodoroStore.isRunning ? 'Focus' : 'Ready' }}
				</span>
			</div>
		</div>

		<div class="timer-controls">
			<template v-if="!pomodoroStore.isRunning && !pomodoroStore.isOnBreak">
				<button class="btn btn-primary btn-lg" @click="$emit('start')" :disabled="!taskId">
					Start Focus
				</button>
			</template>
			<template v-else-if="pomodoroStore.isRunning">
				<button class="btn btn-danger" @click="pomodoroStore.cancelPomodoro()">
					Cancel
				</button>
			</template>
			<template v-else-if="pomodoroStore.isOnBreak">
				<button class="btn btn-primary" @click="pomodoroStore.skipBreak()">
					Skip Break
				</button>
			</template>
		</div>

		<div class="timer-stats">
			<div class="stat-item">
				<span class="stat-value">{{ pomodoroStore.stats.todayCompleted }}</span>
				<span class="stat-label">Today</span>
			</div>
			<div class="stat-item">
				<span class="stat-value">{{ formatMinutes(pomodoroStore.stats.todayTotalTime) }}</span>
				<span class="stat-label">Focus Time</span>
			</div>
			<div class="stat-item">
				<span class="stat-value">{{ pomodoroStore.stats.weeklyCompleted }}</span>
				<span class="stat-label">This Week</span>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import {computed} from 'vue'
import {usePomodoroStore} from '@/stores/pomodoro'

defineProps<{
	taskId?: number
}>()

defineEmits(['start'])

const pomodoroStore = usePomodoroStore()

const circumference = 2 * Math.PI * 54
const strokeDashoffset = computed(() => {
	return circumference - (pomodoroStore.progress / 100) * circumference
})

function formatMinutes(seconds: number): string {
	const mins = Math.floor(seconds / 60)
	if (mins < 60) return `${mins}m`
	const hrs = Math.floor(mins / 60)
	const remainMins = mins % 60
	return `${hrs}h ${remainMins}m`
}
</script>

<style scoped lang="scss">
.pomodoro-timer {
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 1.5rem;
	padding: 2rem;
}

.timer-circle {
	position: relative;
	width: 200px;
	height: 200px;
}

.progress-ring {
	width: 100%;
	height: 100%;
	transform: rotate(-90deg);
}

.progress-ring-bg {
	fill: none;
	stroke: var(--grey-200);
	stroke-width: 6;
}

.progress-ring-fill {
	fill: none;
	stroke: var(--primary);
	stroke-width: 6;
	stroke-linecap: round;
	transition: stroke-dashoffset 0.5s ease;

	.is-break & {
		stroke: var(--success);
	}
}

.timer-display {
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	text-align: center;
}

.timer-time {
	font-size: 2.5rem;
	font-weight: 700;
	color: var(--text);
	font-variant-numeric: tabular-nums;
}

.timer-label {
	display: block;
	font-size: .875rem;
	color: var(--grey-500);
	text-transform: uppercase;
	letter-spacing: 1px;
	margin-top: .25rem;
}

.timer-controls {
	display: flex;
	gap: .75rem;
}

.timer-stats {
	display: flex;
	gap: 2rem;
	padding-top: .75rem;
	border-top: 1px solid var(--grey-200);
}

.stat-item {
	text-align: center;
}

.stat-value {
	display: block;
	font-size: 1.25rem;
	font-weight: 700;
	color: var(--primary);
}

.stat-label {
	font-size: .75rem;
	color: var(--grey-500);
	text-transform: uppercase;
}
</style>
