.PHONY: build
build:
	go build -o bin/md5-request-tool main.go

.PHONY: test
test:
	go test -v -race ./...