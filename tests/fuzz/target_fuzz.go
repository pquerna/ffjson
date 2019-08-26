// +build gofuzz

package fuzz

import (
	_ "github.com/dvyukov/go-fuzz/go-fuzz-dep"
)

// FuzzUnmarshal tests unmarshaling.
func FuzzUnmarshal(fuzz []byte) int {
	data := &Data{}
	err := data.UnmarshalJSON(fuzz)
	if err != nil {
		return 0
	}
	_, err = data.MarshalJSON()
	if err != nil {
		return 0
	}
	return 1
}
