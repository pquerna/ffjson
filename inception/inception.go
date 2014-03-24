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
	"fmt"
	"os"
)

type Inception struct {
	objs []interface{}
	ic   *InceptionContext
}

func NewInception() *Inception {
	return &Inception{
		objs: make([]interface{}, 0),
	}
}

func (i *Inception) Add(obj interface{}) {
	i.objs = append(i.objs, obj)
}

func (i *Inception) generateMarshalJSON() error {
	for _, obj := range i.objs {
		si := NewStructInfo(obj)
		err := CreateMarshalJSON(i.ic, si)
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
	var outputPath = "" // TODO: argv
	i.ic = NewInceptionContext(outputPath)
	err := i.generateMarshalJSON()
	if err != nil {
		i.handleError(err)
		return
	}
}
