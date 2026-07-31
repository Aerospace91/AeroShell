package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
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

		handled, shouldExit := s.handleBuiltin(command)

		if shouldExit {
			return nil
		}

		if handled {
			continue
		}

		s.execute(command)

	}

	return scanner.Err()
}

func (s *Shell) handleBuiltin(command parser.Command) (handled bool, shouldExit bool) {
	switch command.Name {
	case "pwd":
		s.pwd()
		return true, false

	case "cd":
		s.cd(command.Args)
		return true, false
	case "exit":
		if len(command.Args) > 0 {
			fmt.Fprintln(s.Err, "aeroshell: exit: too many arguments")
			return true, false
		}
		return true, true
	default:
		return false, false
	}
}

func (s *Shell) pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(s.Err, "aeroshell: pwd: %v\n", err)
		return
	}
	fmt.Fprintln(s.Out, dir)
}

func (s *Shell) cd(args []string) {
	var dir string

	switch len(args) {
	case 0:
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(s.Err, "aeroshell: cd: %v\n", err)
			return
		}
		dir = homeDir
	case 1:
		dir = args[0]
	default:
		fmt.Fprintln(s.Err, "aeroshell: cd: too many arguments")
		return
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(s.Err, "aeroshell: cd: %v\n", err)
	}
}

func (s *Shell) execute(command parser.Command) {
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdin = s.In
	cmd.Stdout = s.Out
	cmd.Stderr = s.Err

	err := cmd.Run()
	if err == nil {
		return
	}

	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return
	}

	fmt.Fprintf(s.Err, "aeroshell: %s: command not found\n", command.Name)
}
