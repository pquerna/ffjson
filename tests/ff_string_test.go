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
	"testing"
)

// Test data from https://github.com/akheron/jansson/tree/master/test/suites/valid
// jansson, Copyright (c) 2009-2014 Petri Lehtinen <petri@digip.org>
// (MIT Licensed)

func TestString(t *testing.T) {
	testType(t, &Tstring{}, &Xstring{})
}

func TestStringEscapedControlCharacter(t *testing.T) {
	testExpectedXVal(t,
		"\x12 escaped control character",
		`\u0012 escaped control character`,
		&Xstring{})
}

func TestStringOneByteUTF8(t *testing.T) {
	testExpectedXVal(t,
		", one-byte UTF-8",
		`\u002c one-byte UTF-8`,
		&Xstring{})
}

func TestStringEsccapes(t *testing.T) {
	testExpectedXVal(t,
		`"\`+"\b\f\n\r\t",
		`\"\\\b\f\n\r\t`,
		&Xstring{})

	testExpectedXVal(t,
		`/`,
		`\/`,
		&Xstring{})
}

func TestStringSomeUTF8(t *testing.T) {
	testExpectedXVal(t,
		`€þıœəßð some utf-8 ĸʒ×ŋµåäö𝄞`,
		`€þıœəßð some utf-8 ĸʒ×ŋµåäö𝄞`,
		&Xstring{})
}

func TestString4ByteSurrogate(t *testing.T) {
	testExpectedXVal(t,
		"𝄞 surrogate, four-byte UTF-8",
		`\uD834\uDD1E surrogate, four-byte UTF-8`,
		&Xstring{})
}
