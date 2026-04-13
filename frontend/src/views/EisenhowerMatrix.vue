<template>
  <div class="eisenhower-view">
    <div class="eisenhower-header">
      <h1>Eisenhower Matrix</h1>
      <div class="header-actions">
        <button class="btn btn-primary" @click="showCreateTask = true">+ New Task</button>
      </div>
    </div>

    <div class="matrix-grid">
      <!-- Quadrant 1: Urgent & Important - DO -->
      <div class="quadrant q1">
        <div class="quadrant-header">
          <div class="quadrant-title">
            <span class="quadrant-icon">1</span>
            <div>
              <h2>Do First</h2>
              <span class="quadrant-desc">Urgent & Important</span>
            </div>
          </div>
          <span class="task-count">{{ tasksStore.eisenhowerQuadrants.urgentImportant.length }}</span>
        </div>
        <div class="quadrant-tasks">
          <TaskItem
            v-for="task in tasksStore.eisenhowerQuadrants.urgentImportant"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
          <div v-if="tasksStore.eisenhowerQuadrants.urgentImportant.length === 0" class="empty-quadrant">
            No tasks here - great!
          </div>
        </div>
      </div>

      <!-- Quadrant 2: Not Urgent & Important - SCHEDULE -->
      <div class="quadrant q2">
        <div class="quadrant-header">
          <div class="quadrant-title">
            <span class="quadrant-icon">2</span>
            <div>
              <h2>Schedule</h2>
              <span class="quadrant-desc">Not Urgent & Important</span>
            </div>
          </div>
          <span class="task-count">{{ tasksStore.eisenhowerQuadrants.notUrgentImportant.length }}</span>
        </div>
        <div class="quadrant-tasks">
          <TaskItem
            v-for="task in tasksStore.eisenhowerQuadrants.notUrgentImportant"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
          <div v-if="tasksStore.eisenhowerQuadrants.notUrgentImportant.length === 0" class="empty-quadrant">
            Plan important work here
          </div>
        </div>
      </div>

      <!-- Quadrant 3: Urgent & Not Important - DELEGATE -->
      <div class="quadrant q3">
        <div class="quadrant-header">
          <div class="quadrant-title">
            <span class="quadrant-icon">3</span>
            <div>
              <h2>Delegate</h2>
              <span class="quadrant-desc">Urgent & Not Important</span>
            </div>
          </div>
          <span class="task-count">{{ tasksStore.eisenhowerQuadrants.urgentNotImportant.length }}</span>
        </div>
        <div class="quadrant-tasks">
          <TaskItem
            v-for="task in tasksStore.eisenhowerQuadrants.urgentNotImportant"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
          <div v-if="tasksStore.eisenhowerQuadrants.urgentNotImportant.length === 0" class="empty-quadrant">
            Delegate these tasks
          </div>
        </div>
      </div>

      <!-- Quadrant 4: Not Urgent & Not Important - ELIMINATE -->
      <div class="quadrant q4">
        <div class="quadrant-header">
          <div class="quadrant-title">
            <span class="quadrant-icon">4</span>
            <div>
              <h2>Eliminate</h2>
              <span class="quadrant-desc">Not Urgent & Not Important</span>
            </div>
          </div>
          <span class="task-count">{{ tasksStore.eisenhowerQuadrants.notUrgentNotImportant.length }}</span>
        </div>
        <div class="quadrant-tasks">
          <TaskItem
            v-for="task in tasksStore.eisenhowerQuadrants.notUrgentNotImportant"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
          <div v-if="tasksStore.eisenhowerQuadrants.notUrgentNotImportant.length === 0" class="empty-quadrant">
            Consider removing these
          </div>
        </div>
      </div>
    </div>

    <!-- Axis labels -->
    <div class="axis-labels">
      <div class="axis-y-label">
        <span class="axis-arrow">&uarr;</span> IMPORTANT
      </div>
      <div class="axis-x-label">
        URGENT <span class="axis-arrow">&rarr;</span>
      </div>
    </div>

    <CreateTaskDialog v-if="showCreateTask" @close="showCreateTask = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTasksStore } from '@/stores/tasks'
import TaskItem from '@/components/TaskItem.vue'
import CreateTaskDialog from '@/components/CreateTaskDialog.vue'
import type { Task } from '@/types/models'

const router = useRouter()
const tasksStore = useTasksStore()
const showCreateTask = ref(false)

onMounted(() => {
  tasksStore.fetchAll()
})

const navigateToTask = (task: Task) => {
  router.push(`/tasks/${task.id}`)
}
</script>

<style scoped lang="scss">
.eisenhower-view {
  max-width: 1400px;
  margin: 0 auto;
  position: relative;
}

.eisenhower-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);

  h1 { margin: 0; }
}

.matrix-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  gap: var(--spacing-md);
  min-height: 600px;
}

.quadrant {
  background-color: var(--bg-primary);
  border-radius: var(--radius-lg);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: var(--shadow-sm);
  border: 2px solid transparent;

  &.q1 {
    border-color: rgba(239, 68, 68, 0.3);
    .quadrant-icon { background: rgba(239, 68, 68, 0.15); color: var(--color-error); }
  }

  &.q2 {
    border-color: rgba(37, 99, 235, 0.3);
    .quadrant-icon { background: rgba(37, 99, 235, 0.15); color: var(--color-primary); }
  }

  &.q3 {
    border-color: rgba(245, 158, 11, 0.3);
    .quadrant-icon { background: rgba(245, 158, 11, 0.15); color: var(--color-warning); }
  }

  &.q4 {
    border-color: rgba(156, 163, 175, 0.3);
    .quadrant-icon { background: rgba(156, 163, 175, 0.15); color: var(--text-tertiary); }
  }
}

.quadrant-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
}

.quadrant-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);

  h2 {
    margin: 0;
    font-size: var(--font-size-base);
  }
}

.quadrant-icon {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: var(--font-size-sm);
}

.quadrant-desc {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
}

.task-count {
  background-color: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
}

.quadrant-tasks {
  flex: 1;
  padding: var(--spacing-md);
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.empty-quadrant {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  min-height: 100px;
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}

.axis-labels {
  position: relative;
  margin-top: var(--spacing-md);
}

.axis-y-label, .axis-x-label {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  text-transform: uppercase;
  letter-spacing: 1px;
  font-weight: 600;
}

.axis-y-label {
  position: absolute;
  left: -10px;
  top: -320px;
  transform: rotate(-90deg);
  transform-origin: left center;
}

.axis-x-label {
  text-align: right;
}

@media (max-width: 768px) {
  .matrix-grid {
    grid-template-columns: 1fr;
    min-height: auto;
  }

  .axis-y-label {
    display: none;
  }
}
</style>
