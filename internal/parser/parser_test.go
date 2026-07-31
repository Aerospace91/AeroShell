package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Command
		wantErr bool
	}{
		{
			name:  "command without arguments",
			input: "pwd",
			want: Command{
				Name: "pwd",
				Args: []string{},
			},
		},
		{
			name:  "command with arguments",
			input: "echo hello world",
			want: Command{
				Name: "echo",
				Args: []string{"hello", "world"},
			},
		},
		{
			name:  "extra whitespace",
			input: "  ls   -la  ",
			want: Command{
				Name: "ls",
				Args: []string{"-la"},
			},
		},
		{
			name:  "double quoted argument",
			input: `echo "hello world"`,
			want: Command{
				Name: "echo",
				Args: []string{"hello world"},
			},
		},
		{
			name:  "single quoted argument",
			input: `echo 'hello world'`,
			want: Command{
				Name: "echo",
				Args: []string{"hello world"},
			},
		},
		{
			name:  "quoted and unquoted arguments",
			input: `echo "two words" three`,
			want: Command{
				Name: "echo",
				Args: []string{"two words", "three"},
			},
		},
		{
			name:  "empty double quoted argument",
			input: `echo ""`,
			want: Command{
				Name: "echo",
				Args: []string{""},
			},
		},
		{
			name:  "empty single quoted argument",
			input: `echo ''`,
			want: Command{
				Name: "echo",
				Args: []string{""},
			},
		},
		{
			name:  "adjacent quoted and unquoted text",
			input: `echo hello" world"`,
			want: Command{
				Name: "echo",
				Args: []string{"hello world"},
			},
		},
		{
			name:  "double quotes inside single quotes",
			input: `echo 'it has "double quotes"'`,
			want: Command{
				Name: "echo",
				Args: []string{`it has "double quotes"`},
			},
		},
		{
			name:  "single quote inside double quotes",
			input: `echo "it's one argument"`,
			want: Command{
				Name: "echo",
				Args: []string{"it's one argument"},
			},
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "whitespace only",
			input:   "   \t  ",
			wantErr: true,
		},
		{
			name:    "unterminated double quote",
			input:   `echo "unfinished`,
			wantErr: true,
		},
		{
			name:    "unterminated single quote",
			input:   `echo 'unfinished`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Parse(test.input)

			if test.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("Parse() returned unexpected error: %v", err)
			}

			if got.Name != test.want.Name {
				t.Errorf("Name mismatch: got %q, want %q", got.Name, test.want.Name)
			}

			if !reflect.DeepEqual(got.Args, test.want.Args) {
				t.Errorf("Args mismatch: got %#v, want %#v", got.Args, test.want.Args)
			}
		})
	}
}
