package config

type ServerConfig struct {
	Server Server `yaml:"server"`
	// Platform     platform.Type `yaml:"platform"`
	DatabaseType string    `yaml:"database_type"`
	Postgres     Postgres  `yaml:"postgres"`
	Log          LogConfig `yaml:"log"`
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

// LogConfig represents the logger configuration
type LogConfig struct {
	Level string `yaml:"level"`
}
