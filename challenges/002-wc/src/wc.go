// Package wc implements the wc coding challenge:
// https://codingchallenges.fyi/challenges/challenge-wc
package wc

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Options holds the parsed command-line flags and file operands.
// Fields are exported so the tests package can inspect them.
type Options struct {
	Bytes bool
	Lines bool
	Words bool
	Chars bool
	Paths []string
}

// Counts holds the numbers measured from one input.
type Counts struct {
	Bytes int
	Lines int
	Words int
	Chars int
}

// add folds o into c so several inputs can produce a total row.
func (c *Counts) add(o Counts) {
	c.Bytes += o.Bytes
	c.Lines += o.Lines
	c.Words += o.Words
	c.Chars += o.Chars
}

// Run parses flags from args and writes one row of counts per file
// operand to w. With no operands it counts stdin, like wc does.
func Run(args []string, stdin io.Reader, w io.Writer) error {
	opts, err := ParseArgs(args)
	if err != nil {
		return err
	}

	paths := opts.Paths
	if len(paths) == 0 {
		paths = []string{"-"} // piped stdin: count it, print no file name
	}

	var total Counts
	for _, path := range paths {
		c, err := countInput(path, stdin)
		if err != nil {
			return err
		}
		total.add(c)

		name := path
		if len(opts.Paths) == 0 {
			name = ""
		}
		if err := writeResult(w, name, c, opts); err != nil {
			return err
		}
	}

	if len(paths) > 1 {
		return writeResult(w, "total", total, opts)
	}
	return nil
}

// ParseArgs splits args into flags and file operands. Go's flag
// package stops at the first non-flag argument, so "-" and file names
// end up in Paths. Like wc, no flags at all means lines, words, and
// bytes; no operands means read stdin.
func ParseArgs(args []string) (Options, error) {
	fs := flag.NewFlagSet("ccwc", flag.ContinueOnError)
	c := fs.Bool("c", false, "number of bytes")
	l := fs.Bool("l", false, "number of lines")
	m := fs.Bool("m", false, "number of characters")
	w := fs.Bool("w", false, "number of words")
	if err := fs.Parse(expandFlags(args)); err != nil {
		return Options{}, err
	}

	opts := Options{
		Bytes: *c,
		Lines: *l,
		Words: *w,
		Chars: *m,
		Paths: fs.Args(),
	}
	if !opts.Bytes && !opts.Lines && !opts.Words && !opts.Chars {
		// wc's default: no flag means lines, words, and bytes.
		opts.Lines, opts.Words, opts.Bytes = true, true, true
	}
	return opts, nil
}

// expandFlags rewrites combined short flags like -cw into -c -w.
// Go's flag package has no notion of flag clustering (wc accepts
// -lwcm), so we split them before parsing. Safe because every ccwc
// flag is a boolean and takes no value.
func expandFlags(args []string) []string {
	var out []string
	for _, arg := range args {
		if isCombinedFlag(arg) {
			for _, r := range arg[1:] {
				out = append(out, "-"+string(r))
			}
			continue
		}
		out = append(out, arg)
	}
	return out
}

// isCombinedFlag reports whether arg looks like -cw: a single dash
// followed by two or more of ccwc's flag letters. Anything else
// (including -z or -c=maybe) is left untouched for the flag package
// to accept or reject with its own error message.
func isCombinedFlag(arg string) bool {
	if len(arg) < 3 || arg[0] != '-' || arg[1] == '-' || strings.Contains(arg, "=") {
		return false
	}
	for _, r := range arg[1:] {
		if !strings.ContainsRune("clmw", r) {
			return false
		}
	}
	return true
}

// openInput returns a reader for path. An empty path or "-" reads
// from stdin, which lets callers treat files and pipes the same way.
// (macOS wc treats a bare "-" as a file name; we follow GNU and read
// stdin.)
func openInput(path string, stdin io.Reader) (io.ReadCloser, error) {
	if path == "" || path == "-" {
		return io.NopCloser(stdin), nil
	}
	return os.Open(path)
}

// Count reads r in chunks and measures it in a single pass: bytes,
// lines, words, and UTF-8 characters. Because it only needs an
// io.Reader it works for files, stdin, and tests.
func Count(r io.Reader) (Counts, error) {
	var c Counts
	var inWord bool
	buf := make([]byte, 4096)

	for {
		n, err := r.Read(buf)

		for _, b := range buf[:n] {
			c.Bytes++
			// UTF-8 continuation bytes (10xxxxxx) identify themselves,
			// so counting every byte that is NOT one counts characters
			// correctly even when a character straddles the buffer.
			if b&0xC0 != 0x80 {
				c.Chars++
			}
			if b == '\n' {
				c.Lines++
			}
			if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
				inWord = false
			} else if !inWord {
				inWord = true
				c.Words++
			}
		}

		if err == io.EOF {
			return c, nil
		}
		if err != nil {
			return Counts{}, err
		}
	}
}

// countBufio is an alternative bufio-based version of Count(), kept
// for reference. More idiomatic Go — bufio.Reader is the standard
// tool for byte-by-byte reading — but ReadByte() makes ONE function
// call per byte in the file, while Count() above only calls Read once
// per 4096-byte chunk. Same memory (both stream with a ~4 KB buffer),
// same result; the chunk loop is just faster on large files.
//
// To use it: uncomment and add "bufio" to the import block.
//
// func countBufio(r io.Reader) (Counts, error) {
// 	br := bufio.NewReader(r)
// 	var c Counts
// 	var inWord bool
//
// 	for {
// 		b, err := br.ReadByte()
// 		if err == io.EOF {
// 			return c, nil
// 		}
// 		if err != nil {
// 			return Counts{}, err
// 		}
//
// 		c.Bytes++
// 		if b&0xC0 != 0x80 {
// 			c.Chars++
// 		}
// 		if b == '\n' {
// 			c.Lines++
// 		}
// 		if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
// 			inWord = false
// 		} else if !inWord {
// 			inWord = true
// 			c.Words++
// 		}
// 	}
// }

// countInput opens path and counts it. "-" and "" read from stdin.
func countInput(path string, stdin io.Reader) (Counts, error) {
	r, err := openInput(path, stdin)
	if err != nil {
		return Counts{}, err
	}
	defer r.Close()
	return Count(r)
}

// writeResult writes one row: every enabled column in wc's fixed
// order (lines, words, bytes, characters), then the file name. An
// empty name (piped stdin) omits the name and its space.
func writeResult(w io.Writer, name string, c Counts, opts Options) error {
	columns := []struct {
		enabled bool
		value   int
	}{
		{opts.Lines, c.Lines},
		{opts.Words, c.Words},
		{opts.Bytes, c.Bytes},
		{opts.Chars, c.Chars},
	}

	var row strings.Builder
	for _, col := range columns {
		if col.enabled {
			fmt.Fprintf(&row, "%8d", col.value)
		}
	}
	if name != "" {
		row.WriteString(" " + name)
	}
	_, err := fmt.Fprintln(w, row.String())
	return err
}
