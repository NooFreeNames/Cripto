.PHONY: build
build:
	rm -rf build 
	mkdir build
	go build -o build/cripto -v ./cmd

.PHONY: run
run:
	go run cmd/main.go

.DEFAULT_GOAL := build