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

package tff

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"

	_ "encoding/json"
	"testing"
)

// Test data from https://github.com/akheron/jansson/tree/master/test/suites/invalid
// jansson, Copyright (c) 2009-2014 Petri Lehtinen <petri@digip.org>
// (MIT Licensed)

func TestInvalidApostrophe(t *testing.T) {
	testExpectedError(t,
		&fflib.LexerError{},
		`'`,
		&Xstring{})
}

func TestInvalidASCIIUnicodeIdentifier(t *testing.T) {
	testExpectedError(t,
		&fflib.LexerError{},
		`aå`,
		&Xstring{})
}

func TestInvalidBraceComma(t *testing.T) {
	testExpectedError(t,
		&fflib.LexerError{},
		`{,}`,
		&Xstring{})
}

func TestInvalidBracketComma(t *testing.T) {
	testExpectedError(t,
		&fflib.LexerError{},
		`[,]`,
		&Xarray{})
}
