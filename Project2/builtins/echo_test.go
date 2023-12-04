package builtins_test

import (
	"bytes"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
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
