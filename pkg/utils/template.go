package utils

import (
	"bytes"
	"text/template"
)

// ProcessTemplate parses and executes a template with the given data.
// It's useful for creating prompts from predefined structures.
func ProcessTemplate(templateStr string, data interface{}) (string, error) {
	tmpl, err := template.New("prompt").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}