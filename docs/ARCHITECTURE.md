# Architecture

## System Design

```
[Browser]
    |
    | HTTPS
    v
[Caddy]  ← auto-SSL, reverse proxy
  |    |
  |    +---> /api/* → [Go Backend :8080]
  |                        |
  |                        v
  |                   [PostgreSQL]
  |
  +---------> /* → [React Frontend :80]

[Prometheus] ← scrapes /metrics from backend
[Grafana]    ← queries Prometheus + Loki
[Loki]       ← receives logs via Promtail
[Promtail]   ← scrapes container stdout
```

## Architecture Decision Records

### ADR-001: Go + Chi for backend

**Status**: Accepted

**Context**: Need a simple, performant backend. Team knows Go.

**Decision**: Use Go with Chi router. No framework magic — just HTTP handlers and middleware.

**Consequences**:
- Positive: Fast, explicit, easy to read
- Negative: More boilerplate than Rails/Django-style frameworks

---

### ADR-002: No ORM — raw SQL with pgx

**Status**: Accepted

**Context**: Database queries should be readable and predictable.

**Decision**: Use `pgx/v5` directly with parameterized queries. No GORM or similar.

**Consequences**:
- Positive: Explicit queries, no N+1 surprises, full SQL power
- Negative: More code per query

---

### ADR-003: React Query + Zustand for frontend state

**Status**: Accepted

**Context**: Need server state management and local UI state management.

**Decision**: React Query for server state (fetching, caching, invalidation). Zustand for UI state (auth, modals, etc.).

**Consequences**:
- Positive: Clear separation, minimal boilerplate
- Negative: Two libraries to understand instead of one

---

### ADR-004: Caddy for reverse proxy

**Status**: Accepted

**Context**: Need SSL termination and routing.

**Decision**: Caddy with automatic HTTPS (Let's Encrypt).

**Consequences**:
- Positive: Zero config SSL, simple Caddyfile syntax
- Negative: Less widespread than Nginx, fewer tutorials

---

### ADR-005: Sequential SQL migrations

**Status**: Accepted

**Context**: Need reliable database schema versioning.

**Decision**: Sequential `.sql` files applied on backend startup. Tracked in `schema_migrations` table.

**Consequences**:
- Positive: Simple, no extra tooling (no Flyway/Liquibase/goose needed)
- Negative: Must run backend at least once to apply migrations

## Add New ADRs Below

When making a significant technical decision, document it here following the template above.
