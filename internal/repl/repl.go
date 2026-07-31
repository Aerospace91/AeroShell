package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Shell struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (s *Shell) Run() error {

	scanner := bufio.NewScanner(s.In)

	for {
		fmt.Fprint(s.Out, "aeroshell$ ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		fmt.Fprintf(s.Out, "You entered: %s\n", line)
	}

	return scanner.Err()
}
