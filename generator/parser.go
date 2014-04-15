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
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type StructInfo struct {
	Name string
}

func NewStructInfo(name string) *StructInfo {
	return &StructInfo{
		Name: name,
	}
}

func ExtractStructs(inputPath string) (string, []*StructInfo, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, inputPath, nil, 0)

	if err != nil {
		return "", nil, err
	}

	packageName := f.Name.String()
	structs := make([]*StructInfo, 0)

	for k, d := range f.Scope.Objects {
		if d.Kind == ast.Typ {
			ts, ok := d.Decl.(*ast.TypeSpec)
			if !ok {
				return "", nil, errors.New(fmt.Sprintf("Unknown type without TypeSec: %v", d))
			}

			_, ok = ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			stobj := NewStructInfo(k)

			structs = append(structs, stobj)
		}
	}

	return packageName, structs, nil
}
