# MD5 request tool

This tool makes http requests and prints the address of the request along with the MD5 hash of the response.

## Table of Contents

 - [Install](#install)
 - [Makefile usage](#makefile usage)
 - [Usage without Makefile](#usage without Makefile)
 - [Examples](#examples)

## Install
Before you begin you must have Go installed and configured properly for your computer. Please see https://golang.org/doc/install

Clone the git repository in your system

```bash
$ git clone https://github.com/SwirlGit/md5-request-tool.git
$ cd md5-request-tool
```

## Makefile usage

Build tool
```bash
$ make build
```

Test tool
```bash
$ make test
```

Run simple example
```bash
$ make run-example
```

## Usage without Makefile
Build tool
```bash
$ go build -o bin/md5-request-tool main.go
```

Test tool
```bash
$ go test -v -race ./...
```

Run simple example
```bash
$ ./bin/md5-request-tool -parallel=3 google.com facebook.com yandex.com mail.com twitter.com reddit.com/r/funny \
		broken_random_link.c reddit.com/r/notfunny yahoo.com
```

## Examples

```bash
http://yandex.com 46a88738b322d882a026bb23369fce80
http://google.com 91edd2b7b370c5fa27a48260053841c9
http://facebook.com 3175c2ed4ff8b558808f2cb8927b63eb
http://mail.com ef19fcda6acd3a90db606c3341b5117b
http://broken_random_link.c Get "http://broken_random_link.c": dial tcp: lookup broken_random_link.c: no such host
http://twitter.com 4d62222cb39fe06829f23c946abc4ce1
http://reddit.com/r/funny 042264de53fc22a1db23868f0b4e4d53
http://reddit.com/r/notfunny a6ed543f741686959135216b77baf0cc
http://yahoo.com b809d84deffc042a23121d510f320ae7
```

This repository doesn't include dependencies beyond Go's standard libraries.