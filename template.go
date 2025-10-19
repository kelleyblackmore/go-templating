package main

import (
	"bytes"
	"io"
	"text/template"
)

// RenderTemplate renders a template with the given data
func RenderTemplate(tmplContent string, data interface{}) (string, error) {
	tmpl, err := template.New("config").Parse(tmplContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderTemplateToWriter renders a template with the given data to a writer
func RenderTemplateToWriter(w io.Writer, tmplContent string, data interface{}) error {
	tmpl, err := template.New("config").Parse(tmplContent)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}
