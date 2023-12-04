package builtins_test

import (
	"os"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestRmdir(t *testing.T) {
	dir := t.TempDir() + "/toremove"
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
}
