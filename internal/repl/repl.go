package repl

import "io"

type Shell struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}
