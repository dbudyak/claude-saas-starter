#!/usr/bin/env bash
# deploy.sh — Manual deploy script for development server
# Usage: ./services/infra/scripts/deploy.sh <host> <user> <deploy_dir>

set -euo pipefail

HOST="${1:?Usage: deploy.sh <host> <user> <deploy_dir>}"
USER="${2:?Usage: deploy.sh <host> <user> <deploy_dir>}"
DIR="${3:?Usage: deploy.sh <host> <user> <deploy_dir>}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"

echo "Deploying to $USER@$HOST:$DIR"

# Sync compose files, caddy config, and monitoring configs
rsync -az --delete \
    "$REPO_ROOT/services/infra/docker/docker-compose.yml" \
    "$REPO_ROOT/services/infra/caddy/" \
    "$REPO_ROOT/services/infra/monitoring/" \
    "$USER@$HOST:$DIR/"

# Pull latest images and restart
ssh "$USER@$HOST" "
    cd $DIR
    docker compose -f docker-compose.yml pull
    docker compose -f docker-compose.yml up -d
    docker image prune -f
    echo 'Deploy complete'
"
