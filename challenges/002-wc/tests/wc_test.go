package tests

import (
	"bytes"
	"os"
	"strings"
	"testing"

	wc "github.com/defUserName-404/coding-challenges/challenges/002-wc/src"
)

const testFile = "testdata/test.txt"

// Expected counts for testdata/test.txt, cross-checked against the
// system wc: -c → 342190, -l → 7145, -w → 58164, -m → 339292.
// Output layout matches macOS wc: "%8d" per column, then " name".
func TestRunSingleFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"bytes", []string{"-c", testFile}, "  342190 testdata/test.txt\n"},
		{"lines", []string{"-l", testFile}, "    7145 testdata/test.txt\n"},
		{"words", []string{"-w", testFile}, "   58164 testdata/test.txt\n"},
		{"chars", []string{"-m", testFile}, "  339292 testdata/test.txt\n"},
		{
			"default mode: no flag means lines, words, bytes",
			[]string{testFile},
			"    7145   58164  342190 testdata/test.txt\n",
		},
		{
			// Our deliberate difference from macOS wc, which would
			// print only the char count when -c and -m are combined.
			"bytes and chars together",
			[]string{"-cm", testFile},
			"  342190  339292 testdata/test.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := wc.Run(tt.args, strings.NewReader(""), &out); err != nil {
				t.Fatalf("Run(%q) error = %v", tt.args, err)
			}
			if out.String() != tt.want {
				t.Errorf("Run(%q) output = %q, want %q", tt.args, out.String(), tt.want)
			}
		})
	}
}

func TestRunMultipleFilesWithTotal(t *testing.T) {
	var out bytes.Buffer

	args := []string{"-c", testFile, testFile}
	if err := wc.Run(args, strings.NewReader(""), &out); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	want := "  342190 testdata/test.txt\n" +
		"  342190 testdata/test.txt\n" +
		"  684380 total\n"
	if out.String() != want {
		t.Errorf("Run() output = %q, want %q", out.String(), want)
	}
}

func TestRunStdin(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		stdin string
		want  string
	}{
		{
			"no operand reads stdin, default columns",
			nil, "hi\n",
			"       1       1       3\n",
		},
		{
			"no operand with -c",
			[]string{"-c"}, "hi\n",
			"       3\n",
		},
		{
			"dash operand reads stdin",
			[]string{"-c", "-"}, "hi\n",
			"       3 -\n",
		},
		{
			"utf8 character count from stdin",
			[]string{"-m"}, "héllo\n",
			"       6\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := wc.Run(tt.args, strings.NewReader(tt.stdin), &out); err != nil {
				t.Fatalf("Run(%q) error = %v", tt.args, err)
			}
			if out.String() != tt.want {
				t.Errorf("Run(%q) output = %q, want %q", tt.args, out.String(), tt.want)
			}
		})
	}
}

func TestRunFileNotFound(t *testing.T) {
	var out bytes.Buffer

	err := wc.Run([]string{"-c", "testdata/does-not-exist.txt"}, strings.NewReader(""), &out)
	if !os.IsNotExist(err) {
		t.Errorf("Run() error = %v, want a not-exist error", err)
	}
	if out.Len() != 0 {
		t.Errorf("Run() wrote %q before failing, want no output", out.String())
	}
}

func TestRunUnknownFlag(t *testing.T) {
	var out bytes.Buffer

	err := wc.Run([]string{"-z", testFile}, strings.NewReader(""), &out)
	if err == nil {
		t.Fatal("Run() with unknown flag: error = nil, want error")
	}
}
