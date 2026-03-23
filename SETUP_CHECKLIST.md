# Setup Checklist

Complete these steps after cloning the template to customize it for your project.

## 1. Project Identity

- [ ] Choose a project name (e.g., `myapp`)
- [ ] Update `go.mod`: change `module github.com/yourorg/yourapp` to your module path
- [ ] Update `package.json`: change `"name": "saas-starter"` to your app name
- [ ] Update `CLAUDE.md`: replace the template description with your product description
- [ ] Update `docs/PROJECT.md`: describe your product

## 2. Environment Configuration

```bash
cp services/infra/docker/.env.example services/infra/docker/.env
```

Fill in:
- [ ] `DATABASE_URL` — PostgreSQL connection string
- [ ] `JWT_SECRET` — generate with `openssl rand -hex 32`
- [ ] `DOMAIN` — your domain name
- [ ] `APP_URL` — full URL including https://

## 3. Backend Customization

- [ ] Update module name in `services/backend/go.mod`
- [ ] Add your domain models to `services/backend/models/`
- [ ] Create migrations in `services/backend/migrations/`
- [ ] Add handlers to `services/backend/handlers/`
- [ ] Update routes in `services/backend/main.go`

## 4. Frontend Customization

- [ ] Update app name in `services/frontend/package.json`
- [ ] Add your branding (logo, colors, fonts) in `services/frontend/src/`
- [ ] Update page titles and meta tags
- [ ] Add your pages to `services/frontend/src/pages/`
- [ ] Configure API base URL in `services/frontend/src/api/client.ts`

## 5. Infrastructure Customization

- [ ] Update container image names in `services/infra/docker/docker-compose.yml`
  - Replace `ghcr.io/yourorg/yourapp/backend:latest`
  - Replace `ghcr.io/yourorg/yourapp/frontend:latest`
- [ ] Update Caddyfile with your domain
- [ ] Update GitHub Actions workflow with your GHCR org/repo
- [ ] Update project name prefix on Docker containers

## 6. Documentation

- [ ] Fill in `docs/PROJECT.md` with your product description
- [ ] Update `docs/ARCHITECTURE.md` with any architectural decisions
- [ ] Document your data model in `docs/data_model.md`
- [ ] Create API spec in `docs/openapi.yaml`

## 7. CI/CD Setup

Add these secrets to your GitHub repository (`Settings → Secrets → Actions`):

- [ ] `DEPLOY_HOST` — your server IP or hostname
- [ ] `DEPLOY_USER` — SSH user on server
- [ ] `DEPLOY_SSH_KEY` — SSH private key
- [ ] `DATABASE_URL` — production database URL
- [ ] `JWT_SECRET` — production JWT secret
- [ ] `DOMAIN` — your domain
- [ ] `APP_URL` — full app URL

## 8. Server Setup

On your deployment server:
```bash
# Install Docker
curl -fsSL https://get.docker.com | sh

# Create app directory
mkdir -p /opt/myapp

# Set up SSH access for GitHub Actions
# Add your deploy SSH public key to ~/.ssh/authorized_keys
```

## 9. MCP Servers (Optional but Recommended)

Install MCP servers for enhanced Claude Code capabilities:

```bash
# Documentation queries (already configured)
npm install -g @upstash/context7-mcp

# GitHub integration
npm install -g @modelcontextprotocol/server-github

# Brave Search
npm install -g @modelcontextprotocol/server-brave-search
```

Configure API keys in `.claude/settings.json` under `mcpServers`.

## 10. Claude Code Agent Setup

The agents are pre-configured and ready to use. To start:

```bash
claude
```

Verify agents work:
- Ask Claude to "check the backend health handler" — it should use the backend agent
- Ask Claude to "add a button to the home page" — it should use the frontend agent

## Done!

Once you've completed this checklist, you're ready to start building your product. The agents will help you implement features, fix bugs, and maintain quality throughout development.
