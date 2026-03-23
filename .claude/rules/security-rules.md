# Security Rules

## Secrets

- **Never** commit secrets, credentials, or API keys
- **Never** commit `.env` files (only `.env.example`)
- Use environment variables for all configuration
- Use `openssl rand -hex 32` to generate secrets

```bash
# .env.example (commit this)
JWT_SECRET=generate-with-openssl-rand-hex-32
DATABASE_URL=postgresql://user:password@localhost:5432/dbname
RESEND_API_KEY=re_xxxxx

# .env (NEVER commit this)
JWT_SECRET=actual-secret-2Qx9mK8...
```

## Authentication (Go)

```go
// Hash passwords with bcrypt
hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// JWT with expiration
claims := jwt.MapClaims{
    "user_id": user.ID,
    "exp":     time.Now().Add(24 * time.Hour).Unix(),
}

// Validate on every request (middleware)
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(jwtSecret), nil
})
```

## SQL Injection Prevention

```go
// DO: Parameterized queries
query := "SELECT * FROM users WHERE email = $1"
pool.QueryRow(ctx, query, email)

// NEVER: String concatenation
// query := "SELECT * FROM users WHERE email = '" + email + "'"
```

## Input Validation

```go
func validateCreateRequest(req *CreateRequest) error {
    if req.Email == "" || !isValidEmail(req.Email) {
        return errors.New("invalid email")
    }
    if len(req.Password) < 8 {
        return errors.New("password too short")
    }
    return nil
}
```

## Error Messages

```go
// Return generic messages to clients
http.Error(w, "invalid credentials", http.StatusUnauthorized)

// Log detailed errors server-side
log.Error("login failed", "email", email, "err", err)

// Never reveal: user existence, specific failure reasons, stack traces
```

## XSS Prevention (React)

```tsx
// Safe: React escapes by default
<div>{userInput}</div>

// Dangerous: avoid unless sanitized
<div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(input) }} />
```

## CORS (Go)

```go
// Explicit allowed origins
r.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://yourdomain.com", "http://localhost:3000"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
}))
// Never: AllowedOrigins: []string{"*"} in production
```

## Logging

```go
// Log security events
log.Warn("failed login attempt", "email", email, "ip", ip)

// Never log
// log.Info("login", "password", password)  // NEVER
// log.Info("token", "jwt", token)           // NEVER
```

## Security Checklist (Pre-deployment)

- [ ] No secrets in code or git history
- [ ] All inputs validated server-side
- [ ] Parameterized SQL queries only
- [ ] Passwords hashed with bcrypt
- [ ] JWT has expiration
- [ ] CORS explicitly configured
- [ ] Error messages don't leak information
- [ ] HTTPS enforced
- [ ] Dependencies audited (`govulncheck ./...`, `npm audit`)
