del ff_ffjson.go
ffjson ff.go

go test -benchmem -bench MarshalJSON
go test -benchmem -bench MarshalJSONNative -cpuprofile="prof.dat" -benchtime 10s &&go tool pprof -gif tests.test.exe prof.dat >out.gif
