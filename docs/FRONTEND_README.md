# TaskDream Frontend

A modern Vue 3 + TypeScript project management application frontend with Kanban boards, time tracking, notes, and calendar views.

## Architecture

### Tech Stack
- **Framework**: Vue 3 with Composition API
- **Language**: TypeScript
- **Build Tool**: Vite
- **State Management**: Pinia
- **Routing**: Vue Router
- **Styling**: SCSS with CSS Variables
- **HTTP Client**: Axios
- **Utilities**: date-fns, @vueuse/core

### Project Structure

```
src/
├── api/                 # API calls and HTTP client
│   ├── client.ts       # Axios instance with auth interceptor
│   ├── auth.ts         # Authentication endpoints
│   ├── tasks.ts        # Task CRUD operations
│   ├── projects.ts     # Project management
│   └── timeTracking.ts # Time tracking endpoints
├── stores/             # Pinia state management
│   ├── auth.ts         # Authentication state
│   ├── tasks.ts        # Task state
│   ├── projects.ts     # Project state
│   └── timeTracking.ts # Time tracking state
├── types/              # TypeScript interfaces
│   └── models.ts       # Data models
├── views/              # Page components
│   ├── Dashboard.vue       # Main dashboard
│   ├── Login.vue           # Login page
│   ├── Register.vue        # Registration page
│   ├── ProjectView.vue     # Project details
│   ├── TaskDetail.vue      # Individual task view
│   ├── KanbanBoard.vue     # Kanban board view
│   ├── TimeTracking.vue    # Time tracking interface
│   ├── CalendarView.vue    # Calendar view
│   ├── Notes.vue           # Notes management
│   └── Settings.vue        # User settings
├── components/         # Reusable components
│   ├── Sidebar.vue         # Navigation sidebar
│   ├── TaskItem.vue        # Task list item
│   ├── CreateTaskDialog.vue # Task creation modal
│   └── Timer.vue           # Timer component
├── utils/              # Utility functions
│   ├── format.ts       # Date/time formatting
│   └── date.ts         # Date utilities
├── assets/
│   └── main.scss       # Global styles with CSS variables
├── router/             # Vue Router configuration
│   └── index.ts
├── App.vue             # Root component
└── main.ts             # Application entry point
```

## Features

### Core Features
- **Authentication**: Login/Register with JWT token management
- **Dashboard**: Overview of tasks, projects, and statistics
- **Task Management**: Create, read, update, delete tasks with priorities and due dates
- **Projects**: Organize tasks into projects with buckets/columns
- **Kanban Board**: Drag-and-drop task management across status columns
- **Time Tracking**: Track time spent on tasks with timer and history
- **Calendar View**: View tasks by date with monthly navigation
- **Notes**: Create and manage project notes
- **Settings**: User profile and preference management

### UI Features
- Dark/Light mode support (CSS variables)
- Responsive design (mobile-first)
- Smooth animations and transitions
- Loading states and error handling
- Auto-auth token refresh

## Getting Started

### Installation

```bash
# Install dependencies with pnpm
pnpm install

# Or with npm
npm install
```

### Development

```bash
# Start dev server (http://localhost:5173)
pnpm dev

# Build for production
pnpm build

# Preview production build
pnpm preview
```

### Environment Configuration

The application expects a Go backend at `http://localhost:3456`. Configure this in `vite.config.ts`:

```typescript
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:3456',
      changeOrigin: true
    }
  }
}
```

## API Integration

### Authentication Flow

1. User logs in/registers
2. Backend returns JWT token and user object
3. Token stored in localStorage
4. Axios interceptor adds token to all requests
5. 401 responses redirect to login

### Store Usage

```typescript
import { useTasksStore } from '@/stores/tasks'

const tasksStore = useTasksStore()
await tasksStore.fetchAll()
const tasks = tasksStore.tasks
```

## Styling

### CSS Variables

All colors, spacing, and typography are defined as CSS variables in `src/assets/main.scss`:

- **Colors**: primary, secondary, success, warning, error, info
- **Backgrounds**: bg-primary, bg-secondary, bg-tertiary
- **Text**: text-primary, text-secondary, text-tertiary
- **Spacing**: spacing-xs through spacing-2xl
- **Border Radius**: radius-sm through radius-xl
- **Shadows**: shadow-sm through shadow-xl

### Dark Mode

Dark mode is automatically applied based on system preferences via `prefers-color-scheme`. Override CSS variables in dark mode blocks to customize.

### Component Styling

Components use SCSS with:
- CSS variables for theming
- Flexbox/Grid layouts
- BEM-like class naming
- Scoped styles to prevent conflicts

## Performance Optimizations

- Code splitting with route-based components
- Lazy loading of heavy components
- Efficient state management with Pinia
- Debounced API calls
- Optimized re-renders with computed properties

## Development Guidelines

### Component Structure

```vue
<template>
  <!-- HTML -->
</template>

<script setup lang="ts">
// Imports
// Props & Emits
// Reactive State
// Computed Properties
// Methods
// Lifecycle Hooks
</script>

<style scoped lang="scss">
// Styles
</style>
```

### Store Pattern

```typescript
export const useMyStore = defineStore('my', () => {
  // State
  const items = ref([])
  
  // Computed
  const filtered = computed(() => items.value.filter(...))
  
  // Methods
  const fetch = async () => {}
  
  return { items, filtered, fetch }
})
```

## Building & Deployment

### Production Build

```bash
pnpm build
# Output in dist/
```

### Environment Variables

Create `.env.production`:

```
VITE_API_URL=https://api.taskdream.com
```

Access in code:

```typescript
const apiUrl = import.meta.env.VITE_API_URL
```

## Browser Support

- Chrome/Edge: Latest 2 versions
- Firefox: Latest 2 versions
- Safari: Latest 2 versions
- Mobile browsers: iOS Safari 12+, Chrome Android 90+

## Testing

Add test files next to components:

```
components/
├── TaskItem.vue
└── TaskItem.test.ts
```

Run tests:

```bash
pnpm test
```

## Contributing

1. Create feature branch
2. Follow component guidelines
3. Use TypeScript for type safety
4. Add SCSS scoped styles
5. Test responsive design

## License

MIT
