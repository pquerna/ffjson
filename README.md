# ffjson: freaking fast JSON serialization for Go / Golang

`ffjson` generates static `MarshalJSON` and `UnmarshalJSON` functions for  structures. The generated functions reduce the reliance unpon runtime reflection to do serialization.  In cases where `ffjson` doesn't understand a Type involved, it falls back to `encoding/json`, meaning it is a safe drop in replacement.  By using `ffjson` your JSON serialization just gets faster with no additional code changes.

When you change your `struct`, you will need to run `ffjson` again (or make it part of your build tools).

## Getting Started

If `myfile.go` contains the `struct` types you would like to be faster, and assuming GOPATH is set to a reasonable value for an existing project, you can just run:

    go get -u github.com/pquerna/ffjson
    ffjson myfile.go

## Details

`ffjson` generates code based on an existing `struct` in go.  For example, `ffjson mypackage/foo.go` will by default create a new file `mypackage/foo_ffjson.go` that contains serialization funcions for all structs found in `foo.go`.

```sh
Usage of ffjson:

	ffjson [options] [input_file]

ffjson generates Go code for optimized JSON serialization.

  -w="": Write generate code to this path instead of ${input}_ffjson.go.
```

## Status:

* `MarshalJSON` is working and about 25% faster.
* `UnmarshalJSON` has not been started.

## Improving, adding features, taking ffjson new directions!

Please [open issues in Github](https://github.com/pquerna/ffjson/issues) for ideas, bugs, and general thoughts.  Pull requests are of course preferred :)

## License

`ffjson` is licensed under the [Apache License, Version 2.0](./LICENSE)

