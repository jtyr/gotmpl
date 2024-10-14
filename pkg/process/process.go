package process

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Map is a map of any type.
type Map map[string]any

// readInput returns content of the file if the input is a file path otherwise it returns the input as []byte.
func readInput(input string) ([]byte, error) {
	var content []byte

	if _, e := os.Stat(input); e == nil {
		var err error

		content, err = os.ReadFile(input)
		if err != nil {
			return []byte(""), fmt.Errorf("failed to read file: %w", err)
		}
	} else {
		content = []byte(input)
	}

	return content, nil
}

// ProcessTmpl prints out templated output based on the input parameters and template string.
func ProcessTmpl(params, tmpl string) error {
	paramsDoc, err := readInput(params)
	if err != nil {
		return fmt.Errorf("failed to read params file: %w", err)
	}

	tmplDoc, err := readInput(tmpl)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	var data Map

	err = yaml.Unmarshal(paramsDoc, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal YAML to map: %w", err)
	}

	t, err := template.New("test").Parse(string(tmplDoc))
	if err != nil {
		return fmt.Errorf("failed to parse template string: %w", err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		return fmt.Errorf("failed to template data: %w", err)
	}

	return nil
}
