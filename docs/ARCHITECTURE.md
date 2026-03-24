# Architecture

Technical decisions and architectural notes. Use ADR format for significant choices.

## System Overview

```
┌─────────────────────────────────────┐
│         Caddy (port 80/443)         │
│      reverse proxy + auto-SSL       │
└────────────┬────────────────────────┘
             │
      /api/* │  /*
             ▼         ▼
     ┌──────────┐  ┌──────────┐
     │ Backend  │  │ Frontend │
     │ :8080    │  │ nginx:80 │
     └────┬─────┘  └──────────┘
          │
     ┌────▼─────┐
     │ Postgres │
     │ :5432    │
     └──────────┘
```

All services run on a single Docker network. Services communicate via Docker service names (`backend:8080`, `db:5432`), not localhost.

## ADR Format

Record decisions here as they are made.

```markdown
## ADR-001: [Title]

**Date**: YYYY-MM-DD
**Status**: Accepted

### Context
Why did this decision need to be made?

### Decision
What was decided?

### Consequences
- What becomes easier?
- What becomes harder?
```

## ADRs

<!-- Add your decisions below as you make them. Examples:

## ADR-001: PostgreSQL over MySQL

**Date**: ...
**Status**: Accepted

### Context
Needed a relational database. Both are viable options.

### Decision
PostgreSQL. Better JSON support, better full-text search, pgx driver is excellent for Go.

### Consequences
- Team needs PostgreSQL familiarity
- Better query capabilities long-term

-->
