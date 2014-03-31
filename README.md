# ffjson: freaking fast JSON serialization for Go / Golang

`ffjson` generates static `MarshalJSON` functions for structures. The generated functions reduce the reliance unpon runtime reflection to do serialization.  In cases where `ffjson` doesn't understand a Type involved, it falls back to `encoding/json`, meaning it is a safe drop in replacement.  By using `ffjson` your JSON serialization just gets faster with no additional code changes.

When you change your `struct`, you will need to run `ffjson` again (or make it part of your build tools).

### [Blog Post explaining more background](https://journal.paul.querna.org/articles/2014/03/31/ffjson-faster-json-in-go/)

## Getting Started

If `myfile.go` contains the `struct` types you would like to be faster, and assuming `GOPATH` is set to a reasonable value for an existing project, you can just run:

    go get -u github.com/pquerna/ffjson
    ffjson myfile.go

## Details

`ffjson` generates code based upon existing `struct` types.  For example, `ffjson foo.go` will by default create a new file `foo_ffjson.go` that contains serialization funcions for all structs found in `foo.go`.

```sh
Usage of ffjson:

	ffjson [options] [input_file]

ffjson generates Go code for optimized JSON serialization.

  -w="": Write generate code to this path instead of ${input}_ffjson.go.
```

## Performance Status:

* `MarshalJSON` is **2x to 3x** faster than `encoding/json`, depending on the structure.
* `UnmarshalJSON` is not yet implemented.

## Features

* **Drop in Replacement:** Because `ffjson` implements the interfaces already defined by `encoding/json` the performance enhancements are transparent to users of your structures.
* **No additional dependencies:** `ffjson` generated code depends on nothing but standard library provided modules.
* **Supports all types:** `ffjson` has native support for most of Go's types -- for any type it doesn't support with fast paths, it falls back to using `encoding/json`.  This means all structures should work out of the box.

## Improvements, bugs, adding features, and taking ffjson new directions!

Please [open issues in Github](https://github.com/pquerna/ffjson/issues) for ideas, bugs, and general thoughts.  Pull requests are of course preferred :)

## License

`ffjson` is licensed under the [Apache License, Version 2.0](./LICENSE)

