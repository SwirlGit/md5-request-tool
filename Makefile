.PHONY: build
build:
	go build -o bin/md5-request-tool main.go

.PHONY: test
test:
	go test -v -race ./...

.PHONY: run-example
run-example:
	make build
	./bin/md5-request-tool -parallel=3 google.com facebook.com yandex.com mail.com twitter.com reddit.com/r/funny \
		broken_random_link.c reddit.com/r/notfunny yahoo.com