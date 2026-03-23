# Monorepo Conventions

## Service Independence

- Services must NOT import from each other
- Communicate via HTTP API (documented in `docs/openapi.yaml`)
- Each service can be built, tested, and deployed independently

## Directory Structure

```
services/
├── backend/    # Go API
├── frontend/   # React app
└── infra/      # Deployment config
docs/           # Documentation
.claude/        # Claude Code config
.github/        # CI/CD workflows
```

## File Naming

- **Go files**: lowercase_with_underscores (`user_repository.go`)
- **TypeScript/React components**: PascalCase (`UserProfile.tsx`)
- **TypeScript utilities**: camelCase (`authHelpers.ts`)
- **Docs**: UPPERCASE for main docs (`README.md`, `ARCHITECTURE.md`)
- **Config**: lowercase with dots (`.env.example`, `docker-compose.yml`)

## Git Workflow

### Branch Naming
- `feature/description` — New features
- `fix/description` — Bug fixes
- `docs/description` — Documentation
- `infra/description` — Infrastructure

### Commit Messages (Conventional Commits)
```
feat(backend): add user invitation endpoint
fix(frontend): correct mobile navigation overlap
docs: update API authentication docs
infra: add Prometheus scrape config for backend
```

## Source of Truth

| Concern | Source |
|---------|--------|
| Database schema | `docs/data_model.md` + migration files |
| API contract | `docs/openapi.yaml` |
| Environment vars | `services/infra/docker/.env.example` |
| Architecture decisions | `docs/ARCHITECTURE.md` |

## Version Control

- Never commit `.env` files
- Never commit compiled binaries
- Never commit `node_modules/`
- Commit `go.sum` and `package-lock.json` (reproducible builds)

## Documentation Updates

When adding a feature:
1. Update `docs/openapi.yaml` if API changes
2. Update `docs/data_model.md` if schema changes
3. Update `docs/TODO.md` to mark items complete
4. Add migration file if database changes
