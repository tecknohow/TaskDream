<template>
  <div class="settings">
    <div class="settings-header">
      <h1>Settings</h1>
    </div>

    <div class="settings-grid">
      <div class="settings-card">
        <h2>Account Settings</h2>
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
        <h2>Change Password</h2>
        <div class="form-group">
          <label>Current Password</label>
          <input v-model="currentPassword" type="password" />
        </div>

        <div class="form-group">
          <label>New Password</label>
          <input v-model="newPassword" type="password" />
        </div>

        <div class="form-group">
          <label>Confirm Password</label>
          <input v-model="confirmPassword" type="password" />
        </div>

        <button class="btn btn-primary" @click="changePassword">Change Password</button>
      </div>

      <div class="settings-card">
        <h2>Preferences</h2>
        <div class="form-group">
          <label>
            <input type="checkbox" v-model="preferences.emailNotifications" />
            Email Notifications
          </label>
        </div>

        <div class="form-group">
          <label>
            <input type="checkbox" v-model="preferences.taskReminders" />
            Task Reminders
          </label>
        </div>

        <div class="form-group">
          <label>Theme</label>
          <select v-model="preferences.theme">
            <option value="light">Light</option>
            <option value="dark">Dark</option>
            <option value="auto">Auto</option>
          </select>
        </div>

        <button class="btn btn-primary" @click="savePreferences">Save Preferences</button>
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

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const email = ref('')
const currentPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const message = ref('')

const preferences = ref({
  emailNotifications: true,
  taskReminders: true,
  theme: 'auto'
})

onMounted(() => {
  if (authStore.user) {
    username.value = authStore.user.username
    email.value = authStore.user.email
  }

  const savedPreferences = localStorage.getItem('preferences')
  if (savedPreferences) {
    preferences.value = JSON.parse(savedPreferences)
  }
})

const saveProfile = async () => {
  try {
    await authStore.updateProfile({
      username: username.value,
      email: email.value
    })
    message.value = 'Profile updated successfully!'
    setTimeout(() => {
      message.value = ''
    }, 3000)
  } catch (err) {
    console.error('Failed to update profile', err)
  }
}

const changePassword = async () => {
  if (newPassword.value !== confirmPassword.value) {
    alert('Passwords do not match')
    return
  }

  try {
    await authStore.updateProfile({
      username: authStore.user?.username
    })
    message.value = 'Password changed successfully!'
    currentPassword.value = ''
    newPassword.value = ''
    confirmPassword.value = ''
    setTimeout(() => {
      message.value = ''
    }, 3000)
  } catch (err) {
    console.error('Failed to change password', err)
  }
}

const savePreferences = () => {
  localStorage.setItem('preferences', JSON.stringify(preferences.value))
  message.value = 'Preferences saved!'
  setTimeout(() => {
    message.value = ''
  }, 3000)
}

const logout = () => {
  if (confirm('Are you sure you want to logout?')) {
    authStore.logout()
    router.push('/login')
  }
}
</script>

<style scoped lang="scss">
.settings {
  max-width: 1000px;
  margin: 0 auto;
}

.settings-header {
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
  }
}

.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.settings-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);

  h2 {
    margin-top: 0;
    margin-bottom: var(--spacing-lg);
    font-size: var(--font-size-lg);
  }

  p {
    margin: 0 0 var(--spacing-md) 0;
    color: var(--text-secondary);
    font-size: var(--font-size-sm);
  }
}

.danger-zone {
  border-left: 4px solid var(--color-error);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-md);

  label {
    font-weight: 500;
    font-size: var(--font-size-sm);
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);

    input[type='checkbox'] {
      width: 18px;
      height: 18px;
      cursor: pointer;
    }
  }

  input[type='text'],
  input[type='email'],
  input[type='password'],
  select {
    padding: var(--spacing-md);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    font-size: var(--font-size-base);
    background-color: var(--bg-secondary);
    color: var(--text-primary);

    &:focus {
      outline: none;
      border-color: var(--color-primary);
      background-color: var(--bg-primary);
    }
  }
}

.success-message {
  padding: var(--spacing-md) var(--spacing-lg);
  background-color: rgba(16, 185, 129, 0.1);
  color: var(--color-success);
  border-radius: var(--radius-md);
  text-align: center;
  font-size: var(--font-size-sm);
}

@media (max-width: 768px) {
  .settings-grid {
    grid-template-columns: 1fr;
  }
}
</style>
