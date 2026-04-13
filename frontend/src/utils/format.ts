import { format } from 'date-fns'

export const formatDate = (date: Date | string): string => {
  try {
    const d = typeof date === 'string' ? new Date(date) : date
    return format(d, 'MMM dd, yyyy')
  } catch {
    return 'Invalid date'
  }
}

export const formatDateTime = (date: Date | string): string => {
  try {
    const d = typeof date === 'string' ? new Date(date) : date
    return format(d, 'MMM dd, yyyy HH:mm')
  } catch {
    return 'Invalid date'
  }
}

export const formatSeconds = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60

  if (hours > 0) {
    return `${hours}h ${minutes}m ${secs}s`
  }
  if (minutes > 0) {
    return `${minutes}m ${secs}s`
  }
  return `${secs}s`
}

export const formatTime = (seconds: number): string => {
  const pad = (num: number) => String(num).padStart(2, '0')
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60

  return `${pad(hours)}:${pad(minutes)}:${pad(secs)}`
}
