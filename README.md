# SaaS Starter — Go + React + Claude Code

A full-stack SaaS template with Claude Code multi-agent configuration. The goal is to skip the setup and get straight to building your product.

## What's included

**Application scaffold**
- Go backend: Chi router, pgx/v5, Prometheus metrics, sequential SQL migrations, JWT middleware
- React frontend: TypeScript, TailwindCSS 4, React Query, Zustand, React Router v7
- Pattern examples for handler, model, repository — not a full implementation

**Infrastructure**
- Docker Compose stack: backend, frontend, PostgreSQL, Caddy, Mailpit (local email), Prometheus, Grafana, Loki, pgAdmin
- Caddy reverse proxy: routes `/api/*` to backend, everything else to frontend; SSL via Let's Encrypt in production
- GitHub Actions: test → build Docker images → push to GHCR → SSH deploy

**Claude Code configuration**
- Four scoped agents (backend, frontend, infra, product) — each restricted to its own directory
- Rules for Go, React/TypeScript, security, and frontend design
- MCP servers configured: context7 (library docs), GitHub, Brave Search, Postgres
- Agents required to verify builds before reporting done

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

### 4. Start Claude Code

```bash
claude
```

The agents are pre-configured. Ask Claude to add a feature and it will use the appropriate scoped agent for the relevant service.

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

## Claude Code Agents

| Agent | Directory | Handles |
|-------|-----------|---------|
| `backend` | `services/backend/` | handlers, repository, migrations, services |
| `frontend` | `services/frontend/` | components, pages, API client, state |
| `infra` | `services/infra/` | Docker, Caddy, CI/CD, monitoring config |
| `product` | `docs/` | user stories, ADRs, API spec, TODO |

Each agent is restricted to its directory and required to run a build check before reporting done. The agents also require context7 lookups before using any library API — avoids outdated usage from training data.

## Observability

- Backend exposes `/metrics` in Prometheus format
- Prometheus scrapes it; Grafana reads Prometheus and Loki
- Promtail ships container logs to Loki
- Grafana at `http://localhost:3002` (default credentials: admin/admin — change in `.env`)

## Deployment

Push to `main`:
1. GitHub Actions runs backend unit tests + frontend build check
2. Builds Docker images, pushes to `ghcr.io/<your-org>/<your-app>`
3. SSH into your server, pulls new images, restarts containers

Required secrets: `DEPLOY_HOST`, `DEPLOY_USER`, `DEPLOY_SSH_KEY`, plus all app env vars. See `.github/workflows/ci-cd.yml` for the full list.

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
