/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package main

import (
	_ "github.com/pquerna/ffjson/fflib/v1"
	"github.com/pquerna/ffjson/generator"
	_ "github.com/pquerna/ffjson/inception"

	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var outputPathFlag = flag.String("w", "", "Write generate code to this path instead of ${input}_ffjson.go.")
var goCmdFlag = flag.String("go-cmd", "", "Path to go command; Useful for `goapp` support.")
var importNameFlag = flag.String("import-name", "", "Override import name in case it cannot be detected.")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [options] [input_file]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s generates Go code for optimized JSON serialization.\n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

var extRe = regexp.MustCompile(`(.*)(\.go)$`)

func main() {
	flag.Parse()
	extra := flag.Args()

	if len(extra) != 1 {
		usage()
	}

	inputPath := filepath.ToSlash(extra[0])
	inputFiles := make([]string, 0)

	var outputPath string
	if extRe.MatchString(inputPath) {
		if outputPathFlag == nil || *outputPathFlag == "" {
			outputPath = extRe.ReplaceAllString(inputPath, "${1}_ffjson.go")
		} else {
			outputPath = *outputPathFlag
		}
		inputFiles = []string{inputPath}
	} else {
		// No extension, try package
		if inputPath == "." {
			var err error
			inputPath, err = os.Getwd()
			if err != nil {
				panic(err)
			}
			inputPath = filepath.ToSlash(inputPath)
		}
		parts := strings.Split(inputPath, "/")
		pname := parts[len(parts)-1]
		if outputPathFlag == nil || *outputPathFlag == "" {
			outputPath = inputPath + "/" + pname + "_ffjson.go"
		} else {
			outputPath = *outputPathFlag
		}
		fmt.Println("Walking " + inputPath)
		filepath.Walk(inputPath+"/.", func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}
			if info.IsDir() && "." != info.Name() {
				fmt.Println("Skipping " + info.Name())
				return filepath.SkipDir
			}
			if strings.HasSuffix(info.Name(), "_ffjson.go") {
				return nil
			}
			if strings.HasSuffix(info.Name(), "ffjson_expose.go") {
				return nil
			}
			if strings.HasSuffix(info.Name(), "_test.go") {
				return nil
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
				fmt.Println("Found " + filepath.ToSlash(path))
				inputFiles = append(inputFiles, filepath.ToSlash(path))
			}
			return nil
		})
		if len(inputFiles) == 0 {
			fmt.Println("No files found")
			usage()
		}
	}

	var goCmd string
	if goCmdFlag == nil || *goCmdFlag == "" {
		goCmd = "go"
	} else {
		goCmd = *goCmdFlag
	}

	var importName string
	if importNameFlag != nil && *importNameFlag != "" {
		importName = *importNameFlag
	}

	err := generator.GenerateFiles(goCmd, inputFiles, outputPath, importName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s:\n\n", err)
		os.Exit(1)
	}

	println(outputPath)
}
