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
	"reflect"
	"regexp"
)

type StructField struct {
	Name        string
	JsonName    string
	Type        string // TODO (pquerna): fix to be a Type, not a string (?)
	OmitEmpty   bool
	ForceString bool
}

type StructInfo struct {
	Name   string
	Fields []StructField
}

func NewStructInfo(name string) *StructInfo {
	return &StructInfo{
		Name:   name,
		Fields: make([]StructField, 0),
	}
}

var tagRe = regexp.MustCompile("^`(.*)`$")

func (si *StructInfo) AddField(field *ast.Field) error {
	if field.Names == nil || len(field.Names) != 1 {
		return errors.New(fmt.Sprintf("Field contains no name: %v", field))
	}

	jsonName := field.Names[0].Name

	opts := tagOptions("")
	if field.Tag != nil {
		var tagName string
		// the Tag.Value contains wrapping `` which we slice off here. We hope.
		v := tagRe.ReplaceAllString(field.Tag.Value, "$1")
		tag := reflect.StructTag(v).Get("json")
		tagName, opts = parseTag(tag)
		if tagName != "" {
			jsonName = tagName
		}
	}

	si.Fields = append(si.Fields, StructField{
		Name:        field.Names[0].Name,
		JsonName:    jsonName,
		Type:        field.Type.(*ast.Ident).Name, // TODO (pquerna): find a better way to get the Type (?)
		OmitEmpty:   opts.Contains("omitempty"),
		ForceString: opts.Contains("string"),
	})
	return nil
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

			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			stobj := NewStructInfo(k)

			if st.Fields.List != nil {
				for _, field := range st.Fields.List {
					err := stobj.AddField(field)
					if err != nil {
						return "", nil, err
					}
				}
			}
			structs = append(structs, stobj)
		}
	}

	return packageName, structs, nil
}
