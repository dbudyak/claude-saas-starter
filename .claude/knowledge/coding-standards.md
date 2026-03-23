# Coding Standards Index

Quick reference for all standards. Full details in `rules/` files.

## Go (Backend)

See `.claude/rules/go-standards.md` for full details.

**Key rules:**
- Parameterized SQL queries (`$1`, `$2`) — never string concatenation
- Check every error — no `_, _` ignoring
- Context as first parameter in repo/service functions
- Typed string constants for enum-like fields
- Table-driven tests
- bcrypt for passwords, JWT with expiration for auth

## React/TypeScript (Frontend)

See `.claude/rules/react-standards.md` for full details.

**Key rules:**
- Never use `any` — use specific types
- React Query for server state, Zustand for UI state
- API functions in `src/api/` — never fetch directly in components
- Always handle loading/error/empty states
- `memo`, `useMemo`, `useCallback` for performance

## Security

See `.claude/rules/security-rules.md` for full details.

**Key rules:**
- Never commit secrets (`.env` files)
- Parameterized SQL only
- bcrypt passwords
- Generic error messages to clients
- Log security events (failed logins, unauthorized access)

## Design

See `.claude/rules/frontend-aesthetics.md` for full details.

**Key rules:**
- Distinctive fonts (not Inter/Roboto)
- Committed color palette
- Avoid generic "AI slop" patterns
- CSS transitions before JavaScript animations

## Git

See `.claude/rules/monorepo-conventions.md` for full details.

**Commit format**: `feat(scope): description`

**Scopes**: `backend`, `frontend`, `infra`, `docs`
