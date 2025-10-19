# go-templating

A demonstration of Go's powerful templating capabilities through its `text/template` package for generating configuration files.

## Overview

Go offers robust templating features that can be effectively utilized for generating configuration files. This project demonstrates:

- Basic Go template structure for configuration
- Parsing and executing templates with data structures
- Using template logic (conditionals, loops)
- Practical examples with configuration data

## Project Structure

```
.
├── main.go           # Main application demonstrating template usage
├── template.go       # Template utility functions
├── template_test.go  # Tests for template functionality
└── config.tmpl       # Example configuration template
```

## Basic Go Template Structure for Configuration

A typical Go template for configuration involves defining placeholders and logic within a text file (e.g., `.tmpl`, `.gotmpl`). These files are then parsed and executed with a data structure to produce the final configuration output.

### Example Template (config.tmpl)

```
# Configuration File
AppName: {{.AppName}}
Version: {{.Version}}
Environment: {{.Environment}}

Server:
  Host: {{.Server.Host}}
  Port: {{.Server.Port}}
  Timeout: {{.Server.Timeout}}s

Database:
  Driver: {{.Database.Driver}}
  Host: {{.Database.Host}}
  Port: {{.Database.Port}}

{{- if .Features}}
Features:
{{- range .Features}}
  - {{.}}
{{- end}}
{{- end}}

Debug: {{.Debug}}
```

## Usage

### Running the Example

```bash
go run main.go
```

This will parse the `config.tmpl` template and execute it with sample configuration data.

### Running Tests

```bash
go test -v
```

### Using the Template Functions

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Define your configuration structure
    config := Config{
        AppName:     "MyApp",
        Version:     "1.0.0",
        Environment: "production",
        Server: ServerConfig{
            Host:    "localhost",
            Port:    8080,
            Timeout: 30,
        },
        // ... more fields
    }

    // Render template from string
    tmplString := "AppName: {{.AppName}}, Version: {{.Version}}"
    result, err := RenderTemplate(tmplString, config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }
    fmt.Println(result)
}
```

## Template Syntax Features

### Variable Substitution
```
{{.FieldName}}
```

### Nested Fields
```
{{.Server.Host}}
{{.Database.Port}}
```

### Conditionals
```
{{if .Debug}}
Debug mode is enabled
{{else}}
Debug mode is disabled
{{end}}
```

### Loops
```
{{range .Features}}
  - {{.}}
{{end}}
```

### Whitespace Control
```
{{- .Field}}   # Remove whitespace before
{{.Field -}}   # Remove whitespace after
```

## Configuration Structure

The project uses a structured approach to configuration:

```go
type Config struct {
    AppName     string
    Version     string
    Environment string
    Server      ServerConfig
    Database    DatabaseConfig
    Features    []string
    Debug       bool
}
```

## Benefits of Go Templates for Configuration

1. **Type Safety**: Configuration structures are defined in Go code
2. **Logic Support**: Conditionals and loops for dynamic configurations
3. **Reusability**: Templates can be reused across different environments
4. **Maintainability**: Separation of configuration logic from data
5. **Validation**: Go compiler ensures data structure correctness

## License

MIT