package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"orchestrator/internal/generator"
	"orchestrator/internal/models"
)

func main() {
	// Load config
	file, err := os.Open("configs.json")
	if err != nil {
		panic(fmt.Sprintf("failed to open configs.json: %v", err))
	}
	defer file.Close()

	var environments []*models.NextcloudEnvironment
	if err := json.NewDecoder(file).Decode(&environments); err != nil {
		panic(fmt.Sprintf("failed to decode JSON: %v", err))
	}

	// Generate each instance
	for _, env := range environments {
		outputDir := filepath.Join("./environemnts")
		opt := &generator.NextcloudComposeGeneratorOptions{
			Name:            env.Name,
			Domain:          env.Domain,
			NextcloudConfig: env.NextcloudConfig,
			DatabaseConfig:  env.DatabaseConfig,
			RedisConfig:     env.RedisConfig,
			S3Config:        env.S3Config,
			SMTPConfig:      env.SMTPConfig,
		}
		gen := generator.NewNextcloudComposeGenerator(outputDir, opt)

		if err := gen.Generate(); err != nil {
			fmt.Printf("❌ Failed to generate for %s: %v\n", env.Name, err)
		} else {
			fmt.Printf("✅ Successfully generated config for %s\n", env.Name)
		}
	}
}
