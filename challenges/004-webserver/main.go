package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/defUserName-404/coding-challenges/challenges/004-webserver/src"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	root := flag.String("root", ".", "directory to serve")
	flag.Parse()

	if err := webserver.Serve(*addr, *root); err != nil {
		fmt.Fprintln(os.Stderr, "webserver:", err)
		os.Exit(1)
	}
}
