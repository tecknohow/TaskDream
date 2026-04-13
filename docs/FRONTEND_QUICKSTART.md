# TaskDream Frontend - Quick Start Guide

## Installation & Setup

### 1. Install Dependencies
```bash
cd TaskDream/frontend
pnpm install
# or
npm install
```

### 2. Start Development Server
```bash
pnpm dev
```
The app will run at `http://localhost:5173`

### 3. Configure Backend
The frontend expects a Go backend at `http://localhost:3456`. Update in `vite.config.ts` if needed:

```typescript
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:3456',  // Change this
      changeOrigin: true
    }
  }
}
```

## Project Overview

TaskDream is a full-featured project management application with:

- **Dashboard** - Overview of tasks and projects
- **Task Management** - Create, edit, delete tasks with priorities and due dates
- **Kanban Board** - Drag-and-drop task management
- **Time Tracking** - Track time spent on tasks
- **Calendar** - View tasks by date
- **Notes** - Create project notes
- **Authentication** - Secure login/register

## Key Technologies

- **Vue 3** - Progressive framework with Composition API
- **TypeScript** - Type-safe development
- **Pinia** - State management
- **Vite** - Fast build tool
- **SCSS** - Styling with CSS variables
- **Axios** - HTTP client
- **date-fns** - Date utilities

## File Organization

```
frontend/
├── src/
│   ├── api/           # API endpoints
│   ├── stores/        # Pinia stores
│   ├── views/         # Page components
│   ├── components/    # Reusable components
│   ├── types/         # TypeScript interfaces
│   ├── utils/         # Helper functions
│   ├── router/        # Vue Router config
│   ├── assets/        # SCSS styles
│   ├── App.vue        # Root component
│   └── main.ts        # Entry point
├── index.html         # HTML template
├── vite.config.ts     # Build config
├── tsconfig.json      # TypeScript config
└── package.json       # Dependencies
```

## Building for Production

```bash
# Build
pnpm build

# Preview build
pnpm preview
```

Output goes to `dist/` directory.

## Development Workflow

### Creating a New View
1. Create file in `src/views/MyView.vue`
2. Add route in `src/router/index.ts`
3. Create corresponding store in `src/stores/` if needed
4. Add API functions in `src/api/`

### Creating a New Component
1. Create file in `src/components/MyComponent.vue`
2. Use `<script setup lang="ts">` syntax
3. Define props with TypeScript
4. Add scoped SCSS styles

### Adding State Management
1. Create store in `src/stores/myStore.ts`
2. Use Pinia's `defineStore` with Composition API
3. Import and use in components with `const store = useMyStore()`

### Styling
- Use CSS variables for colors: `color: var(--color-primary)`
- Reference spacing: `padding: var(--spacing-md)`
- Responsive: `@media (max-width: 768px)`
- Dark mode: automatic via `prefers-color-scheme`

## Common Tasks

### Fetch Data in Component
```typescript
import { useTasksStore } from '@/stores/tasks'
import { onMounted } from 'vue'

const tasksStore = useTasksStore()

onMounted(() => {
  tasksStore.fetchAll()
})
```

### Emit Events from Component
```typescript
const emit = defineEmits<{
  submit: [value: string]
}>()

const handleSubmit = () => {
  emit('submit', 'value')
}
```

### Navigate Programmatically
```typescript
import { useRouter } from 'vue-router'

const router = useRouter()
router.push('/dashboard')
router.push({ name: 'TaskDetail', params: { id: '123' } })
```

### Format Dates
```typescript
import { formatDate, formatSeconds } from '@/utils/format'

formatDate(new Date())        // "Jan 15, 2026"
formatSeconds(3661)            // "1h 1m 1s"
```

## Testing

Add test files next to components:

```typescript
// TaskItem.test.ts
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import TaskItem from './TaskItem.vue'

describe('TaskItem', () => {
  it('renders task title', () => {
    const wrapper = mount(TaskItem, {
      props: {
        task: { title: 'Test' }
      }
    })
    expect(wrapper.text()).toContain('Test')
  })
})
```

## Debugging

### Enable Vue DevTools
Install Vue DevTools browser extension to inspect components and state.

### Check Network Tab
View API requests in browser DevTools Network tab.

### Console Logging
```typescript
console.log('Debug:', value)
console.table(arrayData)
console.error('Error:', error)
```

### TypeScript Errors
```bash
vue-tsc --noEmit  # Check without building
```

## Performance Tips

1. Use computed properties for derived state
2. Lazy load heavy components
3. Avoid creating functions in templates
4. Use v-if instead of v-show for hidden content
5. Memoize expensive computations

## Troubleshooting

### Port 5173 Already in Use
```bash
pnpm dev -- --port 3000
```

### Backend Connection Issues
- Check backend is running on `http://localhost:3456`
- Check CORS is enabled in backend
- Check browser DevTools Network tab for failed requests

### Build Size Too Large
```bash
# Analyze bundle
pnpm build -- --mode analyze
```

## Resources

- [Vue 3 Documentation](https://vuejs.org/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [Pinia Documentation](https://pinia.vuejs.org/)
- [Vite Documentation](https://vitejs.dev/)
- [Vue Router Documentation](https://router.vuejs.org/)

## Support

For issues or questions:
1. Check the README.md for detailed documentation
2. Review STRUCTURE.md for project organization
3. Check component JSDoc comments
4. Review TypeScript types in `src/types/models.ts`
