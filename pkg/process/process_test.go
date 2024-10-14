package process

import (
	"io"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out), err
}

func TestProcess(t *testing.T) {
	// Initiate Gomega
	g := NewWithT(t)

	tests := map[string]struct {
		params            string
		tmpl              string
		expectedOutput    string
		expectError       bool
		errorRegexp       string
	}{
		"all-good": {
			params:         "name: John Doe",
			tmpl:           "Hello {{ .name }}.",
			expectedOutput: "Hello John Doe.",
		},
		"params-not-yaml": {
			params:         ":",
			expectError:    true,
			errorRegexp:    "failed to unmarshal YAML to map: yaml: did not find expected key",
		},
		"invalid-template": {
			tmpl:           "Hello {{ .name }.",
			expectError:    true,
			errorRegexp:    "failed to parse template string: template: test:1: unexpected \"}\" in operand",
		},
		"failed-templating": {
			params:         "names: [ John Doe ]",
			tmpl:           "Hello {{ index .names 1 }}.",
			expectError:    true,
			errorRegexp:    "failed to template data: template: test:1:9: executing \"test\" at <index .names 1>: error calling index: reflect: slice index out of range",
		},
		"invalid-params-file": {
			params:         "/",
			expectError:    true,
			errorRegexp:    "failed to read file: read /: is a directory",
		},
		"invalid-template-file": {
			tmpl:           "/",
			expectError:    true,
			errorRegexp:    "failed to read file: read /: is a directory",
		},
	}

	for name, test := range tests {
		output, err := captureOutput(func() error {
			err := ProcessTmpl(test.params, test.tmpl)

			return err
		})

		if test.expectError {
			g.Expect(err).To(MatchError(MatchRegexp(test.errorRegexp)), "Test [%s]:", name)
		} else {
			g.Expect(err).ToNot(HaveOccurred(), "Test [%s]:", name)
			g.Expect(output).To(Equal(test.expectedOutput), "Test [%s]:", name)
		}
	}
}
