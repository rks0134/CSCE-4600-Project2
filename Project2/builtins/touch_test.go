package builtins_test

import (
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"os"
	"testing"
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
