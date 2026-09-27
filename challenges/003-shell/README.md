# 003 — Build Your Own Shell

> Challenge: https://codingchallenges.fyi/challenges/challenge-shell

**Status:** 🔄 Scaffolded · **Language:** Go

An interactive shell supporting builtins, external commands, `PATH` lookup, and more.

## Layout

- `main.go` — entry point (thin wrapper)
- `src/shell.go` — core logic (`package shell`)
- `tests/shell_test.go` — tests (`package tests`)
- `tests/testdata/` — test fixtures

## Run

```sh
go run .
```

## Test

```sh
go test ./...
```

## Notes

Approach, learnings, and what I'd improve go here.
