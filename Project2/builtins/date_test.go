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

	got := strings.TrimSpace(w.String())

	// Validate the format
	expectedFormat := "Mon Jan 2 15:04:05 MST 2006"
	parsedTime, err := time.Parse(expectedFormat, got)
	if err != nil {
		t.Errorf("Date() output format error = %v", err)
	}

	// Test individual components
	currentTime := time.Now()
	if parsedTime.Year() != currentTime.Year() ||
		parsedTime.Month() != currentTime.Month() ||
		parsedTime.Day() != currentTime.Day() {
		t.Errorf("Date() components do not match current time")
	}
}
