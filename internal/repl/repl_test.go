package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestShellParsesInput(t *testing.T) {
	input := strings.NewReader("echo hello world\n")
	var output bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &output,
	}

	err := shell.Run()

	if err != nil {
		t.Fatalf("Run() returned an unexpected error: %v", err)
	}

	got := output.String()
	want := "aeroshell$ command: echo\n" +
		"args: [hello world]\n" +
		"aeroshell$ "

	if got != want {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", want, got)
	}
}

func TestShellIgnoresBlankLines(t *testing.T) {
	input := strings.NewReader("\n \nhello\n")
	var output bytes.Buffer

	shell := &Shell{
		In:  input,
		Out: &output,
		Err: &output,
	}

	err := shell.Run()

	if err != nil {
		t.Fatalf("Run() returned an unexpected error: %v", err)
	}

	got := output.String()
	want := "aeroshell$ " +
		"aeroshell$ " +
		"aeroshell$ command: hello\n" +
		"args: []\n" +
		"aeroshell$ "

	if got != want {
		t.Errorf("output mismatch\nwant: %q\ngot:  %q", want, got)
	}
}
