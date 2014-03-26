
all: test install
	@echo "Done"

install:
	go install github.com/pquerna/ffjson

deps:

fmt:
	go fmt github.com/pquerna/ffjson/...

test: clean
	go test -v github.com/pquerna/ffjson github.com/pquerna/ffjson/generator

bench: all
	ffjson tests/goser/ff/goser.go
	go test -v -benchmem -bench MarshalJSON  github.com/pquerna/ffjson/tests/goser

clean:
	go clean -i github.com/pquerna/ffjson/...

.PHONY: deps clean test fmt install all
