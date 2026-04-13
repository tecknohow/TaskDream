# TaskDream

> The dream task & project management app. Offline-first, self-hosted.

TaskDream is a modern, feature-rich task and project management application built with **Go** and **Vue.js**. It combines the best ideas from [Vikunja](https://vikunja.io), [Super Productivity](https://super-productivity.com), and [Tududi](https://tududi.com) into one cohesive, self-hosted solution.

## Features

- **Tasks** - Full CRUD with priorities (1-5), due dates, labels, reminders, attachments, subtasks, and relations
- **Projects** - Organize tasks into projects with custom colors and backgrounds
- **Kanban Board** - Drag-and-drop board view with configurable columns/buckets
- **Time Tracking** - Start/stop timer, log time entries, daily summaries (inspired by Super Productivity)
- **Notes** - Rich notes attached to projects (inspired by Tududi)
- **Areas** - Group related projects into high-level areas (inspired by Tududi)
- **Calendar View** - See tasks by date on a monthly calendar
- **Teams** - Team management with member roles
- **Labels** - Color-coded labels for cross-project organization
- **Dark Mode** - Full light/dark theme support
- **Auth** - JWT-based authentication with registration and login
- **Docker** - Single-container deployment with multi-stage build

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.22, Echo v4, xorm ORM |
| Frontend | Vue 3, TypeScript, Vite, Pinia |
| Database | SQLite (default) or PostgreSQL |
| Styling | SCSS with CSS variables, dark mode |
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

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the planned development phases.

## Inspiration & Credits

TaskDream stands on the shoulders of these excellent open-source projects:

- [Vikunja](https://github.com/go-vikunja/vikunja) - Go-based task management with a Vue frontend (AGPL-3.0)
- [Super Productivity](https://github.com/super-productivity/super-productivity) - Time tracking and task management (MIT)
- [Tududi](https://github.com/chrisvel/tududi) - Calm task/project/notes organizer (MIT)

## License

TaskDream is licensed under the [GNU Affero General Public License v3.0](LICENSE).
