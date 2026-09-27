package tests

import (
	"slices"
	"strings"
	"testing"

	wc "github.com/defUserName-404/coding-challenges/challenges/002-wc/src"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want wc.Counts
	}{
		{"empty", "", wc.Counts{}},
		{"single word", "hello", wc.Counts{Bytes: 5, Words: 1, Chars: 5}},
		{"single line", "hello\n", wc.Counts{Bytes: 6, Lines: 1, Words: 1, Chars: 6}},
		{"two lines", "hello\nworld\n", wc.Counts{Bytes: 12, Lines: 2, Words: 2, Chars: 12}},
		{"multiple spaces", "a  b\tc\n", wc.Counts{Bytes: 7, Lines: 1, Words: 3, Chars: 7}},
		{"leading spaces", "   x\n", wc.Counts{Bytes: 5, Lines: 1, Words: 1, Chars: 5}},
		{"no trailing newline", "a b", wc.Counts{Bytes: 3, Words: 2, Chars: 3}},
		{
			// é and ö are 2 bytes each but 1 character each.
			name: "utf8 characters",
			in:   "héllo wörld\n",
			want: wc.Counts{Bytes: 14, Lines: 1, Words: 2, Chars: 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wc.Count(strings.NewReader(tt.in))
			if err != nil {
				t.Fatalf("Count() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Count(%q) = %+v, want %+v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCountAcrossBufferBoundary(t *testing.T) {
	// 4095 filler bytes put é's first byte at index 4095 and its
	// continuation byte at index 4096 — exactly across the 4096-byte
	// buffer. A naive per-chunk rune count would report 2 characters.
	in := strings.Repeat("a", 4095) + "é\n"

	got, err := wc.Count(strings.NewReader(in))
	if err != nil {
		t.Fatalf("Count() error = %v", err)
	}
	want := wc.Counts{Bytes: 4098, Lines: 1, Words: 1, Chars: 4097}
	if got != want {
		t.Errorf("Count() = %+v, want %+v", got, want)
	}
}

func TestCountLargeInput(t *testing.T) {
	// Multiple full chunks through the reader loop.
	in := strings.Repeat("word ", 2000)

	got, err := wc.Count(strings.NewReader(in))
	if err != nil {
		t.Fatalf("Count() error = %v", err)
	}
	want := wc.Counts{Bytes: 10000, Words: 2000, Chars: 10000}
	if got != want {
		t.Errorf("Count() = %+v, want %+v", got, want)
	}
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    wc.Options
		wantErr bool
	}{
		{
			name: "byte flag and file",
			args: []string{"-c", "foo.txt"},
			want: wc.Options{Bytes: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "line flag",
			args: []string{"-l", "foo.txt"},
			want: wc.Options{Lines: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "word flag",
			args: []string{"-w", "foo.txt"},
			want: wc.Options{Words: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "character flag",
			args: []string{"-m", "foo.txt"},
			want: wc.Options{Chars: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "combined flags",
			args: []string{"-cw", "foo.txt"},
			want: wc.Options{Bytes: true, Words: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "default when no flags given",
			args: []string{"foo.txt"},
			want: wc.Options{Lines: true, Words: true, Bytes: true, Paths: []string{"foo.txt"}},
		},
		{
			name: "no operand means stdin",
			args: []string{"-c"},
			want: wc.Options{Bytes: true},
		},
		{
			name: "several operands",
			args: []string{"-c", "a.txt", "b.txt"},
			want: wc.Options{Bytes: true, Paths: []string{"a.txt", "b.txt"}},
		},
		{
			name: "dash operand",
			args: []string{"-c", "-"},
			want: wc.Options{Bytes: true, Paths: []string{"-"}},
		},
		{
			name:    "unknown flag",
			args:    []string{"-z", "foo.txt"},
			wantErr: true,
		},
		{
			name:    "invalid flag value",
			args:    []string{"-c=maybe", "foo.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wc.ParseArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseArgs(%q) = %+v, want error", tt.args, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseArgs(%q) error = %v", tt.args, err)
			}
			if got.Bytes != tt.want.Bytes || got.Lines != tt.want.Lines ||
				got.Words != tt.want.Words || got.Chars != tt.want.Chars ||
				!slices.Equal(got.Paths, tt.want.Paths) {
				t.Errorf("ParseArgs(%q) = %+v, want %+v", tt.args, got, tt.want)
			}
		})
	}
}
