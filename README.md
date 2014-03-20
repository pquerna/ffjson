# ffjson: freaking fast JSON serialization for Go / Golang

`ffjson` generates static `MarshalJSON` and `UnmarshalJSON` functions for a structure that reduce the reliance unpon runtime reflection to do serialization.  `ffjson` is meant for when you have a few JSON objects that have a similiar object and structure shape all of the time, and don't want all of the features of the main `encoding/json` package.

## Installation

    go install github.com/pquerna/ffjson

## Running

`ffjson` generates code based on an existing `struct` in go.  For example, `ffjson mypackage/foo.go` will by default create a new file `mypackage/foo_ffjson.go` that contains serialization funcions for all structs found in `foo.go`.

```sh
ffjson: generate freaking fast json handling Go code.

Usage:

   ffjson [options] [input_file]

   -w [path]  	Write generate code to this path instead of _ffjson.go.
 ```


## Improving, adding features, taking ffjson new directions!

Please [open issues in Github]() for ideas, bugs, and general thoughts.  Pull requests are of course preferred :)

## License

`ffjson` is licensed under the [Apache License, Version 2.0](./LICENSE)

