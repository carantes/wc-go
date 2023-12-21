# WC Command (Go version)

Replica of Linux wc cli, used to count the number of words, lines, bytes and characters given a file or stdin.

This is part of the Coding Challenges series by John Crickett https://codingchallenges.fyi/challenges/challenge-wc.

## Setup

Build and install the app

```bash
make build
```

```bash
make install
```

## Run

```bash
wc-go -l sample.txt
```

or reading from stdin

```bash
cat sample.txt| wc-go -w
```

Flags:
-c, --bytes Number of bytes
-m, --characters Number of characters
-h, --help help for wc-go
-l, --lines Number of lines
-w, --words Number of words

## Run tests

```bash
make test
```

## About this Project

This project was built using `Go 1.21.5` and `Cobra`, a CLI and interface to build CLI apps.
