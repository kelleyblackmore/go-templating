package main

import (
	"fmt"
	"os"
	"text/template"
)

// Config represents the application configuration structure
type Config struct {
	AppName     string
	Version     string
	Environment string
	Server      ServerConfig
	Database    DatabaseConfig
	Features    []string
	Debug       bool
}

// ServerConfig holds server-related settings
type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

// DatabaseConfig holds database-related settings
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	Username string
}

func main() {
	// Create sample configuration data
	config := Config{
		AppName:     "MyApplication",
		Version:     "1.0.0",
		Environment: "production",
		Server: ServerConfig{
			Host:    "localhost",
			Port:    8080,
			Timeout: 30,
		},
		Database: DatabaseConfig{
			Driver:   "postgresql",
			Host:     "db.example.com",
			Port:     5432,
			Name:     "myapp_db",
			Username: "admin",
		},
		Features: []string{"logging", "monitoring", "caching"},
		Debug:    false,
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("config.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Execute the template with the configuration data
	err = tmpl.Execute(os.Stdout, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		os.Exit(1)
	}
}
