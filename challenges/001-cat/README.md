# 001 — Build Your Own cat

> Challenge: https://codingchallenges.fyi/challenges/challenge-cat

**Status:** ✅ Done · **Language:** Go

Concatenate files and print them to standard output, like the Unix `cat`.
Implements all challenge steps: print a file (1), read stdin with `-` (2),
concatenate files (3), `-n` number all lines (4), `-b` number non-blank
lines (5).

## Layout

- `main.go` — entry point (thin wrapper: args in, errors to stderr, exit codes)
- `src/cat.go` — core logic (`package cat`)
- `tests/cat_test.go` — tests (`package tests`)
- `tests/testdata/` — test fixtures

## Build

```sh
go build -o bin/cccat .
```

`bin/` is gitignored. In the examples below, substitute `go run .` if you
don't want to build.

## Usage

```sh
# step 1: file -> stdout
bin/cccat tests/testdata/quotes1.txt

# step 2: stdin
echo "hello" | bin/cccat -

# step 3: concatenate multiple files
bin/cccat tests/testdata/quotes1.txt tests/testdata/quotes2.txt

# step 4: number all lines
sed G tests/testdata/quotes1.txt | bin/cccat -n

# step 5: number non-blank lines only
sed G tests/testdata/quotes1.txt | bin/cccat -b
```

## Test

```sh
go test ./...
```

## Notes

- Line numbering matches system `cat -n` (`%6d` + tab), so you can verify
  with `diff <(bin/cccat -n file) <(cat -n file)`.
- Unnumbered output streams with `io.Copy`; numbered output reads line by
  line with `bufio.Reader` — neither loads a whole file into memory.
- `-b` overrides `-n` (POSIX behaviour).
- The flag package stops parsing at the first operand, so flags must come
  before file names: `cccat -n file`, not `cccat file -n`.
