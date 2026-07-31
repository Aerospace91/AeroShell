package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/Aerospace91/AeroShell/internal/parser"
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

		command, err := parser.Parse(line)
		if err != nil {
			fmt.Fprintf(s.Err, "aeroshell: %v\n", err)
			continue
		}

		fmt.Fprintf(s.Out, "command: %s\n", command.Name)
		fmt.Fprintf(s.Out, "args: %v\n", command.Args)
	}

	return scanner.Err()
}
