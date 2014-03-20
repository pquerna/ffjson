
all: test install
	@echo "Done"

install:
	go install github.com/pquerna/ffjson

deps:

fmt:
	go fmt github.com/pquerna/ffjson/...

test: clean
	go test -v github.com/pquerna/ffjson/...

clean:
	go clean -i github.com/pquerna/ffjson/...

.PHONY: deps clean test fmt install all
