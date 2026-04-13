<template>
  <div class="kanban-board">
    <div class="kanban-header">
      <h1>Kanban Board</h1>
      <div class="header-actions">
        <select v-model="selectedProjectId" class="project-filter">
          <option :value="0">All Projects</option>
          <option
            v-for="project in projectsStore.activeProjects"
            :key="project.id"
            :value="project.id"
          >
            {{ project.title }}
          </option>
        </select>
        <button class="btn btn-primary" @click="showCreateTask = true">+ New Task</button>
      </div>
    </div>

    <div class="kanban-container">
      <div
        v-for="column in columns"
        :key="column.status"
        class="kanban-column"
        @dragover.prevent
        @drop="onDrop($event, column.status)"
      >
        <div class="column-header" :style="{ borderBottomColor: column.color }">
          <div class="column-title">
            <span class="column-dot" :style="{ backgroundColor: column.color }"></span>
            <h2>{{ column.label }}</h2>
          </div>
          <span class="task-count">{{ getColumnTasks(column.status).length }}</span>
        </div>

        <div class="tasks-container">
          <div
            v-for="task in getColumnTasks(column.status)"
            :key="task.id"
            class="kanban-task"
            draggable="true"
            @dragstart="onDragStart($event, task)"
            @click="navigateToTask(task)"
          >
            <div class="task-top">
              <input
                type="checkbox"
                :checked="task.done"
                @click.stop
                @change="toggleTaskDone(task)"
                class="task-checkbox"
              />
              <h3>{{ task.title }}</h3>
            </div>

            <p v-if="task.description" class="task-description">
              {{ task.description.substring(0, 80) }}
            </p>

            <div class="task-bottom">
              <span
                v-if="task.priority"
                class="priority-badge"
                :class="`priority-${task.priority}`"
              >
                {{ ['Low', 'Med', 'High', '!!!'][task.priority] }}
              </span>

              <span v-if="task.due_date" class="due-date" :class="{ overdue: isOverdue(task) }">
                {{ formatDueDate(task.due_date) }}
              </span>

              <span v-if="task.estimated_time" class="estimate-badge">
                {{ formatDuration(task.estimated_time) }}
              </span>

              <span
                v-if="task.urgency === 1 && task.importance === 1"
                class="eisenhower-mini do"
              >DO</span>
            </div>
          </div>

          <div v-if="getColumnTasks(column.status).length === 0" class="empty-column">
            Drop tasks here
          </div>
        </div>
      </div>
    </div>

    <CreateTaskDialog v-if="showCreateTask" @close="showCreateTask = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useTasksStore } from '@/stores/tasks'
import { useProjectsStore } from '@/stores/projects'
import CreateTaskDialog from '@/components/CreateTaskDialog.vue'
import type { Task } from '@/types/models'

const route = useRoute()
const router = useRouter()
const tasksStore = useTasksStore()
const projectsStore = useProjectsStore()
const showCreateTask = ref(false)
const selectedProjectId = ref(0)

const columns = [
  { status: 'todo', label: 'To Do', color: '#6b7280' },
  { status: 'in_progress', label: 'In Progress', color: '#3b82f6' },
  { status: 'in_review', label: 'In Review', color: '#f59e0b' },
  { status: 'done', label: 'Done', color: '#10b981' }
]

onMounted(async () => {
  const projectId = route.params.projectId
  if (projectId && projectId !== 'all') {
    selectedProjectId.value = Number(projectId)
  }
  await projectsStore.fetchAll()
  loadTasks()
})

watch(selectedProjectId, () => {
  loadTasks()
})

const loadTasks = () => {
  const filters: any = {}
  if (selectedProjectId.value) {
    filters.project_id = selectedProjectId.value
  }
  tasksStore.fetchAll(filters)
}

const getColumnTasks = (status: string) => {
  if (status === 'done') {
    return tasksStore.doneTasks
  }
  // Map non-done tasks by bucket/status
  return tasksStore.openTasks.filter(t => {
    if (status === 'todo') return !t.bucket_id || t.bucket_id === 0
    if (status === 'in_progress') return t.bucket_id === 1
    if (status === 'in_review') return t.bucket_id === 2
    return false
  })
}

let draggedTask: Task | null = null

const onDragStart = (event: DragEvent, task: Task) => {
  draggedTask = task
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/plain', String(task.id))
  }
}

const onDrop = async (event: DragEvent, status: string) => {
  event.preventDefault()
  if (!draggedTask) return

  const bucketMap: Record<string, number> = {
    todo: 0,
    in_progress: 1,
    in_review: 2,
    done: 0
  }

  const updates: Partial<Task> = {
    bucket_id: bucketMap[status],
    done: status === 'done'
  }

  await tasksStore.update(draggedTask.id, updates)
  draggedTask = null
}

const toggleTaskDone = async (task: Task) => {
  await tasksStore.toggleDone(task.id)
}

const navigateToTask = (task: Task) => {
  router.push(`/tasks/${task.id}`)
}

const isOverdue = (task: Task) => {
  return task.due_date && new Date(task.due_date) < new Date() && !task.done
}

const formatDueDate = (date: string) => {
  const d = new Date(date)
  return d.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}

const formatDuration = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m`
  const hrs = Math.floor(mins / 60)
  return `${hrs}h`
}
</script>

<style scoped lang="scss">
.kanban-board {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 60px);
}

.kanban-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);

  h1 { margin: 0; }
}

.header-actions {
  display: flex;
  gap: var(--spacing-md);
  align-items: center;
}

.project-filter {
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
}

.kanban-container {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: var(--spacing-md);
  flex: 1;
  overflow-x: auto;
  min-height: 0;
}

.kanban-column {
  background-color: var(--bg-secondary);
  border-radius: var(--radius-lg);
  display: flex;
  flex-direction: column;
  min-width: 280px;
}

.column-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  border-bottom: 3px solid var(--border-color);
}

.column-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);

  h2 {
    margin: 0;
    font-size: var(--font-size-sm);
    font-weight: 600;
  }
}

.column-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.task-count {
  background-color: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: 2px 8px;
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
}

.tasks-container {
  flex: 1;
  padding: var(--spacing-sm);
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.kanban-task {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
  cursor: grab;
  transition: all var(--transition-fast);

  &:hover {
    box-shadow: var(--shadow-md);
    border-color: var(--color-primary);
  }

  &:active {
    cursor: grabbing;
    opacity: 0.8;
  }
}

.task-top {
  display: flex;
  align-items: flex-start;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-xs);

  h3 {
    margin: 0;
    font-size: var(--font-size-sm);
    font-weight: 500;
    flex: 1;
  }
}

.task-checkbox {
  margin-top: 2px;
  cursor: pointer;
}

.task-description {
  font-size: var(--font-size-xs);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-sm) 0;
  line-height: 1.4;
}

.task-bottom {
  display: flex;
  gap: var(--spacing-xs);
  flex-wrap: wrap;
  font-size: 11px;
}

.priority-badge {
  padding: 1px 5px;
  border-radius: var(--radius-sm);
  font-weight: 600;

  &.priority-0 { background: rgba(16, 185, 129, 0.1); color: var(--color-success); }
  &.priority-1 { background: rgba(245, 158, 11, 0.1); color: var(--color-warning); }
  &.priority-2 { background: rgba(239, 68, 68, 0.15); color: #dc2626; }
  &.priority-3 { background: rgba(239, 68, 68, 0.25); color: var(--color-error); }
}

.due-date {
  color: var(--text-tertiary);
  &.overdue { color: var(--color-error); font-weight: 600; }
}

.estimate-badge {
  color: var(--text-tertiary);
  background: var(--bg-tertiary);
  padding: 1px 4px;
  border-radius: var(--radius-sm);
}

.eisenhower-mini {
  padding: 1px 4px;
  border-radius: var(--radius-sm);
  font-weight: 700;
  font-size: 9px;
  letter-spacing: 0.5px;

  &.do { background: rgba(239, 68, 68, 0.15); color: var(--color-error); }
}

.empty-column {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
  border: 2px dashed var(--border-color);
  border-radius: var(--radius-md);
  margin: var(--spacing-sm);
}

@media (max-width: 1024px) {
  .kanban-container {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .kanban-container {
    grid-template-columns: 1fr;
  }
}
</style>
