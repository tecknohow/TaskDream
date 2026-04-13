<template>
  <div class="focus-mode" :class="{ 'is-active': pomodoroStore.isRunning || pomodoroStore.isOnBreak }">
    <div class="focus-container">
      <!-- Task selector when no task selected -->
      <div v-if="!selectedTask" class="task-selector">
        <h1>Focus Mode</h1>
        <p class="text-secondary">Select a task to start a focused work session</p>

        <div class="task-search">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search tasks..."
            class="search-input"
          />
        </div>

        <div class="task-list">
          <div
            v-for="task in filteredTasks"
            :key="task.id"
            class="focus-task-item"
            @click="selectTask(task)"
          >
            <div class="task-info">
              <div class="task-title">{{ task.title }}</div>
              <div class="task-meta-row">
                <span v-if="task.estimated_time" class="meta-badge">
                  Est: {{ formatDuration(task.estimated_time) }}
                </span>
                <span v-if="task.priority" class="meta-badge" :class="`priority-${task.priority}`">
                  {{ ['Low', 'Medium', 'High', 'Urgent'][task.priority] }}
                </span>
              </div>
            </div>
            <button class="btn btn-primary btn-sm">Focus</button>
          </div>

          <div v-if="filteredTasks.length === 0" class="empty-state">
            {{ searchQuery ? 'No matching tasks' : 'No open tasks available' }}
          </div>
        </div>
      </div>

      <!-- Active focus session -->
      <div v-else class="focus-session">
        <button class="btn btn-ghost back-btn" @click="deselectTask">
          &larr; Back to task list
        </button>

        <div class="focus-task-header">
          <h2>{{ selectedTask.title }}</h2>
          <p v-if="selectedTask.description" class="task-description">
            {{ selectedTask.description }}
          </p>
        </div>

        <PomodoroTimer
          :taskId="selectedTask.id"
          @start="startFocus"
        />

        <!-- Subtasks -->
        <div v-if="subtasks.length > 0" class="subtasks-section">
          <h3>Subtasks</h3>
          <div class="subtask-list">
            <div v-for="st in subtasks" :key="st.id" class="subtask-item">
              <input
                type="checkbox"
                :checked="st.done"
                @change="toggleSubtask(st)"
              />
              <span :class="{ done: st.done }">{{ st.title }}</span>
            </div>
          </div>
        </div>

        <!-- Time tracking info -->
        <div class="focus-stats-panel">
          <div class="stat-card">
            <div class="stat-label">Estimated</div>
            <div class="stat-value">
              {{ selectedTask.estimated_time ? formatDuration(selectedTask.estimated_time) : 'N/A' }}
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-label">Time Spent</div>
            <div class="stat-value">
              {{ formatDuration(selectedTask.total_time_spent || 0) }}
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-label">Progress</div>
            <div class="stat-value">{{ selectedTask.percent_done }}%</div>
          </div>
        </div>

        <div class="focus-actions">
          <button class="btn btn-primary" @click="markDone" v-if="!selectedTask.done">
            Mark as Done
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import { usePomodoroStore } from '@/stores/pomodoro'
import PomodoroTimer from '@/components/PomodoroTimer.vue'
import type { Task } from '@/types/models'

const tasksStore = useTasksStore()
const pomodoroStore = usePomodoroStore()

const selectedTask = ref<Task | null>(null)
const subtasks = ref<Task[]>([])
const searchQuery = ref('')

const filteredTasks = computed(() => {
  const tasks = tasksStore.openTasks
  if (!searchQuery.value) return tasks
  const q = searchQuery.value.toLowerCase()
  return tasks.filter(t =>
    t.title.toLowerCase().includes(q) ||
    (t.description && t.description.toLowerCase().includes(q))
  )
})

onMounted(async () => {
  await tasksStore.fetchAll({ status: 'undone' })
  await pomodoroStore.fetchSettings()
  await pomodoroStore.fetchStats()
})

const selectTask = async (task: Task) => {
  selectedTask.value = task
  // Fetch subtasks
  try {
    const resp = await tasksStore.fetchById(task.id)
    if (resp) {
      subtasks.value = resp.subtasks || []
    }
  } catch {
    subtasks.value = []
  }
}

const deselectTask = () => {
  if (pomodoroStore.isRunning) {
    pomodoroStore.cancelPomodoro()
  }
  selectedTask.value = null
  subtasks.value = []
}

const startFocus = () => {
  if (selectedTask.value) {
    pomodoroStore.startPomodoro(selectedTask.value.id)
  }
}

const toggleSubtask = async (st: Task) => {
  await tasksStore.update(st.id, { done: !st.done })
  st.done = !st.done
}

const markDone = async () => {
  if (selectedTask.value) {
    await tasksStore.toggleDone(selectedTask.value.id)
    selectedTask.value = null
  }
}

const formatDuration = (seconds: number) => {
  if (!seconds) return '0m'
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m`
  const hrs = Math.floor(mins / 60)
  const remainMins = mins % 60
  return remainMins > 0 ? `${hrs}h ${remainMins}m` : `${hrs}h`
}
</script>

<style scoped lang="scss">
.focus-mode {
  max-width: 700px;
  margin: 0 auto;
  min-height: 80vh;

  &.is-active {
    .focus-container {
      background: linear-gradient(135deg, var(--bg-primary) 0%, var(--bg-secondary) 100%);
    }
  }
}

.focus-container {
  background-color: var(--bg-primary);
  border-radius: var(--radius-xl);
  padding: var(--spacing-xl);
  box-shadow: var(--shadow-md);
}

.task-selector {
  h1 {
    text-align: center;
    margin-bottom: var(--spacing-xs);
  }

  p {
    text-align: center;
    margin-bottom: var(--spacing-xl);
  }
}

.search-input {
  width: 100%;
  padding: var(--spacing-md);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  font-size: var(--font-size-base);
  margin-bottom: var(--spacing-lg);
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  max-height: 500px;
  overflow-y: auto;
}

.focus-task-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--spacing-md);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    border-color: var(--color-primary);
    box-shadow: var(--shadow-sm);
  }
}

.task-info {
  flex: 1;
}

.task-title {
  font-weight: 500;
  margin-bottom: 2px;
}

.task-meta-row {
  display: flex;
  gap: var(--spacing-sm);
}

.meta-badge {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  padding: 1px 4px;
  background: var(--bg-tertiary);
  border-radius: var(--radius-sm);

  &.priority-1 { color: var(--color-warning); }
  &.priority-2 { color: #dc2626; }
  &.priority-3 { color: var(--color-error); }
}

.back-btn {
  margin-bottom: var(--spacing-lg);
}

.focus-task-header {
  text-align: center;
  margin-bottom: var(--spacing-lg);

  h2 { margin-bottom: var(--spacing-sm); }

  .task-description {
    color: var(--text-secondary);
    font-size: var(--font-size-sm);
  }
}

.subtasks-section {
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--border-color);

  h3 {
    margin-bottom: var(--spacing-md);
    font-size: var(--font-size-base);
  }
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
  padding: var(--spacing-sm);
  font-size: var(--font-size-sm);

  .done {
    text-decoration: line-through;
    color: var(--text-tertiary);
  }
}

.focus-stats-panel {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: var(--spacing-md);
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--border-color);
}

.stat-card {
  text-align: center;
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--radius-md);
}

.stat-label {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  text-transform: uppercase;
  margin-bottom: var(--spacing-xs);
}

.stat-value {
  font-size: var(--font-size-lg);
  font-weight: 700;
  color: var(--color-primary);
}

.focus-actions {
  display: flex;
  justify-content: center;
  margin-top: var(--spacing-lg);
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-tertiary);
}
</style>
