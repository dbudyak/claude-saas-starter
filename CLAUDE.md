# SaaS Starter

**Template**: Go + React SaaS Starter
**Purpose**: Fast full-stack SaaS development with Claude Code multi-agent orchestration

## Project Overview

This is a starter template for building B2B/B2C SaaS applications. Replace this section with your product description once you've cloned the template.

**Stack**: Go (Chi router) backend + React/TypeScript frontend + PostgreSQL + Docker

## Key Design Decisions

- Chi router for idiomatic Go HTTP handling
- pgx/v5 for PostgreSQL (no ORM — SQL is explicit and readable)
- JWT authentication with bcrypt password hashing
- React Query for server state, Zustand for UI state
- TailwindCSS 4 for styling
- Docker Compose for local development and deployment
- Caddy for reverse proxy with auto-SSL in production
- Prometheus + Grafana for observability

## Documentation Map

| Document | Purpose |
|----------|---------|
| `docs/PROJECT.md` | Product overview, goals, user types |
| `docs/ARCHITECTURE.md` | Technical decisions (ADRs) |
| `docs/data_model.md` | Database schema (source of truth) |
| `docs/openapi.yaml` | API specification (source of truth) |
| `docs/TODO.md` | Current tasks and backlog |
| `docs/features/` | Feature specifications by domain |

## Project Structure

```
saas-starter/
├── .claude/
│   ├── agents/           # Agent definitions
│   ├── knowledge/        # Shared knowledge
│   ├── rules/            # Development rules
│   └── settings.json     # Permissions + MCP config
├── services/
│   ├── backend/          # Go backend service
│   │   ├── config/       # Configuration loading
│   │   ├── handlers/     # HTTP request handlers
│   │   ├── middleware/   # Auth, logging, metrics
│   │   ├── migrations/   # SQL migration files
│   │   ├── models/       # Data structures
│   │   ├── repository/   # Database access layer
│   │   ├── services/     # Business logic
│   │   └── main.go       # Entry point, router setup
│   ├── frontend/         # React/TypeScript application
│   │   └── src/
│   │       ├── api/      # API client functions
│   │       ├── components/ # Reusable UI components
│   │       ├── pages/    # Route pages
│   │       ├── store/    # Zustand state management
│   │       └── types/    # TypeScript type definitions
│   └── infra/            # Docker Compose, Caddy, monitoring, scripts
├── docs/                 # All project documentation
└── Makefile              # Development commands
```

## Technology Stack

### Backend
- **Language**: Go 1.24+
- **Framework**: Chi (`github.com/go-chi/chi/v5`)
- **Database**: PostgreSQL 15 via `pgx/v5`
- **Auth**: JWT (`golang-jwt/jwt/v5`), bcrypt passwords
- **Metrics**: Prometheus (`prometheus/client_golang`)
- **Testing**: Standard `testing` + `testcontainers-go`

### Frontend
- **Framework**: React 19 + TypeScript
- **Build Tool**: Vite 6
- **Styling**: TailwindCSS 4
- **State**: React Query + Zustand
- **i18n**: i18next

### Infrastructure
- **Hosting**: Any Linux server with Docker
- **Container**: Docker / Docker Compose
- **Reverse Proxy**: Caddy (auto-SSL)
- **Monitoring**: Prometheus + Grafana + Loki + Promtail
- **CI/CD**: GitHub Actions + GHCR

## Available Agents

| Agent | Scope | Use For |
|-------|-------|---------|
| `backend` | `services/backend/` | Go API development, DB queries, business logic |
| `frontend` | `services/frontend/` | React UI, components, pages, API integration |
| `infra` | `services/infra/` | Docker, deployment, server config, monitoring |
| `product` | `docs/` | Requirements, user stories, documentation |

## Agent Usage

Agents are self-contained with full domain knowledge. Each agent:
1. Has complete knowledge of its domain embedded in the agent file
2. Must verify changes with `git status` before reporting success
3. Stays within its defined scope
4. Reads and follows all rules in `.claude/rules/`

**Note**: Tool names in agent frontmatter MUST use PascalCase (`Read, Write, Edit, Bash, Glob, Grep`), not lowercase.

## Development Workflow

### Local Development (Docker)
```bash
make local-up          # Start all services
make local-down        # Stop all services
make local-logs        # View logs
```

### Local Development (No Docker)
```bash
make backend-run       # Run backend locally
make frontend-run      # Run frontend locally
```

### Testing
```bash
make backend-test          # Unit tests
make backend-integration   # Integration tests (requires Docker)
make frontend-test         # Frontend tests
```

### Deployment
```bash
make dev-deploy        # Deploy to development server
make dev-logs          # View dev server logs
```

## Superpowers Skills

If the [superpowers](https://github.com/southbridgeai/superpowers) skill pack is installed, the following skills apply to this project:

| Skill | When it triggers |
|-------|-----------------|
| `brainstorming` | Before building any new feature |
| `test-driven-development` | Before writing implementation code |
| `systematic-debugging` | When encountering a bug or unexpected behavior |
| `writing-plans` | When given a multi-step spec before touching code |
| `executing-plans` | When running a written plan in a new session |
| `verification-before-completion` | Before claiming any task is done |
| `finishing-a-development-branch` | When implementation is complete and ready to integrate |

Install: follow instructions at the superpowers repository.

## Important Context

- **Source of truth for schema**: `docs/data_model.md` and migration files in `services/backend/migrations/`
- **Source of truth for API**: `docs/openapi.yaml`
- **Authentication**: JWT tokens, verify on every protected route
- **Error handling**: Return generic errors to clients, log detailed errors server-side
- **Secrets**: Never commit `.env` files or credentials
