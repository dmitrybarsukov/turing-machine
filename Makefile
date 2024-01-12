TARGET=bin/turing-machine

build:
	go build -o ${TARGET} cmd/turing-machine/main.go

all: build
