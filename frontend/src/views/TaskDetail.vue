<template>
  <div class="task-detail" v-if="task">
    <div class="task-detail-header">
      <button class="btn btn-ghost" @click="goBack">&larr; Back</button>
      <h1>{{ task.title }}</h1>
    </div>

    <div class="task-detail-grid">
      <div class="task-content">
        <div class="task-section">
          <h2>Description</h2>
          <div v-if="task.description" class="task-description">
            {{ task.description }}
          </div>
          <div v-else class="empty-text">No description</div>
        </div>

        <!-- Subtasks -->
        <div class="task-section">
          <div class="section-header">
            <h2>Subtasks</h2>
            <button class="btn btn-ghost btn-sm" @click="showAddSubtask = !showAddSubtask">+ Add</button>
          </div>

          <div v-if="showAddSubtask" class="add-subtask-form">
            <input
              v-model="newSubtaskTitle"
              type="text"
              placeholder="Subtask title"
              @keyup.enter="addSubtask"
              autofocus
            />
            <button class="btn btn-primary btn-sm" @click="addSubtask" :disabled="!newSubtaskTitle.trim()">Add</button>
          </div>

          <div class="subtask-list" v-if="subtasks.length > 0">
            <div v-for="st in subtasks" :key="st.id" class="subtask-item">
              <input type="checkbox" :checked="st.done" @change="toggleSubtask(st)" />
              <span :class="{ done: st.done }">{{ st.title }}</span>
            </div>
          </div>
          <div v-else class="empty-text">No subtasks</div>
        </div>

        <!-- Time Tracking Entries -->
        <div class="task-section">
          <h2>Time Entries</h2>
          <div class="time-entries" v-if="timeEntries.length > 0">
            <div v-for="entry in timeEntries" :key="entry.id" class="time-entry">
              <span>{{ formatDuration(entry.duration) }}</span>
              <span class="text-xs text-tertiary">{{ entry.comment || 'No note' }}</span>
              <span class="text-xs text-tertiary">{{ formatDateTime(entry.start) }}</span>
            </div>
          </div>
          <div v-else class="empty-text">No time entries</div>
        </div>

        <div class="task-section">
          <h2>Activity</h2>
          <div class="activity-list">
            <div class="activity-item">
              <span class="activity-label">Created</span>
              <span class="activity-value">{{ formatDateTime(task.created) }}</span>
            </div>
            <div class="activity-item">
              <span class="activity-label">Updated</span>
              <span class="activity-value">{{ formatDateTime(task.updated) }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="task-sidebar">
        <div class="task-card">
          <h3>Status</h3>
          <div class="status-toggle">
            <button
              class="btn btn-sm"
              :class="task.done ? 'btn-success' : 'btn-secondary'"
              @click="toggleDone"
            >
              {{ task.done ? 'Completed' : 'Mark Done' }}
            </button>
          </div>
        </div>

        <div class="task-card">
          <h3>Priority</h3>
          <div class="priority-buttons">
            <button
              v-for="p in [0, 1, 2, 3]"
              :key="p"
              class="priority-btn"
              :class="{ active: task.priority === p, [`p-${p}`]: true }"
              @click="updateField('priority', p)"
            >
              {{ ['Low', 'Med', 'High', 'Urgent'][p] }}
            </button>
          </div>
        </div>

        <div class="task-card">
          <h3>Eisenhower</h3>
          <div class="eisenhower-toggles">
            <label class="toggle-label">
              <input type="checkbox" :checked="task.urgency === 1" @change="updateField('urgency', task.urgency === 1 ? 0 : 1)" />
              <span>Urgent</span>
            </label>
            <label class="toggle-label">
              <input type="checkbox" :checked="task.importance === 1" @change="updateField('importance', task.importance === 1 ? 0 : 1)" />
              <span>Important</span>
            </label>
          </div>
        </div>

        <div class="task-card">
          <h3>Due Date</h3>
          <input :value="formattedDueDate" type="date" @change="updateDueDate($event)" />
        </div>

        <div class="task-card">
          <h3>Estimation</h3>
          <div class="estimate-row">
            <div>
              <span class="text-xs text-tertiary">Estimated</span>
              <div class="estimate-value">{{ formatDuration(task.estimated_time || 0) }}</div>
            </div>
            <div>
              <span class="text-xs text-tertiary">Actual</span>
              <div class="estimate-value">{{ formatDuration(task.total_time_spent || 0) }}</div>
            </div>
          </div>
          <div class="estimate-input">
            <input
              v-model.number="estimateMinutes"
              type="number"
              min="0"
              placeholder="Minutes"
              class="estimate-field"
            />
            <button class="btn btn-sm btn-ghost" @click="saveEstimate">Set</button>
          </div>
        </div>

        <div class="task-card">
          <h3>Time Tracking</h3>
          <div class="tracking-controls">
            <button
              v-if="!timeTrackingStore.isTracking || timeTrackingStore.currentTaskId !== task.id"
              class="btn btn-primary btn-sm"
              @click="startTracking"
            >
              Start Timer
            </button>
            <div v-else class="active-timer">
              <span class="timer-display">{{ timeTrackingStore.formattedTime }}</span>
              <button class="btn btn-danger btn-sm" @click="stopTracking">Stop</button>
            </div>
          </div>
        </div>

        <div class="task-card">
          <h3>Pomodoro</h3>
          <button
            class="btn btn-secondary btn-sm"
            @click="startFocus"
          >
            Start Focus Session
          </button>
        </div>

        <div class="task-actions">
          <button class="btn btn-danger btn-sm" @click="deleteTask">Delete Task</button>
        </div>
      </div>
    </div>
  </div>

  <div v-else class="loading">
    Loading task...
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useTasksStore } from '@/stores/tasks'
import { useTimeTrackingStore } from '@/stores/timeTracking'
import { tasksAPI } from '@/api/tasks'
import type { Task, TimeTrackingEntry } from '@/types/models'

const route = useRoute()
const router = useRouter()
const tasksStore = useTasksStore()
const timeTrackingStore = useTimeTrackingStore()

const task = ref<Task | null>(null)
const subtasks = ref<Task[]>([])
const timeEntries = ref<TimeTrackingEntry[]>([])
const showAddSubtask = ref(false)
const newSubtaskTitle = ref('')
const estimateMinutes = ref(0)

const formattedDueDate = computed(() => {
  if (!task.value?.due_date) return ''
  return new Date(task.value.due_date).toISOString().split('T')[0]
})

onMounted(async () => {
  const taskId = Number(route.params.id)
  const result = await tasksStore.fetchById(taskId)
  if (result) {
    task.value = result.task
    subtasks.value = result.subtasks || []
    estimateMinutes.value = Math.floor((task.value.estimated_time || 0) / 60)
  }

  // Fetch time entries
  try {
    const resp = await tasksAPI.getTimeTracking(taskId)
    timeEntries.value = resp.data
  } catch {
    timeEntries.value = []
  }
})

const updateField = async (field: string, value: any) => {
  if (!task.value) return
  ;(task.value as any)[field] = value
  await tasksStore.update(task.value.id, { [field]: value } as any)
}

const toggleDone = async () => {
  if (!task.value) return
  task.value.done = !task.value.done
  await tasksStore.update(task.value.id, { done: task.value.done })
}

const updateDueDate = async (event: Event) => {
  if (!task.value) return
  const val = (event.target as HTMLInputElement).value
  task.value.due_date = val || undefined
  await tasksStore.update(task.value.id, { due_date: val || undefined } as any)
}

const saveEstimate = async () => {
  if (!task.value) return
  const seconds = estimateMinutes.value * 60
  task.value.estimated_time = seconds
  await tasksStore.update(task.value.id, { estimated_time: seconds })
}

const addSubtask = async () => {
  if (!task.value || !newSubtaskTitle.value.trim()) return
  try {
    const resp = await tasksAPI.createSubtask(task.value.id, {
      title: newSubtaskTitle.value.trim()
    })
    subtasks.value.push(resp.data)
    newSubtaskTitle.value = ''
  } catch {
    // error
  }
}

const toggleSubtask = async (st: Task) => {
  await tasksStore.update(st.id, { done: !st.done })
  st.done = !st.done
}

const startTracking = () => {
  if (task.value) {
    timeTrackingStore.startTracking(task.value.id)
  }
}

const stopTracking = async () => {
  await timeTrackingStore.stopTracking()
}

const startFocus = () => {
  router.push('/focus')
}

const deleteTask = async () => {
  if (!task.value) return
  if (confirm('Are you sure you want to delete this task?')) {
    await tasksStore.deleteTask(task.value.id)
    router.back()
  }
}

const goBack = () => {
  router.back()
}

const formatDuration = (seconds: number) => {
  if (!seconds) return '0m'
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m`
  const hrs = Math.floor(mins / 60)
  const remainMins = mins % 60
  return remainMins > 0 ? `${hrs}h ${remainMins}m` : `${hrs}h`
}

const formatDateTime = (dateStr: string) => {
  if (!dateStr) return 'N/A'
  return new Date(dateStr).toLocaleDateString('en-US', {
    month: 'short', day: 'numeric', year: 'numeric',
    hour: '2-digit', minute: '2-digit'
  })
}
</script>

<style scoped lang="scss">
.task-detail {
  max-width: 1000px;
  margin: 0 auto;
}

.task-detail-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);

  h1 { margin: 0; flex: 1; }
}

.task-detail-grid {
  display: grid;
  grid-template-columns: 1fr 300px;
  gap: var(--spacing-lg);
}

.task-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.task-section {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);

  h2 { margin-bottom: var(--spacing-md); font-size: var(--font-size-base); }
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-md);

  h2 { margin: 0; }
}

.task-description {
  color: var(--text-secondary);
  line-height: 1.6;
  white-space: pre-wrap;
}

.empty-text {
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}

.add-subtask-form {
  display: flex;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-md);

  input { flex: 1; }
}

.subtask-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.subtask-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-xs) 0;
  font-size: var(--font-size-sm);

  .done { text-decoration: line-through; color: var(--text-tertiary); }
}

.time-entries {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.time-entry {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm);
  background: var(--bg-secondary);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.activity-item {
  display: flex;
  justify-content: space-between;
  padding: var(--spacing-sm);
  background: var(--bg-secondary);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
}

.activity-label { color: var(--text-secondary); }
.activity-value { font-weight: 500; }

.task-sidebar {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.task-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);

  h3 {
    font-size: var(--font-size-xs);
    margin-bottom: var(--spacing-sm);
    text-transform: uppercase;
    color: var(--text-tertiary);
    letter-spacing: 0.5px;
  }

  input[type="date"] { width: 100%; }
}

.priority-buttons {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-xs);
}

.priority-btn {
  padding: var(--spacing-xs) var(--spacing-sm);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  background: var(--bg-secondary);
  cursor: pointer;
  font-size: var(--font-size-xs);
  font-weight: 500;
  transition: all var(--transition-fast);

  &.active { background: var(--color-primary); color: white; border-color: var(--color-primary); }
  &:hover:not(.active) { border-color: var(--color-primary); }
}

.eisenhower-toggles {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.toggle-label {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  font-size: var(--font-size-sm);
  cursor: pointer;
}

.estimate-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-sm);
}

.estimate-value {
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-primary);
}

.estimate-input {
  display: flex;
  gap: var(--spacing-sm);

  .estimate-field {
    flex: 1;
    padding: var(--spacing-xs) var(--spacing-sm);
  }
}

.tracking-controls {
  display: flex;
  align-items: center;
}

.active-timer {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  width: 100%;
}

.timer-display {
  font-family: monospace;
  font-weight: 600;
  font-size: var(--font-size-lg);
  color: var(--color-error);
  flex: 1;
}

.task-actions {
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}

.loading {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-secondary);
}

@media (max-width: 768px) {
  .task-detail-grid { grid-template-columns: 1fr; }
}
</style>
