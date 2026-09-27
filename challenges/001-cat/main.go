package main

import (
	"fmt"
	"os"

	"github.com/defUserName-404/coding-challenges/challenges/001-cat/src"
)

func main() {
	if err := cat.Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "cccat:", err)
		os.Exit(1)
	}
}
