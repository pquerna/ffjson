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
	"errors"
	"strconv"
	"testing"
)

func scanAll(ffl *FFLexer) []FFTok {
	rv := make([]FFTok, 0, 0)
	for {
		tok := ffl.Scan()
		rv = append(rv, tok)
		if tok == FFTok_eof || tok == FFTok_error {
			break
		}
	}

	return rv
}

func assertTokensEqual(t *testing.T, a []FFTok, b []FFTok) {

	if len(a) != len(b) {
		t.Fatalf("Token lists of mixed length: expected=%v found=%v", a, b)
		return
	}

	for i, v := range a {
		if b[i] != v {
			t.Fatalf("Invalid Token: expected=%d found=%d token=%d",
				v, b, i)
			return
		}
	}
}

func scanToTok(ffl *FFLexer, targetTok FFTok) error {
	_, err := scanToTokCount(ffl, targetTok)
	return err
}

func scanToTokCount(ffl *FFLexer, targetTok FFTok) (int, error) {
	c := 0
	for {
		tok := ffl.Scan()
		c++

		if tok == targetTok {
			return c, nil
		}

		if tok == FFTok_error {
			return c, errors.New("Hit error before target token")
		}
		if tok == FFTok_eof {
			return c, errors.New("Hit EOF before target token")
		}
	}

	return c, errors.New("Could not find target token.")
}

func TestBasicLexing(t *testing.T) {
	ffl := NewFFLexer(bytes.NewBufferString(`{}`))
	toks := scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)
}

func TestHelloWorld(t *testing.T) {
	ffl := NewFFLexer(bytes.NewBufferString(`{"hello":"world"}`))
	toks := scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_string,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": 1}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_integer,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": 1.0}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_double,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": 1e2}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_double,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": {}}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_left_bracket,
		FFTok_right_bracket,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": {"blah": null}}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_null,
		FFTok_right_bracket,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": /* comment */ 0}`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_comment,
		FFTok_integer,
		FFTok_right_bracket,
		FFTok_eof,
	}, toks)

	ffl = NewFFLexer(bytes.NewBufferString(`{"hello": / comment`))
	toks = scanAll(ffl)
	assertTokensEqual(t, []FFTok{
		FFTok_left_bracket,
		FFTok_string,
		FFTok_colon,
		FFTok_error,
	}, toks)
}

func tDouble(t *testing.T, input string, target float64) {
	ffl := NewFFLexer(bytes.NewBufferString(input))
	err := scanToTok(ffl, FFTok_double)
	if err != nil {
		t.Fatalf("scanToTok failed, couldnt find double: %v input: %v", err, input)
	}

	f64, err := strconv.ParseFloat(string(ffl.Output), 64)
	if err != nil {
		t.Fatalf("ParseFloat failed, shouldnt of: %v input: %v", err, input)
	}

	if int64(f64*1000) != int64(target*1000) {
		t.Fatalf("ffl.Output: expected f64 '%v', got: %v from: %v input: %v",
			target, f64, string(ffl.Output), input)
	}

	err = scanToTok(ffl, FFTok_eof)
	if err != nil {
		t.Fatal("Failed to find EOF after double. input: %v", input)
	}
}

func TestDouble(t *testing.T) {
	tDouble(t, `{"a": 1.2}`, 1.2)
	tDouble(t, `{"a": 1.2e2}`, 1.2e2)
	tDouble(t, `{"a": -1.2e2}`, -1.2e2)
	tDouble(t, `{"a": 1.2e-2}`, 1.2e-2)
	tDouble(t, `{"a": -1.2e-2}`, -1.2e-2)
}

func tInt(t *testing.T, input string, target int64) {
	ffl := NewFFLexer(bytes.NewBufferString(input))
	err := scanToTok(ffl, FFTok_integer)
	if err != nil {
		t.Fatalf("scanToTok failed, couldnt find int: %v input: %v", err, input)
	}

	// Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64.
	i64, err := strconv.ParseInt(string(ffl.Output), 10, 64)
	if err != nil {
		t.Fatalf("ParseInt failed, shouldnt of: %v input: %v", err, input)
	}

	if i64 != target {
		t.Fatalf("ffl.Output: expected i64 '%v', got: %v from: %v", target, i64, string(ffl.Output))
	}

	err = scanToTok(ffl, FFTok_eof)
	if err != nil {
		t.Fatal("Failed to find EOF after int. input: %v", input)
	}
}

func TestInt(t *testing.T) {
	tInt(t, `{"a": 2000}`, 2000)
	tInt(t, `{"a": -2000}`, -2000)
	tInt(t, `{"a": 0}`, 0)
	tInt(t, `{"a": -0}`, -0)
}

func tError(t *testing.T, input string, targetCount int, targetError FFErr) {
	ffl := NewFFLexer(bytes.NewBufferString(input))
	count, err := scanToTokCount(ffl, FFTok_error)
	if err != nil {
		t.Fatalf("scanToTok failed, couldnt find error token: %v input: %v", err, input)
	}

	if count != targetCount {
		t.Fatalf("Expected error token at offset %v, but found it at %v input: %v",
			count, targetCount, input)
	}

	if ffl.Error != targetError {
		t.Fatalf("Expected error token %v, but got %v input: %v",
			targetError, ffl.Error, input)
	}
}

func TestBroken(t *testing.T) {
	tError(t, `{"a": nul}`, 4, FFErr_invalid_string)
	tError(t, `{"a": 1.a}`, 4, FFErr_missing_integer_after_decimal)
}
