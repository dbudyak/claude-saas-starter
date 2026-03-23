# SaaS Starter — Go + React + Claude Code

A battle-tested template for shipping full-stack SaaS applications fast. Built around Claude Code's multi-agent system, this starter gives you a production-ready foundation with Go backend, React frontend, observability, and CI/CD — all wired up and ready to customize.

## What's included

- **Go backend** (Chi router, pgx/PostgreSQL, JWT auth, migrations)
- **React frontend** (TypeScript, TailwindCSS 4, React Query, Zustand, i18n)
- **Infrastructure** (Docker Compose, Caddy reverse proxy, Prometheus + Grafana + Loki)
- **CI/CD** (GitHub Actions: test → build → deploy)
- **Claude Code agents** (backend, frontend, infra, product — scoped and optimized)
- **Development rules** (Go standards, React standards, security, aesthetics)
- **Observability** (metrics, logs, dashboards out of the box)

## Quick Start

### Prerequisites

- Docker + Docker Compose
- Go 1.24+
- Node 20+
- Claude Code CLI (`npm install -g @anthropic-ai/claude-code`)

### 1. Clone and configure

```bash
git clone <this-repo> myapp
cd myapp
cp services/infra/docker/.env.example services/infra/docker/.env
# Edit .env with your values
```

### 2. Start everything

```bash
make local-up
```

This starts: backend, frontend, PostgreSQL, Caddy, Prometheus, Grafana, Loki, pgAdmin.

### 3. Access your app

| Service | URL |
|---------|-----|
| App | http://localhost |
| API | http://localhost/api |
| Grafana | http://localhost:3002 |
| pgAdmin | http://localhost:5050 |

### 4. Customize with Claude Code

```bash
claude
```

The multi-agent system is pre-configured. Claude Code knows your stack and can operate each service independently.

## Project Structure

```
saas-starter/
├── .claude/
│   ├── agents/           # Scoped AI agents per service
│   ├── rules/            # Coding standards and guidelines
│   ├── knowledge/        # Shared architectural knowledge
│   └── settings.json     # Agent permissions + MCP servers
├── services/
│   ├── backend/          # Go API (Chi, pgx, JWT)
│   ├── frontend/         # React/TypeScript (Vite, Tailwind)
│   └── infra/            # Docker, Caddy, monitoring, scripts
├── docs/                 # Project documentation
├── .github/workflows/    # CI/CD pipelines
├── Makefile              # All dev commands
└── CLAUDE.md             # Claude Code project context
```

## Development Commands

```bash
make help              # Show all commands
make local-up          # Start all services (Docker)
make local-down        # Stop all services
make local-logs        # Tail all logs
make backend-run       # Run backend locally (no Docker)
make frontend-run      # Run frontend locally (no Docker)
make backend-test      # Run backend unit tests
make backend-integration  # Run integration tests (needs Docker)
make frontend-test     # Run frontend tests
make dev-deploy        # Deploy to dev server
```

## Environment Variables

Copy `.env.example` to `.env` and configure:

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgresql://user:pass@db:5432/myapp` |
| `JWT_SECRET` | JWT signing secret (min 32 chars) | `$(openssl rand -hex 32)` |
| `DOMAIN` | Your domain | `myapp.com` |
| `APP_URL` | Full app URL | `https://myapp.com` |

## Customization Checklist

After cloning, update these for your project:

- [ ] Replace `myapp` with your project name throughout
- [ ] Update `go.mod` module name (`github.com/yourorg/yourapp`)
- [ ] Update `package.json` name field
- [ ] Configure domain in Caddy and environment variables
- [ ] Set up GitHub repository secrets for CI/CD
- [ ] Add your logo and branding to frontend
- [ ] Update `docs/PROJECT.md` with your product description
- [ ] Set up your deployment server

See [SETUP_CHECKLIST.md](SETUP_CHECKLIST.md) for the complete guide.

## Claude Code Multi-Agent System

This template ships with four specialized agents:

| Agent | Scope |
|-------|-------|
| `backend` | `services/backend/` |
| `frontend` | `services/frontend/` |
| `infra` | `services/infra/` |
| `product` | `docs/` |

Agents are automatically scoped — the backend agent cannot accidentally modify frontend files. Each agent has deep knowledge of its domain and follows verification protocols.

## Observability

Pre-configured out of the box:

- **Prometheus** — scrapes backend metrics at `/metrics`
- **Grafana** — dashboards at `:3002` (admin/admin by default)
- **Loki + Promtail** — log aggregation from all containers

## Deployment

Deployment via GitHub Actions to any Linux server with Docker:

1. Push to `main` → tests run
2. Tests pass → Docker images built and pushed to GHCR
3. Images ready → SSH deploy to your server

Required GitHub secrets: `DEPLOY_HOST`, `DEPLOY_USER`, `DEPLOY_SSH_KEY`, plus all app environment variables.

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.24, Chi, pgx/v5, JWT |
| Frontend | React 19, TypeScript, Vite 6, TailwindCSS 4 |
| Database | PostgreSQL 15 |
| Reverse Proxy | Caddy (auto SSL) |
| Monitoring | Prometheus, Grafana, Loki |
| Container | Docker, Docker Compose |
| CI/CD | GitHub Actions, GHCR |

## License

MIT
