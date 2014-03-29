
all: test install
	@echo "Done"

install:
	go install github.com/pquerna/ffjson

deps:

fmt:
	go fmt github.com/pquerna/ffjson/...

test: ffize
	go test -v github.com/pquerna/ffjson github.com/pquerna/ffjson/generator github.com/pquerna/ffjson/inception github.com/pquerna/ffjson/pills github.com/pquerna/ffjson/tests/...

ffize: install
	ffjson tests/goser/ff/goser.go

bench: ffize all
	ffjson tests/goser/ff/goser.go
	go test -v -benchmem -bench MarshalJSON  github.com/pquerna/ffjson/tests/goser

clean:
	go clean -i github.com/pquerna/ffjson/...
	rm -f tests/*/ff/*_ffjson.go

.PHONY: deps clean test fmt install all
