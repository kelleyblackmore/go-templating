package main

import (
	"strings"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     interface{}
		expected string
		wantErr  bool
	}{
		{
			name:     "Simple variable substitution",
			template: "Hello {{.Name}}!",
			data:     map[string]string{"Name": "World"},
			expected: "Hello World!",
			wantErr:  false,
		},
		{
			name:     "Nested struct fields",
			template: "Host: {{.Server.Host}}, Port: {{.Server.Port}}",
			data: struct {
				Server struct {
					Host string
					Port int
				}
			}{
				Server: struct {
					Host string
					Port int
				}{
					Host: "localhost",
					Port: 8080,
				},
			},
			expected: "Host: localhost, Port: 8080",
			wantErr:  false,
		},
		{
			name:     "Range over slice",
			template: "{{range .Items}}{{.}} {{end}}",
			data:     map[string][]string{"Items": {"one", "two", "three"}},
			expected: "one two three ",
			wantErr:  false,
		},
		{
			name:     "Conditional rendering",
			template: "{{if .Debug}}Debug mode enabled{{else}}Debug mode disabled{{end}}",
			data:     map[string]bool{"Debug": true},
			expected: "Debug mode enabled",
			wantErr:  false,
		},
		{
			name:     "Invalid template syntax",
			template: "{{.Name",
			data:     map[string]string{"Name": "Test"},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RenderTemplate(tt.template, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("RenderTemplate() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRenderTemplateToWriter(t *testing.T) {
	template := "AppName: {{.AppName}}, Version: {{.Version}}"
	data := map[string]string{
		"AppName": "TestApp",
		"Version": "1.0.0",
	}

	var buf strings.Builder
	err := RenderTemplateToWriter(&buf, template, data)
	if err != nil {
		t.Fatalf("RenderTemplateToWriter() error = %v", err)
	}

	expected := "AppName: TestApp, Version: 1.0.0"
	if buf.String() != expected {
		t.Errorf("RenderTemplateToWriter() = %q, want %q", buf.String(), expected)
	}
}

func TestConfigStructTemplate(t *testing.T) {
	config := Config{
		AppName:     "TestApp",
		Version:     "1.0.0",
		Environment: "test",
		Server: ServerConfig{
			Host:    "localhost",
			Port:    3000,
			Timeout: 10,
		},
		Database: DatabaseConfig{
			Driver:   "sqlite",
			Host:     "localhost",
			Port:     0,
			Name:     "test.db",
			Username: "test",
		},
		Features: []string{"feature1", "feature2"},
		Debug:    true,
	}

	template := "App: {{.AppName}}, Env: {{.Environment}}, Server: {{.Server.Host}}:{{.Server.Port}}"
	result, err := RenderTemplate(template, config)
	if err != nil {
		t.Fatalf("RenderTemplate() error = %v", err)
	}

	expected := "App: TestApp, Env: test, Server: localhost:3000"
	if result != expected {
		t.Errorf("RenderTemplate() = %q, want %q", result, expected)
	}
}
