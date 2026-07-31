package parser

import (
	"errors"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func Parse(line string) (Command, error) {
	fields := strings.Fields(line)

	if len(fields) == 0 {
		return Command{
			Name: "",
			Args: nil,
		}, errors.New("Empty Command or Inputs")
	}

	return Command{
		Name: fields[0],
		Args: fields[1:],
	}, nil

}
