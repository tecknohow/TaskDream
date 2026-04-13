import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authAPI } from '@/api/auth'
import type { User, LoginRequest, RegisterRequest } from '@/types/models'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const initializeAuth = () => {
    const savedToken = localStorage.getItem('token')
    const savedUser = localStorage.getItem('user')

    if (savedToken && savedUser) {
      token.value = savedToken
      try {
        user.value = JSON.parse(savedUser)
      } catch {
        localStorage.removeItem('user')
        localStorage.removeItem('token')
      }
    }
  }

  const login = async (credentials: LoginRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await authAPI.login(credentials)
      const { token: newToken, user: newUser } = response.data

      token.value = newToken
      user.value = newUser

      localStorage.setItem('token', newToken)
      localStorage.setItem('user', JSON.stringify(newUser))

      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Login failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const register = async (data: RegisterRequest) => {
    loading.value = true
    error.value = null

    try {
      const response = await authAPI.register(data)
      const { token: newToken, user: newUser } = response.data

      token.value = newToken
      user.value = newUser

      localStorage.setItem('token', newToken)
      localStorage.setItem('user', JSON.stringify(newUser))

      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Registration failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const logout = () => {
    authAPI.logout()
    token.value = null
    user.value = null
  }

  const updateProfile = async (updates: Partial<User>) => {
    loading.value = true
    error.value = null

    try {
      const response = await authAPI.updateProfile(updates)
      user.value = response.data
      localStorage.setItem('user', JSON.stringify(response.data))
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Update failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  return {
    user,
    token,
    loading,
    error,
    isAuthenticated,
    initializeAuth,
    login,
    register,
    logout,
    updateProfile
  }
})
