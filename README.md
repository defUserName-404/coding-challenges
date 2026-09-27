# Coding Challenges

Solutions to [Coding Challenges by John Crickett](https://codingchallenges.fyi/challenges/intro) — a weekly series of projects where you build real tools (wc, grep, a shell, a Redis server, …) to level up as a software engineer.

**Primary language: Go.** A rare challenge may use a different language if it fits better — each folder stays self-contained either way.

## Ground rules

- Every challenge is fully self-contained: its own `go.mod`, README, tests, and fixtures live **inside** the challenge folder.
- Always run build/test commands **from inside the challenge folder** — never from the repo root.
- No binaries or build output are committed (`bin/`, `*.test`, etc. are gitignored).
- No clutter: delete abandoned experiments instead of leaving partial solutions behind.
- Every challenge README documents what it does and how to build/run/test it.

## Layout

```
challenges/
├── 001-cat/
│   ├── README.md        # what it does, how to run/test
│   ├── go.mod           # its own module — self-contained
│   ├── main.go          # entry point (thin wrapper)
│   ├── src/
│   │   └── cat.go       # core logic (package cat)
│   └── tests/
│       ├── cat_test.go  # tests (package tests)
│       └── testdata/    # test fixtures
├── 002-wc/
│   └── ...
└── ...
```

Core logic lives in `src/` as a library package; `main.go` at the challenge root is only a thin wrapper; tests live in their own `tests/` package and import `src`.

Folders are numbered in **solve order** (the beginner roadmap first), not the site's listing order.

Suggested order (John's [beginner roadmap](https://codingchallenges.fyi/blog/junior-developer-challenges)): **cat → wc → shell → web server → redis**.

## Running a challenge

```sh
cd challenges/001-cat
go run .                # run
go test ./...           # test
go build -o bin/ .      # build binary into bin/ (gitignored)
```

## Progress

⬜ not started · 🔄 scaffolded / in progress · ✅ done

| Challenge | Status | Solution |
| --- | --- | --- |
| [cat](https://codingchallenges.fyi/challenges/challenge-cat) | ✅ | [challenges/001-cat](challenges/001-cat/) |
| [wc](https://codingchallenges.fyi/challenges/challenge-wc) | 🔄 | [challenges/002-wc](challenges/002-wc/) |
| [shell](https://codingchallenges.fyi/challenges/challenge-shell) | 🔄 | [challenges/003-shell](challenges/003-shell/) |
| [web server](https://codingchallenges.fyi/challenges/challenge-webserver) | 🔄 | [challenges/004-webserver](challenges/004-webserver/) |
| [redis server](https://codingchallenges.fyi/challenges/challenge-redis) | 🔄 | [challenges/005-redis](challenges/005-redis/) |
| [JSON parser](https://codingchallenges.fyi/challenges/challenge-json-parser) | ⬜ | — |
| [compression tool (Huffman)](https://codingchallenges.fyi/challenges/challenge-huffman) | ⬜ | — |
| [cut](https://codingchallenges.fyi/challenges/challenge-cut) | ⬜ | — |
| [load balancer](https://codingchallenges.fyi/challenges/challenge-load-balancer) | ⬜ | — |
| [sort](https://codingchallenges.fyi/challenges/challenge-sort) | ⬜ | — |
| [calculator](https://codingchallenges.fyi/challenges/challenge-calculator) | ⬜ | — |
| [grep](https://codingchallenges.fyi/challenges/challenge-grep) | ⬜ | — |
| [uniq](https://codingchallenges.fyi/challenges/challenge-uniq) | ⬜ | — |
| [URL shortener](https://codingchallenges.fyi/challenges/challenge-url-shortener) | ⬜ | — |
| [diff](https://codingchallenges.fyi/challenges/challenge-diff) | ⬜ | — |
| [IRC client](https://codingchallenges.fyi/challenges/challenge-irc) | ⬜ | — |
| [memcached server](https://codingchallenges.fyi/challenges/challenge-memcached) | ⬜ | — |
| [Spotify client](https://codingchallenges.fyi/challenges/challenge-spotify) | ⬜ | — |
| [Discord bot](https://codingchallenges.fyi/challenges/challenge-discord) | ⬜ | — |
| [LinkedIn carousel generator](https://codingchallenges.fyi/challenges/challenge-licq) | ⬜ | — |
| [sed](https://codingchallenges.fyi/challenges/challenge-sed) | ⬜ | — |
| [DNS resolver](https://codingchallenges.fyi/challenges/challenge-dns-resolver) | ⬜ | — |
| [traceroute](https://codingchallenges.fyi/challenges/challenge-traceroute) | ⬜ | — |
| [realtime chat client & server](https://codingchallenges.fyi/challenges/challenge-realtime-chat) | ⬜ | — |
| [NATS message broker](https://codingchallenges.fyi/challenges/challenge-nats) | ⬜ | — |
| [git](https://codingchallenges.fyi/challenges/challenge-git) | ⬜ | — |
| [rate limiter](https://codingchallenges.fyi/challenges/challenge-rate-limiter) | ⬜ | — |
| [NTP client](https://codingchallenges.fyi/challenges/challenge-ntp) | ⬜ | — |
| [scheduling automation app](https://codingchallenges.fyi/challenges/challenge-scheduler) | ⬜ | — |
| [Lisp interpreter](https://codingchallenges.fyi/challenges/challenge-lisp) | ⬜ | — |
| [QR code generator](https://codingchallenges.fyi/challenges/challenge-qr-generator) | ⬜ | — |
| [crontab tool](https://codingchallenges.fyi/challenges/challenge-cron) | ⬜ | — |
| [head](https://codingchallenges.fyi/challenges/challenge-head) | ⬜ | — |
| [jq](https://codingchallenges.fyi/challenges/challenge-jq) | ⬜ | — |
| [Google Keep](https://codingchallenges.fyi/challenges/challenge-keep) | ⬜ | — |
| [Pong](https://codingchallenges.fyi/challenges/challenge-pong) | ⬜ | — |
| [Redis CLI tool](https://codingchallenges.fyi/challenges/challenge-redis-cli) | ⬜ | — |
| [network modelling tool](https://codingchallenges.fyi/challenges/challenge-network-modeller) | ⬜ | — |
| [social media tool](https://codingchallenges.fyi/challenges/challenge-sm-tool) | ⬜ | — |
| [curl](https://codingchallenges.fyi/challenges/challenge-curl) | ⬜ | — |
| [HTTP(S) load tester](https://codingchallenges.fyi/challenges/challenge-load-tester) | ⬜ | — |
| [tr](https://codingchallenges.fyi/challenges/challenge-tr) | ⬜ | — |
| [Tetris](https://codingchallenges.fyi/challenges/challenge-tetris) | ⬜ | — |
| [DNS forwarder](https://codingchallenges.fyi/challenges/challenge-dns-forwarder) | ⬜ | — |
| [port scanner](https://codingchallenges.fyi/challenges/challenge-port-scanner) | ⬜ | — |
| [yq](https://codingchallenges.fyi/challenges/challenge-yq) | ⬜ | — |
| [Chrome extension](https://codingchallenges.fyi/challenges/challenge-chrome-extension) | ⬜ | — |
| [data privacy vault](https://codingchallenges.fyi/challenges/challenge-data-privacy-vault) | ⬜ | — |
| [password cracker](https://codingchallenges.fyi/challenges/challenge-password-cracker) | ⬜ | — |
| [xargs](https://codingchallenges.fyi/challenges/challenge-xargs) | ⬜ | — |
| [HTTP forward proxy server](https://codingchallenges.fyi/challenges/challenge-forward-proxy) | ⬜ | — |
| [Docker](https://codingchallenges.fyi/challenges/challenge-docker) | ⬜ | — |
| [spell checker (bloom filter)](https://codingchallenges.fyi/challenges/challenge-bloom) | ⬜ | — |
| [tar](https://codingchallenges.fyi/challenges/challenge-tar) | ⬜ | — |
| [xxd](https://codingchallenges.fyi/challenges/challenge-xxd) | ⬜ | — |
| [chess game](https://codingchallenges.fyi/challenges/challenge-chess) | ⬜ | — |
| [snake game](https://codingchallenges.fyi/challenges/challenge-snake) | ⬜ | — |
| [password manager](https://codingchallenges.fyi/challenges/challenge-password-manager) | ⬜ | — |
| [netcat](https://codingchallenges.fyi/challenges/challenge-netcat) | ⬜ | — |
| [pastebin](https://codingchallenges.fyi/challenges/challenge-pastebin) | ⬜ | — |
| [Dropbox](https://codingchallenges.fyi/challenges/challenge-dropbox) | ⬜ | — |
| [git contributions visualisation](https://codingchallenges.fyi/challenges/challenge-contrib-vis) | ⬜ | — |
| [Space Invaders](https://codingchallenges.fyi/challenges/challenge-space-invaders) | ⬜ | — |
| [Spotify playlist backup](https://codingchallenges.fyi/challenges/challenge-spotify-backup) | ⬜ | — |
| [Minesweeper](https://codingchallenges.fyi/challenges/challenge-minesweeper) | ⬜ | — |
| [zip file cracker](https://codingchallenges.fyi/challenges/challenge-zip-cracker) | ⬜ | — |
| [YAML parser](https://codingchallenges.fyi/challenges/challenge-yaml) | ⬜ | — |
| [blogging software](https://codingchallenges.fyi/challenges/challenge-blog) | ⬜ | — |
| [Notion](https://codingchallenges.fyi/challenges/challenge-notion) | ⬜ | — |
| [memcached CLI tool](https://codingchallenges.fyi/challenges/challenge-memcached-client) | ⬜ | — |
| [wheel of names](https://codingchallenges.fyi/challenges/challenge-wheel) | ⬜ | — |
| [Sudoku](https://codingchallenges.fyi/challenges/challenge-sudoku) | ⬜ | — |
| [text editor](https://codingchallenges.fyi/challenges/challenge-text-editor) | ⬜ | — |
| [Asteroids](https://codingchallenges.fyi/challenges/challenge-asteroids) | ⬜ | — |
| [duplicate file finder](https://codingchallenges.fyi/challenges/challenge-duplicate-files) | ⬜ | — |
| [video chat application](https://codingchallenges.fyi/challenges/challenge-video-chat) | ⬜ | — |
| [static site generator](https://codingchallenges.fyi/challenges/challenge-ssg) | ⬜ | — |
| [uptime monitoring service](https://codingchallenges.fyi/challenges/challenge-uptime-monitoring) | ⬜ | — |
| [socat](https://codingchallenges.fyi/challenges/challenge-socat) | ⬜ | — |
| [optical character recognition](https://codingchallenges.fyi/challenges/challenge-ocr) | ⬜ | — |
| [Brainfuck interpreter](https://codingchallenges.fyi/challenges/challenge-brainfuck) | ⬜ | — |
| [markdown to PDF editor](https://codingchallenges.fyi/challenges/challenge-md-to-pdf) | ⬜ | — |
| [markdown presentation tool](https://codingchallenges.fyi/challenges/challenge-md-to-slides) | ⬜ | — |
| [Mandelbrot set explorer](https://codingchallenges.fyi/challenges/challenge-mandelbrot) | ⬜ | — |
| [time zone converter](https://codingchallenges.fyi/challenges/timezone-converter) | ⬜ | — |
| [strace](https://codingchallenges.fyi/challenges/challenge-strace) | ⬜ | — |
| [code comment remover](https://codingchallenges.fyi/challenges/challenge-code-comment-remover) | ⬜ | — |
| [top](https://codingchallenges.fyi/challenges/challenge-top) | ⬜ | — |
| [ELIZA](https://codingchallenges.fyi/challenges/challenge-eliza) | ⬜ | — |
| [SMTP server](https://codingchallenges.fyi/challenges/challenge-smtp) | ⬜ | — |
| [Monkeytype](https://codingchallenges.fyi/challenges/challenge-monkeytype) | ⬜ | — |
| [LOC counter](https://codingchallenges.fyi/challenges/challenge-loc-counter) | ⬜ | — |
| [which](https://codingchallenges.fyi/challenges/challenge-which) | ⬜ | — |
| [Forth interpreter](https://codingchallenges.fyi/challenges/challenge-forth) | ⬜ | — |
| [whois](https://codingchallenges.fyi/challenges/challenge-whois) | ⬜ | — |
| [Bitcask](https://codingchallenges.fyi/challenges/challenge-bitcask) | ⬜ | — |
| [spelling correction tool](https://codingchallenges.fyi/challenges/challenge-spelling-correction) | ⬜ | — |
| [language server (LSP)](https://codingchallenges.fyi/challenges/challenge-lsp) | ⬜ | — |
| [BitTorrent client](https://codingchallenges.fyi/challenges/challenge-bittorrent) | ⬜ | — |
| [echo server](https://codingchallenges.fyi/challenges/challenge-echo) | ⬜ | — |
| [LLM-powered AI chatbot](https://codingchallenges.fyi/challenges/challenge-llm.chatbot) | ⬜ | — |
| [software teleprompter](https://codingchallenges.fyi/challenges/challenge-teleprompter) | ⬜ | — |
| [MCP server for AI agents](https://codingchallenges.fyi/challenges/challenge-mcp-server) | ⬜ | — |
| [top programming stories dashboard](https://codingchallenges.fyi/challenges/challenge-top-stories) | ⬜ | — |
| [JSON validator and prettier](https://codingchallenges.fyi/challenges/challenge-json-validator) | ⬜ | — |
| [Loom clone](https://codingchallenges.fyi/challenges/challenge-loom) | ⬜ | — |
| [online Python playground](https://codingchallenges.fyi/challenges/challenge-online-python-playground) | ⬜ | — |
| [ebook reader](https://codingchallenges.fyi/challenges/challenge-ebook-reader) | ⬜ | — |
| [RTFM-for-me agent](https://codingchallenges.fyi/challenges/challenge-rtfm-agent) | ⬜ | — |
