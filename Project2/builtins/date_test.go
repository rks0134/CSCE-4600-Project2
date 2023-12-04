package builtins_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/rks0134/CSCE4600/Project2/builtins"
)

func TestDate(t *testing.T) {
	var w bytes.Buffer
	err := builtins.Date(&w)
	if err != nil {
		t.Fatalf("Date() error = %v", err)
	}

	// Read the output and trim the newline character
	got := strings.TrimSpace(w.String())

	// Define the expected date format
	expectedFormat := "Mon Jan 2 15:04:05 MST 2006"

	// Parse the output to validate the format
	_, err = time.Parse(expectedFormat, got)
	if err != nil {
		t.Errorf("Date() output format error = %v", err)
	}
}
