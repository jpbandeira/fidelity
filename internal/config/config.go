package config

import (
	"fmt"

	"github.com/jp/fidelity/internal/infra/platform"
)

type ServerConfig struct {
	Server       Server        `yaml:"server"`
	Platform     platform.Type `yaml:"platform"`
	DatabaseType string        `yaml:"database_type"`
	Postgres     Postgres      `yaml:"postgres"`
	Log          LogConfig     `yaml:"log"`
}

// Server defines the host address
type Server struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	SpireAudience string `yaml:"spire_audience"`
}

// Postgres represents the Postgres database configuration.
type Postgres struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
}

// DSN returns data source name for connection string from Postgres configuration.
func (p *Postgres) DSN() string {
	if len(p.Password) == 0 {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", p.Host, p.Port, p.Username, p.Database, p.SSLMode)
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Host, p.Port, p.Username, p.Password, p.Database, p.SSLMode)
}

// LogConfig represents the logger configuration
type LogConfig struct {
	Level string `yaml:"level"`
}
