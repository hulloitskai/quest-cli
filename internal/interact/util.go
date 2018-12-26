package interact

import (
	"fmt"
	"os"
)

// Errf is a shortcut for fmt.Fprintf(os.Stderr, ...).
func Errf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

// Errln is a shortcut for fmt.Fprintln(os.Stderr, ...).
func Errln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}
