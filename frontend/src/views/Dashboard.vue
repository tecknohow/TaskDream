<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h1>Dashboard</h1>
      <button class="btn btn-primary" @click="showCreateTask = true">+ New Task</button>
    </div>

    <div class="dashboard-grid">
      <div class="card">
        <h2>Today's Tasks</h2>
        <div class="task-list">
          <TaskItem
            v-for="task in tasksStore.todayTasks"
            :key="task.id"
            :task="task"
            @click="selectTask(task)"
          />
          <div v-if="tasksStore.todayTasks.length === 0" class="empty-state">
            No tasks for today
          </div>
        </div>
      </div>

      <div class="card">
        <h2>Overdue Tasks</h2>
        <div class="task-list">
          <TaskItem
            v-for="task in tasksStore.overdueTasks"
            :key="task.id"
            :task="task"
            @click="selectTask(task)"
          />
          <div v-if="tasksStore.overdueTasks.length === 0" class="empty-state">
            No overdue tasks
          </div>
        </div>
      </div>

      <div class="card">
        <h2>Recent Projects</h2>
        <div class="project-list">
          <RouterLink
            v-for="project in projectsStore.activeProjects.slice(0, 5)"
            :key="project.id"
            :to="`/projects/${project.id}`"
            class="project-item-link"
          >
            <div class="project-item-header">
              <div class="project-dot" :style="{ backgroundColor: project.color || 'var(--color-primary)' }"></div>
              <span>{{ project.name }}</span>
            </div>
            <span class="text-xs text-tertiary">{{ project.buckets?.length || 0 }} buckets</span>
          </RouterLink>
        </div>
      </div>

      <div class="card">
        <h2>Quick Stats</h2>
        <div class="stats-grid">
          <div class="stat">
            <div class="stat-value">{{ tasksStore.tasks.length }}</div>
            <div class="stat-label">Total Tasks</div>
          </div>
          <div class="stat">
            <div class="stat-value">{{ tasksStore.todayTasks.length }}</div>
            <div class="stat-label">Today</div>
          </div>
          <div class="stat">
            <div class="stat-value">{{ tasksStore.overdueTasks.length }}</div>
            <div class="stat-label">Overdue</div>
          </div>
          <div class="stat">
            <div class="stat-value">{{ projectsStore.activeProjects.length }}</div>
            <div class="stat-label">Projects</div>
          </div>
        </div>
      </div>
    </div>

    <CreateTaskDialog v-if="showCreateTask" @close="showCreateTask = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import { useProjectsStore } from '@/stores/projects'
import TaskItem from '@/components/TaskItem.vue'
import CreateTaskDialog from '@/components/CreateTaskDialog.vue'
import type { Task } from '@/types/models'

const tasksStore = useTasksStore()
const projectsStore = useProjectsStore()
const showCreateTask = ref(false)

onMounted(() => {
  tasksStore.fetchAll()
  projectsStore.fetchAll()
})

const selectTask = (task: Task) => {
  tasksStore.selectTask(task)
}
</script>

<style scoped lang="scss">
.dashboard {
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
  }
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: var(--spacing-lg);
}

.card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);

  h2 {
    font-size: var(--font-size-lg);
    margin-bottom: var(--spacing-md);
  }
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.empty-state {
  text-align: center;
  padding: var(--spacing-lg);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}

.project-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.project-item-link {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  transition: background-color var(--transition-fast);
  text-decoration: none;
  color: inherit;

  &:hover {
    background-color: var(--bg-tertiary);
  }
}

.project-item-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  font-weight: 500;
}

.project-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--spacing-md);
}

.stat {
  padding: var(--spacing-md);
  background-color: var(--bg-secondary);
  border-radius: var(--radius-md);
  text-align: center;
}

.stat-value {
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--color-primary);
  margin-bottom: var(--spacing-xs);
}

.stat-label {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}
</style>
