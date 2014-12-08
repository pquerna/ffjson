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
	"time"
)

type FFFoo struct {
	Blah int
}

type FFRecord struct {
	Timestamp int64 `json:"id,omitempty"`
	OriginId  uint32
	Bar       FFFoo
	Method    string `json:"meth"`
	ReqId     string
	ServerIp  string
	RemoteIp  string
	BytesSent uint64
}

// fjson: skip
type Tobj struct {
	X Tint
}
type Xobj struct {
	X Xint
}

// fjson: skip
type Tduration struct {
	X time.Duration
}
type Xduration struct {
	X time.Duration
}

// fjson: skip
type Tarray struct {
	X []int
}
type Xarray struct {
	X []int
}

// fjson: skip
type TarrayPtr struct {
	X []*int
}
type XarrayPtr struct {
	X []*int
}

// fjson: skip
type Tstring struct {
	X string
}
type Xstring struct {
	X string
}

// fjson: skip
type Tbool struct {
	X bool
}
type Xbool struct {
	Tbool
}

// fjson: skip
type Tint struct {
	X int
}
type Xint struct {
	Tint
}

// fjson: skip
type Tint8 struct {
	X int8
}
type Xint8 struct {
	Tint8
}

// fjson: skip
type Tint16 struct {
	X int16
}
type Xint16 struct {
	Tint16
}

// fjson: skip
type Tint32 struct {
	X int32
}
type Xint32 struct {
	Tint32
}

// fjson: skip
type Tint64 struct {
	X int64
}
type Xint64 struct {
	Tint64
}

// fjson: skip
type Tuint struct {
	X uint
}
type Xuint struct {
	Tuint
}

// fjson: skip
type Tuint8 struct {
	X uint8
}
type Xuint8 struct {
	Tuint8
}

// fjson: skip
type Tuint16 struct {
	X uint16
}
type Xuint16 struct {
	Tuint16
}

// fjson: skip
type Tuint32 struct {
	X uint32
}
type Xuint32 struct {
	Tuint32
}

// fjson: skip
type Tuint64 struct {
	X uint64
}
type Xuint64 struct {
	Tuint64
}

// fjson: skip
type Tuintptr struct {
	X uintptr
}
type Xuintptr struct {
	Tuintptr
}

// fjson: skip
type Tfloat32 struct {
	X float32
}
type Xfloat32 struct {
	Tfloat32
}

// fjson: skip
type Tfloat64 struct {
	X float64
}
type Xfloat64 struct {
	Tfloat64
}
