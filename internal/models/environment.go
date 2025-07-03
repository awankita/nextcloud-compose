package models

import "time"

type NextcloudEnvironment struct {
	ID        string    `json:"id" yaml:"id"`
	Name      string    `json:"name" yaml:"name"`
	Domain    string    `json:"domain" yaml:"domain"`
	Port      int       `json:"port" yaml:"port"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`

	NextcloudConfig NextcloudConfig `json:"nextcloud_config" yaml:"nextcloud_config"`
	DatabaseConfig  DatabaseConfig  `json:"database_config" yaml:"database_config"`
	RedisConfig     RedisConfig     `json:"redis_config" yaml:"redis_config"`
	S3Config        S3Config        `json:"s3_config" yaml:"s3_config"`
	SMTPConfig      SMTPConfig      `json:"smtp_config" yaml:"smtp_config"`
}
