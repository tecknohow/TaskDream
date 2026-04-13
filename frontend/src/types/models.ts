export interface User {
  id: string
  username: string
  email: string
  avatar?: string
  createdAt: Date
  updatedAt: Date
}

export interface Task {
  id: string
  title: string
  description?: string
  status: TaskStatus
  priority: TaskPriority
  dueDate?: Date
  createdAt: Date
  updatedAt: Date
  projectId: string
  assigneeId?: string
  labels?: Label[]
  attachments?: Attachment[]
  subtasks?: Task[]
  timeTracked?: number
  bucketId?: string
  order?: number
}

export interface Project {
  id: string
  name: string
  description?: string
  color?: string
  icon?: string
  owner: User
  members?: User[]
  buckets?: Bucket[]
  createdAt: Date
  updatedAt: Date
  archived?: boolean
}

export interface Bucket {
  id: string
  projectId: string
  title: string
  tasks: Task[]
  order: number
}

export interface Label {
  id: string
  name: string
  color: string
  projectId?: string
}

export interface Team {
  id: string
  name: string
  members: User[]
  createdAt: Date
}

export interface TimeTrackingEntry {
  id: string
  taskId: string
  userId: string
  startTime: Date
  endTime?: Date
  duration: number
  note?: string
}

export interface Note {
  id: string
  projectId?: string
  title: string
  content: string
  createdAt: Date
  updatedAt: Date
  userId: string
}

export interface Area {
  id: string
  name: string
  icon?: string
  color?: string
  projects: Project[]
  createdAt: Date
}

export enum TaskStatus {
  TODO = 'todo',
  IN_PROGRESS = 'in_progress',
  IN_REVIEW = 'in_review',
  DONE = 'done'
}

export enum TaskPriority {
  LOW = 0,
  MEDIUM = 1,
  HIGH = 2,
  URGENT = 3
}

export interface Attachment {
  id: string
  taskId: string
  filename: string
  url: string
  uploadedAt: Date
}

export interface AuthResponse {
  token: string
  user: User
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}
