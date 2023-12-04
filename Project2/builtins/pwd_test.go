package builtins_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestPwd(t *testing.T) {
	var w bytes.Buffer
	err := builtins.Pwd(&w)
	if err != nil {
		t.Fatalf("Pwd() error = %v", err)
	}

	expected, err := os.Getwd()
	if err != nil {
		t.Skip("Skipping TestPwd as working directory is not accessible: ", err)
	}

	got := strings.TrimSpace(w.String())
	if got != expected {
		t.Errorf("Pwd() = %v, want %v", got, expected)
	}
}
