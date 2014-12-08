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
	"github.com/stretchr/testify/require"

	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"
	"testing"
)

func newLogRecord() *Record {
	return &Record{
		OriginId: 11,
		Method:   "POST",
	}
}

func newLogFFRecord() *FFRecord {
	return &FFRecord{
		OriginId: 11,
		Method:   "POST",
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	record := newLogRecord()

	buf, err := json.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&record)
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkMarshalJSONNative(b *testing.B) {
	record := newLogFFRecord()

	buf, err := json.Marshal(&record)
	if err != nil {
		b.Fatalf("Marshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := record.MarshalJSON()
		if err != nil {
			b.Fatalf("Marshal: %v", err)
		}
	}
}

func BenchmarkSimpleUnmarshal(b *testing.B) {
	record := newLogFFRecord()
	buf := []byte(`{"id": 123213, "OriginId": 22, "meth": "GET"}`)
	err := record.UnmarshalJSON(buf)
	if err != nil {
		b.Fatalf("UnmarshalJSON: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := record.UnmarshalJSON(buf)
		if err != nil {
			b.Fatalf("UnmarshalJSON: %v", err)
		}
	}
}

func BenchmarkSXimpleUnmarshalNative(b *testing.B) {
	record := newLogRecord()
	buf := []byte(`{"id": 123213, "OriginId": 22, "meth": "GET"}`)
	err := json.Unmarshal(buf, record)
	if err != nil {
		b.Fatalf("json.Unmarshal: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(buf, record)
		if err != nil {
			b.Fatalf("json.Unmarshal: %v", err)
		}
	}
}

func TestSimpleUnmarshal(t *testing.T) {
	record := newLogFFRecord()

	err := record.UnmarshalJSON([]byte(`{"id": 123213, "OriginId": 22, "meth": "GET"}`))
	if err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}

	if record.Timestamp != 123213 {
		t.Fatalf("record.Timestamp: expected: 0 got: %v", record.Timestamp)
	}

	if record.OriginId != 22 {
		t.Fatalf("record.OriginId: expected: 22 got: %v", record.OriginId)
	}

	if record.Method != "GET" {
		t.Fatalf("record.Method: expected: GET got: %v", record.Method)
	}
}

func testType(t *testing.T, base interface{}, ff interface{}) {
	testSameMarshal(t, base, ff)
	testCycle(t, base, ff)
}

func testSameMarshal(t *testing.T, base interface{}, ff interface{}) {
	bufbase, err := json.Marshal(base)
	require.NoError(t, err, "base[%T] failed to Marshal", base)

	bufff, err := json.Marshal(ff)
	require.NoError(t, err, "ff[%T] failed to Marshal", ff)

	require.Equal(t, bufbase, bufff, "json.Marshal of base[%T] != ff[%T]", base, ff)
}

func testCycle(t *testing.T, base interface{}, ff interface{}) {
	setXValue(t, base)

	buf, err := json.Marshal(base)
	require.NoError(t, err, "base[%T] failed to Marshal", base)

	err = json.Unmarshal(buf, ff)
	require.NoError(t, err, "ff[%T] failed to Unmarshal", ff)

	require.Equal(t, getXValue(base), getXValue(ff), "json.Unmarshal of base[%T] into ff[%T]", base, ff)
}

func testExpectedX(t *testing.T, expected interface{}, base interface{}, ff interface{}) {
	buf, err := json.Marshal(base)
	require.NoError(t, err, "base[%T] failed to Marshal", base)

	err = json.Unmarshal(buf, ff)
	require.NoError(t, err, "ff[%T] failed to Unmarshal", ff)

	require.Equal(t, expected, getXValue(ff), "json.Unmarshal of base[%T] into ff[%T]", base, ff)
}

func testExpectedXValBare(t *testing.T, expected interface{}, xval string, ff interface{}) {
	buf := []byte(`{"X":` + xval + `}`)
	err := json.Unmarshal(buf, ff)
	require.NoError(t, err, "ff[%T] failed to Unmarshal", ff)

	require.Equal(t, expected, getXValue(ff), "json.Unmarshal of %T into ff[%T]", xval, ff)
}

func testExpectedXVal(t *testing.T, expected interface{}, xval string, ff interface{}) {
	testExpectedXValBare(t, expected, `"`+xval+`"`, ff)
}

func testExpectedError(t *testing.T, expected error, xval string, ff json.Unmarshaler) {
	buf := []byte(`{"X":` + xval + `}`)
	err := ff.UnmarshalJSON(buf)
	require.Error(t, err, "ff[%T] failed to Unmarshal", ff)
	require.IsType(t, expected, err)
}

func setXValue(t *testing.T, thing interface{}) {
	v := reflect.ValueOf(thing)
	v = reflect.Indirect(v)
	f := v.FieldByName("X")
	switch f.Kind() {
	case reflect.Bool:
		f.SetBool(true)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.SetInt(-42)
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f.SetUint(42)
	case reflect.Float32, reflect.Float64:
		f.SetFloat(3.141592653)
	case reflect.String:
		f.SetString("hello world")
	}
}

func getXValue(thing interface{}) interface{} {
	v := reflect.ValueOf(thing)
	v = reflect.Indirect(v)
	f := v.FieldByName("X")
	switch f.Kind() {
	case reflect.Bool:
		return f.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int()
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return f.Uint()
	case reflect.Float32, reflect.Float64:
		return f.Float()
	case reflect.String:
		return f.String()
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(f)
	return buf.String()
}

func TestArray(t *testing.T) {
	testType(t, &Tarray{X: []int{}}, &Xarray{X: []int{}})
	testCycle(t, &Tarray{X: []int{42, -42, 44}}, &Xarray{X: []int{}})
}

func TestArrayPtr(t *testing.T) {
	testType(t, &TarrayPtr{X: []*int{}}, &XarrayPtr{X: []*int{}})
	v := 33
	testCycle(t, &TarrayPtr{X: []*int{&v}}, &XarrayPtr{X: []*int{}})
}

func TestTimeDuration(t *testing.T) {
	testType(t, &Tduration{}, &Xduration{})
}

func TestBool(t *testing.T) {
	testType(t, &Tbool{}, &Xbool{})
}

func TestInt(t *testing.T) {
	testType(t, &Tint{}, &Xint{})
}

func TestInt8(t *testing.T) {
	testType(t, &Tint8{}, &Xint8{})
}

func TestInt16(t *testing.T) {
	testType(t, &Tint16{}, &Xint16{})
}

func TestInt32(t *testing.T) {
	testType(t, &Tint32{}, &Xint32{})
}

func TestInt64(t *testing.T) {
	testType(t, &Tint64{}, &Xint64{})
}

func TestUint(t *testing.T) {
	testType(t, &Tuint{}, &Xuint{})
}

func TestUint8(t *testing.T) {
	testType(t, &Tuint8{}, &Xuint8{})
}

func TestUint16(t *testing.T) {
	testType(t, &Tuint16{}, &Xuint16{})
}

func TestUint32(t *testing.T) {
	testType(t, &Tuint32{}, &Xuint32{})
}

func TestUint64(t *testing.T) {
	testType(t, &Tuint64{}, &Xuint64{})
}

func TestUintptr(t *testing.T) {
	testType(t, &Tuintptr{}, &Xuintptr{})
}

func TestFloat32(t *testing.T) {
	testType(t, &Tfloat32{}, &Xfloat32{})
}

func TestFloat64(t *testing.T) {
	testType(t, &Tfloat64{}, &Xfloat64{})
}
