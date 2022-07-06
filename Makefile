.PHONY: all build clean
.SILENT:

filename = "TransactionAPI"

build:
	go build -o ./build/$(filename)__linux-x86_64 ./cmd/TransactionAPI/main.go

	#GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H linux" -o ./build/$(filename)__linux-x86_64 ./cmd/TransactionAPI
	#GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H linux" -o ./build/$(filename)__linux-x86 ./cmd/TransactionAPI

run: build
	./build/$(filename)__linux-x86_64

test:
	go test -v -race -timeout 30s ./...

gen:
	protoc --go_out=. --go_opt=paths=import \
        --go-grpc_out=. --go-grpc_opt=paths=import \
        proto/extPaySys.proto

.DEFAULT_GOAL := run
