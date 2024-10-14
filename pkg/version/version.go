package version

import (
        "fmt"
)

// Version specifies the current version.
var Version = "source"

// String returns the version information.
func String() string {
    return fmt.Sprintf("gotmpl version: %s", Version)
}
