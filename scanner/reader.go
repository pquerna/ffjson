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

package scanner

import (
	"bytes"
	"fmt"
	"io"
)

type FFReader struct {
	s []byte
	i int
	l int
}

func NewFFReader(d []byte) *FFReader {
	return &FFReader{
		s: d,
		i: 0,
		l: len(d),
	}
}

func (r *FFReader) Pos() int {
	return r.i
}

func (r *FFReader) PosWithLine() (int, int) {
	currentLine := 1
	currentChar := 0

	for i := 0; i <= r.i; i++ {
		c := r.s[i]
		currentChar++
		if c == '\n' {
			currentLine++
			currentChar = 0
		}
	}

	return currentLine, currentChar
}

func (r *FFReader) ReadByte() (byte, error) {
	if r.i >= r.l {
		return 0, io.EOF
	}

	r.i++

	return r.s[r.i-1], nil
}

func (r *FFReader) UnreadByte() {
	if r.i <= 0 {
		panic("FFReader.UnreadByte: at beginning of slice")
	}
	r.i--
}

func (r *FFReader) SliceString(out *bytes.Buffer, mask int8, lt [255]int8) error {
	// TODO(pquerna): string_with_escapes?  escape here?
	j := r.i

	for {
		if j >= r.l || r.i >= r.l {
			return io.EOF
		}

		c := r.s[j]
		j++
		if lt[c]&mask == 0 {
			continue
		}

		if c == '"' {
			if j != r.i {
				out.Write(r.s[r.i : j-1])
				r.i = j
			}
			return nil
		}

		// TODO(pquerna): rest of string parsing.
		fmt.Printf("FFTok_error lexString char=%d\n", c)
		return nil
	}

	panic("ffjson: SliceString unreached exit")
}
