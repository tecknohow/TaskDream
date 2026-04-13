<template>
  <div class="task-detail" v-if="task">
    <div class="task-detail-header">
      <h1>{{ task.title }}</h1>
      <button class="btn btn-ghost" @click="goBack">← Back</button>
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

        <div class="task-section">
          <h2>Activity</h2>
          <div class="activity-list">
            <div class="activity-item">
              <span class="activity-label">Created at</span>
              <span class="activity-value">{{ formatDate(task.createdAt) }}</span>
            </div>
            <div class="activity-item">
              <span class="activity-label">Updated at</span>
              <span class="activity-value">{{ formatDate(task.updatedAt) }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="task-sidebar">
        <div class="task-card">
          <h3>Status</h3>
          <select v-model="task.status" @change="updateTask" class="status-select">
            <option value="todo">To Do</option>
            <option value="in_progress">In Progress</option>
            <option value="in_review">In Review</option>
            <option value="done">Done</option>
          </select>
        </div>

        <div class="task-card">
          <h3>Priority</h3>
          <div class="priority-buttons">
            <button
              v-for="p in [0, 1, 2, 3]"
              :key="p"
              class="priority-btn"
              :class="{ active: task.priority === p }"
              @click="updatePriority(p)"
            >
              {{ priorityLabel(p) }}
            </button>
          </div>
        </div>

        <div class="task-card">
          <h3>Due Date</h3>
          <input v-model="dueDate" type="date" @change="updateDueDate" />
        </div>

        <div v-if="task.labels && task.labels.length > 0" class="task-card">
          <h3>Labels</h3>
          <div class="labels-list">
            <span v-for="label in task.labels" :key="label.id" class="label-badge">
              {{ label.name }}
            </span>
          </div>
        </div>

        <div class="task-card">
          <h3>Time Tracked</h3>
          <div class="time-tracked">
            {{ formatSeconds(task.timeTracked || 0) }}
          </div>
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
import { formatDate, formatSeconds } from '@/utils/format'
import type { Task } from '@/types/models'

const route = useRoute()
const router = useRouter()
const tasksStore = useTasksStore()

const task = ref<Task | null>(null)
const dueDate = computed({
  get: () => task.value?.dueDate ? new Date(task.value.dueDate).toISOString().split('T')[0] : '',
  set: (val: string) => {
    if (task.value) {
      task.value.dueDate = val ? new Date(val) : undefined
    }
  }
})

onMounted(async () => {
  const taskId = route.params.id as string
  await tasksStore.fetchById(taskId)
  task.value = tasksStore.selectedTask
})

const updateTask = async () => {
  if (!task.value) return
  await tasksStore.update(task.value.id, task.value)
}

const updatePriority = async (priority: number) => {
  if (!task.value) return
  task.value.priority = priority
  await updateTask()
}

const updateDueDate = async () => {
  if (!task.value) return
  await updateTask()
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

const priorityLabel = (p: number) => {
  const labels = ['Low', 'Medium', 'High', 'Urgent']
  return labels[p] || 'Unknown'
}
</script>

<style scoped lang="scss">
.task-detail {
  max-width: 1000px;
  margin: 0 auto;
}

.task-detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
  }
}

.task-detail-grid {
  display: grid;
  grid-template-columns: 1fr 300px;
  gap: var(--spacing-lg);
}

.task-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xl);
}

.task-section {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);

  h2 {
    margin-bottom: var(--spacing-md);
  }
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

.activity-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.activity-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm);
  background-color: var(--bg-secondary);
  border-radius: var(--radius-md);
}

.activity-label {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

.activity-value {
  font-size: var(--font-size-sm);
  font-weight: 500;
}

.task-sidebar {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.task-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);

  h3 {
    font-size: var(--font-size-sm);
    margin-bottom: var(--spacing-sm);
    text-transform: uppercase;
    color: var(--text-tertiary);
  }
}

.status-select {
  width: 100%;
  padding: var(--spacing-sm);
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
}

.priority-buttons {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-sm);
}

.priority-btn {
  padding: var(--spacing-sm);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  cursor: pointer;
  transition: all var(--transition-fast);
  font-size: var(--font-size-xs);
  font-weight: 500;

  &.active {
    background-color: var(--color-primary);
    color: white;
    border-color: var(--color-primary);
  }

  &:hover {
    border-color: var(--color-primary);
  }
}

.time-tracked {
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-primary);
}

.labels-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.label-badge {
  display: inline-flex;
  align-items: center;
  padding: var(--spacing-xs) var(--spacing-sm);
  background-color: var(--color-primary);
  color: white;
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
  width: fit-content;
}

.task-actions {
  margin-top: auto;
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}

.loading {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-secondary);
}

@media (max-width: 768px) {
  .task-detail-grid {
    grid-template-columns: 1fr;
  }
}
</style>
