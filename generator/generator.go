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

package generator

import (
	"io/ioutil"
	"os"
)

type GenContext struct {
	OutputFuncs  []string
	IsEmptyValue bool
	WriteString  bool
}

func NewGenContext() *GenContext {
	return &GenContext{
		OutputFuncs: make([]string, 0),
	}
}

func (gc *GenContext) AddFunc(out string) {
	gc.OutputFuncs = append(gc.OutputFuncs, out)
}

func GenerateFiles(inputPath string, outputPath string) error {
	packageName, structs, err := ExtractStructs(inputPath)
	if err != nil {
		return err
	}

	gc := NewGenContext()

	for _, st := range structs {
		err := CreateMarshalJSON(gc, st)
		if err != nil {
			return err
		}
	}

	data, err := RenderTemplate(inputPath, packageName, gc)

	if err != nil {
		return err
	}

	stat, err := os.Stat(inputPath)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath, data, stat.Mode())

	if err != nil {
		return err
	}

	return nil
}
