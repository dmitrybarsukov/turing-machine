TARGET=bin/turing-machine

build:
	mkdir -p bin
	go build -o ${TARGET} cmd/turing-machine/main.go

all: build
