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
	"github.com/pquerna/ffjson/pills"

	"bytes"
	"fmt"
	"io"
	"unicode/utf16"
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

// Calcuates the Position with line and line offset,
// because this isn't counted for performance reasons,
// it will iterate the buffer from the begining, and should
// only be used in error-paths.
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

func (r *FFReader) ReadByteNoWS() (byte, error) {
	if r.i >= r.l {
		return 0, io.EOF
	}

	j := r.i

	for {
		c := r.s[j]
		j++

		// inline whitespace parsing gives another ~8% performance boost
		// for many kinds of nicely indented JSON.
		// ... and using a [255]bool instead of multiple ifs, gives another 2%
		/*
			if c != '\t' &&
				c != '\n' &&
				c != '\v' &&
				c != '\f' &&
				c != '\r' &&
				c != ' ' {
				r.i = j
				return c, nil
			}
		*/
		if whitespaceLookupTable[c] == false {
			r.i = j
			return c, nil
		}

		if j >= r.l {
			return 0, io.EOF
		}
	}
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

func (r *FFReader) readU4(j int) (rune, error) {

	var u4 [4]byte
	for i := 0; i < 4; i++ {
		if j >= r.l {
			return -1, io.EOF
		}
		c := r.s[j]
		if byteLookupTable[c]&VHC != 0 {
			u4[i] = c
			j++
			continue
		} else {
			// TODO(pquerna): handle errors better.
			return -1, fmt.Errorf("lex_string_invalid_hex_char: %v %v", c, string(u4[:]))
		}
	}

	// TODO(pquerna): utf16.IsSurrogate
	rr, err := pills.ParseUint(u4[:], 16, 64)
	if err != nil {
		return -1, err
	}
	return rune(rr), nil
}

func (r *FFReader) SliceString(out *bytes.Buffer) error {
	mask := IJC | NFP

	// TODO(pquerna): string_with_escapes? de-escape here?
	j := r.i

	for {
		if j >= r.l || r.i >= r.l {
			return io.EOF
		}

		c := r.s[j]
		j++
		if byteLookupTable[c]&mask == 0 {
			continue
		}

		if c == '"' {
			if j != r.i {
				out.Write(r.s[r.i : j-1])
				r.i = j
			}
			return nil
		} else if c == '\\' {
			if j >= r.l {
				return io.EOF
			}

			c = r.s[j]
			j++

			if c == 'u' {
				ru, err := r.readU4(j)
				if err != nil {
					return err
				}
				if utf16.IsSurrogate(ru) {
					ru2, err := r.readU4(j + 6)
					if err != nil {
						return err
					}
					out.Write(r.s[r.i : j-2])
					r.i = j + 10
					j = r.i
					out.WriteRune(utf16.DecodeRune(ru, ru2))
				} else {
					out.Write(r.s[r.i : j-2])
					r.i = j + 4
					j = r.i
					out.WriteRune(ru)
				}

				continue
			} else if byteLookupTable[c]&VEC != 0 {
				// yajl_lex_string_invalid_escaped_char;
			}
		}

		/**
		 * VEC - valid escaped control char
		 * note.  the solidus '/' may be escaped or not.
		 * IJC - invalid json char
		 * VHC - valid hex char
		 * NFP - needs further processing (from a string scanning perspective)
		 * NUC - needs utf8 checking when enabled (from a string scanning perspective)
		 */

		// TODO(pquerna): rest of string parsing.
		// fmt.Printf("FFTok_error lexString char=%d string=%s\n", c, string(r.s[r.i:j-1]))
		continue
	}

	panic("ffjson: SliceString unreached exit")
}

// TODO(pquerna): consider combining wibth the normal byte mask.
var whitespaceLookupTable [255]bool = [255]bool{
	false, /* 0 */
	false, /* 1 */
	false, /* 2 */
	false, /* 3 */
	false, /* 4 */
	false, /* 5 */
	false, /* 6 */
	false, /* 7 */
	false, /* 8 */
	true,  /* 9 */
	true,  /* 10 */
	true,  /* 11 */
	true,  /* 12 */
	true,  /* 13 */
	false, /* 14 */
	false, /* 15 */
	false, /* 16 */
	false, /* 17 */
	false, /* 18 */
	false, /* 19 */
	false, /* 20 */
	false, /* 21 */
	false, /* 22 */
	false, /* 23 */
	false, /* 24 */
	false, /* 25 */
	false, /* 26 */
	false, /* 27 */
	false, /* 28 */
	false, /* 29 */
	false, /* 30 */
	false, /* 31 */
	true,  /* 32 */
	false, /* 33 */
	false, /* 34 */
	false, /* 35 */
	false, /* 36 */
	false, /* 37 */
	false, /* 38 */
	false, /* 39 */
	false, /* 40 */
	false, /* 41 */
	false, /* 42 */
	false, /* 43 */
	false, /* 44 */
	false, /* 45 */
	false, /* 46 */
	false, /* 47 */
	false, /* 48 */
	false, /* 49 */
	false, /* 50 */
	false, /* 51 */
	false, /* 52 */
	false, /* 53 */
	false, /* 54 */
	false, /* 55 */
	false, /* 56 */
	false, /* 57 */
	false, /* 58 */
	false, /* 59 */
	false, /* 60 */
	false, /* 61 */
	false, /* 62 */
	false, /* 63 */
	false, /* 64 */
	false, /* 65 */
	false, /* 66 */
	false, /* 67 */
	false, /* 68 */
	false, /* 69 */
	false, /* 70 */
	false, /* 71 */
	false, /* 72 */
	false, /* 73 */
	false, /* 74 */
	false, /* 75 */
	false, /* 76 */
	false, /* 77 */
	false, /* 78 */
	false, /* 79 */
	false, /* 80 */
	false, /* 81 */
	false, /* 82 */
	false, /* 83 */
	false, /* 84 */
	false, /* 85 */
	false, /* 86 */
	false, /* 87 */
	false, /* 88 */
	false, /* 89 */
	false, /* 90 */
	false, /* 91 */
	false, /* 92 */
	false, /* 93 */
	false, /* 94 */
	false, /* 95 */
	false, /* 96 */
	false, /* 97 */
	false, /* 98 */
	false, /* 99 */
	false, /* 100 */
	false, /* 101 */
	false, /* 102 */
	false, /* 103 */
	false, /* 104 */
	false, /* 105 */
	false, /* 106 */
	false, /* 107 */
	false, /* 108 */
	false, /* 109 */
	false, /* 110 */
	false, /* 111 */
	false, /* 112 */
	false, /* 113 */
	false, /* 114 */
	false, /* 115 */
	false, /* 116 */
	false, /* 117 */
	false, /* 118 */
	false, /* 119 */
	false, /* 120 */
	false, /* 121 */
	false, /* 122 */
	false, /* 123 */
	false, /* 124 */
	false, /* 125 */
	false, /* 126 */
	false, /* 127 */
	false, /* 128 */
	false, /* 129 */
	false, /* 130 */
	false, /* 131 */
	false, /* 132 */
	false, /* 133 */
	false, /* 134 */
	false, /* 135 */
	false, /* 136 */
	false, /* 137 */
	false, /* 138 */
	false, /* 139 */
	false, /* 140 */
	false, /* 141 */
	false, /* 142 */
	false, /* 143 */
	false, /* 144 */
	false, /* 145 */
	false, /* 146 */
	false, /* 147 */
	false, /* 148 */
	false, /* 149 */
	false, /* 150 */
	false, /* 151 */
	false, /* 152 */
	false, /* 153 */
	false, /* 154 */
	false, /* 155 */
	false, /* 156 */
	false, /* 157 */
	false, /* 158 */
	false, /* 159 */
	false, /* 160 */
	false, /* 161 */
	false, /* 162 */
	false, /* 163 */
	false, /* 164 */
	false, /* 165 */
	false, /* 166 */
	false, /* 167 */
	false, /* 168 */
	false, /* 169 */
	false, /* 170 */
	false, /* 171 */
	false, /* 172 */
	false, /* 173 */
	false, /* 174 */
	false, /* 175 */
	false, /* 176 */
	false, /* 177 */
	false, /* 178 */
	false, /* 179 */
	false, /* 180 */
	false, /* 181 */
	false, /* 182 */
	false, /* 183 */
	false, /* 184 */
	false, /* 185 */
	false, /* 186 */
	false, /* 187 */
	false, /* 188 */
	false, /* 189 */
	false, /* 190 */
	false, /* 191 */
	false, /* 192 */
	false, /* 193 */
	false, /* 194 */
	false, /* 195 */
	false, /* 196 */
	false, /* 197 */
	false, /* 198 */
	false, /* 199 */
	false, /* 200 */
	false, /* 201 */
	false, /* 202 */
	false, /* 203 */
	false, /* 204 */
	false, /* 205 */
	false, /* 206 */
	false, /* 207 */
	false, /* 208 */
	false, /* 209 */
	false, /* 210 */
	false, /* 211 */
	false, /* 212 */
	false, /* 213 */
	false, /* 214 */
	false, /* 215 */
	false, /* 216 */
	false, /* 217 */
	false, /* 218 */
	false, /* 219 */
	false, /* 220 */
	false, /* 221 */
	false, /* 222 */
	false, /* 223 */
	false, /* 224 */
	false, /* 225 */
	false, /* 226 */
	false, /* 227 */
	false, /* 228 */
	false, /* 229 */
	false, /* 230 */
	false, /* 231 */
	false, /* 232 */
	false, /* 233 */
	false, /* 234 */
	false, /* 235 */
	false, /* 236 */
	false, /* 237 */
	false, /* 238 */
	false, /* 239 */
	false, /* 240 */
	false, /* 241 */
	false, /* 242 */
	false, /* 243 */
	false, /* 244 */
	false, /* 245 */
	false, /* 246 */
	false, /* 247 */
	false, /* 248 */
	false, /* 249 */
	false, /* 250 */
	false, /* 251 */
	false, /* 252 */
	false, /* 253 */
	false, /* 254 */
}
