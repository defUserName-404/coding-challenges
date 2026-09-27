// Package cat implements the cat coding challenge:
// https://codingchallenges.fyi/challenges/challenge-cat
package cat

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// options holds the parsed command line flags.
type options struct {
	numberAll      bool // -n: number all lines
	numberNonBlank bool // -b: number non-blank lines, overrides -n
}

// Run implements the cat tool. args are the command line arguments
// without the program name: flags followed by file operands. Each
// operand is written to w in order; the operand "-" or no operands at
// all reads from stdin.
func Run(args []string, stdin io.Reader, w io.Writer) error {
	opts, paths, err := parseArgs(args)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		paths = []string{"-"}
	}

	lineNo := 0
	for _, path := range paths {
		lineNo, err = process(path, stdin, w, opts, lineNo)
		if err != nil {
			return err
		}
	}
	return nil
}

// parseArgs splits args into flags and file operands. Go's flag
// package stops at the first non-flag argument, so "-" and file names
// end up in the returned operands.
func parseArgs(args []string) (options, []string, error) {
	fs := flag.NewFlagSet("cccat", flag.ContinueOnError)
	numberAll := fs.Bool("n", false, "number all output lines")
	numberNonBlank := fs.Bool("b", false, "number only nonempty output lines")
	if err := fs.Parse(args); err != nil {
		return options{}, nil, err
	}
	opts := options{
		numberAll:      *numberAll && !*numberNonBlank, // POSIX: -b wins over -n
		numberNonBlank: *numberNonBlank,
	}
	return opts, fs.Args(), nil
}

// process streams a single operand to w. The line counter is threaded
// through so numbering continues across files.
func process(path string, stdin io.Reader, w io.Writer, opts options, lineNo int) (int, error) {
	r, closer, err := openInput(path, stdin)
	if err != nil {
		return lineNo, err
	}
	defer closer()

	return copyInput(r, w, opts, lineNo)
}

// openInput resolves an operand to a reader. "-" is stdin, anything
// else is opened as a file.
func openInput(path string, stdin io.Reader) (io.Reader, func(), error) {
	if path == "-" {
		return stdin, func() {}, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return f, func() { f.Close() }, nil
}

// copyInput writes r to w, numbering lines when a numbering flag is
// set. Unnumbered output streams straight through with io.Copy so
// arbitrarily large inputs never have to fit in memory.
func copyInput(r io.Reader, w io.Writer, opts options, lineNo int) (int, error) {
	if !opts.numberAll && !opts.numberNonBlank {
		_, err := io.Copy(w, r)
		return lineNo, err
	}
	return copyNumbered(r, w, opts, lineNo)
}

// copyNumbered reads r line by line and writes numbered lines to w.
// ReadString grows as needed, so long lines are handled, and reading
// line by line keeps memory usage flat regardless of file size.
func copyNumbered(r io.Reader, w io.Writer, opts options, lineNo int) (int, error) {
	bw := bufio.NewWriter(w)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if line != "" {
			body := strings.TrimSuffix(line, "\n")
			output := body
			if opts.numberAll || (opts.numberNonBlank && body != "") {
				lineNo++
				output = fmt.Sprintf("%6d\t%s", lineNo, body)
			}
			if strings.HasSuffix(line, "\n") {
				output += "\n"
			}
			if _, werr := io.WriteString(bw, output); werr != nil {
				return lineNo, werr
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return lineNo, err
		}
	}
	return lineNo, bw.Flush()
}
