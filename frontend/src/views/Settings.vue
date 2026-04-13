<template>
  <div class="settings">
    <div class="settings-header">
      <h1>Settings</h1>
    </div>

    <div class="settings-grid">
      <div class="settings-card">
        <h2>Account</h2>
        <div class="form-group">
          <label>Username</label>
          <input v-model="username" type="text" />
        </div>
        <div class="form-group">
          <label>Email</label>
          <input v-model="email" type="email" />
        </div>
        <button class="btn btn-primary" @click="saveProfile">Save Profile</button>
      </div>

      <div class="settings-card">
        <h2>Appearance</h2>
        <div class="form-group">
          <label>Theme</label>
          <div class="theme-selector">
            <button
              v-for="t in themes"
              :key="t.value"
              class="theme-btn"
              :class="{ active: theme === t.value }"
              @click="setTheme(t.value)"
            >
              {{ t.label }}
            </button>
          </div>
        </div>
      </div>

      <div class="settings-card">
        <h2>Pomodoro Timer</h2>
        <div class="form-group">
          <label>Work Duration (minutes)</label>
          <input v-model.number="pomodoroWork" type="number" min="1" max="120" />
        </div>
        <div class="form-group">
          <label>Short Break (minutes)</label>
          <input v-model.number="pomodoroShortBreak" type="number" min="1" max="30" />
        </div>
        <div class="form-group">
          <label>Long Break (minutes)</label>
          <input v-model.number="pomodoroLongBreak" type="number" min="1" max="60" />
        </div>
        <div class="form-group">
          <label>Long Break After (sessions)</label>
          <input v-model.number="pomodoroInterval" type="number" min="2" max="10" />
        </div>
        <div class="form-group">
          <label class="toggle-label">
            <input type="checkbox" v-model="autoStartBreaks" />
            Auto-start breaks
          </label>
        </div>
        <button class="btn btn-primary" @click="savePomodoroSettings">Save Pomodoro Settings</button>
      </div>

      <div class="settings-card">
        <h2>Keyboard Shortcuts</h2>
        <div class="shortcuts-list">
          <div class="shortcut-item">
            <span class="shortcut-desc">Quick Add Task</span>
            <kbd>Shift + A</kbd>
          </div>
          <div class="shortcut-item">
            <span class="shortcut-desc">Close Dialog</span>
            <kbd>Esc</kbd>
          </div>
        </div>
      </div>

      <div class="settings-card danger-zone">
        <h2>Danger Zone</h2>
        <p>These actions cannot be undone.</p>
        <button class="btn btn-danger" @click="logout">Logout</button>
      </div>
    </div>

    <div v-if="message" class="success-message">
      {{ message }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePomodoroStore } from '@/stores/pomodoro'

const router = useRouter()
const authStore = useAuthStore()
const pomodoroStore = usePomodoroStore()

const username = ref('')
const email = ref('')
const message = ref('')
const theme = ref('auto')

const pomodoroWork = ref(25)
const pomodoroShortBreak = ref(5)
const pomodoroLongBreak = ref(15)
const pomodoroInterval = ref(4)
const autoStartBreaks = ref(false)

const themes = [
  { value: 'light', label: 'Light' },
  { value: 'dark', label: 'Dark' },
  { value: 'auto', label: 'Auto' }
]

onMounted(async () => {
  if (authStore.user) {
    username.value = authStore.user.username
    email.value = authStore.user.email
  }

  theme.value = localStorage.getItem('theme') || 'auto'

  await pomodoroStore.fetchSettings()
  const s = pomodoroStore.settings
  pomodoroWork.value = Math.floor(s.work_duration / 60)
  pomodoroShortBreak.value = Math.floor(s.short_break / 60)
  pomodoroLongBreak.value = Math.floor(s.long_break / 60)
  pomodoroInterval.value = s.long_break_interval
  autoStartBreaks.value = s.auto_start_breaks
})

const setTheme = (t: string) => {
  theme.value = t
  localStorage.setItem('theme', t)

  if (t === 'dark') {
    document.documentElement.classList.add('dark')
    document.documentElement.classList.remove('light')
  } else if (t === 'light') {
    document.documentElement.classList.remove('dark')
    document.documentElement.classList.add('light')
  } else {
    document.documentElement.classList.remove('dark', 'light')
  }
}

const saveProfile = async () => {
  try {
    await authStore.updateProfile({ username: username.value, email: email.value })
    showMessage('Profile updated!')
  } catch {
    // error
  }
}

const savePomodoroSettings = async () => {
  await pomodoroStore.updateSettings({
    work_duration: pomodoroWork.value * 60,
    short_break: pomodoroShortBreak.value * 60,
    long_break: pomodoroLongBreak.value * 60,
    long_break_interval: pomodoroInterval.value,
    auto_start_breaks: autoStartBreaks.value
  })
  showMessage('Pomodoro settings saved!')
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}

const showMessage = (msg: string) => {
  message.value = msg
  setTimeout(() => { message.value = '' }, 3000)
}
</script>

<style scoped lang="scss">
.settings {
  max-width: 800px;
  margin: 0 auto;
}

.settings-header {
  margin-bottom: var(--spacing-xl);
  h1 { margin: 0; }
}

.settings-grid {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-xl);
}

.settings-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);

  h2 { margin: 0 0 var(--spacing-lg); font-size: var(--font-size-lg); }
  p { margin: 0 0 var(--spacing-md); color: var(--text-secondary); font-size: var(--font-size-sm); }
}

.danger-zone {
  border-left: 4px solid var(--color-error);
}

.form-group {
  margin-bottom: var(--spacing-md);

  > label:not(.toggle-label) {
    display: block;
    font-size: var(--font-size-sm);
    font-weight: 500;
    color: var(--text-secondary);
    margin-bottom: var(--spacing-xs);
  }

  input[type='text'],
  input[type='email'],
  input[type='password'],
  input[type='number'],
  select {
    width: 100%;
  }
}

.toggle-label {
  display: flex !important;
  align-items: center;
  gap: var(--spacing-sm);
  cursor: pointer;
  font-size: var(--font-size-sm);

  input[type="checkbox"] { width: auto; }
}

.theme-selector {
  display: flex;
  gap: var(--spacing-sm);
  background: var(--bg-tertiary);
  border-radius: var(--radius-md);
  padding: 3px;
}

.theme-btn {
  flex: 1;
  padding: var(--spacing-sm);
  border: none;
  background: transparent;
  border-radius: var(--radius-sm);
  cursor: pointer;
  font-size: var(--font-size-sm);
  font-weight: 500;
  color: var(--text-secondary);
  transition: all var(--transition-fast);

  &.active {
    background: var(--bg-primary);
    color: var(--text-primary);
    box-shadow: var(--shadow-sm);
  }
}

.shortcuts-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.shortcut-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-sm) 0;
}

.shortcut-desc {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

kbd {
  padding: 3px 8px;
  background: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-family: monospace;
}

.success-message {
  padding: var(--spacing-md);
  background: rgba(16, 185, 129, 0.1);
  color: var(--color-success);
  border-radius: var(--radius-md);
  text-align: center;
  font-size: var(--font-size-sm);
}
</style>
