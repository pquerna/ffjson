package tff

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestUnmarshalPointerTimes(t *testing.T) {
	for _, input := range []string{
		`{"Slice":["2020-01-02T03:04:05Z",null]}`,
		`{"Array":["2020-01-02T03:04:05Z",null]}`,
		`{"Map":{"date":"2020-01-02T03:04:05Z","empty":null}}`,
		`{"Direct":"2020-01-02T03:04:05Z"}`,
		`{"Slice":[null],"Array":[null,null],"Map":{"date":null},"Direct":null}`,
	} {
		t.Run(input, func(t *testing.T) {
			defer func() {
				if p := recover(); p != nil {
					t.Errorf("unexpected panic: %v", p)
				}
			}()
			now := time.Now()
			got := PointerTimes{Array: [2]*time.Time{&now, &now}, Direct: &now}
			type plain PointerTimes
			want := plain(got)
			if err := json.Unmarshal([]byte(input), &want); err != nil {
				t.Fatal(err)
			}
			if err := got.UnmarshalJSON([]byte(input)); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(plain(got), want) {
				t.Errorf("got %+v; want %+v", got, want)
			}
		})
	}
}
