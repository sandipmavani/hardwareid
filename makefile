.PHONY: build clean default test

build: clean
	@go build -o hardwareid ./cmd/hardwareid/main.go

clean:
	@rm -rf ./hardwareid

test:
	go test ./...

default: build
