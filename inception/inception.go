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

package ffjsoninception

import (
	"errors"
	"fmt"
	"github.com/pquerna/ffjson/pills"
	"io/ioutil"
	"os"
)

type Inception struct {
	objs          []*StructInfo
	InputPath     string
	OutputPath    string
	PackageName   string
	OutputImports map[string]bool
	OutputFuncs   []string
	OutputPills   map[pills.Pill]bool
}

func NewInception(inputPath string, packageName string, outputPath string) *Inception {
	return &Inception{
		objs:          make([]*StructInfo, 0),
		InputPath:     inputPath,
		OutputPath:    outputPath,
		PackageName:   packageName,
		OutputFuncs:   make([]string, 0),
		OutputImports: make(map[string]bool),
		OutputPills:   make(map[pills.Pill]bool),
	}
}

func (i *Inception) Add(obj interface{}) {
	i.objs = append(i.objs, NewStructInfo(obj))
}

func (i *Inception) generateCode() error {
	for _, si := range i.objs {
		err := CreateMarshalJSON(i, si)
		if err != nil {
			return err
		}
		err = CreateUnmarshalJSON(i, si)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Inception) handleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s:\n\n", err)
	os.Exit(1)
}

func (i *Inception) Execute() {
	if len(os.Args) != 1 {
		i.handleError(errors.New(fmt.Sprintf("Internal ffjson error: inception executable takes no args: %v", os.Args)))
		return
	}

	err := i.generateCode()
	if err != nil {
		i.handleError(err)
		return
	}

	data, err := RenderTemplate(i)
	if err != nil {
		i.handleError(err)
		return
	}

	stat, err := os.Stat(i.InputPath)

	if err != nil {
		i.handleError(err)
		return
	}

	err = ioutil.WriteFile(i.OutputPath, data, stat.Mode())

	if err != nil {
		i.handleError(err)
		return
	}

}
