# 📦 Nextcloud Production Stack (Docker Compose)

This project contains a production-grade, modular Docker Compose setup to run **Nextcloud** with optional services like:

- MariaDB, Redis, Cron
- Nginx (FPM reverse proxy)
- Elasticsearch (Fulltext Search)
- Prometheus + Node Exporter (Monitoring)
- Watchtower (Auto-updates)
- Backup job (manual or scheduled)

> 🛡 Uses a `.env` file for configuration and a helper script to manage the stack.


## 📁 Folder Structure

```
.
├── .env # Your environment variables
├── nextcloud-stack.sh # Control script for up/down/fresh
├── docker-compose.base.yml # Core services: app, db, redis
├── docker-compose.web.yml # Nginx reverse proxy
├── docker-compose.cron.yml # Background jobs
├── docker-compose.elasticsearch.yml
├── docker-compose.monitoring.yml
├── docker-compose.backup.yml
├── docker-compose.watchtower.yml
├── config/ # Config overrides (nginx, php, etc.)
├── logs/ # Log mounts
├── backups/ # MariaDB backups
├── data/, mariadb/, redis/, etc. # Volume mounts
```


## 🚀 Getting Started

### 1. Prepare `.env`

Edit the `.env` file to match your production secrets and settings.

### 2. Make the script executable

```bash
chmod +x nextcloud-stack.sh
```

### 3. Start the entire stack

```bash
./nextcloud-stack.sh up
```

---

## 🔧 Commands

| Command                        | Description                                    |
| ------------------------------ | ---------------------------------------------- |
| `./nextcloud-stack.sh up`      | Start the full stack in background (detached)  |
| `./nextcloud-stack.sh down`    | Stop and remove all containers                 |
| `./nextcloud-stack.sh restart` | Full restart of the stack                      |
| `./nextcloud-stack.sh logs`    | Follow logs of all containers                  |
| `./nextcloud-stack.sh ps`      | Show status of running containers              |
| `./nextcloud-stack.sh --fresh` | ⚠ Full reset: remove containers, volumes, data |

---

## 📡 Network

This stack uses a shared external Docker network:
```yaml
networks:
  shared:
    external: true
```
If it doesn't exist yet, the script will create it automatically on --fresh.


## 📦 Backups

Backups are handled by:
```bash
./nextcloud-stack.sh --profile backup up db-backup
```
Customize scripts/backup.sh to your needs.


## 📊 Monitoring

- Prometheus runs on port 9090

- Node Exporter monitors host metrics

- Easily integrate into your own Grafana setup


## 🔄 Auto Updates

Watchtower monitors containers and auto-updates them (daily at 2:00 AM).

Configure email alerts using environment variables in .env.


## ✅ Requirements

- Docker Engine

- Linux or WSL (tested on Ubuntu 22+)

- .env file with all required variables


## 📌 Notes

- This stack assumes you're using an external reverse proxy, like Nginx Proxy Manager or Traefik, to expose web to the public.

- You can disable unused services by removing Compose files or commenting out services.

- You can use this with systemd for auto-start on boot.