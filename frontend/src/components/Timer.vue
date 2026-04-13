<template>
  <div class="timer">
    <div class="timer-display" v-if="timeTrackingStore.isTracking">
      <div class="timer-time">{{ timeTrackingStore.formattedTime }}</div>
      <p class="timer-label">Active Timer</p>
      <button class="btn btn-danger btn-sm" @click="stop">Stop</button>
    </div>
    <div v-else class="timer-inactive">
      <p>No active timer</p>
      <button class="btn btn-primary btn-sm" @click="start">Start Timer</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTimeTrackingStore } from '@/stores/timeTracking'

const timeTrackingStore = useTimeTrackingStore()

const start = () => {
  // Implementation for starting timer
}

const stop = async () => {
  if (timeTrackingStore.currentEntry) {
    await timeTrackingStore.stopTracking(timeTrackingStore.currentEntry.id)
  }
}
</script>

<style scoped lang="scss">
.timer {
  padding: var(--spacing-md);
  background-color: var(--bg-secondary);
  border-radius: var(--radius-md);
  text-align: center;
}

.timer-display {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  align-items: center;
}

.timer-time {
  font-size: 32px;
  font-weight: 700;
  font-family: 'Courier New', monospace;
  color: var(--color-primary);
}

.timer-label {
  margin: 0;
  color: var(--text-secondary);
  font-size: var(--font-size-sm);
}

.timer-inactive {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  align-items: center;

  p {
    margin: 0;
    color: var(--text-tertiary);
    font-size: var(--font-size-sm);
  }
}
</style>
