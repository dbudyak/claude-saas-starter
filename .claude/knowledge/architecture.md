# System Architecture

## Overview

```
┌─────────────────────────────────────────────────────────┐
│                    Caddy (Port 80/443)                  │
│              Reverse Proxy + Auto-SSL                    │
└────────────────────┬────────────────────┬───────────────┘
                     │                    │
              /api/* │                /* │
                     ▼                    ▼
          ┌──────────────────┐  ┌──────────────────┐
          │   Go Backend     │  │  React Frontend  │
          │   (Chi Router)   │  │  (Nginx static)  │
          │   Port 8080      │  │  Port 80         │
          └────────┬─────────┘  └──────────────────┘
                   │
          ┌────────▼─────────┐
          │   PostgreSQL 15  │
          │   Port 5432      │
          └──────────────────┘
```

## Service Communication

- **Frontend → Backend**: HTTP via Caddy proxy (`/api/*` → `backend:8080`)
- **Backend → Database**: pgx connection pool (`db:5432`)
- **Services communicate via Docker network names** (not localhost)

## Environment Separation

| Environment | How | Notes |
|-------------|-----|-------|
| Local Dev | Docker Compose | No SSL |
| Dev Server | Docker Compose + SSH | Real domain, Caddy SSL |
| Production | Docker Compose + CI/CD | Full secrets, monitoring |

## Monitoring Stack

```
Backend → /metrics → Prometheus → Grafana dashboards
Containers → Promtail → Loki → Grafana logs
```

## Key Coordination Points

### API Contract
`docs/openapi.yaml` is the source of truth. Frontend and backend must agree.

### Database Schema
`docs/data_model.md` describes the schema. Migration files are the implementation.

### Environment Variables
`services/infra/docker/.env.example` documents all required env vars.

## Common Pitfalls

### CORS Issues
Backend CORS must allow the frontend origin. In Docker: frontend origin is the Caddy domain.

### Docker Networking
Services communicate via service names (`backend:8080`), not `localhost:8080`.

### Migration Ordering
Migrations run in filename order. Always use sequential numbering: `001_`, `002_`, etc.

### JWT Token Flow
1. User logs in → backend returns JWT
2. Frontend stores in localStorage
3. All API requests: `Authorization: Bearer <token>`
4. Backend middleware validates on protected routes

## Scalability Notes

This template is designed for MVP → early growth. When you need to scale:
- Add read replicas (PostgreSQL streaming replication)
- Move to container orchestration (Kubernetes or Fly.io)
- Add Redis for caching and sessions
- Move email to background workers (already partially implemented)
- Add CDN for frontend assets
