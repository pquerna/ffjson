package v1

import (
	"encoding/json"
	"errors"
	"reflect"
)

type marshalerFaster interface {
	MarshalJSONBuf(buf EncodingBuffer) error
}

type unmarshalFaster interface {
	UnmarshalJSONFFLexer(l *FFLexer, state FFParseState) error
}

// Marshal will act the same way as json.Marshal, except
// it will choose the ffjson marshal function before falling
// back to using json.Marshal.
// Using this function will bypass the internal copying and parsing
// the json library normally does, which greatly speeds up encoding time.
// It is ok to call this function even if no ffjson code has been
// generated for the data type you pass in the interface.
func Marshal(v interface{}) ([]byte, error) {
	f, ok := v.(marshalerFaster)
	if ok {
		buf := Buffer{}
		err := f.MarshalJSONBuf(&buf)
		b := buf.Bytes()
		if err != nil {
			// TODO: Enable when we can pool
			//if len(b) > 0 {
			//	Pool(b)
			//}
			return nil, err
		}
		return b, nil
	}

	j, ok := v.(json.Marshaler)
	if ok {
		return j.MarshalJSON()
	}
	return json.Marshal(v)
}

// MarshalFast will unmarshal the data if fast marshall is available.
// This function can be used if you want to be sure the fast
// marshal is used or in testing.
// If you would like to have fallback to encoding/json you can use the
// Marshal() method.
func MarshalFast(v interface{}) ([]byte, error) {
	_, ok := v.(marshalerFaster)
	if !ok {
		return nil, errors.New("ffjson marshal not available for type " + reflect.TypeOf(v).String())
	}
	return Marshal(v)
}

// Unmarshal will act the same way as json.Unmarshal, except
// it will choose the ffjson unmarshal function before falling
// back to using json.Unmarshal.
// The overhead of unmarshal is lower than on Marshal,
// however this should still provide a speedup for your encoding.
// It is ok to call this function even if no ffjson code has been
// generated for the data type you pass in the interface.
func Unmarshal(data []byte, v interface{}) error {
	f, ok := v.(unmarshalFaster)
	if ok {
		fs := NewFFLexer(data)
		return f.UnmarshalJSONFFLexer(fs, FFParse_map_start)
	}

	j, ok := v.(json.Unmarshaler)
	if ok {
		return j.UnmarshalJSON(data)
	}
	return json.Unmarshal(data, v)
}

// UnmarshalFast will unmarshal the data if fast marshall is available.
// This function can be used if you want to be sure the fast
// unmarshal is used or in testing.
// If you would like to have fallback to encoding/json you can use the
// Unmarshal() method.
func UnmarshalFast(data []byte, v interface{}) error {
	_, ok := v.(unmarshalFaster)
	if !ok {
		return errors.New("ffjson unmarshal not available for type " + reflect.TypeOf(v).String())
	}
	return Unmarshal(data, v)
}
