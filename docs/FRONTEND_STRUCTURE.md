# TaskDream Frontend - File Structure

## Configuration Files
- `package.json` - Project dependencies and scripts
- `vite.config.ts` - Vite build configuration with Vue plugin and backend proxy
- `tsconfig.json` - TypeScript compiler options
- `index.html` - HTML entry point
- `env.d.ts` - Vite environment type declarations
- `.editorconfig` - Editor configuration
- `.gitignore` - Git ignore rules

## Source Files (src/)

### Root
- `main.ts` - Application bootstrap with Pinia and Router
- `App.vue` - Root component with sidebar and router-view

### API Layer (src/api/)
- `client.ts` - Axios instance with auth interceptor and 401 redirect
- `auth.ts` - Login, register, profile update endpoints
- `tasks.ts` - Task CRUD, filtering, bulk operations
- `projects.ts` - Project management and bucket operations
- `timeTracking.ts` - Timer control and time entry management

### State Management (src/stores/)
- `auth.ts` - Auth store (login, register, logout, profile)
- `tasks.ts` - Task store (CRUD, filtering, sorting, selection)
- `projects.ts` - Project store (CRUD, buckets management)
- `timeTracking.ts` - Time tracking store (timer control, entries, history)

### Type Definitions (src/types/)
- `models.ts` - TypeScript interfaces for all data models:
  - User, Task, Project, Bucket, Label, Team
  - TimeTrackingEntry, Note, Area
  - TaskStatus, TaskPriority enums
  - Auth request/response types

### Routing (src/router/)
- `index.ts` - Vue Router configuration with auth guard
  - Routes for all views with requiresAuth metadata

### Views (src/views/)
- `Dashboard.vue` - Main dashboard with today's tasks, overdue, projects, stats
- `Login.vue` - Login form with authentication
- `Register.vue` - Registration form
- `ProjectView.vue` - Project detail with task buckets
- `TaskDetail.vue` - Individual task detail panel
- `KanbanBoard.vue` - Kanban board with drag-drop status columns
- `TimeTracking.vue` - Active timer and time entry history
- `CalendarView.vue` - Monthly calendar with task visualization
- `Notes.vue` - Note creation and management
- `Settings.vue` - User profile and preferences

### Components (src/components/)
- `Sidebar.vue` - Navigation sidebar with projects and views
- `TaskItem.vue` - Reusable task list item component
- `CreateTaskDialog.vue` - Task creation modal dialog
- `Timer.vue` - Active timer display component

### Utilities (src/utils/)
- `format.ts` - Date/time formatting with date-fns
- `date.ts` - Date utility functions

### Styles (src/assets/)
- `main.scss` - Global stylesheet with:
  - CSS variables for light/dark mode
  - Typography, colors, spacing
  - Form elements and buttons
  - Cards, badges, utilities
  - Responsive breakpoints
  - Animations and transitions

## Key Features Implemented

### Authentication
- JWT token management with localStorage
- Auto-refresh on 401 responses
- Protected routes with auth guard
- Login/Register pages

### Task Management
- Full CRUD operations
- Priority levels (Low, Medium, High, Urgent)
- Due date tracking with overdue highlighting
- Task status workflow (To Do → In Progress → In Review → Done)
- Checkbox toggles for completion
- Bulk operations support

### Projects
- Project creation and management
- Color-coded project indicators
- Bucket/column organization
- Project-specific task filtering
- Member management API ready

### Kanban Board
- Status-based columns (To Do, In Progress, In Review, Done)
- Task count per column
- Quick status updates via checkbox
- Priority-based color coding
- Task preview with due dates

### Time Tracking
- Active timer with formatted display (HH:MM:SS)
- Start/stop timer controls
- Time entry history with date/time
- Daily summary statistics
- Time per task aggregation
- Entry deletion

### Calendar View
- Monthly calendar grid
- Task display on due dates
- Today highlighting
- Month navigation
- Task detail sidebar on selection
- Responsive design

### Notes
- Note creation and deletion
- Rich text editing
- Note list with preview
- Updated timestamp tracking
- Auto-save functionality

### Settings
- User profile management
- Password change
- Notification preferences
- Theme selection
- Account security

## Styling Architecture

### CSS Variables
All styling uses CSS custom properties for:
- **Theme Colors**: primary, secondary, success, warning, error, info
- **Backgrounds**: Three-level hierarchy (primary, secondary, tertiary)
- **Text Colors**: Three levels of emphasis
- **Spacing**: 8 levels from xs to 2xl
- **Border Radius**: 4 levels
- **Shadows**: 4 levels
- **Typography**: 6 font sizes
- **Transitions**: Fast, base, slow durations

### Dark Mode
- Automatic via `prefers-color-scheme` media query
- Variables redefined in dark mode block
- No component changes needed

### Components
- Scoped styles per component
- Flexbox/Grid layouts
- BEM-like naming conventions
- Consistent spacing and rhythm
- Interactive hover/focus states
- Loading and disabled states

## Development Setup

### Installation
```bash
pnpm install
```

### Development Server
```bash
pnpm dev
# Runs on http://localhost:5173
```

### Production Build
```bash
pnpm build
pnpm preview
```

### Backend Integration
The frontend expects a Go backend at `http://localhost:3456`. All `/api` requests are proxied to this address.

## Browser Compatibility
- Modern browsers with ES2020+ support
- Chrome, Firefox, Safari, Edge
- Mobile browsers (iOS Safari 12+, Chrome Android 90+)
- Dark mode CSS support required

## Performance Optimizations
- Code splitting with async routes
- Lazy component imports
- Efficient Pinia state management
- Computed properties for derived state
- Scoped styles prevent CSS conflicts
- CSS variables enable fast theme switching

## Testing Ready
Test files can be added alongside components using:
- Vitest
- Vue Test Utils
- jsdom or happy-dom

## Documentation
- `README.md` - Complete project documentation
- Component comments for complex logic
- TypeScript types for API safety
