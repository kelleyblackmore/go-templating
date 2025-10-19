package main

import (
	"fmt"
	"strings"
	"testing"
	"text/template"
)

// Example_basicTemplate demonstrates basic variable substitution
func Example_basicTemplate() {
	tmpl := "Hello {{.Name}}!"
	data := map[string]string{"Name": "World"}
	
	result, _ := RenderTemplate(tmpl, data)
	fmt.Println(result)
	// Output: Hello World!
}

// Example_nestedStructs demonstrates accessing nested struct fields
func Example_nestedStructs() {
	config := Config{
		AppName: "MyApp",
		Server: ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
	}
	
	tmpl := "{{.AppName}} running on {{.Server.Host}}:{{.Server.Port}}"
	result, _ := RenderTemplate(tmpl, config)
	fmt.Println(result)
	// Output: MyApp running on localhost:8080
}

// Example_rangeLoop demonstrates iterating over a slice
func Example_rangeLoop() {
	data := map[string][]string{
		"Items": {"apple", "banana", "cherry"},
	}
	
	tmpl := "{{range .Items}}{{.}} {{end}}"
	result, _ := RenderTemplate(tmpl, data)
	fmt.Println(result)
	// Output: apple banana cherry 
}

// TestCustomTemplateFunctions demonstrates using custom functions in templates
func TestCustomTemplateFunctions(t *testing.T) {
	// Define custom functions
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"title": strings.Title,
	}
	
	tmplText := "{{upper .Name}} - {{lower .Status}} - {{title .Message}}"
	
	tmpl, err := template.New("custom").Funcs(funcMap).Parse(tmplText)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}
	
	data := map[string]string{
		"Name":    "john",
		"Status":  "ACTIVE",
		"Message": "hello world",
	}
	
	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}
	
	expected := "JOHN - active - Hello World"
	if buf.String() != expected {
		t.Errorf("Expected %q, got %q", expected, buf.String())
	}
}

// TestConditionalRendering demonstrates if-else conditions in templates
func TestConditionalRendering(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     map[string]interface{}
		expected string
	}{
		{
			name:     "If condition true",
			template: "{{if .Active}}Service is active{{end}}",
			data:     map[string]interface{}{"Active": true},
			expected: "Service is active",
		},
		{
			name:     "If condition false",
			template: "{{if .Active}}Service is active{{end}}",
			data:     map[string]interface{}{"Active": false},
			expected: "",
		},
		{
			name:     "If-else condition",
			template: "{{if .Debug}}Debug mode{{else}}Production mode{{end}}",
			data:     map[string]interface{}{"Debug": false},
			expected: "Production mode",
		},
		{
			name:     "Check non-empty slice",
			template: "{{if .Items}}Has items{{else}}No items{{end}}",
			data:     map[string]interface{}{"Items": []string{"a", "b"}},
			expected: "Has items",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RenderTemplate(tt.template, tt.data)
			if err != nil {
				t.Fatalf("RenderTemplate() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestWhitespaceControl demonstrates whitespace trimming in templates
func TestWhitespaceControl(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     map[string][]string
		expected string
	}{
		{
			name: "Without whitespace control",
			template: `{{range .Items}}
{{.}}
{{end}}`,
			data:     map[string][]string{"Items": {"a", "b", "c"}},
			expected: "\na\n\nb\n\nc\n",
		},
		{
			name: "With whitespace control",
			template: `{{range .Items -}}
{{.}}
{{end -}}`,
			data:     map[string][]string{"Items": {"a", "b", "c"}},
			expected: "a\nb\nc\n",
		},
		{
			name:     "Trim both sides",
			template: "{{- range .Items}}{{.}}{{- end}}",
			data:     map[string][]string{"Items": {"a", "b", "c"}},
			expected: "abc",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RenderTemplate(tt.template, tt.data)
			if err != nil {
				t.Fatalf("RenderTemplate() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
