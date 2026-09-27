# 005 — Build Your Own Redis Server

> Challenge: https://codingchallenges.fyi/challenges/challenge-redis

**Status:** 🔄 Scaffolded · **Language:** Go

A Redis-compatible server implementing the RESP protocol with key-value commands.

## Layout

- `main.go` — entry point (thin wrapper)
- `src/redis.go` — core logic (`package redis`)
- `tests/redis_test.go` — tests (`package tests`)
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
