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
