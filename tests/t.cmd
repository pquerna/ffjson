go install github.com/jborozdina/ffjson
del ff_ffjson.go
del goser\ff\goser_ffjson.go
del go.stripe\ff\customer_ffjson.go
del types\ff\everything_ffjson.go

go test -v github.com/jborozdina/ffjson/fflib/v1 github.com/jborozdina/ffjson/generator github.com/jborozdina/ffjson/inception && ffjson ff.go && go test -v
ffjson goser/ff/goser.go && go test github.com/jborozdina/ffjson/tests/goser
ffjson go.stripe/ff/customer.go && go test github.com/jborozdina/ffjson/tests/go.stripe
ffjson types/ff/everything.go && go test github.com/jborozdina/ffjson/tests/types


