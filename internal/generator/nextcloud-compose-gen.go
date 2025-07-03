package generator

import (
	"fmt"
	"orchestrator/internal/models"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type NextcloudComposeGenerator struct {
	basePath string
}

func NewNextcloudComposeGenerator(basePath string) *NextcloudComposeGenerator {
	return &NextcloudComposeGenerator{
		basePath: basePath,
	}
}

func (g *NextcloudComposeGenerator) templateHeader() string {
	return `services:
`
}

func (g *NextcloudComposeGenerator) templateAppService() string {
	return `  app-standalone:
    image: nextcloud:production-apache
    container_name: {{.Name}}-nextcloud-app
    restart: unless-stopped
    ports:
      - "{{.Port}}:80"
    environment:
      - MYSQL_HOST={{.Name}}-mariadb
      - MYSQL_USER={{.DatabaseConfig.User}}
      - MYSQL_PASSWORD={{.DatabaseConfig.Password}}
      - MYSQL_DATABASE={{.DatabaseConfig.Database}}
      - NEXTCLOUD_ADMIN_USER={{.NextcloudConfig.AdminUser}}
      - NEXTCLOUD_ADMIN_PASSWORD={{.NextcloudConfig.AdminPassword}}
      - NEXTCLOUD_TRUSTED_DOMAINS={{.Domain}}
      - REDIS_HOST={{.Name}}-redis
      - REDIS_HOST_PORT=6379
      - REDIS_HOST_PASSWORD={{.RedisConfig.Password}}
      - OBJECTSTORE_S3_BUCKET={{.S3Config.Bucket}}
      - OBJECTSTORE_S3_KEY={{.S3Config.AccessKey}}
      - OBJECTSTORE_S3_SECRET={{.S3Config.SecretKey}}
      - OBJECTSTORE_S3_REGION={{.S3Config.Region}}
      - OBJECTSTORE_S3_HOST={{.S3Config.Endpoint}}
      - OBJECTSTORE_S3_PORT={{.S3Config.Port}}
      - OBJECTSTORE_S3_SSL={{.S3Config.SSL}}
      - OBJECTSTORE_S3_USEPATH_STYLE=true
      - PHP_MEMORY_LIMIT=1024M
      - PHP_UPLOAD_LIMIT=10G
      - APACHE_DISABLE_REWRITE_IP=1
      - TRUSTED_PROXIES=172.16.0.0/12
      - OVERWRITEPROTOCOL=https
      - OVERWRITEHOST={{.Domain}}
      - MAIL_FROM_ADDRESS={{.SMTPConfig.FromAddress}}
      - MAIL_DOMAIN={{.SMTPConfig.Domain}}
      - SMTP_HOST={{.SMTPConfig.Host}}
      - SMTP_SECURE={{.SMTPConfig.Secure}}
      - SMTP_PORT={{.SMTPConfig.Port}}
      - SMTP_AUTHTYPE={{.SMTPConfig.AuthType}}
      - SMTP_NAME={{.SMTPConfig.Name}}
      - SMTP_PASSWORD={{.SMTPConfig.Password}}
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/status.php"]
      interval: 30s
      timeout: 10s
      retries: 3
    volumes:
      - {{.Name}}-nextcloud-data:/var/www/html
      - ./config/php/custom.ini:/usr/local/etc/php/conf.d/custom.ini:ro
    networks:
      - {{.Name}}-nextcloud
    depends_on:
      - mariadb
      - redis
`
}

func (g *NextcloudComposeGenerator) templateMariaDB() string {
	return `  mariadb:
    image: mariadb:lts
    container_name: {{.Name}}-mariadb
    environment:
      - MYSQL_ROOT_PASSWORD={{.DatabaseConfig.RootPassword}}
      - MYSQL_DATABASE={{.DatabaseConfig.Database}}
      - MYSQL_USER={{.DatabaseConfig.User}}
      - MYSQL_PASSWORD={{.DatabaseConfig.Password}}
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 30s
      timeout: 10s
      retries: 3
    restart: unless-stopped
    volumes:
      - {{.Name}}-mariadb-data:/var/lib/mysql
      - ./config/mariadb/my.cnf:/etc/mysql/my.cnf:ro
    networks:
      - {{.Name}}-nextcloud
`
}

func (g *NextcloudComposeGenerator) templateRedis() string {
	return `  redis:
    image: redis:7-alpine
    container_name: {{.Name}}-redis
    command: >
      redis-server 
        --requirepass {{.RedisConfig.Password}}
        --maxmemory 512mb
        --maxmemory-policy allkeys-lru
        --save 900 1
        --save 300 10
        --save 60 10000
        --appendonly yes
        --appendfsync everysec
    volumes:
      - {{.Name}}-redis-data:/data
    networks:
      - {{.Name}}-nextcloud
`
}

func (g *NextcloudComposeGenerator) templateVolumesAndNetworks() string {
	volumes := []string{
		"  " + "{{.Name}}-nextcloud-data:",
		"  " + "{{.Name}}-mariadb-data:",
		"  " + "{{.Name}}-redis-data:",
	}

	return fmt.Sprintf(`
volumes:
%s

networks:
  shared:
	external: true
  {{.Name}}-nextcloud:
    driver: bridge
`, strings.Join(volumes, "\n"))
}

func (g *NextcloudComposeGenerator) Generate(env *models.NextcloudEnvironment) error {
	var parts []string

	parts = append(parts, g.templateHeader())
	parts = append(parts, g.templateAppService())
	parts = append(parts, g.templateMariaDB())
	parts = append(parts, g.templateRedis())
	parts = append(parts, g.templateVolumesAndNetworks())

	composeTemplate := strings.Join(parts, "\n")

	tmpl, err := template.New("docker-compose").Parse(composeTemplate)
	if err != nil {
		return err
	}

	environmentDir := filepath.Join(g.basePath, env.Name)
	if err := os.MkdirAll(environmentDir, 0755); err != nil {
		return err
	}
	if err := g.createDefaultConfigs(environmentDir); err != nil {
		return err
	}

	composePath := filepath.Join(environmentDir, "docker-compose.yml")
	file, err := os.Create(composePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, env)
}

func (g *NextcloudComposeGenerator) createDefaultConfigs(companyDir string) error {
	phpIniPath := filepath.Join(companyDir, "config/php/custom.ini")
	if err := os.MkdirAll(filepath.Dir(phpIniPath), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(phpIniPath, []byte(defaultPHPConfig), 0644); err != nil {
		return err
	}

	myCnfPath := filepath.Join(companyDir, "config/mariadb/my.cnf")
	if err := os.MkdirAll(filepath.Dir(myCnfPath), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(myCnfPath, []byte(defaultMariaDBConfig), 0644); err != nil {
		return err
	}

	return nil
}
