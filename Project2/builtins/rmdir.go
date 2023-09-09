// builtins/rmdir.go

package builtins

import (
	"fmt"
	"os"
)

// Rmdir handles the "rmdir" built-in command.
func Rmdir(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: rmdir <directory>")
	}
	dir := args[0]
	if err := os.Remove(dir); err != nil {
		return err
	}
	return nil
}
