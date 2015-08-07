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
)

func GenerateFiles(goCmd string, inputPath []string, outputPath string, importName string) error {
	var packageName string
	var structs []*StructInfo
	for _, input := range inputPath {
		p, s, err := ExtractStructs(input)
		if err != nil {
			return err
		}
		if packageName != "" && p != packageName {
			return fmt.Errorf("Multiple package names in package: %s and %s", packageName, p)
		}
		packageName = p
		structs = append(structs, s...)
	}

	im := NewInceptionMain(goCmd, inputPath[0], outputPath)

	err := im.Generate(packageName, structs, importName)
	if err != nil {
		return errors.New(fmt.Sprintf("error=%v path=%q", err, im.TempMainPath))
	}

	err = im.Run()
	if err != nil {
		return err
	}

	return nil
}
