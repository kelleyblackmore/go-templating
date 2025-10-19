package config

// Config represents the data structure for your configuration
type Config struct {
	AppName     string
	Version     string
	Environment string
	Database    struct {
		Host     string
		Port     int
		User     string
		Password string
	}
	Features struct {
		EnableFeatureA bool
		EnableFeatureB bool
	}
}

// NewConfig creates a new Config with sample data
func NewConfig() Config {
	cfg := Config{
		AppName:     "MyApplication",
		Version:     "1.0.0",
		Environment: "production",
	}
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5432
	cfg.Database.User = "admin"
	cfg.Database.Password = "secure_password"
	cfg.Features.EnableFeatureA = true
	cfg.Features.EnableFeatureB = false
	return cfg
}
