# Data Model

> **This file is a template. Replace the example tables below with your actual schema.**
>
> This document is the source of truth for the database schema — keep it in sync with migration files in `services/backend/migrations/`.

## Tables

### Example: users

Registered application users.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| `id` | UUID | PK, default `gen_random_uuid()` | Unique identifier |
| `email` | TEXT | NOT NULL, UNIQUE | Login email address |
| `password_hash` | TEXT | NOT NULL | bcrypt hash of password |
| `role` | TEXT | NOT NULL, default `'member'` | User role: `admin` or `member` |
| `created_at` | TIMESTAMPTZ | NOT NULL, default `NOW()` | Creation timestamp |

**Indexes**:
- `idx_users_email` on `(email)`

---

## Migration Files

| File | Description |
|------|-------------|
| `001_init.sql` | Create users table |

## Guidelines

- When adding a table, document it here first (or update after migration)
- Use `UUID` for all primary keys (`gen_random_uuid()`)
- All timestamps use `TIMESTAMPTZ` (not `TIMESTAMP`)
- Use `TEXT` for string columns (not `VARCHAR(n)`)
- Add indexes for all foreign keys and frequently filtered columns
- Soft deletes: use `deleted_at TIMESTAMPTZ` column if records should be recoverable
