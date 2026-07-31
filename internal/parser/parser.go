package parser

import (
	"errors"
	"strings"
	"unicode"
)

type Command struct {
	Name string
	Args []string
}

type parseState int

const (
	unquoted parseState = iota
	singleQuoted
	doubleQuoted
)

func Parse(line string) (Command, error) {
	tokens, err := tokenize(line)

	if err != nil {
		return Command{}, err
	}

	if len(tokens) == 0 {
		return Command{}, errors.New("empty command")
	}

	return Command{
		Name: tokens[0],
		Args: tokens[1:],
	}, nil

}

func tokenize(line string) ([]string, error) {
	tokens := []string{}
	var current strings.Builder

	state := unquoted
	inToken := false

	for _, char := range line {
		switch state {
		case unquoted:
			switch {
			case unicode.IsSpace(char):
				if inToken {
					tokens = append(tokens, current.String())
					current.Reset()
					inToken = false
				}
			case char == '\'':
				state = singleQuoted
				inToken = true
			case char == '"':
				state = doubleQuoted
				inToken = true
			default:
				current.WriteRune(char)
				inToken = true
			}
		case singleQuoted:
			if char == '\'' {
				state = unquoted
			} else {
				current.WriteRune(char)
			}
		case doubleQuoted:
			if char == '"' {
				state = unquoted
			} else {
				current.WriteRune(char)
			}
		}
	}

	if state != unquoted {
		return nil, errors.New("unterminated quote")
	}

	if inToken {
		tokens = append(tokens, current.String())
	}

	return tokens, nil
}
