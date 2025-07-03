package main

import (
	"fmt"
	"orchestrator/internal/generator"
	"orchestrator/internal/models"
)

func main() {
	gen := generator.NewNextcloudComposeGenerator("./output")

	env := &models.NextcloudEnvironment{
		Name:   "awankita",
		Port:   8080,
		Domain: "awankita.local",
		NextcloudConfig: models.NextcloudConfig{
			AdminUser:     "admin",
			AdminPassword: "admin123",
		},
		DatabaseConfig: models.DatabaseConfig{
			RootPassword: "rootpassword",
			Database:     "nextcloud",
			User:         "nextclouduser",
			Password:     "nextcloudpassword",
		},
		RedisConfig: models.RedisConfig{
			Password: "redispassword",
		},
		S3Config: models.S3Config{
			Bucket:    "nextcloud-bucket",
			AccessKey: "s3accesskey",
			SecretKey: "s3secretkey",
			Region:    "us-west-1",
			Endpoint:  "s3.amazonaws.com",
			Port:      "443",
			SSL:       "true",
		},
		SMTPConfig: models.SMTPConfig{
			FromAddress: "",
			Domain:      "awankita.local",
			Host:        "smtp.awankita.local",
			Secure:      "tls",
			Port:        "587",
			AuthType:    "login",
			Name:        "awankita",
			Password:    "smtppassword",
		},
	}

	if err := gen.Generate(env); err != nil {
		panic(fmt.Sprintf("Failed to generate docker-compose.yml: %v", err))
	}
}
