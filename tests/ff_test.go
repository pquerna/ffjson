package tff

import (
	"encoding/json"
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
	err := record.XUnmarshalJSON(buf)
	if err != nil {
		b.Fatalf("XUnmarshalJSON: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := record.XUnmarshalJSON(buf)
		if err != nil {
			b.Fatalf("XUnmarshalJSON: %v", err)
		}
	}
}

func BenchmarkSXimpleUnmarshalNative(b *testing.B) {
	record := newLogRecord()
	buf := []byte(`{"id": 123213, "OriginId": 22, "meth": "GET"}`)
	err := json.Unmarshal(buf, record)
	if err != nil {
		b.Fatalf("XUnmarshalJSON: %v", err)
	}
	b.SetBytes(int64(len(buf)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(buf, record)
		if err != nil {
			b.Fatalf("XUnmarshalJSON: %v", err)
		}
	}
}

func TestSimpleUnmarshal(t *testing.T) {
	record := newLogFFRecord()

	err := record.XUnmarshalJSON([]byte(`{"id": 123213, "OriginId": 22, "meth": "GET"}`))
	if err != nil {
		t.Fatalf("XUnmarshalJSON: %v", err)
	}

	t.Logf("record.Timestamp: %v", record.Timestamp)
	t.Logf("record.OriginId: %v", record.OriginId)
	t.Logf("record.Method: %v", record.Method)
}
