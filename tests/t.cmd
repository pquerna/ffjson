go install github.com/pquerna/ffjson
del ff_ffjson.go
go test -v github.com/pquerna/ffjson/fflib/v1 github.com/pquerna/ffjson/generator github.com/pquerna/ffjson/inception && ffjson ff.go && go test -v
