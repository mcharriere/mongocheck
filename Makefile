
.PHONY: build install

build:
	go build -o bin/mongocheck

install:
	go install
