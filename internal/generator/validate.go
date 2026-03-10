package generator

// GenerateValidate returns the source for internal/validate/validate.go in the
// generated CLI project. The validate package provides a shared Enum helper
// used by all verb commands that have enum-constrained flags, replacing the
// inline multi-line validation blocks with a single call site.
func GenerateValidate() (string, error) {
	src := `package validate

import (
	"fmt"
	"strings"
)

// Enum checks that val is one of the allowed values.
// Returns nil if val is empty (flag not set). Call only after cmd.Flags().Changed().
func Enum(flagName, val string, allowed []string) error {
	for _, a := range allowed {
		if a == val {
			return nil
		}
	}
	return fmt.Errorf("invalid value %q for --%s: must be one of: %s", val, flagName, strings.Join(allowed, ", "))
}
`
	return validateGoSource(src)
}
