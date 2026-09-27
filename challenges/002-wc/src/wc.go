// Package wc implements the wc coding challenge:
// https://codingchallenges.fyi/challenges/challenge-wc
package wc

import (
	"errors"
	"io"
)

// Run parses flags from args and writes the wc output for each
// operand to w. The operand "-" (or no operands at all) reads from stdin.
// TODO: implement - currently returns an error.
func Run(args []string, stdin io.Reader, w io.Writer) error {
	return errors.New("not implemented yet")
}
