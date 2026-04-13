import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import '@/assets/main.scss'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

// Initialize auth state before mounting
const { useAuthStore } = await import('./stores/auth')
const authStore = useAuthStore()
authStore.initializeAuth()

app.mount('#app')
