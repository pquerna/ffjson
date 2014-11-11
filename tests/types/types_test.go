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

package types

import (
	"encoding/json"
	ff "github.com/pquerna/ffjson/tests/types/ff"
	"reflect"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	var record ff.Everything
	var recordTripped ff.Everything
	ff.NewEverything(&record)

	buf1, err := json.Marshal(&record)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	err = json.Unmarshal(buf1, &recordTripped)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	good := reflect.DeepEqual(record.FooStruct, recordTripped.FooStruct)
	if !good {
		t.Fatalf("Expected: %v\n Got: %v", *record.FooStruct, *recordTripped.FooStruct)
	}

	record.FooStruct = nil
	recordTripped.FooStruct = nil

	good = reflect.DeepEqual(record, recordTripped)
	if !good {
		t.Fatalf("Expected: %v\n Got: %v", record, recordTripped)
	}
}

func TestUnmarshalEmpty(t *testing.T) {
	record := ff.Everything{}
	err := record.XUnmarshalJSON([]byte(`{}`))
	if err != nil {
		t.Fatalf("XUnmarshalJSON: %v", err)
	}
}

func TestUnmarshalFull(t *testing.T) {
	record := ff.Everything{}
	// TODO(pquerna): add unicode snowman
	// TODO(pquerna): handle arrays
	// TODO(pquerna): handle Bar subtype
	err := record.XUnmarshalJSON([]byte(`{
    "Bool": true,
    "Int": 1,
    "Int8": 2,
    "Int16": 3,
    "Int32": -4,
    "Int64": 57,
    "Uint": 100,
    "Uint8": 101,
    "Uint16": 102,
    "Uint32": 0,
    "Uint64": 103,
    "Uintptr": 104,
    "Float32": 3.14,
    "Float64": 3.15,
    "Array": null,
    "Map": {
        "bar": 2,
        "foo": 1
    },
    "String": "snowman-not-here-yet",
    "StringPointer": null,
    "Int64Pointer": null
}`))
	if err != nil {
		t.Fatalf("XUnmarshalJSON: %v", err)
	}
}
