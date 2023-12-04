package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestRmdir(t *testing.T) {
	// Test case for removing an existing directory
	t.Run("Remove existing directory", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "toremove")
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			t.Fatalf("Setup error: %v", err)
		}

		err = builtins.Rmdir(dir)
		if err != nil {
			t.Fatalf("Rmdir() error = %v", err)
		}

		_, err = os.Stat(dir)
		if !os.IsNotExist(err) {
			t.Errorf("Rmdir() did not remove directory")
		}
	})

	// Test case for no arguments
	t.Run("No arguments", func(t *testing.T) {
		err := builtins.Rmdir()
		if err == nil {
			t.Errorf("Rmdir() expected error, got nil")
		}
	})

	// Test case for non-existent directory
	t.Run("Non-existent directory", func(t *testing.T) {
		nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")
		err := builtins.Rmdir(nonExistentDir)
		if !os.IsNotExist(err) {
			t.Errorf("Rmdir() expected os.IsNotExist error, got %v", err)
		}
	})

	// Additional test cases can be added here as needed
}
