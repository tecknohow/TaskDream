<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="logo">
        <span class="logo-icon">TD</span>
        <span class="logo-text">TaskDream</span>
      </div>
    </div>

    <nav class="sidebar-nav">
      <div class="nav-section">
        <h3 class="nav-section-title">Main</h3>
        <RouterLink to="/" class="nav-item" :class="{ active: $route.path === '/' }">
          <span class="nav-icon">D</span>
          <span class="nav-label">Dashboard</span>
        </RouterLink>
        <RouterLink to="/focus" class="nav-item" :class="{ active: $route.name === 'FocusMode' }">
          <span class="nav-icon">F</span>
          <span class="nav-label">Focus Mode</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Projects</h3>
        <button class="nav-item nav-item-btn" @click="showCreateProject = !showCreateProject">
          <span class="nav-icon">+</span>
          <span class="nav-label">New Project</span>
        </button>

        <div v-if="showCreateProject" class="inline-form">
          <input
            v-model="newProjectTitle"
            type="text"
            placeholder="Project name"
            @keyup.enter="createProject"
            @keyup.esc="showCreateProject = false"
            autofocus
            class="inline-input"
          />
        </div>

        <RouterLink
          v-for="project in projectsStore.activeProjects"
          :key="project.id"
          :to="`/projects/${project.id}`"
          class="nav-item project-item"
          :class="{ active: $route.params.id == project.id }"
        >
          <span class="project-color" :style="{ backgroundColor: project.color || 'var(--color-primary)' }"></span>
          <span class="nav-label">{{ project.title }}</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Views</h3>
        <RouterLink to="/kanban" class="nav-item" :class="{ active: $route.name === 'KanbanBoard' }">
          <span class="nav-icon">K</span>
          <span class="nav-label">Kanban</span>
        </RouterLink>
        <RouterLink to="/eisenhower" class="nav-item" :class="{ active: $route.name === 'EisenhowerMatrix' }">
          <span class="nav-icon">E</span>
          <span class="nav-label">Eisenhower</span>
        </RouterLink>
        <RouterLink to="/calendar" class="nav-item" :class="{ active: $route.name === 'CalendarView' }">
          <span class="nav-icon">C</span>
          <span class="nav-label">Calendar</span>
        </RouterLink>
        <RouterLink to="/time-tracking" class="nav-item" :class="{ active: $route.name === 'TimeTracking' }">
          <span class="nav-icon">T</span>
          <span class="nav-label">Time Tracking</span>
        </RouterLink>
        <RouterLink to="/notes" class="nav-item" :class="{ active: $route.name === 'Notes' }">
          <span class="nav-icon">N</span>
          <span class="nav-label">Notes</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Account</h3>
        <RouterLink to="/integrations/github" class="nav-item" :class="{ active: $route.name === 'GithubIntegration' }">
          <span class="nav-icon">G</span>
          <span class="nav-label">GitHub</span>
        </RouterLink>
        <RouterLink to="/settings" class="nav-item" :class="{ active: $route.name === 'Settings' }">
          <span class="nav-icon">S</span>
          <span class="nav-label">Settings</span>
        </RouterLink>
        <button class="nav-item nav-item-btn" @click="logout">
          <span class="nav-icon">X</span>
          <span class="nav-label">Logout</span>
        </button>
      </div>
    </nav>

    <!-- Active Timer -->
    <div v-if="timeTrackingStore.isTracking" class="sidebar-timer">
      <div class="timer-active">
        <div class="timer-pulse"></div>
        <span class="timer-time">{{ timeTrackingStore.formattedTime }}</span>
        <button class="btn btn-sm btn-ghost" @click="timeTrackingStore.stopTracking()">Stop</button>
      </div>
    </div>

    <div class="sidebar-footer" v-if="authStore.user">
      <div class="user-info">
        <div class="user-avatar">{{ authStore.user.username.charAt(0).toUpperCase() }}</div>
        <div class="user-details">
          <div class="user-name">{{ authStore.user.username }}</div>
          <div class="user-email">{{ authStore.user.email }}</div>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useProjectsStore } from '@/stores/projects'
import { useTimeTrackingStore } from '@/stores/timeTracking'

const router = useRouter()
const authStore = useAuthStore()
const projectsStore = useProjectsStore()
const timeTrackingStore = useTimeTrackingStore()
const showCreateProject = ref(false)
const newProjectTitle = ref('')

onMounted(() => {
  projectsStore.fetchAll()
})

const createProject = async () => {
  if (!newProjectTitle.value.trim()) return
  await projectsStore.create({
    title: newProjectTitle.value.trim()
  } as any)
  newProjectTitle.value = ''
  showCreateProject.value = false
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<style scoped lang="scss">
.sidebar {
  width: var(--sidebar-width);
  background-color: var(--bg-primary);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  padding: var(--spacing-lg) 0;
}

.sidebar-header {
  padding: var(--spacing-md) var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
  margin-bottom: var(--spacing-lg);
}

.logo {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  font-weight: 700;
  font-size: var(--font-size-lg);
  color: var(--text-primary);
}

.logo-icon {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  color: white;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--font-size-xs);
  font-weight: 800;
}

.sidebar-nav {
  flex: 1;
  padding: 0 var(--spacing-md);
}

.nav-section {
  margin-bottom: var(--spacing-xl);
}

.nav-section-title {
  font-size: var(--font-size-xs);
  font-weight: 600;
  color: var(--text-tertiary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  padding: 0 var(--spacing-md);
  margin-bottom: var(--spacing-sm);
}

.nav-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  color: var(--text-secondary);
  transition: all var(--transition-fast);
  background: none;
  border: none;
  font-size: var(--font-size-base);
  cursor: pointer;
  width: 100%;
  text-align: left;
  text-decoration: none;

  &:hover {
    background-color: var(--bg-tertiary);
    color: var(--text-primary);
  }

  &.active {
    background-color: var(--color-primary);
    color: white;

    .nav-label { font-weight: 600; }
    .nav-icon { color: white; }
  }
}

.nav-item-btn {
  background: none;
  border: none;
}

.nav-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  background: var(--bg-tertiary);
  border-radius: var(--radius-sm);
  flex-shrink: 0;

  .active & {
    background: rgba(255, 255, 255, 0.2);
    color: white;
  }
}

.nav-label {
  flex: 1;
  font-size: var(--font-size-sm);
}

.project-item {
  .project-color {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    flex-shrink: 0;
  }
}

.inline-form {
  padding: var(--spacing-xs) var(--spacing-md);
}

.inline-input {
  width: 100%;
  padding: var(--spacing-xs) var(--spacing-sm);
  font-size: var(--font-size-sm);
  border-radius: var(--radius-sm);
}

.sidebar-timer {
  padding: var(--spacing-md) var(--spacing-lg);
  border-top: 1px solid var(--border-color);
}

.timer-active {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  background: rgba(239, 68, 68, 0.1);
  border-radius: var(--radius-md);
}

.timer-pulse {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--color-error);
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}

.timer-time {
  flex: 1;
  font-family: monospace;
  font-weight: 600;
  font-size: var(--font-size-sm);
}

.sidebar-footer {
  padding: var(--spacing-md) var(--spacing-lg);
  border-top: 1px solid var(--border-color);
  margin-top: auto;
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--color-primary), var(--color-secondary));
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: var(--font-size-sm);
}

.user-details {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: var(--font-size-sm);
  font-weight: 600;
  color: var(--text-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-email {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

@media (max-width: 768px) {
  .sidebar {
    display: none;
  }
}
</style>
