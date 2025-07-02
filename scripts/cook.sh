#!/bin/bash

set -e

PROJECT_NAME="nextcloud"

# Compose files in order
COMPOSE_FILES=(
  docker-compose.base.yml
  docker-compose.web.yml
  # docker-compose.cron.yml
  # docker-compose.elasticsearch.yml
  # docker-compose.observability.yml
  # docker-compose.backup.yml
  # docker-compose.watchtower.yml
)

# Compose file args
COMPOSE_ARGS=""
for file in "${COMPOSE_FILES[@]}"; do
  COMPOSE_ARGS="$COMPOSE_ARGS -f $file"
done

# Colors
GREEN='\033[0;32m'
NC='\033[0m'

function up() {
  echo -e "${GREEN}Starting Nextcloud stack...${NC}"
  docker compose -p $PROJECT_NAME $COMPOSE_ARGS up -d
}

function down() {
  echo -e "${GREEN}Stopping and removing Nextcloud stack...${NC}"
  docker compose -p $PROJECT_NAME $COMPOSE_ARGS down
}

function restart() {
  echo -e "${GREEN}Restarting Nextcloud stack...${NC}"
  down
  up
}

function logs() {
  echo -e "${GREEN}Showing logs... (Ctrl+C to stop)${NC}"
  docker compose -p $PROJECT_NAME $COMPOSE_ARGS logs -f
}

function ps() {
  docker compose -p $PROJECT_NAME $COMPOSE_ARGS ps
}

function fresh() {
  echo -e "${RED}⚠️ WARNING: This will REMOVE all containers, volumes, and data.${NC}"
  read -p "Are you sure? [y/N]: " confirm
  if [[ "$confirm" =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}Bringing down stack and removing volumes...${NC}"
    docker compose -p "$PROJECT_NAME" $COMPOSE_ARGS down -v

    echo -e "${GREEN}Removing dangling volumes (if any)...${NC}"
    docker volume prune -f

    echo -e "${GREEN}Cleaning bind-mounted folders...${NC}"
    rm -rf ./data ./mariadb ./redis ./elasticsearch ./prometheus ./logs ./backups

    echo -e "${GREEN}Recreating shared Docker network (if missing)...${NC}"
    if ! docker network ls | grep -q "$SHARED_NETWORK"; then
      docker network create --driver bridge "$SHARED_NETWORK"
      echo -e "${GREEN}Created network: $SHARED_NETWORK${NC}"
    else
      echo -e "${GREEN}Shared network already exists.${NC}"
    fi

    echo -e "${GREEN}✔ Fresh reset complete. You can now run './nextcloud-stack.sh up'${NC}"
  else
    echo "Aborted."
  fi
}

function help() {
  echo -e "${GREEN}Nextcloud Docker Stack Helper${NC}"
  echo "Usage: $0 {up|down|restart|logs|ps|fresh|help}"
}

case "$1" in
  up) up ;;
  down) down ;;
  restart) restart ;;
  logs) logs ;;
  ps) ps ;;
  fresh) fresh ;;
  help|*) help ;;
esac
