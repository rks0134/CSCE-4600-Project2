package builtins_test

import (
	"bytes"
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"testing"
)

func TestEcho(t *testing.T) {
	args := []string{"hello", "world"}
	want := "hello world\n"

	var w bytes.Buffer
	err := builtins.Echo(&w, args...)
	if err != nil {
		t.Fatalf("Echo() error = %v", err)
	}

	if got := w.String(); got != want {
		t.Errorf("Echo() = %v, want %v", got, want)
	}
}
