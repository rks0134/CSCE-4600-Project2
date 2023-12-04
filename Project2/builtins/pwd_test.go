package builtins_test

import (
	"bytes"
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
	want, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() error = %v", err)
	}

	var w bytes.Buffer
	err = builtins.Pwd(&w)
	if err != nil {
		t.Fatalf("Pwd() error = %v", err)
	}

	if got := w.String(); got != want+"\n" {
		t.Errorf("Pwd() = %v, want %v", got, want)
	}
}
