package version

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestVersion(t *testing.T) {
	// Initiate Gomega
	g := NewWithT(t)

	tests := map[string]struct {
		version            string
		expectedOutput    string
	}{
		"no-build-version": {
			expectedOutput: "gotmpl version: source",
		},
		"params-not-yaml": {
			version:         "1.2.3-abcdef",
			expectedOutput: "gotmpl version: 1.2.3-abcdef",
		},
	}

	for name, test := range tests {
		if len(test.version) > 0 {
			Version = test.version
		}

		version := String()

		g.Expect(version).To(Equal(test.expectedOutput), "Test [%s]:", name)
	}
}
