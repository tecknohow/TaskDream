<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="logo">
        <span class="logo-icon">✓</span>
        <span class="logo-text">TaskDream</span>
      </div>
    </div>

    <nav class="sidebar-nav">
      <div class="nav-section">
        <h3 class="nav-section-title">Main</h3>
        <RouterLink to="/" class="nav-item" :class="{ active: $route.path === '/' }">
          <span class="nav-icon">📊</span>
          <span class="nav-label">Dashboard</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Projects</h3>
        <button class="nav-item nav-item-btn" @click="showCreateProject = true">
          <span class="nav-icon">+</span>
          <span class="nav-label">New Project</span>
        </button>

        <RouterLink
          v-for="project in projectsStore.activeProjects"
          :key="project.id"
          :to="`/projects/${project.id}`"
          class="nav-item project-item"
          :class="{ active: $route.params.id === project.id }"
        >
          <span class="project-color" :style="{ backgroundColor: project.color || 'var(--color-primary)' }"></span>
          <span class="nav-label">{{ project.name }}</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Views</h3>
        <RouterLink to="/kanban/all" class="nav-item" :class="{ active: $route.name === 'KanbanBoard' }">
          <span class="nav-icon">🎯</span>
          <span class="nav-label">Kanban</span>
        </RouterLink>
        <RouterLink to="/calendar" class="nav-item" :class="{ active: $route.name === 'CalendarView' }">
          <span class="nav-icon">📅</span>
          <span class="nav-label">Calendar</span>
        </RouterLink>
        <RouterLink to="/time-tracking" class="nav-item" :class="{ active: $route.name === 'TimeTracking' }">
          <span class="nav-icon">⏱️</span>
          <span class="nav-label">Time Tracking</span>
        </RouterLink>
        <RouterLink to="/notes" class="nav-item" :class="{ active: $route.name === 'Notes' }">
          <span class="nav-icon">📝</span>
          <span class="nav-label">Notes</span>
        </RouterLink>
      </div>

      <div class="nav-section">
        <h3 class="nav-section-title">Account</h3>
        <RouterLink to="/settings" class="nav-item" :class="{ active: $route.name === 'Settings' }">
          <span class="nav-icon">⚙️</span>
          <span class="nav-label">Settings</span>
        </RouterLink>
        <button class="nav-item nav-item-btn" @click="logout">
          <span class="nav-icon">🚪</span>
          <span class="nav-label">Logout</span>
        </button>
      </div>
    </nav>

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

const router = useRouter()
const authStore = useAuthStore()
const projectsStore = useProjectsStore()
const showCreateProject = ref(false)

onMounted(() => {
  projectsStore.fetchAll()
})

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
  background-color: var(--color-primary);
  color: white;
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: var(--font-size-lg);
}

.sidebar-nav {
  flex: 1;
  padding: 0 var(--spacing-md);
}

.nav-section {
  margin-bottom: var(--spacing-xl);

  &:last-child {
    margin-bottom: 0;
  }
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

  &:hover {
    background-color: var(--bg-tertiary);
    color: var(--text-primary);
  }

  &.active {
    background-color: var(--color-primary);
    color: white;

    .nav-label {
      font-weight: 600;
    }
  }
}

.nav-item-btn {
  background: none;
  border: none;
  padding: var(--spacing-sm) var(--spacing-md);
}

.nav-icon {
  font-size: var(--font-size-lg);
  display: flex;
  align-items: center;
  justify-content: center;
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
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: var(--color-primary);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: var(--font-size-base);
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
