package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestShellReadsandEchoesInput(t *testing.T) {
	input := strings.NewReader("hello\n")
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
	want := "aeroshell$ You entered: hello\n" + "aeroshell$ "

	if got != want {
		t.Errorf("output mismatch\nwant: %q\ngot: %q", want, got)
	}
}

func TestShellReadsandEchoesInputWithNewLine(t *testing.T) {
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
	want := "aeroshell$ " + "aeroshell$ " + "aeroshell$ You entered: hello\n" + "aeroshell$ "

	if got != want {
		t.Errorf("output mismatch\nwant: %q\ngot: %q", want, got)
	}
}
