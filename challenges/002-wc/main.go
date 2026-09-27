package main

import (
	"fmt"
	"os"

	"github.com/defUserName-404/coding-challenges/challenges/002-wc/src"
)

func main() {
	if err := wc.Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "wc:", err)
		os.Exit(1)
	}
}
