package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestTouch(t *testing.T) {
	// Test case for creating a new file
	t.Run("Create new file", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "newfile")
		err := builtins.Touch(file)
		if err != nil {
			t.Fatalf("Touch() error = %v", err)
		}

		_, err = os.Stat(file)
		if os.IsNotExist(err) {
			t.Errorf("Touch() did not create file")
		}
	})

	// Test case for no arguments
	t.Run("No arguments", func(t *testing.T) {
		err := builtins.Touch()
		if err == nil {
			t.Errorf("Touch() expected error, got nil")
		}
	})

	// Test case for creating a file in non-existent directory
	t.Run("Non-existent directory", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "nonexistent", "file")
		err := builtins.Touch(file)
		if !os.IsNotExist(err) && err != nil {
			t.Errorf("Touch() expected os.IsNotExist error, got %v", err)
		}
	})

	// Test case for invalid file name
	t.Run("Invalid file name", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "/invalid/name/\000")
		err := builtins.Touch(file)
		if err == nil {
			t.Errorf("Touch() expected error, got nil")
		}
	})
}
