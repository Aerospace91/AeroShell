package repl

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestShellPwd(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned unexpected error: %v", err)
	}

	input := strings.NewReader("pwd\n")
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err = shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := fmt.Sprintf("aeroshell$ %s\naeroshell$ ", currentDir)

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}

func TestShellCdChangesDirectory(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned unexpected error: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("failed to restore directory: %v", err)
		}
	})

	tempDir := t.TempDir()

	input := strings.NewReader(fmt.Sprintf("cd %s\npwd\n", tempDir))
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err = shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := fmt.Sprintf(
		"aeroshell$ aeroshell$ %s\naeroshell$ ",
		tempDir,
	)

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}

func TestShellCdWithoutArgumentsUsesHomeDirectory(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() returned unexpected error: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Fatalf("failed to restore directory: %v", err)
		}
	})

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir() returned unexpected error: %v", err)
	}

	input := strings.NewReader("cd\npwd\n")
	var output bytes.Buffer
	var errorOutput bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &errorOutput,
	}

	err = shell.Run()
	if err != nil {
		t.Fatalf("Run() returned unexpected error: %v", err)
	}

	wantOutput := fmt.Sprintf(
		"aeroshell$ aeroshell$ %s\naeroshell$ ",
		homeDir,
	)

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}

func TestShellCdInvalidDirectory(t *testing.T) {
	input := strings.NewReader("cd /definitely/not/a/real/path\n")
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

	if !strings.Contains(errorOutput.String(), "aeroshell: cd:") {
		t.Errorf("expected cd error, got: %q", errorOutput.String())
	}
}

func TestShellCdTooManyArguments(t *testing.T) {
	input := strings.NewReader("cd one two\n")
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

	wantError := "aeroshell: cd: too many arguments\n"

	if errorOutput.String() != wantError {
		t.Errorf("stderr mismatch\nwant: %q\ngot:  %q", wantError, errorOutput.String())
	}
}

func TestShellExitStopsLoop(t *testing.T) {
	input := strings.NewReader("exit\npwd\n")
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

	wantOutput := "aeroshell$ "

	if output.String() != wantOutput {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", wantOutput, output.String())
	}

	if errorOutput.Len() != 0 {
		t.Errorf("expected no stderr output, got: %q", errorOutput.String())
	}
}
