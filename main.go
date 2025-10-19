package main

import (
	"log"
	"os"
	"text/template"
)

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

func main() {
	// Sample configuration data
	cfg := Config{
		AppName:     "MyApplication",
		Version:     "1.0.0",
		Environment: "production",
		Database: struct {
			Host     string
			Port     int
			User     string
			Password string
		}{
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "secure_password",
		},
		Features: struct {
			EnableFeatureA bool
			EnableFeatureB bool
		}{
			EnableFeatureA: true,
			EnableFeatureB: false,
		},
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("config.tmpl")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Create an output file for the generated configuration
	outputFile, err := os.Create("generated_config.conf")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	// Execute the template with the configuration data and write to the output file
	err = tmpl.Execute(outputFile, cfg)
	if err != nil {
		outputFile.Close()
		log.Fatalf("Error executing template: %v", err)
	}

	// Close the file and check for errors
	if err := outputFile.Close(); err != nil {
		log.Fatalf("Error closing output file: %v", err)
	}

	log.Println("Configuration file generated successfully!")
}
