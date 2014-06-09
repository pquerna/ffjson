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

package pills

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type Pill int32

const (
	Pill_WriteJsonString Pill = 0
	Pill_FormatBits      Pill = 1
)

var PillFiles = map[Pill]string{
	Pill_WriteJsonString: "jsonstring.go",
	Pill_FormatBits:      "iota.go",
}

var PillNames = map[Pill]string{
	Pill_WriteJsonString: "WriteJsonString",
	Pill_FormatBits:      "FormatBits",
}

func extractFunc(funcName string, inputPath string) ([]string, string, error) {
	fset := token.NewFileSet()

	fp, err := parser.ParseFile(fset, inputPath, nil, 0)

	if err != nil {
		return nil, "", err
	}

	imports := make([]string, 0)

	for _, imp := range fp.Imports {
		imports = append(imports, imp.Path.Value)
	}

	var buf bytes.Buffer

	for _, decl := range fp.Decls {
		f, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if f.Name.Name != funcName {
			continue
		}

		f.Name.Name = "ffjson_" + f.Name.Name
		printer.Fprint(&buf, fset, f)
		break
	}

	return imports, string(buf.Bytes()), nil
}

func getPath(p Pill) (string, error) {
	gopaths := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
	rvs := make([]string, 0)
	for _, path := range gopaths {
		gpath, err := filepath.Abs(path)
		if err != nil {
			continue
		}

		rv := filepath.Join(strings.Replace(gpath, "\\", "/", -1), "src", "github.com", "pquerna", "ffjson", "pills", PillFiles[p])

		if _, err := os.Stat(rv); os.IsNotExist(err) {
			rvs = append(rvs, rv)
			continue
		}

		return rv, nil
	}
	return "", errors.New(fmt.Sprintf("no such file or directory: %s  GOPATH=%s", rvs, gopaths))
}

func GetPill(p Pill) ([]string, string, error) {

	inputPath, err := getPath(p)

	if err != nil {
		return nil, "", err
	}

	return extractFunc(PillNames[p], inputPath)
}
