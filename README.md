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
├── config.tmpl           # Template file with configuration placeholders
├── main.go              # Go code to process the template
├── go.mod               # Go module definition
└── generated_config.conf # Output file (generated at runtime)
```

## How It Works

### 1. Template File (`config.tmpl`)

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

### 2. Data Structure

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

### 3. Template Processing

The application:
1. Parses the template file using `template.ParseFiles()`
2. Creates an output file
3. Executes the template with configuration data using `tmpl.Execute()`
4. Writes the result to the output file

## Usage

### Build the Application

```bash
go build
```

### Run the Application

```bash
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

1. **Modify the template**: Edit `config.tmpl` to change the structure
2. **Update data**: Modify the `cfg` variable in `main.go` to use different values
3. **Change output**: Update the output filename in `main.go`

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

- Go 1.16 or higher

## Security Note

This is a demonstration project. In production environments:
- Never hardcode passwords or sensitive data in source code
- Use environment variables, configuration files, or secure credential management systems
- Consider using `html/template` instead of `text/template` when generating HTML to prevent injection attacks

## License

This is a demonstration project for educational purposes.