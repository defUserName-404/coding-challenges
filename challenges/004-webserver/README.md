# 004 — Build Your Own Web Server

> Challenge: https://codingchallenges.fyi/challenges/challenge-webserver

**Status:** 🔄 Scaffolded · **Language:** Go

Serve static files over HTTP, implementing enough of the HTTP protocol by hand to understand how it works.

## Layout

- `main.go` — entry point (thin wrapper)
- `src/webserver.go` — core logic (`package webserver`)
- `tests/webserver_test.go` — tests (`package tests`)
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
