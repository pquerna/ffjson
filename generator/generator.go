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
	"os"
)

func GenerateFiles(goCmd string, inputPath string, outputPath string, importName string, forceRegenerate bool, resetFields bool, strict bool) error {
	if _, StatErr := os.Stat(outputPath); !os.IsNotExist(StatErr) {
		inputFileInfo, inputFileErr := os.Stat(inputPath)
		outputFileInfo, outputFileErr := os.Stat(outputPath)

		if nil == outputFileErr && nil == inputFileErr {
			if !forceRegenerate && inputFileInfo.ModTime().Before(outputFileInfo.ModTime()) {
				fmt.Println("File " + outputPath + " already exists.")

				return nil
			}
		}
	}

	packageName, structs, err := ExtractStructs(inputPath)
	if err != nil {
		return err
	}

	im := NewInceptionMain(goCmd, inputPath, outputPath, resetFields, strict)

	err = im.Generate(packageName, structs, importName)
	if err != nil {
		return fmt.Errorf("error=%v path=%q", err, im.TempMainPath)
	}

	err = im.Run()
	if err != nil {
		return err
	}

	return nil
}
