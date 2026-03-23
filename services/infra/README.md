# Infrastructure

Docker Compose-based infrastructure for local development and deployment.

## Structure

```
infra/
├── docker/
│   ├── docker-compose.yml   # Full application stack
│   └── .env.example          # Required environment variables
├── caddy/
│   └── Caddyfile             # Reverse proxy configuration
├── monitoring/
│   ├── prometheus.yml        # Metrics scraping config
│   ├── grafana/              # Dashboard provisioning
│   ├── loki/                 # Log aggregation config
│   └── promtail/             # Log shipping config
└── scripts/
    └── deploy.sh             # SSH deployment script
```

## Services

| Service | Port (host) | Description |
|---------|------------|-------------|
| App | 80 | Main application (via Caddy) |
| pgAdmin | 5050 | PostgreSQL GUI |
| Grafana | 3002 | Metrics dashboards |
| Mailpit | 8025 | Email catcher (dev only) |

## Quick Start

```bash
# From project root
cp services/infra/docker/.env.example services/infra/docker/.env
# Edit .env with your values
make local-up
```

## Environment Variables

See `.env.example` for all required variables with descriptions.

**Required:**
- `DATABASE_URL` — full PostgreSQL connection string
- `DB_USER`, `DB_PASSWORD`, `DB_NAME` — database credentials
- `JWT_SECRET` — JWT signing secret (min 32 chars, generate with `openssl rand -hex 32`)
- `APP_URL` — full URL of your app (e.g. `http://localhost` for dev)

**Optional:**
- `RESEND_API_KEY` — for production email (leave blank to use Mailpit in dev)
- `FROM_EMAIL` — sender email address
- `GF_SECURITY_ADMIN_PASSWORD` — Grafana admin password (default: `admin`)

## Monitoring

Grafana is pre-provisioned with:
- **Prometheus** datasource for backend metrics
- **Loki** datasource for container logs

Access Grafana at `http://localhost:3002` (admin/admin by default — change in `.env`).

## Deployment

The `deploy.sh` script deploys via SSH:

```bash
# Set your server details in Makefile or pass as args
make dev-deploy DEPLOY_HOST=your-server-ip DEPLOY_USER=ubuntu
```

For production, use GitHub Actions (see `.github/workflows/ci-cd.yml`).

## Production: Enabling SSL

Edit `caddy/Caddyfile` and uncomment the production section:

```caddyfile
yourdomain.com {
    reverse_proxy /api/* backend:8080
    reverse_proxy frontend:80
}
```

Caddy handles SSL certificates automatically via Let's Encrypt.
