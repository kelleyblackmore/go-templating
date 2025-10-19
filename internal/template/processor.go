package template

import (
	"io"
	"text/template"
)

// Processor handles template processing
type Processor struct {
	templatePath string
}

// NewProcessor creates a new Processor
func NewProcessor(templatePath string) *Processor {
	return &Processor{
		templatePath: templatePath,
	}
}

// Process parses the template file and executes it with the given data
func (p *Processor) Process(writer io.Writer, data interface{}) error {
	tmpl, err := template.ParseFiles(p.templatePath)
	if err != nil {
		return err
	}

	return tmpl.Execute(writer, data)
}
