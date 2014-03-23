#!/bin/bash

set -e

make -C ..
ffjson ff.go
go test -benchmem -bench MarshalJSON
go test -benchmem -bench MarshalJSONNative -cpuprofile="prof.dat" -benchtime 10s
go tool pprof -gif tests.test prof.dat >out.gif
