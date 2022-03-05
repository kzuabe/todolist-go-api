.PHONY: run
run:
	go run ./cmd/todolist

.PHONY: build
build:
	go build ./cmd/todolist

.PHONY: test
test:
	go test ./...

.PHONY: wire
wire:
	wire ./cmd/todolist
