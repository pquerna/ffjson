package tff

import (
	"github.com/stretchr/testify/require"

	"encoding/json"
	"fmt"
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

func setXValue(t *testing.T, thing interface{}) {
	v := reflect.ValueOf(thing)
	v = reflect.Indirect(v)
	f := v.FieldByName("X")
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.SetInt(-42)
	}
}

func getXValue(thing interface{}) interface{} {
	v := reflect.ValueOf(thing)
	v = reflect.Indirect(v)
	f := v.FieldByName("X")
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return f.Int()
	}

	fmt.Printf("%v\n", v.FieldByName("X"))
	return nil
}

func testCycle(t *testing.T, base interface{}, ff interface{}) {
	setXValue(t, base)

	buf, err := json.Marshal(base)
	require.NoError(t, err, "base[%T] failed to Marshal", base)

	err = json.Unmarshal(buf, ff)
	require.NoError(t, err, "ff[%T] failed to Unmarshal", ff)

	require.Equal(t, getXValue(base), getXValue(ff), "json.Unmarshal of base[%T] into ff[%T]", base, ff)
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
