# ============================================================
# SaaS Starter — Development Commands
# ============================================================
# Usage: make <target>
# Run 'make help' to see all available commands.

# Configuration — override these for your project
APP_NAME ?= myapp
BACKEND_DIR := services/backend
FRONTEND_DIR := services/frontend
INFRA_DIR := services/infra
COMPOSE_FILE := $(INFRA_DIR)/docker/docker-compose.yml
ENV_FILE := $(INFRA_DIR)/docker/.env

# Server configuration — set these for your deployment
DEPLOY_HOST ?= your-server-ip
DEPLOY_USER ?= ubuntu
DEPLOY_DIR ?= /opt/$(APP_NAME)
SSH_KEY ?= ~/.ssh/id_rsa

.PHONY: help
help: ## Show this help message
	@echo "SaaS Starter — Available Commands"
	@echo "================================="
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# ============================================================
# Local Development (Docker)
# ============================================================

.PHONY: local-up
local-up: ## Start all services with Docker Compose
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) up -d --build
	@echo "Services started. Visit http://localhost"

.PHONY: local-down
local-down: ## Stop all services
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) down

.PHONY: local-restart
local-restart: ## Restart all services
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) restart

.PHONY: local-logs
local-logs: ## Tail logs from all services
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) logs -f

.PHONY: local-logs-backend
local-logs-backend: ## Tail backend logs
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) logs -f backend

.PHONY: local-logs-frontend
local-logs-frontend: ## Tail frontend logs
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) logs -f frontend

.PHONY: local-status
local-status: ## Show status of all services
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) ps

.PHONY: local-reset
local-reset: ## Stop and remove all containers + volumes (destructive!)
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) down -v
	@echo "All containers and volumes removed."

# ============================================================
# Local Development (Without Docker)
# ============================================================

.PHONY: backend-run
backend-run: ## Run backend locally (requires local PostgreSQL)
	cd $(BACKEND_DIR) && go run .

.PHONY: frontend-run
frontend-run: ## Run frontend dev server locally
	cd $(FRONTEND_DIR) && npm run dev

# ============================================================
# Testing
# ============================================================

.PHONY: backend-test
backend-test: ## Run backend unit tests
	cd $(BACKEND_DIR) && go test ./handlers/ ./services/ ./config/ -v

.PHONY: backend-integration
backend-integration: ## Run backend integration tests (requires Docker)
	cd $(BACKEND_DIR) && go test -tags=integration ./repository/ -v

.PHONY: backend-test-all
backend-test-all: backend-test backend-integration ## Run all backend tests

.PHONY: frontend-test
frontend-test: ## Run frontend tests
	cd $(FRONTEND_DIR) && npm test

.PHONY: frontend-build
frontend-build: ## Build frontend for production
	cd $(FRONTEND_DIR) && npm run build

.PHONY: backend-lint
backend-lint: ## Lint Go code
	cd $(BACKEND_DIR) && go vet ./...

.PHONY: test-all
test-all: backend-test frontend-build ## Run all tests (CI-ready)

# ============================================================
# Database
# ============================================================

.PHONY: db-migrate
db-migrate: ## Run database migrations (via backend on start)
	@echo "Migrations run automatically on backend startup."
	@echo "To run manually: docker compose exec backend ./server migrate"

.PHONY: db-shell
db-shell: ## Open PostgreSQL shell
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) exec db psql -U $${DB_USER:-postgres} -d $${DB_NAME:-myapp}

.PHONY: db-backup
db-backup: ## Backup database to file
	docker compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) exec db \
		pg_dump -U $${DB_USER:-postgres} $${DB_NAME:-myapp} > backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "Backup saved."

# ============================================================
# Deployment
# ============================================================

.PHONY: dev-deploy
dev-deploy: ## Deploy to development server
	@echo "Deploying to $(DEPLOY_HOST)..."
	./$(INFRA_DIR)/scripts/deploy.sh $(DEPLOY_HOST) $(DEPLOY_USER) $(DEPLOY_DIR)

.PHONY: dev-logs
dev-logs: ## View logs on development server
	ssh -i $(SSH_KEY) $(DEPLOY_USER)@$(DEPLOY_HOST) "cd $(DEPLOY_DIR) && docker compose logs -f"

.PHONY: dev-status
dev-status: ## Check service status on development server
	ssh -i $(SSH_KEY) $(DEPLOY_USER)@$(DEPLOY_HOST) "cd $(DEPLOY_DIR) && docker compose ps"

.PHONY: dev-ssh
dev-ssh: ## SSH into development server
	ssh -i $(SSH_KEY) $(DEPLOY_USER)@$(DEPLOY_HOST)

# ============================================================
# Code Quality
# ============================================================

.PHONY: deps-backend
deps-backend: ## Update Go dependencies
	cd $(BACKEND_DIR) && go get -u ./... && go mod tidy

.PHONY: deps-frontend
deps-frontend: ## Update npm dependencies
	cd $(FRONTEND_DIR) && npm update

.PHONY: security-check
security-check: ## Check for security vulnerabilities
	@echo "Checking Go vulnerabilities..."
	cd $(BACKEND_DIR) && govulncheck ./... 2>/dev/null || echo "Install: go install golang.org/x/vuln/cmd/govulncheck@latest"
	@echo "Checking npm vulnerabilities..."
	cd $(FRONTEND_DIR) && npm audit

# ============================================================
# Setup
# ============================================================

.PHONY: setup
setup: ## Initial project setup (copy env, install deps)
	@if [ ! -f $(ENV_FILE) ]; then \
		cp $(INFRA_DIR)/docker/.env.example $(ENV_FILE); \
		echo "Created $(ENV_FILE) — please edit it with your values"; \
	else \
		echo "$(ENV_FILE) already exists"; \
	fi
	cd $(FRONTEND_DIR) && npm install
	@echo "Setup complete. Run 'make local-up' to start."
