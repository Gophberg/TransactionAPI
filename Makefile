.PHONY:
.SILENT:

filename = "TransactionAPI"

build:
	go build -o ./build/$(filename) ./cmd/TransactionAPI/main.go

#	GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./.build/$(filename)__linux-x86_64 cmd/TransactionAPI
#	GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./.build/$(filename)__linux-x86 cmd/TransactionAPI

run: build
	./build/$(filename)

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run
