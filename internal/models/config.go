package models

type DatabaseConfig struct {
	RootPassword string `json:"root_password" yaml:"root_password"`
	Database     string `json:"database" yaml:"database"`
	User         string `json:"user" yaml:"user"`
	Password     string `json:"password" yaml:"password"`
}

type RedisConfig struct {
	Password string `json:"password" yaml:"password"`
}

type NextcloudConfig struct {
	AdminUser     string `json:"admin_user" yaml:"admin_user"`
	AdminPassword string `json:"admin_password" yaml:"admin_password"`
}

type S3Config struct {
	Bucket    string `json:"bucket" yaml:"bucket"`
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	Region    string `json:"region" yaml:"region"`
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	Port      string `json:"port" yaml:"port"`
	SSL       string `json:"ssl" yaml:"ssl"`
}

type SMTPConfig struct {
	FromAddress string `json:"from_address" yaml:"from_address"`
	Domain      string `json:"domain" yaml:"domain"`
	Host        string `json:"host" yaml:"host"`
	Secure      string `json:"secure" yaml:"secure"`
	Port        string `json:"port" yaml:"port"`
	AuthType    string `json:"auth_type" yaml:"auth_type"`
	Name        string `json:"name" yaml:"name"`
	Password    string `json:"password" yaml:"password"`
}
