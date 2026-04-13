<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <div>
        <h1>Dashboard</h1>
        <p class="text-secondary text-sm">Welcome back. Here's your productivity overview.</p>
      </div>
      <div class="header-actions">
        <button class="btn btn-ghost" @click="showQuickAdd = true" title="Quick Add (Shift+A)">
          + Quick Add
        </button>
        <button class="btn btn-primary" @click="showCreateTask = true">+ New Task</button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="stats-row" v-if="dashStats">
      <div class="stat-card">
        <div class="stat-icon" style="background: rgba(37, 99, 235, 0.1); color: var(--color-primary);">T</div>
        <div>
          <div class="stat-value">{{ dashStats.open_tasks }}</div>
          <div class="stat-label">Open Tasks</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: rgba(16, 185, 129, 0.1); color: var(--color-success);">D</div>
        <div>
          <div class="stat-value">{{ dashStats.completed_today }}</div>
          <div class="stat-label">Done Today</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: rgba(239, 68, 68, 0.1); color: var(--color-error);">!</div>
        <div>
          <div class="stat-value">{{ dashStats.overdue_tasks }}</div>
          <div class="stat-label">Overdue</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: rgba(139, 92, 246, 0.1); color: var(--color-secondary);">P</div>
        <div>
          <div class="stat-value">{{ dashStats.pomodoros_today }}</div>
          <div class="stat-label">Pomodoros</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon" style="background: rgba(245, 158, 11, 0.1); color: var(--color-warning);">T</div>
        <div>
          <div class="stat-value">{{ formatDuration(dashStats.today_time_tracked) }}</div>
          <div class="stat-label">Time Today</div>
        </div>
      </div>
    </div>

    <div class="dashboard-grid">
      <!-- Today's Tasks -->
      <div class="card">
        <div class="card-header">
          <h2>Today's Tasks</h2>
          <span class="badge badge-primary">{{ tasksStore.todayTasks.length }}</span>
        </div>
        <div class="task-list">
          <TaskItem
            v-for="task in tasksStore.todayTasks"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
          <div v-if="tasksStore.todayTasks.length === 0" class="empty-state">
            No tasks due today
          </div>
        </div>
      </div>

      <!-- Overdue Tasks -->
      <div class="card" v-if="tasksStore.overdueTasks.length > 0">
        <div class="card-header">
          <h2>Overdue</h2>
          <span class="badge badge-error">{{ tasksStore.overdueTasks.length }}</span>
        </div>
        <div class="task-list">
          <TaskItem
            v-for="task in tasksStore.overdueTasks.slice(0, 5)"
            :key="task.id"
            :task="task"
            @click="navigateToTask(task)"
          />
        </div>
      </div>

      <!-- Productivity Trend -->
      <div class="card trend-card">
        <div class="card-header">
          <h2>7-Day Trend</h2>
        </div>
        <div class="trend-chart" v-if="trendData.length > 0">
          <div class="chart-bars">
            <div
              v-for="day in trendData"
              :key="day.date"
              class="chart-bar-group"
            >
              <div class="bar-container">
                <div
                  class="bar tasks-bar"
                  :style="{ height: getBarHeight(day.tasks_completed, maxTasks) + '%' }"
                  :title="`${day.tasks_completed} tasks`"
                ></div>
              </div>
              <div class="bar-label">{{ formatShortDate(day.date) }}</div>
            </div>
          </div>
        </div>
        <div v-else class="empty-state">No data yet</div>
      </div>

      <!-- Estimation Accuracy -->
      <div class="card" v-if="dashStats && dashStats.tasks_with_estimates > 0">
        <div class="card-header">
          <h2>Estimation Accuracy</h2>
        </div>
        <div class="accuracy-display">
          <div class="accuracy-circle" :class="accuracyClass">
            <span class="accuracy-value">{{ Math.round(dashStats.estimation_accuracy) }}%</span>
          </div>
          <p class="accuracy-description">
            Based on {{ dashStats.tasks_with_estimates }} completed tasks with estimates.
            {{ accuracyMessage }}
          </p>
        </div>
      </div>

      <!-- Recent Projects -->
      <div class="card">
        <div class="card-header">
          <h2>Projects</h2>
          <span class="badge badge-primary">{{ projectsStore.activeProjects.length }}</span>
        </div>
        <div class="project-list">
          <RouterLink
            v-for="project in projectsStore.activeProjects.slice(0, 5)"
            :key="project.id"
            :to="`/projects/${project.id}`"
            class="project-item-link"
          >
            <div class="project-item-header">
              <div class="project-dot" :style="{ backgroundColor: project.color || 'var(--color-primary)' }"></div>
              <span>{{ project.title }}</span>
            </div>
          </RouterLink>
          <div v-if="projectsStore.activeProjects.length === 0" class="empty-state">
            No projects yet
          </div>
        </div>
      </div>
    </div>

    <CreateTaskDialog v-if="showCreateTask" @close="showCreateTask = false" />
    <QuickAddTask v-if="showQuickAdd" @close="showQuickAdd = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTasksStore } from '@/stores/tasks'
import { useProjectsStore } from '@/stores/projects'
import { analyticsAPI } from '@/api/analytics'
import TaskItem from '@/components/TaskItem.vue'
import CreateTaskDialog from '@/components/CreateTaskDialog.vue'
import QuickAddTask from '@/components/QuickAddTask.vue'
import type { Task, DashboardStats, ProductivityTrend } from '@/types/models'

const router = useRouter()
const tasksStore = useTasksStore()
const projectsStore = useProjectsStore()
const showCreateTask = ref(false)
const showQuickAdd = ref(false)
const dashStats = ref<DashboardStats | null>(null)
const trendData = ref<ProductivityTrend[]>([])

const maxTasks = computed(() => {
  return Math.max(1, ...trendData.value.map(d => d.tasks_completed))
})

const accuracyClass = computed(() => {
  if (!dashStats.value) return ''
  const acc = dashStats.value.estimation_accuracy
  if (acc >= 80 && acc <= 120) return 'good'
  if (acc >= 60 && acc <= 150) return 'fair'
  return 'poor'
})

const accuracyMessage = computed(() => {
  if (!dashStats.value) return ''
  const acc = dashStats.value.estimation_accuracy
  if (acc >= 80 && acc <= 120) return 'Great job! Your estimates are on track.'
  if (acc > 120) return 'You tend to underestimate. Consider adding buffer time.'
  return 'You tend to overestimate. Try tighter estimates.'
})

onMounted(async () => {
  tasksStore.fetchAll()
  projectsStore.fetchAll()

  try {
    const [statsResp, trendResp] = await Promise.all([
      analyticsAPI.getDashboardStats(),
      analyticsAPI.getProductivityTrend(7)
    ])
    dashStats.value = statsResp.data
    trendData.value = trendResp.data
  } catch {
    // Analytics are optional
  }

  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
})

const handleKeydown = (e: KeyboardEvent) => {
  if (e.shiftKey && e.key === 'A' && !e.ctrlKey && !e.metaKey) {
    const target = e.target as HTMLElement
    if (target.tagName === 'INPUT' || target.tagName === 'TEXTAREA' || target.tagName === 'SELECT') return
    e.preventDefault()
    showQuickAdd.value = true
  }
}

const navigateToTask = (task: Task) => {
  router.push(`/tasks/${task.id}`)
}

const formatDuration = (seconds: number) => {
  if (!seconds) return '0m'
  const mins = Math.floor(seconds / 60)
  if (mins < 60) return `${mins}m`
  const hrs = Math.floor(mins / 60)
  const remainMins = mins % 60
  return remainMins > 0 ? `${hrs}h ${remainMins}m` : `${hrs}h`
}

const getBarHeight = (value: number, max: number) => {
  return Math.max(5, (value / max) * 100)
}

const formatShortDate = (dateStr: string) => {
  const d = new Date(dateStr)
  return d.toLocaleDateString('en-US', { weekday: 'short' }).substring(0, 2)
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
  align-items: flex-start;
  margin-bottom: var(--spacing-xl);

  h1 { margin-bottom: var(--spacing-xs); }
}

.header-actions {
  display: flex;
  gap: var(--spacing-sm);
}

.stats-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
}

.stat-card {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
}

.stat-icon {
  width: 40px;
  height: 40px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: var(--font-size-lg);
}

.stat-value {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--text-primary);
}

.stat-label {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
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
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-md);

  h2 {
    font-size: var(--font-size-lg);
    margin: 0;
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
  align-items: center;
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  transition: background-color var(--transition-fast);
  text-decoration: none;
  color: inherit;

  &:hover { background-color: var(--bg-tertiary); }
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

// Trend chart
.trend-chart {
  padding: var(--spacing-md) 0;
}

.chart-bars {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  height: 120px;
  gap: var(--spacing-sm);
}

.chart-bar-group {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-xs);
}

.bar-container {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: flex-end;
  justify-content: center;
}

.bar {
  width: 70%;
  border-radius: var(--radius-sm) var(--radius-sm) 0 0;
  transition: height var(--transition-base);
  min-height: 4px;
}

.tasks-bar {
  background: linear-gradient(180deg, var(--color-primary), var(--color-primary-dark));
}

.bar-label {
  font-size: 10px;
  color: var(--text-tertiary);
  text-transform: uppercase;
}

// Estimation accuracy
.accuracy-display {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: var(--spacing-lg) 0;
}

.accuracy-circle {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--spacing-md);
  border: 4px solid;

  &.good { border-color: var(--color-success); }
  &.fair { border-color: var(--color-warning); }
  &.poor { border-color: var(--color-error); }
}

.accuracy-value {
  font-size: var(--font-size-lg);
  font-weight: 700;
}

.accuracy-description {
  text-align: center;
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  max-width: 300px;
}
</style>
