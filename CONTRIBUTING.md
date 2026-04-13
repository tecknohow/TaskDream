# Contributing to TaskDream

Thanks for your interest in contributing to TaskDream!

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/TaskDream.git`
3. Create a feature branch: `git checkout -b feature/my-feature`
4. Make your changes
5. Test your changes: `make test`
6. Commit with a descriptive message: `git commit -m "feat: add cool feature"`
7. Push to your fork: `git push origin feature/my-feature`
8. Open a Pull Request

## Development Setup

### Backend (Go)
- Go 1.22+
- `go mod download` to install deps
- `go run . web` to start the server on :3456

### Frontend (Vue)
- Node.js 20+, pnpm
- `cd frontend && pnpm install && pnpm dev`
- Runs on :5173 with proxy to backend

## Code Style

- Go: follow `gofmt` and `golangci-lint`
- TypeScript/Vue: follow ESLint and Prettier configs
- Commit messages: use conventional commits (feat:, fix:, docs:, etc.)

## Reporting Issues

Use GitHub Issues with the provided templates for bugs and feature requests.

## License

By contributing, you agree that your contributions will be licensed under the AGPL-3.0 license.
