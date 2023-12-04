package main

import (
	"bytes"
	"io"
	"os"
	"os/user"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/require"
)

func TestPrintPrompt(t *testing.T) {
	// Prepare a buffer to capture the output.
	var buf bytes.Buffer

	// Call the printPrompt function, passing the buffer as the io.Writer.
	err := printPrompt(&buf)
	if err != nil {
		t.Fatalf("printPrompt returned an error: %v", err)
	}

	// Get the current user and working directory, as the function would.
	u, err := user.Current()
	if err != nil {
		t.Fatalf("failed to get current user: %v", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	// Construct the expected prompt string.
	expected := wd + " [" + u.Username + "] $ "

	// Check if the output matches the expected prompt.
	if got := buf.String(); got != expected {
		t.Errorf("printPrompt = %q, want %q", got, expected)
	}
}
func Test_handleInput(t *testing.T) {
	t.Parallel()
	exit := make(chan struct{}, 2)

	tests := []struct {
		name     string
		input    string
		exitChan chan struct{}
		wantErr  bool
	}{
		{
			name:     "Valid built-in command",
			input:    "pwd\n",
			exitChan: exit,
			wantErr:  false,
		},
		{
			name:     "Invalid built-in command",
			input:    "invalid\n",
			exitChan: exit,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			err := handleInput(&w, tt.input, tt.exitChan)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_executeCommand(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		cmdName string
		cmdArgs []string
		wantErr bool
	}{
		{
			name:    "Valid command",
			cmdName: "ls",
			cmdArgs: []string{"-l"},
			wantErr: false,
		},
		{
			name:    "Invalid command",
			cmdName: "invalidcommand",
			cmdArgs: []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := executeCommand(tt.cmdName, tt.cmdArgs...)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}
