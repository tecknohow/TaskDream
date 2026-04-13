# TaskDream

> A productivity-focused task & project management app. Built on [Vikunja](https://vikunja.io) with [Super Productivity](https://super-productivity.com) inspired features.

TaskDream is a self-hosted task and project management application that extends Vikunja with deep productivity features inspired by Super Productivity and Tududi.

## What's Added on Top of Vikunja

TaskDream inherits all of Vikunja's features (tasks, projects, kanban, labels, teams, CalDAV, sharing, etc.) and adds:

- **Pomodoro Timer** - Built-in focus timer with customizable work/break durations and session tracking
- **Focus Mode** - Distraction-free view showing only the current task with timer and subtask tracking
- **Eisenhower Matrix** - 4-quadrant urgency/importance grid for task prioritization
- **GitHub Integration** - Connect repositories to sync issues and pull requests as tasks
- **Productivity Analytics** - Dashboard with trends, estimation accuracy, and daily summaries
- **Quick Add** - Global keyboard shortcut (Shift+A) for rapid task capture

## Upstream

This project tracks [go-vikunja/vikunja](https://github.com/go-vikunja/vikunja) as an upstream remote. To pull latest Vikunja updates:

```bash
git fetch upstream
git merge upstream/main
```

## Development

See the [Vikunja development docs](https://vikunja.io/docs/development/) for backend and frontend setup.

```bash
# Backend
go run . web

# Frontend
cd frontend
pnpm install
pnpm dev
```

## License

Licensed under the [GNU Affero General Public License v3.0](LICENSE), same as Vikunja.

## Credits

- [Vikunja](https://github.com/go-vikunja/vikunja) - The foundation (AGPL-3.0)
- [Super Productivity](https://github.com/super-productivity/super-productivity) - Pomodoro, time tracking, focus mode inspiration (MIT)
- [Tududi](https://github.com/chrisvel/tududi) - Area/notes organization inspiration (MIT)
