all: test bench

.PHONY: test
test:
	go generate ./...
	go test -v -timeout 30s ./...

.PHONY:
bench: bench
	go generate ./...
	go test -benchmem -run=^$ -bench=... -v ./...

lint:
	golangci-lint run ./...
