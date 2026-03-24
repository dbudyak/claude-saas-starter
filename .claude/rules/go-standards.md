# Go Development Standards

## Library API Lookups (context7)

Always use context7 to verify package APIs before writing code. Go modules evolve — method signatures, options structs, and interfaces change between minor versions.

```
mcp__context7__resolve-library-id("pgx")          # then query for specific topic
mcp__context7__resolve-library-id("go-chi/chi")
mcp__context7__resolve-library-id("golang-jwt/jwt")
```

**When to use**: any time you're using a package method you haven't verified in the current session, adding a new dependency, or getting a compile error on a package call.

## Code Style

### Formatting
- Use `gofmt` for all Go code (automatic in most editors)
- Use `goimports` to organize imports automatically
- Follow standard Go conventions

### Naming
- **Packages**: Short, lowercase, no underscores (`auth`, `database`, not `user_auth`)
- **Files**: Lowercase with underscores (`user_repository.go`)
- **Exported identifiers**: PascalCase (`ValidateJWT`, `UserRepository`)
- **Unexported identifiers**: camelCase (`validateToken`, `userRepo`)
- **Constants**: PascalCase or SCREAMING_SNAKE_CASE for package-level
- **Interfaces**: `-er` suffix when single-method (`Reader`, `Writer`), or descriptive names

### Enums (Typed String Constants)

Use typed string constants instead of plain `string` for any field with a fixed set of valid values:

```go
// DO: typed constant enum
type UserRole string

const (
    UserRoleAdmin  UserRole = "admin"
    UserRoleMember UserRole = "member"
)

// DON'T: plain string with no type safety
type User struct {
    Role string // "admin" | "member" — comment is the only contract
}
```

- DB columns stay as `TEXT` — cast at the boundary: `string(role)` when writing, `UserRole(row)` when reading.
- JSON serialisation works automatically since the underlying type is `string`.

## Error Handling

```go
// Standard sentinel errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
)

// Error wrapping
if err := db.QueryRow(...); err != nil {
    return fmt.Errorf("get user: %w", err)
}

// Check every error — never ignore
result, err := doSomething()
if err != nil {
    return err
}
```

## Context Usage

Always accept context as first parameter in repository and service functions:

```go
func (r *Repository) FindUser(ctx context.Context, id string) (*User, error) {
    // ...
}
```

## Database Access (pgx/v5)

```go
// Always use parameterized queries
query := "SELECT id, email FROM users WHERE email = $1"
row := pool.QueryRow(ctx, query, email)

// NEVER string concatenation
// query := "SELECT * FROM users WHERE email = '" + email + "'"  // SQL INJECTION

// Transaction pattern
tx, err := pool.Begin(ctx)
if err != nil {
    return err
}
defer tx.Rollback(ctx) // safe no-op if committed

// ... operations ...

return tx.Commit(ctx)
```

## Testing

### Table-Driven Tests
```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name  string
        email string
        want  bool
    }{
        {"valid", "test@example.com", true},
        {"invalid", "notanemail", false},
        {"empty", "", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := ValidateEmail(tt.email)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Security

```go
// Password hashing
hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
err = bcrypt.CompareHashAndPassword(hash, []byte(password))

// JWT with expiration
claims := jwt.MapClaims{
    "user_id": user.ID,
    "exp":     time.Now().Add(24 * time.Hour).Unix(),
}
```

## Anti-Patterns

- ❌ `result, _ := doSomething()` — ignoring errors
- ❌ Global mutable state
- ❌ Panic in library code (return errors)
- ❌ String concatenation in SQL
- ❌ Plain text passwords
- ❌ Missing context propagation
- ❌ Over-abstraction (YAGNI)

## Project Structure

```
services/backend/
├── config/       # Config loading
├── handlers/     # HTTP handlers
├── middleware/   # Auth, logging, metrics
├── migrations/   # SQL files
├── models/       # Data structures
├── repository/   # DB access
├── services/     # Business logic
└── main.go       # Entry point + router
```

Do NOT use `cmd/` or `internal/` structure for SaaS apps of this scale.
