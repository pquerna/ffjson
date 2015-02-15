del ff_ffjson.go
ffjson ff.go

go test -benchmem -test.run="Nothing" -bench MarshalJSON
go test -benchmem -test.run="Nothing" -test.bench="MarshalJSONNative" -cpuprofile="prof.dat" -benchtime 10s &&go tool pprof -gif tests.test.exe prof.dat >cpu.gif
go test -benchmem -test.run="Nothing" -test.bench="MarshalJSONNativeReuse" -memprofile="prof.dat" -test.memprofilerate=1 -benchtime 10s &&go tool pprof -lines -alloc_objects -gif tests.test.exe prof.dat  >mem-reuse.gif

go test -benchmem -test.run="Nothing" -test.bench="MarshalJSONNativePool" -memprofile="prof.dat" -test.memprofilerate=1 -benchtime 10s &&go tool pprof -lines -alloc_objects -gif tests.test.exe prof.dat  >mem-pool.gif
