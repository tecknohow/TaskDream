export interface User {
  id: number
  username: string
  email: string
  name?: string
  avatar?: string
  created: string
  updated: string
}

export interface Task {
  id: number
  title: string
  description?: string
  done: boolean
  priority: number       // 0=low, 1=medium, 2=high, 3=urgent
  urgency: number        // 0=not urgent, 1=urgent (Eisenhower)
  importance: number     // 0=not important, 1=important (Eisenhower)
  due_date?: string
  project_id: number
  bucket_id?: number
  parent_id?: number
  labels?: number[]
  reminders?: Reminder[]
  repeat_after?: string
  repeat_mode?: string
  attachments?: Attachment[]
  position: number
  percent_done: number
  start_date?: string
  end_date?: string
  estimated_time: number   // in seconds
  total_time_spent: number // in seconds
  assignee_id?: number
  created_by_id: number
  created: string
  updated: string
}

export interface TaskWithSubtasks {
  task: Task
  subtasks: Task[]
}

export interface Project {
  id: number
  title: string
  description?: string
  color?: string
  is_archived: boolean
  background_file_id?: number
  position: number
  area_id?: number
  owner_id: number
  created: string
  updated: string
}

export interface Bucket {
  id: number
  title: string
  project_id: number
  limit: number
  position: number
  created: string
  updated: string
}

export interface Label {
  id: number
  title: string
  description?: string
  hex_color: string
  created_by_id: number
  created: string
  updated: string
}

export interface Team {
  id: number
  name: string
  description?: string
  created: string
  updated: string
}

export interface TimeTrackingEntry {
  id: number
  task_id: number
  user_id: number
  start: string
  end?: string
  duration: number // in seconds
  comment?: string
  created: string
  updated: string
}

export interface Note {
  id: number
  title: string
  content: string
  project_id?: number
  created_by_id: number
  created: string
  updated: string
}

export interface Area {
  id: number
  title: string
  description?: string
  color?: string
  created_by_id: number
  created: string
  updated: string
}

export interface Reminder {
  relative_to: string
  duration: number
  reminder: string
}

export interface Attachment {
  id: number
  filename: string
  file_size: number
  created_at: string
}

// Pomodoro types
export interface PomodoroSession {
  id: number
  task_id: number
  user_id: number
  duration: number       // in seconds
  break_duration: number // in seconds
  status: 'pending' | 'running' | 'completed' | 'cancelled'
  started_at?: string
  completed_at?: string
  created: string
  updated: string
}

export interface PomodoroSettings {
  id?: number
  user_id: number
  work_duration: number       // in seconds
  short_break: number         // in seconds
  long_break: number          // in seconds
  long_break_interval: number // sessions before long break
  auto_start_breaks: boolean
  auto_start_pomodoro: boolean
}

export interface PomodoroStats {
  today_completed: number
  today_total_time: number
  weekly_completed: number
}

// GitHub integration types
export interface GithubIntegration {
  connected: boolean
  username?: string
  avatar_url?: string
  repos?: RepoConfig[]
  sync_enabled?: boolean
  last_synced_at?: string
}

export interface RepoConfig {
  owner: string
  name: string
  full_name: string
  project_id: number
  sync_issues: boolean
  sync_prs: boolean
}

export interface GithubIssueSync {
  id: number
  task_id: number
  github_issue_id: number
  repo_full_name: string
  issue_number: number
  issue_title: string
  issue_state: string
  issue_url: string
  is_pull_request: boolean
  last_synced_at: string
  created: string
  updated: string
}

// Analytics types
export interface DashboardStats {
  total_tasks: number
  open_tasks: number
  completed_tasks: number
  completed_today: number
  completed_this_week: number
  overdue_tasks: number
  due_today: number
  today_time_tracked: number
  week_time_tracked: number
  pomodoros_today: number
  active_projects: number
  estimation_accuracy: number
  tasks_with_estimates: number
}

export interface ProductivityTrend {
  date: string
  tasks_completed: number
  time_tracked: number
  pomodoros: number
}

// Search types
export interface SearchResults {
  tasks: Task[]
  projects: Project[]
  notes: Note[]
}

// Enums
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

// Auth types
export interface AuthResponse {
  access_token: string
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
  name?: string
}
