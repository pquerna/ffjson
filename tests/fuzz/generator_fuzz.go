// +build gofuzz

package fuzz

import (
	"io/ioutil"
	"os"

	_ "github.com/dvyukov/go-fuzz/go-fuzz-dep"
	"github.com/pquerna/ffjson/generator"
)

// Fuzz tests code generation.
func FuzzGenerate(fuzz []byte) int {
	err := os.MkdirAll("fuzzing", os.ModePerm)
	if err != nil {
		panic("could not make fuzzing dir")
	}
	err = ioutil.WriteFile("fuzzing/input.go", fuzz, 0644)
	if err != nil {
		panic("could not write input file")
	}
	err = generator.GenerateFiles(
		"go",
		"fuzzing/input.go",
		"fuzzing/output.go",
		"",
		true,
		true,
	)
	if err != nil {
		return 0
	}
	return 1
}
