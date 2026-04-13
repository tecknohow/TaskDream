import client from './client'
import type { AuthResponse, LoginRequest, RegisterRequest, User } from '@/types/models'

export const authAPI = {
  login: (credentials: LoginRequest) => {
    return client.post<AuthResponse>('/auth/login', credentials)
  },

  register: (data: RegisterRequest) => {
    return client.post<AuthResponse>('/auth/register', data)
  },

  logout: () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  },

  getCurrentUser: () => {
    return client.get<User>('/users/me')
  },

  updateProfile: (user: Partial<User>) => {
    return client.put<User>('/users/me', user)
  },

  changePassword: (oldPassword: string, newPassword: string) => {
    return client.post('/auth/change-password', { oldPassword, newPassword })
  }
}
