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
			wantErr: false,
		},
		{
			name:  "command with arguments",
			input: "echo hello world",
			want: Command{
				Name: "echo",
				Args: []string{"hello", "world"},
			},
			wantErr: false,
		},
		{
			name:  "extra whitespace",
			input: "  ls  -la  ",
			want: Command{
				Name: "ls",
				Args: []string{"-la"},
			},
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "whitespace only",
			input:   "  \t  ",
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
				t.Errorf("Args mismatch: got%#v, want %#v", got.Args, test.want.Args)
			}
		})
	}
}
