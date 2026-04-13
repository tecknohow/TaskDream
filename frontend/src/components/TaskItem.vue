<template>
  <div class="task-item" @click="$emit('click')">
    <div class="task-checkbox">
      <input
        type="checkbox"
        :checked="task.status === 'done'"
        @change="updateStatus"
      />
    </div>

    <div class="task-content">
      <div class="task-title-row">
        <h3 class="task-title" :class="{ 'task-done': task.status === 'done' }">
          {{ task.title }}
        </h3>
        <div v-if="task.priority" class="priority-badge" :class="`priority-${task.priority}`">
          {{ priorityLabel(task.priority) }}
        </div>
      </div>

      <div v-if="task.description" class="task-description">
        {{ task.description.substring(0, 60) }}
      </div>

      <div class="task-meta">
        <div v-if="task.dueDate" class="due-date" :class="{ 'overdue': isOverdue(task.dueDate) }">
          📅 {{ formatDate(task.dueDate) }}
        </div>

        <div v-if="task.labels && task.labels.length > 0" class="task-labels">
          <span v-for="label in task.labels.slice(0, 2)" :key="label.id" class="label-dot">
            {{ label.name }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { formatDate } from '@/utils/format'
import { useTasksStore } from '@/stores/tasks'
import type { Task } from '@/types/models'

const tasksStore = useTasksStore()

defineProps<{
  task: Task
}>()

const emit = defineEmits<{
  click: []
}>()

const isOverdue = (dueDate: Date | string) => {
  const due = new Date(dueDate)
  due.setHours(0, 0, 0, 0)
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  return due < today
}

const updateStatus = async (event: Event) => {
  event.stopPropagation()
  const checkbox = event.target as HTMLInputElement
  const newStatus = checkbox.checked ? 'done' : 'todo'
  await tasksStore.updateStatus(tasksStore.selectedTask?.id || '', newStatus)
}

const priorityLabel = (p: number) => {
  const labels = ['Low', 'Medium', 'High', 'Urgent']
  return labels[p] || ''
}
</script>

<style scoped lang="scss">
.task-item {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    background-color: var(--bg-tertiary);
    border-color: var(--color-primary);
    box-shadow: var(--shadow-sm);
  }
}

.task-checkbox {
  display: flex;
  align-items: flex-start;
  padding-top: 4px;

  input[type='checkbox'] {
    cursor: pointer;
    width: 18px;
    height: 18px;
  }
}

.task-content {
  flex: 1;
  min-width: 0;
}

.task-title-row {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-xs);
}

.task-title {
  margin: 0;
  font-size: var(--font-size-base);
  font-weight: 500;
  color: var(--text-primary);
  flex: 1;

  &.task-done {
    text-decoration: line-through;
    color: var(--text-tertiary);
  }
}

.priority-badge {
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
  white-space: nowrap;
  flex-shrink: 0;

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

.task-description {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xs);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.task-meta {
  display: flex;
  gap: var(--spacing-md);
  flex-wrap: wrap;
  font-size: var(--font-size-xs);
}

.due-date {
  color: var(--text-secondary);

  &.overdue {
    color: var(--color-error);
    font-weight: 600;
  }
}

.task-labels {
  display: flex;
  gap: var(--spacing-xs);
}

.label-dot {
  display: inline-block;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: var(--color-primary);
}
</style>
