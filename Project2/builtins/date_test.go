package builtins_test

import (
	"bytes"
	"github.com/jh125486/CSCE4600/Project2/builtins"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	var w bytes.Buffer
	err := builtins.Date(&w)
	if err != nil {
		t.Errorf("Date() error = %v, wantErr nil", err)
	}

	// Verify the output format
	expectedFormat := "Mon Jan 2 15:04:05 MST 2006"
	_, err = time.Parse(expectedFormat, w.String())
	if err != nil {
		t.Errorf("Date() output format error = %v", err)
	}
}