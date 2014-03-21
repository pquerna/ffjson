#!/bin/bash

set -e

make -C ..
ffjson ff.go
go test -benchmem -bench MarshalJSON

