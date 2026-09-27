# 002 — Build Your Own wc Tool

> Challenge: https://codingchallenges.fyi/challenges/challenge-wc

**Status:** ✅ Done · **Language:** Go

Count lines, words, and bytes (or characters) in files or from standard input, like the Unix `wc`.
Implements all challenge steps: `-c` bytes (1), `-l` lines (2), `-w` words (3),
`-m` characters (4), default lines+words+bytes (5), read stdin (6).

## Layout

- `main.go` — entry point (thin wrapper: args in, errors to stderr, exit codes)
- `src/wc.go` — core logic (`package wc`)
- `tests/count_test.go` — unit tests for `Count` / `ParseArgs` (`package tests`)
- `tests/wc_test.go` — end-to-end tests (`package tests`)
- `tests/testdata/` — test fixtures

## Build

```sh
go build -o bin/ccwc .
```

`bin/` is gitignored. In the examples below, substitute `go run .` if you
don't want to build.

## Usage

```sh
# step 1: bytes
bin/ccwc -c tests/testdata/test.txt

# steps 2-4: lines, words, characters
bin/ccwc -l tests/testdata/test.txt
bin/ccwc -w tests/testdata/test.txt
bin/ccwc -m tests/testdata/test.txt

# step 5: default (lines, words, bytes)
bin/ccwc tests/testdata/test.txt

# step 6: stdin
echo "hello" | bin/ccwc -c

# multiple files print one row each plus a total
bin/ccwc -c tests/testdata/test.txt tests/testdata/test.txt

# combined short flags work too
bin/ccwc -lwcm tests/testdata/test.txt
```

## Test

```sh
go test ./...
```

## Notes

- Output matches macOS `wc` byte-for-byte (verified with `diff`) for every
  flag, the default mode, multi-file `total`, and stdin. One intentional
  difference: `-c -m` together prints both columns, where macOS prints only
  the character count.
- One pass over the input: `Count` reads 4096-byte chunks and tallies bytes,
  lines, words, and UTF-8 characters — no file is ever fully loaded.
- Characters are counted as "bytes that are not UTF-8 continuation bytes",
  so a multi-byte character split across the chunk boundary still counts once.
- Flags are parsed only in `ParseArgs`, and `writeResult` prints from a table
  of enabled columns — adding a flag never touches the counting loop.
- Combined short flags like `-lwcm` are expanded before parsing; Go's `flag`
  package doesn't cluster flags itself.
- The flag package stops parsing at the first operand, so flags must come
  before file names: `ccwc -c file`, not `ccwc file -c`.
