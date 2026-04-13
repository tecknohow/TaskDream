# TaskDream Backend Structure

## Overview
Complete Go backend for TaskDream - a task/project management application inspired by Vikunja's architecture, with features from Super Productivity and Tududi.

Module: `github.com/tecknohow/TaskDream`
Go Version: 1.22

## Project Structure

```
TaskDream/
├── main.go                           # Application entry point
├── go.mod                            # Go module definition
│
├── pkg/
│   ├── cmd/                          # CLI commands using Cobra
│   │   ├── root.go                   # Root command
│   │   └── web.go                    # HTTP server startup command
│   │
│   ├── config/                       # Configuration management (Viper)
│   │   └── config.go                 # Config loading, defaults, env vars
│   │
│   ├── db/                           # Database initialization (xorm)
│   │   └── db.go                     # xorm engine setup (SQLite/PostgreSQL)
│   │
│   ├── models/                       # Data models (12 models)
│   │   ├── user.go                   # User accounts with auth fields
│   │   ├── area.go                   # Areas (collection of projects, Tududi-inspired)
│   │   ├── project.go                # Projects with ownership and archiving
│   │   ├── bucket.go                 # Kanban columns for projects
│   │   ├── task.go                   # Tasks with rich metadata (priority, labels, reminders, attachments)
│   │   ├── task_comment.go           # Comments on tasks
│   │   ├── task_relation.go          # Task relationships (parent/child/blocking/related)
│   │   ├── label.go                  # Reusable task labels/tags
│   │   ├── team.go                   # Team management
│   │   ├── time_tracking.go          # Time tracking entries (Super Productivity-inspired)
│   │   └── note.go                   # Notes in projects (Tududi-inspired)
│   │
│   ├── routes/                       # REST API routes
│   │   ├── routes.go                 # Main router setup with Echo, CORS, JWT middleware
│   │   └── api/v1/                   # API v1 handlers (functional implementations)
│   │       ├── auth.go               # Login, Register, Token Refresh
│   │       ├── user.go               # Get/Update current user
│   │       ├── project.go            # List, Create, Get, Update, Delete projects
│   │       ├── task.go               # List, Create, Get, Update, Delete tasks
│   │       ├── task_comment.go       # List, Create, Delete task comments
│   │       ├── label.go              # List, Create, Get, Update, Delete labels
│   │       ├── area.go               # List, Create, Get, Update, Delete areas
│   │       ├── bucket.go             # List, Create, Update, Delete buckets
│   │       ├── note.go               # List, Create, Get, Update, Delete notes
│   │       └── time_tracking.go      # List, Create, Update, Delete time tracking entries
│   │
│   ├── migration/                    # Database migrations
│   │   └── migration.go              # Auto-migration using xorm Sync2
│   │
│   └── utils/                        # Utility functions
│       └── hash.go                   # bcrypt password hashing and verification
```

## Key Features

### Models (12 total)
- **User**: Authentication, profiles, avatars
- **Area**: Project grouping (Tududi-inspired)
- **Project**: Workspace containers with ownership
- **Bucket**: Kanban columns within projects
- **Task**: Full-featured with priority, due dates, labels, reminders, attachments, relations
- **TaskComment**: Discussion on tasks
- **TaskRelation**: Parent/child/blocking/related task relationships
- **Label**: Reusable tags for tasks
- **Team**: Team management with members
- **TimeTracking**: Billable time tracking (Super Productivity-inspired)
- **Note**: Project notes (Tududi-inspired)

### API Routes (RESTful)
All routes follow `/api/v1/` prefix:
- Auth: `/auth/login`, `/auth/register`, `/auth/refresh`
- Users: `/users/me`
- Projects: `/projects`, `/projects/:id`
- Tasks: `/tasks`, `/tasks/:id`, `/tasks/:id/comments`
- Labels: `/labels`, `/labels/:id`
- Areas: `/areas`, `/areas/:id`
- Buckets: `/buckets`, `/buckets/:id`
- Notes: `/notes`, `/notes/:id`
- Time Tracking: `/tasks/:id/time-tracking`, `/time-tracking/:id`

### Handler Implementation
All handlers are fully functional:
- Request binding and validation
- Database queries using xorm
- JSON responses with proper HTTP status codes
- Error handling with meaningful messages
- Authentication via JWT middleware

### Technologies Used
- **Framework**: Echo v4 (lightweight HTTP server)
- **ORM**: xorm (database abstraction)
- **Database**: SQLite or PostgreSQL support
- **Auth**: JWT (golang-jwt/jwt)
- **CLI**: Cobra (command framework)
- **Config**: Viper (config management)
- **Crypto**: bcrypt (password hashing)
- **API Docs**: Swagger/Swaggo

### Configuration
Supports configuration via:
- YAML files (config.yaml)
- Environment variables (TASKDREAM_* prefix)
- Sensible defaults

Example env vars:
- `TASKDREAM_SERVER_PORT` (default: 8080)
- `TASKDREAM_DATABASE_DRIVER` (default: sqlite)
- `TASKDREAM_JWT_SECRET`

### Database Support
- SQLite (file-based, default)
- PostgreSQL (production)

Auto-migration on startup using xorm Sync2.

## Quick Start

1. Initialize module:
```bash
go mod download
```

2. Run the server:
```bash
go run main.go web
```

3. API will be available at `http://localhost:8080/api/v1`

4. Health check: `GET /health`

## File Statistics
- 29 Go source files
- 12 data models
- 10 API handler modules
- 1 main entry point
- 1 go.mod with all dependencies
