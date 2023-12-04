package builtins_test

import (
	"os"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestMkdir(t *testing.T) {
	dir := t.TempDir() + "/newdir"
	err := builtins.Mkdir(dir)
	if err != nil {
		t.Fatalf("Mkdir() error = %v", err)
	}

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		t.Errorf("Mkdir() did not create directory")
	}
}
