# ffjson: faster JSON serialization for Go / Golang

[![Build Status](https://travis-ci.org/pquerna/ffjson.svg?branch=master)](https://travis-ci.org/pquerna/ffjson)

`ffjson` generates static `MarshalJSON` and `UnmarshalJSON` functions for structures in Go. The generated functions reduce the reliance unpon runtime reflection to do serialization and are generally 2 to 3 times faster.  In cases where `ffjson` doesn't understand a Type involved, it falls back to `encoding/json`, meaning it is a safe drop in replacement.  By using `ffjson` your JSON serialization just gets faster with no additional code changes.

When you change your `struct`, you will need to run `ffjson` again (or make it part of your build tools).

## Blog Posts

* 2014-03-31: [First Release and Background](https://journal.paul.querna.org/articles/2014/03/31/ffjson-faster-json-in-go/)

## Getting Started

If `myfile.go` contains the `struct` types you would like to be faster, and assuming `GOPATH` is set to a reasonable value for an existing project (meaning that in this particular example if `myfile.go` is in the `myproject` directory, the project should be under `$GOPATH/src/myproject`), you can just run:

    go get -u github.com/pquerna/ffjson
    ffjson myfile.go
    git add myfile_ffjson.go

## Details

`ffjson` generates code based upon existing `struct` types.  For example, `ffjson foo.go` will by default create a new file `foo_ffjson.go` that contains serialization functions for all structs found in `foo.go`.

```sh
Usage of ffjson:

	ffjson [options] [input_file]

ffjson generates Go code for optimized JSON serialization.

  -w="": Write generate code to this path instead of ${input}_ffjson.go.
```

## Performance Status:

* `MarshalJSON` is **2x to 3x** faster than `encoding/json`.
* `UnmarshalJSON` is **2x to 3x** faster than `encoding/json`.

## Features

* **Unmarshal Support:** Since v0.9, `ffjson` supports Unmarshaling of structures.
* **Drop in Replacement:** Because `ffjson` implements the interfaces already defined by `encoding/json` the performance enhancements are transparent to users of your structures.
* **Supports all types:** `ffjson` has native support for most of Go's types -- for any type it doesn't support with fast paths, it falls back to using `encoding/json`.  This means all structures should work out of the box. If they don't, [open a issue!](https://github.com/pquerna/ffjson/issues)
* **ffjson: skip**: If you have a structure you want `ffjson` to ignore, add `ffjson: skip` to the doc string for this structure.
* **Extensive Tests:** `ffjson` contains an extensive test suite including fuzz'ing against the JSON parser.

## Improvements, bugs, adding features, and taking ffjson new directions!

Please [open issues in Github](https://github.com/pquerna/ffjson/issues) for ideas, bugs, and general thoughts.  Pull requests are of course preferred :)

## Credits

`ffjson` has recieved significant contributions from:

* [Klaus Post](https://github.com/klauspost)
* [Paul Querna](https://github.com/pquerna) 

## License

`ffjson` is licensed under the [Apache License, Version 2.0](./LICENSE)

