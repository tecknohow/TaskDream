<template>
  <div class="task-item" :class="{ 'is-done': task.done }" @click="$emit('click', task)">
    <input
      type="checkbox"
      :checked="task.done"
      @click.stop
      @change="toggleDone"
      class="task-checkbox"
    />

    <div class="task-content">
      <div class="task-title">{{ task.title }}</div>
      <div class="task-meta" v-if="task.due_date || task.priority || task.estimated_time">
        <span
          v-if="task.priority"
          class="priority-badge"
          :class="`priority-${task.priority}`"
        >
          {{ priorityLabel }}
        </span>
        <span v-if="task.due_date" class="due-date" :class="{ overdue: isOverdue }">
          {{ formatDueDate(task.due_date) }}
        </span>
        <span v-if="task.estimated_time" class="estimate-badge">
          Est: {{ formatDuration(task.estimated_time) }}
        </span>
        <span v-if="task.total_time_spent" class="time-badge">
          {{ formatDuration(task.total_time_spent) }}
        </span>
      </div>
      <div class="task-labels" v-if="task.description">
        <span class="description-preview">{{ task.description.substring(0, 80) }}</span>
      </div>
    </div>

    <div class="task-indicators">
      <span v-if="task.urgency === 1 && task.importance === 1" class="eisenhower-badge do" title="Do First">DO</span>
      <span v-else-if="task.importance === 1" class="eisenhower-badge schedule" title="Schedule">PLAN</span>
      <span v-else-if="task.urgency === 1" class="eisenhower-badge delegate" title="Delegate">ASAP</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import type { Task } from '@/types/models'

const props = defineProps<{
  task: Task
}>()

defineEmits(['click'])

const tasksStore = useTasksStore()

const priorityLabel = computed(() => {
  const labels = ['Low', 'Medium', 'High', 'Urgent']
  return labels[props.task.priority] || ''
})

const isOverdue = computed(() => {
  if (!props.task.due_date) return false
  return new Date(props.task.due_date) < new Date() && !props.task.done
})

const toggleDone = () => {
  tasksStore.toggleDone(props.task.id)
}

const formatDueDate = (date: string) => {
  const d = new Date(date)
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  const taskDate = new Date(d)
  taskDate.setHours(0, 0, 0, 0)

  const diff = Math.ceil((taskDate.getTime() - today.getTime()) / (1000 * 60 * 60 * 24))

  if (diff === 0) return 'Today'
  if (diff === 1) return 'Tomorrow'
  if (diff === -1) return 'Yesterday'
  if (diff > 0 && diff <= 7) return `In ${diff} days`

  return d.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}

const formatDuration = (seconds: number) => {
  if (seconds < 60) return `${seconds}s`
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m`
  const hrs = Math.floor(mins / 60)
  const remainMins = mins % 60
  return remainMins > 0 ? `${hrs}h ${remainMins}m` : `${hrs}h`
}
</script>

<style scoped lang="scss">
.task-item {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    border-color: var(--color-primary);
    box-shadow: var(--shadow-sm);
  }

  &.is-done {
    opacity: 0.6;

    .task-title {
      text-decoration: line-through;
      color: var(--text-tertiary);
    }
  }
}

.task-checkbox {
  margin-top: 3px;
  cursor: pointer;
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

.task-content {
  flex: 1;
  min-width: 0;
}

.task-title {
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 2px;
}

.task-meta {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
  margin-top: var(--spacing-xs);
}

.priority-badge {
  padding: 1px 6px;
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;

  &.priority-0 { background: rgba(16, 185, 129, 0.1); color: var(--color-success); }
  &.priority-1 { background: rgba(245, 158, 11, 0.1); color: var(--color-warning); }
  &.priority-2 { background: rgba(239, 68, 68, 0.15); color: #dc2626; }
  &.priority-3 { background: rgba(239, 68, 68, 0.25); color: var(--color-error); }
}

.due-date {
  font-size: var(--font-size-xs);
  color: var(--text-secondary);

  &.overdue {
    color: var(--color-error);
    font-weight: 600;
  }
}

.estimate-badge, .time-badge {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  padding: 1px 4px;
  background: var(--bg-tertiary);
  border-radius: var(--radius-sm);
}

.description-preview {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  margin-top: 2px;
}

.task-indicators {
  flex-shrink: 0;
}

.eisenhower-badge {
  padding: 2px 6px;
  border-radius: var(--radius-sm);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.5px;

  &.do { background: rgba(239, 68, 68, 0.15); color: var(--color-error); }
  &.schedule { background: rgba(37, 99, 235, 0.1); color: var(--color-primary); }
  &.delegate { background: rgba(245, 158, 11, 0.1); color: var(--color-warning); }
}
</style>
