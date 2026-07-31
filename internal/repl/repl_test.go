package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestShellRunsExternalCommand(t *testing.T) {
	input := strings.NewReader("echo hello\n")
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err := shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := "aeroshell$ hello\naeroshell$ "

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}

func TestShellReportsMissingCommand(t *testing.T) {
	input := strings.NewReader("not-a-real-command\n")
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err := shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := "aeroshell$ aeroshell$ "

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	wantError := "aeroshell: not-a-real-command: command not found\n"

	if errorOutput.String() != wantError {
		t.Errorf("stderr mismatch\nwant: %q\ngot:  %q", wantError, errorOutput.String())
	}
}

func TestShellContinuesAfterExternalCommandFailure(t *testing.T) {
	input := strings.NewReader("false\necho still-running\n")
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err := shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := "aeroshell$ aeroshell$ still-running\naeroshell$ "

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}
