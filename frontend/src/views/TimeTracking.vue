<template>
  <div class="time-tracking">
    <div class="tt-header">
      <h1>Time Tracking</h1>
    </div>

    <div class="tt-grid">
      <div class="tt-card timer-card">
        <h2>Active Timer</h2>
        <div v-if="timeTrackingStore.isTracking && timeTrackingStore.currentEntry" class="timer-display">
          <div class="timer-time">{{ timeTrackingStore.formattedTime }}</div>
          <div class="timer-task">
            Task: {{ currentTaskTitle }}
          </div>
          <button class="btn btn-danger" @click="stopTimer">
            Stop Timer
          </button>
        </div>
        <div v-else class="timer-inactive">
          <p>No active timer</p>
          <button class="btn btn-primary" @click="selectTaskForTracking = true">
            Start Timer
          </button>
        </div>
      </div>

      <div class="tt-card stats-card">
        <h2>Today's Summary</h2>
        <div class="stats-list">
          <div class="stat-row">
            <span>Total Time Tracked</span>
            <strong>{{ todayTotal }}</strong>
          </div>
          <div class="stat-row">
            <span>Active Sessions</span>
            <strong>{{ timeTrackingStore.entries.length }}</strong>
          </div>
        </div>
      </div>
    </div>

    <div class="tt-history">
      <h2>Recent Entries</h2>
      <div class="entries-table">
        <div class="table-header">
          <div class="col-time">Time</div>
          <div class="col-task">Task</div>
          <div class="col-duration">Duration</div>
          <div class="col-note">Note</div>
          <div class="col-action">Action</div>
        </div>

        <div
          v-for="entry in timeTrackingStore.entries"
          :key="entry.id"
          class="table-row"
        >
          <div class="col-time">
            {{ formatDateTime(entry.startTime) }}
          </div>
          <div class="col-task">
            Task {{ entry.taskId.substring(0, 8) }}...
          </div>
          <div class="col-duration">
            {{ formatSeconds(entry.duration) }}
          </div>
          <div class="col-note">
            {{ entry.note || '-' }}
          </div>
          <div class="col-action">
            <button class="btn btn-sm btn-ghost" @click="deleteEntry(entry.id)">
              Delete
            </button>
          </div>
        </div>

        <div v-if="timeTrackingStore.entries.length === 0" class="empty-message">
          No time tracking entries yet
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useTimeTrackingStore } from '@/stores/timeTracking'
import { useTasksStore } from '@/stores/tasks'
import { formatSeconds, formatDateTime } from '@/utils/format'

const timeTrackingStore = useTimeTrackingStore()
const tasksStore = useTasksStore()
const selectTaskForTracking = ref(false)

const currentTaskTitle = computed(() => {
  if (timeTrackingStore.currentEntry) {
    const task = tasksStore.tasks.find(t => t.id === timeTrackingStore.currentEntry!.taskId)
    return task?.title || 'Unknown'
  }
  return ''
})

const todayTotal = computed(() => {
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  const total = timeTrackingStore.entries
    .filter(e => {
      const entryDate = new Date(e.startTime)
      entryDate.setHours(0, 0, 0, 0)
      return entryDate.getTime() === today.getTime()
    })
    .reduce((sum, e) => sum + e.duration, 0)

  return formatSeconds(total)
})

onMounted(() => {
  timeTrackingStore.fetchEntries()
  tasksStore.fetchAll()
  timeTrackingStore.getCurrentTimer()
})

const stopTimer = async () => {
  if (timeTrackingStore.currentEntry) {
    await timeTrackingStore.stopTracking(timeTrackingStore.currentEntry.id)
  }
}

const deleteEntry = async (entryId: string) => {
  if (confirm('Delete this entry?')) {
    await timeTrackingStore.deleteEntry(entryId)
  }
}
</script>

<style scoped lang="scss">
.time-tracking {
  max-width: 1200px;
  margin: 0 auto;
}

.tt-header {
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
  }
}

.tt-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.tt-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);

  h2 {
    margin-top: 0;
    margin-bottom: var(--spacing-lg);
  }
}

.timer-card {
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-secondary) 100%);
  color: white;
  border: none;

  h2 {
    color: rgba(255, 255, 255, 0.9);
  }
}

.timer-display {
  text-align: center;
}

.timer-time {
  font-size: 48px;
  font-weight: 700;
  font-family: 'Courier New', monospace;
  margin-bottom: var(--spacing-md);
}

.timer-task {
  font-size: var(--font-size-sm);
  margin-bottom: var(--spacing-lg);
  opacity: 0.9;
}

.timer-inactive {
  text-align: center;

  p {
    margin-bottom: var(--spacing-md);
  }
}

.stats-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.stat-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  background-color: var(--bg-secondary);
  border-radius: var(--radius-md);

  span {
    color: var(--text-secondary);
    font-size: var(--font-size-sm);
  }

  strong {
    font-size: var(--font-size-lg);
    color: var(--color-primary);
  }
}

.tt-history {
  h2 {
    margin-bottom: var(--spacing-lg);
  }
}

.entries-table {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.table-header {
  display: grid;
  grid-template-columns: 150px 1fr 120px 150px 100px;
  gap: var(--spacing-md);
  padding: var(--spacing-md) var(--spacing-lg);
  background-color: var(--bg-secondary);
  border-bottom: 1px solid var(--border-color);
  font-weight: 600;
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

.table-row {
  display: grid;
  grid-template-columns: 150px 1fr 120px 150px 100px;
  gap: var(--spacing-md);
  padding: var(--spacing-md) var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
  align-items: center;

  &:last-child {
    border-bottom: none;
  }

  &:hover {
    background-color: var(--bg-secondary);
  }
}

.col-time,
.col-task,
.col-duration,
.col-note {
  font-size: var(--font-size-sm);
}

.col-action {
  text-align: right;
}

.empty-message {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}

@media (max-width: 768px) {
  .table-header,
  .table-row {
    grid-template-columns: 1fr;
  }

  .col-action {
    text-align: left;
  }
}
</style>
