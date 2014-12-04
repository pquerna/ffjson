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
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"regexp"
)

type StructField struct {
	Name string
}

type StructInfo struct {
	Name string
}

func NewStructInfo(name string) *StructInfo {
	return &StructInfo{
		Name: name,
	}
}

var skipre = regexp.MustCompile("(.*)ffjson:(\\s*)((skip)|(ignore))(.*)")

func ExtractStructs(inputPath string) (string, []*StructInfo, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)

	if err != nil {
		return "", nil, err
	}

	packageName := f.Name.String()
	structs := make(map[string]*StructInfo)

	for k, d := range f.Scope.Objects {
		if d.Kind == ast.Typ {
			ts, ok := d.Decl.(*ast.TypeSpec)
			if !ok {
				return "", nil, fmt.Errorf("Unknown type without TypeSec: %v", d)
			}

			_, ok = ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// TODO(pquerna): Add // ffjson:skip or similiar tagging.
			stobj := NewStructInfo(k)

			structs[k] = stobj
		}
	}

	files := map[string]*ast.File{
		inputPath: f,
	}

	pkg, _ := ast.NewPackage(fset, files, nil, nil)

	d := doc.New(pkg, f.Name.String(), doc.AllDecls)
	for _, t := range d.Types {
		if skipre.MatchString(t.Doc) {
			delete(structs, t.Name)
		}
	}

	rv := make([]*StructInfo, 0)
	for _, v := range structs {
		rv = append(rv, v)
	}
	return packageName, rv, nil
}
