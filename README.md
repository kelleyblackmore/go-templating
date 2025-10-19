# go-templating

A demonstration of Go's powerful templating capabilities for generating configuration files using the `text/template` package.

## Overview

This project demonstrates how to use Go templates to dynamically generate configuration files based on data structures. It's particularly useful for managing configurations across different environments or deployment scenarios.

## Features

- **Dynamic Configuration Generation**: Generate configuration files from Go data structures
- **Template-based Approach**: Use `.tmpl` files with placeholders for flexible configuration
- **Type-safe Data**: Leverage Go structs to ensure data integrity
- **Reusable Templates**: Create templates that can be used with different data inputs

## Project Structure

```
.
├── cmd/
│   └── go-templating/         # Main application entry point
│       └── main.go
├── internal/
│   ├── config/                # Configuration data structures
│   │   └── config.go
│   └── template/              # Template processing logic
│       └── processor.go
├── templates/
│   └── config.tmpl            # Template file with configuration placeholders
├── .github/
│   └── workflows/             # GitHub Actions workflows
│       ├── ci.yml             # Build and test workflow
│       └── lint.yml           # Linting workflow
├── Makefile                   # Build automation
├── .golangci.yml             # Linter configuration
├── go.mod                     # Go module definition
└── generated_config.conf      # Output file (generated at runtime)
```

## How It Works

### 1. Template File (`templates/config.tmpl`)

The template file uses Go template syntax with placeholders:

```
# Configuration for {{.AppName}}
Version: {{.Version}}
Environment: {{.Environment}}

[Database]
Host: {{.Database.Host}}
Port: {{.Database.Port}}
...
```

### 2. Data Structure (`internal/config/config.go`)

A Go struct defines the configuration data:

```go
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
```

### 3. Template Processing (`internal/template/processor.go`)

The application:
1. Parses the template file using `template.ParseFiles()`
2. Creates an output file
3. Executes the template with configuration data using `tmpl.Execute()`
4. Writes the result to the output file

## Usage

### Using Make (Recommended)

```bash
# Display available commands
make help

# Build the application
make build

# Run tests
make test

# Run linters
make lint

# Format code
make fmt

# Build and run
make run

# Clean build artifacts
make clean
```

### Build the Application

```bash
make build
# or
go build -o go-templating ./cmd/go-templating
```

### Run the Application

```bash
make run
# or
./go-templating
```

This will generate a `generated_config.conf` file with the configuration based on the data defined in `main.go`.

### Example Output

```
# Configuration for MyApplication
Version: 1.0.0
Environment: production

[Database]
Host: localhost
Port: 5432
User: admin
Password: secure_password

[Features]
EnableFeatureA: true
EnableFeatureB: false
```

## Customization

To generate different configurations:

1. **Modify the template**: Edit `templates/config.tmpl` to change the structure
2. **Update data**: Modify the configuration in `internal/config/config.go`
3. **Change output**: Update the output filename in `cmd/go-templating/main.go`

## Key Concepts

- **`text/template` Package**: Used for generating textual output (for HTML, use `html/template` for security)
- **Template Syntax**: Use `{{.FieldName}}` for placeholders, with support for control structures like `{{if}}` and `{{range}}`
- **Data Injection**: Templates are executed with a data structure that fills in the placeholders
- **Type Safety**: Go's type system ensures data integrity throughout the process

## Advanced Usage

The template engine supports:
- Conditional logic with `{{if}}...{{end}}`
- Loops with `{{range}}...{{end}}`
- Custom functions
- Nested data structures
- Pipelines and variables

## Requirements

- Go 1.22 or higher
- Make (optional, for using Makefile targets)

## Development

### Project Structure

This project follows Go best practices with a production-grade folder structure:

- **`cmd/`**: Contains application entry points
- **`internal/`**: Private application code not meant to be imported by other projects
- **`templates/`**: Template files used by the application
- **`.github/workflows/`**: GitHub Actions CI/CD workflows

### Continuous Integration

The project includes GitHub Actions workflows for:

- **Linting**: Runs `go vet`, `gofmt`, and `golangci-lint` on every push and PR
- **CI**: Builds and tests the code across multiple Go versions

### Code Quality

Run linters before committing:

```bash
make lint
```

Format your code:

```bash
make fmt
```

## Security Note

This is a demonstration project. In production environments:
- Never hardcode passwords or sensitive data in source code
- Use environment variables, configuration files, or secure credential management systems
- Consider using `html/template` instead of `text/template` when generating HTML to prevent injection attacks

## License

This is a demonstration project for educational purposes.