<template>
  <div class="app-layout" :class="{ 'dark-mode': isDarkMode }">
    <Sidebar v-if="authStore.isAuthenticated" />
    <main class="main-content">
      <RouterView />
    </main>
    <QuickAddTask v-if="showQuickAdd" @close="showQuickAdd = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import Sidebar from '@/components/Sidebar.vue'
import QuickAddTask from '@/components/QuickAddTask.vue'

const authStore = useAuthStore()
const showQuickAdd = ref(false)
const isDarkMode = ref(false)

onMounted(() => {
  authStore.initializeAuth()

  // Load theme preference
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme === 'dark') {
    isDarkMode.value = true
    document.documentElement.classList.add('dark')
  } else if (savedTheme === 'light') {
    isDarkMode.value = false
    document.documentElement.classList.remove('dark')
  } else {
    // Auto-detect
    isDarkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    if (isDarkMode.value) {
      document.documentElement.classList.add('dark')
    }
  }

  // Global keyboard shortcuts
  window.addEventListener('keydown', handleGlobalKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleGlobalKeydown)
})

const handleGlobalKeydown = (e: KeyboardEvent) => {
  const target = e.target as HTMLElement
  const isInput = target.tagName === 'INPUT' || target.tagName === 'TEXTAREA' || target.tagName === 'SELECT'

  // Shift+A: Quick add task (global)
  if (e.shiftKey && e.key === 'A' && !e.ctrlKey && !e.metaKey && !isInput) {
    e.preventDefault()
    showQuickAdd.value = true
    return
  }

  // Escape: Close quick add
  if (e.key === 'Escape' && showQuickAdd.value) {
    showQuickAdd.value = false
    return
  }
}
</script>

<style scoped lang="scss">
.app-layout {
  display: flex;
  min-height: 100vh;
  background-color: var(--bg-secondary);
}

.main-content {
  flex: 1;
  overflow-y: auto;
  padding: var(--spacing-lg);
}

@media (max-width: 768px) {
  .app-layout {
    flex-direction: column;
  }

  .main-content {
    padding: var(--spacing-md);
  }
}
</style>
