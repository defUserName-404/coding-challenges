package tests

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/001-cat/src"
)

const (
	fixture1 = "testdata/quotes1.txt"
	fixture2 = "testdata/quotes2.txt"
)

func readFixture(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read fixture %s: %v", path, err)
	}
	return string(data)
}

func TestRun(t *testing.T) {
	quotes1 := readFixture(t, fixture1)
	quotes2 := readFixture(t, fixture2)

	tests := []struct {
		name  string
		args  []string
		stdin string
		want  string
	}{
		{
			name: "step 1: prints a file",
			args: []string{fixture1},
			want: quotes1,
		},
		{
			name:  "step 2: reads stdin with dash",
			args:  []string{"-"},
			stdin: "first line\n",
			want:  "first line\n",
		},
		{
			name:  "step 2: reads stdin when no operands",
			args:  nil,
			stdin: "piped input\n",
			want:  "piped input\n",
		},
		{
			name: "step 3: concatenates files in order",
			args: []string{fixture1, fixture2},
			want: quotes1 + quotes2,
		},
		{
			name:  "step 4: numbers all lines",
			args:  []string{"-n", "-"},
			stdin: "alpha\n\nbeta\n",
			want:  "     1\talpha\n     2\t\n     3\tbeta\n",
		},
		{
			name:  "step 5: numbers non-blank lines only",
			args:  []string{"-b", "-"},
			stdin: "alpha\n\nbeta\n",
			want:  "     1\talpha\n\n     2\tbeta\n",
		},
		{
			name:  "step 5: -b wins over -n",
			args:  []string{"-n", "-b", "-"},
			stdin: "alpha\n\nbeta\n",
			want:  "     1\talpha\n\n     2\tbeta\n",
		},
		{
			name:  "keeps a file without trailing newline intact",
			args:  []string{"-n", "-"},
			stdin: "no newline",
			want:  "     1\tno newline",
		},
		{
			name: "numbering continues across files",
			args: []string{"-n", fixture1, fixture2},
			want: "     1\t\"Life isn't about getting and having, it's about giving and being.\"\n" +
				"     2\t\"Whatever the mind of man can conceive and believe, it can achieve.\"\n" +
				"     3\t\"Strive not to be a success, but rather to be of value.\"\n" +
				"     4\t\"We must balance conspicuous consumption with conscious capitalism.\"\n" +
				"     5\t\"Life is what happens to you while you're busy making other plans.\"\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out strings.Builder
			if err := cat.Run(tt.args, strings.NewReader(tt.stdin), &out); err != nil {
				t.Fatalf("Run() error = %v", err)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("Run() output mismatch\n got: %q\nwant: %q", got, tt.want)
			}
		})
	}
}

func TestRunMissingFile(t *testing.T) {
	err := cat.Run([]string{fixture1, "testdata/does-not-exist.txt"}, strings.NewReader(""), io.Discard)
	if err == nil {
		t.Fatal("Run() error = nil, want an error for a missing file")
	}
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Run() error = %v, want fs.ErrNotExist", err)
	}
}

func TestRunBadFlag(t *testing.T) {
	if err := cat.Run([]string{"-z"}, strings.NewReader(""), io.Discard); err == nil {
		t.Fatal("Run() error = nil, want an error for an unknown flag")
	}
}
