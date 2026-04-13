# TaskDream

> The dream task & project management app. Offline-first, self-hosted.

TaskDream is a modern, feature-rich task and project management application built with **Go** and **Vue.js**. It combines the best ideas from [Vikunja](https://vikunja.io), [Super Productivity](https://super-productivity.com), and [Tududi](https://tududi.com) into one cohesive, self-hosted solution.

## Features

### Task Management
- **Tasks** - Full CRUD with priorities (Low/Medium/High/Urgent), due dates, labels, reminders, attachments, and relations
- **Subtasks** - Break tasks into manageable sub-items with independent tracking
- **Recurring Tasks** - Set tasks to repeat on schedules with cron expressions or durations
- **Task Estimation** - Estimate task duration and compare against actual time spent
- **Eisenhower Matrix** - Categorize tasks by urgency and importance for smart prioritization

### Project Organization
- **Projects** - Organize tasks into projects with custom colors and backgrounds
- **Areas** - Group related projects into high-level areas
- **Kanban Board** - Drag-and-drop board view with configurable columns (To Do, In Progress, In Review, Done)
- **Calendar View** - See tasks by date on a monthly calendar

### Productivity
- **Pomodoro Timer** - Built-in focus timer with customizable work/break durations, session tracking, and statistics
- **Focus Mode** - Distraction-free view showing only the current task with timer and subtasks
- **Time Tracking** - Start/stop timer, log time entries, daily summaries
- **Quick Add** - Press `Shift+A` anywhere to quickly capture tasks without context switching
- **Keyboard Shortcuts** - Keyboard-first design for power users

### Analytics & Insights
- **Dashboard** - Productivity overview with stats, trends, and estimation accuracy
- **7-Day Trend** - Visual chart of daily task completion
- **Estimation Accuracy** - Track how well you estimate task durations over time

### Integrations
- **GitHub** - Connect repositories to sync issues and pull requests as tasks
- **Notes** - Rich notes attached to projects

### Platform
- **Dark Mode** - Full light/dark/auto theme support with toggle
- **Auth** - JWT-based authentication with registration and login
- **Search** - Global search across tasks, projects, and notes
- **Docker** - Single-container deployment with multi-stage build
- **Teams** - Team management with member roles
- **Labels** - Color-coded labels for cross-project organization

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.22, Echo v4, xorm ORM |
| Frontend | Vue 3, TypeScript, Vite, Pinia |
| Database | SQLite (default) or PostgreSQL |
| Styling | SCSS with CSS variables, dark/light mode |
| Build | Docker multi-stage, GitHub Actions CI |

## Quick Start

### Docker (recommended)

```bash
docker-compose up -d
```

Then open http://localhost:3456

### Manual Build

```bash
# Backend
go mod download
go build -o taskdream .

# Frontend
cd frontend
pnpm install
pnpm build
cd ..

# Run
./taskdream web
```

## Development

```bash
# Start backend (port 3456)
go run . web

# Start frontend dev server (port 5173, proxies API to backend)
cd frontend
pnpm install
pnpm dev
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full development guide.

## Configuration

Copy `config.example.yml` to `config.yml` and edit to your needs. Environment variables are also supported with the `TASKDREAM_` prefix (e.g., `TASKDREAM_DATABASE_TYPE=postgres`).

## Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Shift + A` | Quick Add Task |
| `Esc` | Close dialog/modal |

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the planned development phases.

## Inspiration & Credits

TaskDream stands on the shoulders of these excellent open-source projects:

- [Vikunja](https://github.com/go-vikunja/vikunja) - Go-based task management with a Vue frontend (AGPL-3.0)
- [Super Productivity](https://github.com/super-productivity/super-productivity) - Time tracking and task management (MIT)
- [Tududi](https://github.com/chrisvel/tududi) - Calm task/project/notes organizer (MIT)

## License

TaskDream is licensed under the [GNU Affero General Public License v3.0](LICENSE).
