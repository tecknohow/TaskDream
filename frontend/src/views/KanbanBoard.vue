<template>
  <div class="kanban-board">
    <div class="kanban-header">
      <h1>Kanban Board</h1>
    </div>

    <div class="kanban-container">
      <div
        v-for="status in statuses"
        :key="status.value"
        class="kanban-column"
      >
        <div class="column-header">
          <h2>{{ status.label }}</h2>
          <span class="task-count">{{ getTasksByStatus(status.value).length }}</span>
        </div>

        <div class="tasks-container">
          <div
            v-for="task in getTasksByStatus(status.value)"
            :key="task.id"
            class="kanban-task"
            @click="selectTask(task)"
          >
            <div class="task-header">
              <input
                type="checkbox"
                :checked="task.status === 'done'"
                @change="updateStatus(task, 'done')"
                class="task-checkbox"
              />
              <h3>{{ task.title }}</h3>
            </div>

            <p v-if="task.description" class="task-description">
              {{ task.description.substring(0, 100) }}...
            </p>

            <div class="task-footer">
              <span v-if="task.priority" class="priority-badge" :class="`priority-${task.priority}`">
                {{ priorityLabel(task.priority) }}
              </span>
              <span v-if="task.dueDate" class="due-date">
                {{ formatDate(task.dueDate) }}
              </span>
            </div>
          </div>

          <div v-if="getTasksByStatus(status.value).length === 0" class="empty-column">
            No tasks
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import { formatDate } from '@/utils/format'
import type { Task, TaskStatus } from '@/types/models'

const tasksStore = useTasksStore()

const statuses = [
  { value: 'todo' as TaskStatus, label: 'To Do' },
  { value: 'in_progress' as TaskStatus, label: 'In Progress' },
  { value: 'in_review' as TaskStatus, label: 'In Review' },
  { value: 'done' as TaskStatus, label: 'Done' }
]

onMounted(() => {
  tasksStore.fetchAll()
})

const getTasksByStatus = (status: TaskStatus) => {
  return tasksStore.tasksByStatus(status)
}

const updateStatus = async (task: Task, status: TaskStatus) => {
  await tasksStore.updateStatus(task.id, status)
}

const selectTask = (task: Task) => {
  tasksStore.selectTask(task)
}

const priorityLabel = (p: number) => {
  const labels = ['Low', 'Medium', 'High', 'Urgent']
  return labels[p] || 'Unknown'
}
</script>

<style scoped lang="scss">
.kanban-board {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.kanban-header {
  margin-bottom: var(--spacing-lg);

  h1 {
    margin: 0;
  }
}

.kanban-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
  flex: 1;
  overflow-x: auto;
  padding: var(--spacing-md);
  background-color: var(--bg-secondary);
  border-radius: var(--radius-lg);
}

.kanban-column {
  background-color: var(--bg-primary);
  border-radius: var(--radius-lg);
  display: flex;
  flex-direction: column;
  min-height: 500px;
  box-shadow: var(--shadow-sm);
}

.column-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  border-bottom: 2px solid var(--border-color);

  h2 {
    margin: 0;
    font-size: var(--font-size-base);
  }
}

.task-count {
  background-color: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
}

.tasks-container {
  flex: 1;
  padding: var(--spacing-md);
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.kanban-task {
  background-color: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    box-shadow: var(--shadow-md);
    border-color: var(--color-primary);
  }
}

.task-header {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-sm);

  h3 {
    margin: 0;
    font-size: var(--font-size-base);
    flex: 1;
  }
}

.task-checkbox {
  margin-top: 2px;
  cursor: pointer;
}

.task-description {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-sm) 0;
  line-height: 1.4;
}

.task-footer {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
  font-size: var(--font-size-xs);
}

.priority-badge {
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-weight: 600;

  &.priority-0 {
    background-color: rgba(16, 185, 129, 0.1);
    color: var(--color-success);
  }

  &.priority-1 {
    background-color: rgba(245, 158, 11, 0.1);
    color: var(--color-warning);
  }

  &.priority-2 {
    background-color: rgba(239, 68, 68, 0.15);
    color: #dc2626;
  }

  &.priority-3 {
    background-color: rgba(239, 68, 68, 0.25);
    color: var(--color-error);
  }
}

.due-date {
  color: var(--text-secondary);
}

.empty-column {
  text-align: center;
  padding: var(--spacing-lg);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}
</style>
