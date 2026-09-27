package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/defUserName-404/coding-challenges/challenges/005-redis/src"
)

func main() {
	addr := flag.String("addr", ":6379", "address to listen on")
	flag.Parse()

	if err := redis.Serve(*addr); err != nil {
		fmt.Fprintln(os.Stderr, "redis:", err)
		os.Exit(1)
	}
}
