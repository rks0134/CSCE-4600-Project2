package builtins_test

import (
	"os"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestTouch(t *testing.T) {
	file := t.TempDir() + "/newfile"
	err := builtins.Touch(file)
	if err != nil {
		t.Fatalf("Touch() error = %v", err)
	}

	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		t.Errorf("Touch() did not create file")
	}
}
