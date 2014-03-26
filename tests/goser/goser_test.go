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

package goser

import (
	"encoding/json"
	base "github.com/pquerna/ffjson/tests/goser/base"
	ff "github.com/pquerna/ffjson/tests/goser/ff"
	"testing"
)

func BenchmarkMarshalJSON(b *testing.B) {
	var record base.Log
	base.NewLog(&record)

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

func BenchmarkFFMarshalJSON(b *testing.B) {
	var record ff.Log
	ff.NewLog(&record)

	buf, err := record.MarshalJSON()
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
