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
        <span class="stat-value">{{ pomodoroStore.stats.today_completed }}</span>
        <span class="stat-label">Today</span>
      </div>
      <div class="stat-item">
        <span class="stat-value">{{ formatMinutes(pomodoroStore.stats.today_total_time) }}</span>
        <span class="stat-label">Focus Time</span>
      </div>
      <div class="stat-item">
        <span class="stat-value">{{ pomodoroStore.stats.weekly_completed }}</span>
        <span class="stat-label">This Week</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { usePomodoroStore } from '@/stores/pomodoro'

const props = defineProps<{
  taskId?: number
}>()

defineEmits(['start'])

const pomodoroStore = usePomodoroStore()

const circumference = 2 * Math.PI * 54
const strokeDashoffset = computed(() => {
  return circumference - (pomodoroStore.progress / 100) * circumference
})

const formatMinutes = (seconds: number) => {
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
  gap: var(--spacing-lg);
  padding: var(--spacing-xl);
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
  stroke: var(--border-color);
  stroke-width: 6;
}

.progress-ring-fill {
  fill: none;
  stroke: var(--color-primary);
  stroke-width: 6;
  stroke-linecap: round;
  transition: stroke-dashoffset 0.5s ease;

  .is-break & {
    stroke: var(--color-success);
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
  color: var(--text-primary);
  font-variant-numeric: tabular-nums;
}

.timer-label {
  display: block;
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-top: var(--spacing-xs);
}

.timer-controls {
  display: flex;
  gap: var(--spacing-md);
}

.timer-stats {
  display: flex;
  gap: var(--spacing-xl);
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}

.stat-item {
  text-align: center;
}

.stat-value {
  display: block;
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--color-primary);
}

.stat-label {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  text-transform: uppercase;
}
</style>
