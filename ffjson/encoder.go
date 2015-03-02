package ffjson

/**
 *  Copyright 2015 Paul Querna, Klaus Post
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

import (
	"encoding/json"
	"errors"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"io"
	"reflect"
	"sync"
)

// This is a reusable encoder.
// It allows to encode many objects to a single writer.
// This should not be used by more than one goroutine at the time.
type Encoder struct {
	buf fflib.Buffer
	w   io.Writer
	enc *json.Encoder
}

// NewEncoder returns a reusable Encoder.
// Output will be written to the supplied writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, enc: json.NewEncoder(w)}
}

// Encode the data in the supplied value to the stream
// given on creation.
func (e *Encoder) Encode(v interface{}) error {
	f, ok := v.(marshalerFaster)
	if ok {
		e.buf.Reset()
		err := f.MarshalJSONBuf(&e.buf)
		if err != nil {
			return err
		}

		_, err = io.Copy(e.w, &e.buf)
		return err
	}

	return e.enc.Encode(v)
}

// EncodeFast will unmarshal the data if fast marshall is available.
// This function can be used if you want to be sure the fast
// marshal is used or in testing.
// If you would like to have fallback to encoding/json you can use the
// regular Encode() method.
func (e *Encoder) EncodeFast(v interface{}) error {
	_, ok := v.(marshalerFaster)
	if !ok {
		return errors.New("ffjson marshal not available for type " + reflect.TypeOf(v).String())
	}
	return e.Encode(v)
}

// Error returned on async encodings.
// It contains the object and the error received to assist in debugging.
type ErrEncoderAsync struct {
	Err        error
	Object     interface{}
	ObjectType string
}

// Returns the Error as readable string.
func (e ErrEncoderAsync) Error() string {
	return "encoding object of type '" + e.ObjectType + "':" + e.Err.Error()
}

func newErrEncoderAsync(err error, v interface{}) error {
	return ErrEncoderAsync{Err: err, Object: v, ObjectType: reflect.TypeOf(v).String()}
}

// This is a reusable asynchronous encoder.
// It allows to encode many objects to a single writer.
// It will spawn a goroutine that will encode incoming
// objects.
//
// This is *safe* to use by several goroutines at one,
// although the order will of course not be predictable.
//
// If used by a single goroutine, the order written will be the order
// given to Encode.
type EncoderAsync struct {
	enc      Encoder          // The encoder used
	incoming chan interface{} // Channel of incoming objects
	flush    chan bool        // flush signal
	asyncerr error            // errors received during encoding
	mu       sync.RWMutex     // protection for 'asyncerr'
}

// NewEncoderAsync returns a reusable asynchronous encoder.
//
// You can specify the number of buffered objects.
// Output will be written to the supplied writer.
func NewEncoderAsync(w io.Writer, buffer uint) *EncoderAsync {
	e := EncoderAsync{enc: Encoder{w: w, enc: json.NewEncoder(w)}}

	e.asyncerr = nil
	e.incoming = make(chan interface{}, buffer)
	e.flush = make(chan bool, 0)

	go func() {
		var err error
		for {
			select {
			case v := <-e.incoming:
				// If we already have an error, skip this.
				if e.asyncerr == nil {
					err = e.enc.Encode(v)
					if err != nil {
						e.mu.Lock()
						e.asyncerr = newErrEncoderAsync(err, v)
						e.mu.Unlock()
					}
				}
			case <-e.flush:
				end := false
				for !end {
					select {
					case v := <-e.incoming:
						err = e.enc.Encode(v)
						if err != nil {
							e.asyncerr = newErrEncoderAsync(err, v)
							end = true
						}
					default:
						end = true
					}
				}
				e.flush <- true
				return
			}
		}
	}()

	return &e
}

// Encode the data sent.
//
// The function will return immediately.
//
// If an error has occurred from encoding other objects,
// this error will be returned, so a call might receive an error
// that occurred while encoding a previous object.
func (e *EncoderAsync) Encode(v interface{}) error {
	e.mu.RLock()
	if e.asyncerr != nil {
		err := e.asyncerr
		e.mu.RUnlock()
		return err
	}
	e.mu.RUnlock()
	e.incoming <- v
	return nil
}

// Close will flush all pending encodes and write to output.
// If an error occurred during encode, it will be returned.
// This should only be called once for an encoder objects.
func (e *EncoderAsync) Close() error {
	e.mu.RLock()
	if e.asyncerr != nil {
		err := e.asyncerr
		e.mu.RUnlock()
		return err
	}
	e.mu.RUnlock()
	// Send signal that we want to flush
	// Will block until we are in "case e.flush"
	e.flush <- true

	// Wait for flush to finish.
	<-e.flush
	err := e.asyncerr

	// Set this error, so it is returned, if the objct is re-used.
	if err == nil {
		e.asyncerr = errors.New("encoder closed")
	}
	return err
}
