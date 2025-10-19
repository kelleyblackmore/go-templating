package main

import (
	"log"
	"os"

	"github.com/kelleyblackmore/go-templating/internal/config"
	"github.com/kelleyblackmore/go-templating/internal/template"
)

func main() {
	// Create sample configuration data
	cfg := config.NewConfig()

	// Create template processor
	processor := template.NewProcessor("templates/config.tmpl")

	// Create an output file for the generated configuration
	outputFile, err := os.Create("generated_config.conf")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	// Process the template with the configuration data
	err = processor.Process(outputFile, cfg)
	if err != nil {
		log.Fatalf("Error processing template: %v", err)
	}

	log.Println("Configuration file generated successfully!")
}
