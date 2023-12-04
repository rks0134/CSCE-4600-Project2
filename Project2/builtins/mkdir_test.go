package builtins_test

import (
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"os"
	"testing"
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
