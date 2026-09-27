package main

import (
	"fmt"
	"os"

	"github.com/defUserName-404/coding-challenges/challenges/003-shell/src"
)

func main() {
	if err := shell.Run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "shell:", err)
		os.Exit(1)
	}
}
