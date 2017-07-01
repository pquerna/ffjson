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
	"errors"
	"math"
	"time"
)

// FFFoo struc... just  blah
type FFFoo struct {
	Blah int
}

// FFRecord struct
type FFRecord struct {
	Timestamp int64 `json:"id,omitempty"`
	OriginID  uint32
	Bar       FFFoo
	Method    string `json:"meth"`
	ReqID     string
	ServerIP  string
	RemoteIP  string
	BytesSent uint64
}

// TI18nName struct
// ffjson: skip
type TI18nName struct {
	Ændret   int64
	Aוההקלדה uint32
	Позната  string
}

// XI18nName struct
type XI18nName struct {
	Ændret   int64
	Aוההקלדה uint32
	Позната  string
}

type mystring string

// TsortName struct
// ffjson: skip
type TsortName struct {
	C string
	B int `json:"A"`
}

// XsortName struct
type XsortName struct {
	C string
	B int `json:"A"`
}

// Tobj struct
// ffjson: skip
type Tobj struct {
	X Tint
}

// Xobj struct
type Xobj struct {
	X Xint
}

// Tduration struct
// ffjson: skip
type Tduration struct {
	X time.Duration
}

// Xduration struct
type Xduration struct {
	X time.Duration
}

// TtimePtr struct
// ffjson: skip
type TtimePtr struct {
	X *time.Time
}

// XtimePtr struct
type XtimePtr struct {
	X *time.Time
}

// Tarray struct
// ffjson: skip
type Tarray struct {
	X [3]int
}

// Xarray struct
type Xarray struct {
	X [3]int
}

// TarrayPtr struct
// ffjson: skip
type TarrayPtr struct {
	X [3]*int
}

// XarrayPtr struct
type XarrayPtr struct {
	X [3]*int
}

// Tslice struct
// ffjson: skip
type Tslice struct {
	X []int
}

//Xslice struct
type Xslice struct {
	X []int
}

// TslicePtr struct
// ffjson: skip
type TslicePtr struct {
	X []*int
}

// XslicePtr struct
type XslicePtr struct {
	X []*int
}

// Tstring struct
// ffjson: skip
type Tstring struct {
	X string
}

// Xstring struct
type Xstring struct {
	X string
}

// Tmystring struct
// ffjson: skip
type Tmystring struct {
	X mystring
}

// Xmystring struct
type Xmystring struct {
	X mystring
}

// TmystringPtr struct
// ffjson: skip
type TmystringPtr struct {
	X *mystring
}

// XmystringPtr struct
type XmystringPtr struct {
	X *mystring
}

// TstringTagged struct
// ffjson: skip
type TstringTagged struct {
	X string `json:",string"`
}

// XstringTagged struct
type XstringTagged struct {
	X string `json:",string"`
}

// TstringTagedPtr struct
// ffjson: skip
type TstringTaggedPtr struct {
	X *string `json:",string"`
}

// XstringTaggedPtr struct
type XstringTaggedPtr struct {
	X *string `json:",string"`
}

// TintTagged struct
// ffjson: skip
type TintTagged struct {
	X int `json:",string"`
}

//XintTagged struct
type XintTagged struct {
	X int `json:",string"`
}

// TboolTagged struct
// ffjson: skip
type TboolTagged struct {
	X int `json:",string"`
}

// XboolTagged struct
type XboolTagged struct {
	X int `json:",string"`
}

// TMapStringString struct
// ffjson: skip
type TMapStringString struct {
	X map[string]string
}

// XMapStringString struct
type XMapStringString struct {
	X map[string]string
}

// Tbool struct
// ffjson: skip
type Tbool struct {
	X bool
}

// Xbool struct
type Xbool struct {
	X bool
}

// Tint struct
// ffjson: skip
type Tint struct {
	X int
}

// Xint struct
type Xint struct {
	X int
}

// Tbyte struct
// ffjson: skip
type Tbyte struct {
	X byte
}

// Xbyte struct
type Xbyte struct {
	X byte
}

// Tint8 struct
// ffjson: skip
type Tint8 struct {
	X int8
}

// Xint8 struct
type Xint8 struct {
	X int8
}

// Tint16 struct
// ffjson: skip
type Tint16 struct {
	X int16
}

// Xint16 stuct
type Xint16 struct {
	X int16
}

// Tint32 struct
// ffjson: skip
type Tint32 struct {
	X int32
}

// Xint32 struct
type Xint32 struct {
	X int32
}

// Tint64 struct
// ffjson: skip
type Tint64 struct {
	X int64
}

// Xint64 struct
type Xint64 struct {
	X int64
}

// Tuiny struct
// ffjson: skip
type Tuint struct {
	X uint
}

// Xuint struct
type Xuint struct {
	X uint
}

// Tuint8 struct
// ffjson: skip
type Tuint8 struct {
	X uint8
}

// Xuint8 struct
type Xuint8 struct {
	X uint8
}

// Tuint16 struct
// ffjson: skip
type Tuint16 struct {
	X uint16
}

// Xuint16 struct
type Xuint16 struct {
	X uint16
}

// Tuint32 struct
// ffjson: skip
type Tuint32 struct {
	X uint32
}

// Xuint32 struct
type Xuint32 struct {
	X uint32
}

// Tuint64 struct
// ffjson: skip
type Tuint64 struct {
	X uint64
}

// Xuint64 struct
type Xuint64 struct {
	X uint64
}

// Tuintptr struct
// ffjson: skip
type Tuintptr struct {
	X uintptr
}

// Xuintptr struct
type Xuintptr struct {
	X uintptr
}

// Tfloat32 struct
// ffjson: skip
type Tfloat32 struct {
	X float32
}

// Xfloat32 struct
type Xfloat32 struct {
	X float32
}

// Tfloat64 struct
// ffjson: skip
type Tfloat64 struct {
	X float64
}

// Xfloat64 struct
type Xfloat64 struct {
	X float64
}

// ATduration struct
// Arrays
// ffjson: skip
type ATduration struct {
	X [3]time.Duration
}

// AXduration struct
type AXduration struct {
	X [3]time.Duration
}

// ATbool struct
// ffjson: skip
type ATbool struct {
	X [3]bool
}

// AXbool struct
type AXbool struct {
	X [3]bool
}

// ATint struct
// ffjson: skip
type ATint struct {
	X [3]int
}

// AXint struct
type AXint struct {
	X [3]int
}

// ATbyte struct
// ffjson: skip
type ATbyte struct {
	X [3]byte
}

// AXbyte struct
type AXbyte struct {
	X [3]byte
}

// ATint8
// ffjson: skip
type ATint8 struct {
	X [3]int8
}

// AXint8 struct
type AXint8 struct {
	X [3]int8
}

// ATint16 struct
// ffjson: skip
type ATint16 struct {
	X [3]int16
}

// AXint16 struct
type AXint16 struct {
	X [3]int16
}

// ATint32 struct
// ffjson: skip
type ATint32 struct {
	X [3]int32
}

// AXint32 struct
type AXint32 struct {
	X [3]int32
}

// ATint64 struct
// ffjson: skip
type ATint64 struct {
	X [3]int64
}

// AXint64 struct
type AXint64 struct {
	X [3]int64
}

// ATuint struct
// ffjson: skip
type ATuint struct {
	X [3]uint
}

// AXuint struct
type AXuint struct {
	X [3]uint
}

// ATuint8 struct
// ffjson: skip
type ATuint8 struct {
	X [3]uint8
}

// AXuint8 struct
type AXuint8 struct {
	X [3]uint8
}

// ATuint16 struct
// ffjson: skip
type ATuint16 struct {
	X [3]uint16
}

// AXuint16 struct
type AXuint16 struct {
	X [3]uint16
}

// ATuint32 struct
// ffjson: skip
type ATuint32 struct {
	X [3]uint32
}

// AXuint32 struct
type AXuint32 struct {
	X [3]uint32
}

// ATuint64 struct
// ffjson: skip
type ATuint64 struct {
	X [3]uint64
}

// AXuint64 struct
type AXuint64 struct {
	X [3]uint64
}

// ATuintptr
// ffjson: skip
type ATuintptr struct {
	X [3]uintptr
}

// AXuintptr
type AXuintptr struct {
	X [3]uintptr
}

// ATfloat32 struct
// ffjson: skip
type ATfloat32 struct {
	X [3]float32
}

// AXfloat32 struct
type AXfloat32 struct {
	X [3]float32
}

// ATfloat54 struct
// ffjson: skip
type ATfloat64 struct {
	X [3]float64
}

// AXfloat64 struct
type AXfloat64 struct {
	X [3]float64
}

// ATtime struct
// ffjson: skip
type ATtime struct {
	X [3]time.Time
}

// AXtime struct
type AXtime struct {
	X [3]time.Time
}

// STduration struct
// Slices
// ffjson: skip
type STduration struct {
	X []time.Duration
}

// SXduration struct
type SXduration struct {
	X []time.Duration
}

// STbool struct
// ffjson: skip
type STbool struct {
	X []bool
}

// SXbool struct
type SXbool struct {
	X []bool
}

// STint struct
// ffjson: skip
type STint struct {
	X []int
}

// SXint struct
type SXint struct {
	X []int
}

// STbyte struct
// ffjson: skip
type STbyte struct {
	X []byte
}

// SXbyte struct
type SXbyte struct {
	X []byte
}

// STint8 struct
// ffjson: skip
type STint8 struct {
	X []int8
}

// SXint8 struct
type SXint8 struct {
	X []int8
}

// STint16 struct
// ffjson: skip
type STint16 struct {
	X []int16
}

// SXint16 struct
type SXint16 struct {
	X []int16
}

// STint32 struct
// ffjson: skip
type STint32 struct {
	X []int32
}

// SXint32 struct
type SXint32 struct {
	X []int32
}

// STint64 struct
// ffjson: skip
type STint64 struct {
	X []int64
}

// SXint64 struct
type SXint64 struct {
	X []int64
}

// STuint struct
// ffjson: skip
type STuint struct {
	X []uint
}

// SXuint struct
type SXuint struct {
	X []uint
}

// STuint8 struct
// ffjson: skip
type STuint8 struct {
	X []uint8
}

// SXuint8 struct
type SXuint8 struct {
	X []uint8
}

// STuint16 struct
// ffjson: skip
type STuint16 struct {
	X []uint16
}

// SXuint16 struct
type SXuint16 struct {
	X []uint16
}

// STuint32 struct
// ffjson: skip
type STuint32 struct {
	X []uint32
}

// SXuint32 struct
type SXuint32 struct {
	X []uint32
}

// STuint64 struct
// ffjson: skip
type STuint64 struct {
	X []uint64
}

// SXuint64 struct
type SXuint64 struct {
	X []uint64
}

// STuintptr struct
// ffjson: skip
type STuintptr struct {
	X []uintptr
}

// SXuintptr struct
type SXuintptr struct {
	X []uintptr
}

// STfloat32 struct
// ffjson: skip
type STfloat32 struct {
	X []float32
}

// SXfloat32 struct
type SXfloat32 struct {
	X []float32
}

// STfloat64 struct
// ffjson: skip
type STfloat64 struct {
	X []float64
}

// SXfloat64 struct
type SXfloat64 struct {
	X []float64
}

// STtime struct
// ffjson: skip
type STtime struct {
	X []time.Time
}

// SXtime struct
type SXtime struct {
	X []time.Time
}

// TMapStringMapString struct
// Nested
// ffjson: skip
type TMapStringMapString struct {
	X map[string]map[string]string
}

// XMapStringMapString struct
type XMapStringMapString struct {
	X map[string]map[string]string
}

// TMapStringAString struct
// ffjson: skip
type TMapStringAString struct {
	X map[string][3]string
}

// XMapStringAString struct
type XMapStringAString struct {
	X map[string][3]string
}

// TSAAtring struct
// ffjson: skip
type TSAAtring struct {
	X [2][3]string
}

// XSAAtring struct
type XSAAtring struct {
	X [2][3]string
}

// TSAString struct
// ffjson: skip
type TSAString struct {
	X [][3]string
}

// XSAString struct
type XSAString struct {
	X [][3]string
}

// Optionals tests from golang test suite
type Optionals struct {
	Sr string `json:"sr"`
	So string `json:"so,omitempty"`
	Sw string `json:"-"`

	Ir int `json:"omitempty"` // actually named omitempty, not an option
	Io int `json:"io,omitempty"`

	Slr []string `json:"slr,random"`
	Slo []string `json:"slo,omitempty"`

	Mr map[string]interface{} `json:"mr"`
	Mo map[string]interface{} `json:",omitempty"`

	Fr float64 `json:"fr"`
	Fo float64 `json:"fo,omitempty"`

	Br bool `json:"br"`
	Bo bool `json:"bo,omitempty"`

	Ur uint `json:"ur"`
	Uo uint `json:"uo,omitempty"`

	Str struct{} `json:"str"`
	Sto struct{} `json:"sto,omitempty"`
}

var unsupportedValues = []interface{}{
	math.NaN(),
	math.Inf(-1),
	math.Inf(1),
}

var optionalsExpected = `{
 "sr": "",
 "omitempty": 0,
 "slr": null,
 "mr": {},
 "fr": 0,
 "br": false,
 "ur": 0,
 "str": {},
 "sto": {}
}`

// StringTag struct
type StringTag struct {
	BoolStr bool    `json:",string"`
	IntStr  int64   `json:",string"`
	FltStr  float64 `json:",string"`
	StrStr  string  `json:",string"`
}

var stringTagExpected = `{
 "BoolStr": "true",
 "IntStr": "42",
 "FltStr": "0",
 "StrStr": "\"xzbit\""
}`

// OmitAll struct
type OmitAll struct {
	Ostr    string                 `json:",omitempty"`
	Oint    int                    `json:",omitempty"`
	Obool   bool                   `json:",omitempty"`
	Odouble float64                `json:",omitempty"`
	Ointer  interface{}            `json:",omitempty"`
	Omap    map[string]interface{} `json:",omitempty"`
	OstrP   *string                `json:",omitempty"`
	OintP   *int                   `json:",omitempty"`
	// TODO: Re-enable when issue #55 is fixed.
	OboolP  *bool                   `json:",omitempty"`
	OmapP   *map[string]interface{} `json:",omitempty"`
	Astr    []string                `json:",omitempty"`
	Aint    []int                   `json:",omitempty"`
	Abool   []bool                  `json:",omitempty"`
	Adouble []float64               `json:",omitempty"`
}

var omitAllExpected = `{}`

// NoExported struct
type NoExported struct {
	field1 string
	field2 string
	field3 string
}

var noExportedExpected = `{}`

// OmitFirst struct
type OmitFirst struct {
	Ostr string `json:",omitempty"`
	Str  string
}

var omitFirstExpected = `{
 "Str": ""
}`

// OmitLast struct
type OmitLast struct {
	Xstr string `json:",omitempty"`
	Str  string
}

var omitLastExpected = `{
 "Str": ""
}`

// byte slices are special even if they're renamed types.
type renamedByte byte
type renamedByteSlice []byte
type renamedRenamedByteSlice []renamedByte

// ByteSliceNormal struct
type ByteSliceNormal struct {
	X []byte
}

// ByteSliceRenamed stuct
type ByteSliceRenamed struct {
	X renamedByteSlice
}

// ByteSliceDoubleRenamned struct
type ByteSliceDoubleRenamed struct {
	X renamedRenamedByteSlice
}

// Ref has Marshaler and Unmarshaler methods with pointer receiver.
type Ref int

// MarshalJSON func
func (*Ref) MarshalJSON() ([]byte, error) {
	return []byte(`"ref"`), nil
}

// UnmarshalJSON func
func (r *Ref) UnmarshalJSON([]byte) error {
	*r = 12
	return nil
}

// Val has Marshaler methods with value receiver.
type Val int

// MarshalJSON var
func (Val) MarshalJSON() ([]byte, error) {
	return []byte(`"val"`), nil
}

// RefText has Marshaler and Unmarshaler methods with pointer receiver.
type RefText int

// MarshalText func
func (*RefText) MarshalText() ([]byte, error) {
	return []byte(`"ref"`), nil
}

// UnmarshalText func
func (r *RefText) UnmarshalText([]byte) error {
	*r = 13
	return nil
}

// ValText has Marshaler methods with value receiver.
type ValText int

// MarshalText val
func (ValText) MarshalText() ([]byte, error) {
	return []byte(`"val"`), nil
}

// C implements Marshaler and returns unescaped JSON.
type C int

// MarshalJSON func
func (C) MarshalJSON() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

// CText implements Marshaler and returns unescaped text.
type CText int

// MarshalText func
func (CText) MarshalText() ([]byte, error) {
	return []byte(`"<&>"`), nil
}

// ErrGiveError generates error
var ErrGiveError = errors.New("GiveError error")

// GiveError always returns an ErrGiveError on Marshal/Unmarshal.
type GiveError struct{}

// MarshalJSON func
func (r GiveError) MarshalJSON() ([]byte, error) {
	return nil, ErrGiveError
}

// UnmarshalJSON func
func (r *GiveError) UnmarshalJSON([]byte) error {
	return ErrGiveError
}

// IntType
type IntType int

// MyStruct struc
type MyStruct struct {
	IntType
}

// BugA struct
type BugA struct {
	S string
}

// BugB struct
type BugB struct {
	BugA
	S string
}

type BugC struct {
	S string
}

// Legal Go: We never use the repeated embedded field (S).
type BugX struct {
	A int
	BugA
	BugB
}

type BugD struct { // Same as BugA after tagging.
	XXX string `json:"S"`
}

// BugD's tagged S field should dominate BugA's.
type BugY struct {
	BugA
	BugD
}

// There are no tags here, so S should not appear.
type BugZ struct {
	BugA
	BugC
	BugY // Contains a tagged S field through BugD; should not dominate.
}

type FfFuzz struct {
	A uint8
	B uint16
	C uint32
	D uint64

	E int8
	F int16
	G int32
	H int64

	I float32
	J float64

	M byte
	N rune

	O int
	P uint
	Q string
	R bool
	S time.Time

	Ap *uint8
	Bp *uint16
	Cp *uint32
	Dp *uint64

	Ep *int8
	Fp *int16
	Gp *int32
	Hp *int64

	Ip *float32
	Jp *float64

	Mp *byte
	Np *rune

	Op *int
	Pp *uint
	Qp *string
	Rp *bool
	Sp *time.Time

	Aa []uint8
	Ba []uint16
	Ca []uint32
	Da []uint64

	Ea []int8
	Fa []int16
	Ga []int32
	Ha []int64

	Ia []float32
	Ja []float64

	Ma []byte
	Na []rune

	Oa []int
	Pa []uint
	Qa []string
	Ra []bool

	Aap []*uint8
	Bap []*uint16
	Cap []*uint32
	Dap []*uint64

	Eap []*int8
	Fap []*int16
	Gap []*int32
	Hap []*int64

	Iap []*float32
	Jap []*float64

	Map []*byte
	Nap []*rune

	Oap []*int
	Pap []*uint
	Qap []*string
	Rap []*bool
}

// ffjson: skip
type FuzzOmitEmpty struct {
	A uint8  `json:",omitempty"`
	B uint16 `json:",omitempty"`
	C uint32 `json:",omitempty"`
	D uint64 `json:",omitempty"`

	E int8  `json:",omitempty"`
	F int16 `json:",omitempty"`
	G int32 `json:",omitempty"`
	H int64 `json:",omitempty"`

	I float32 `json:",omitempty"`
	J float64 `json:",omitempty"`

	M byte `json:",omitempty"`
	N rune `json:",omitempty"`

	O int       `json:",omitempty"`
	P uint      `json:",omitempty"`
	Q string    `json:",omitempty"`
	R bool      `json:",omitempty"`
	S time.Time `json:",omitempty"`

	Ap *uint8  `json:",omitempty"`
	Bp *uint16 `json:",omitempty"`
	Cp *uint32 `json:",omitempty"`
	Dp *uint64 `json:",omitempty"`

	Ep *int8  `json:",omitempty"`
	Fp *int16 `json:",omitempty"`
	Gp *int32 `json:",omitempty"`
	Hp *int64 `json:",omitempty"`

	Ip *float32 `json:",omitempty"`
	Jp *float64 `json:",omitempty"`

	Mp *byte `json:",omitempty"`
	Np *rune `json:",omitempty"`

	Op *int       `json:",omitempty"`
	Pp *uint      `json:",omitempty"`
	Qp *string    `json:",omitempty"`
	Rp *bool      `json:",omitempty"`
	Sp *time.Time `json:",omitempty"`

	Aa []uint8  `json:",omitempty"`
	Ba []uint16 `json:",omitempty"`
	Ca []uint32 `json:",omitempty"`
	Da []uint64 `json:",omitempty"`

	Ea []int8  `json:",omitempty"`
	Fa []int16 `json:",omitempty"`
	Ga []int32 `json:",omitempty"`
	Ha []int64 `json:",omitempty"`

	Ia []float32 `json:",omitempty"`
	Ja []float64 `json:",omitempty"`

	Ma []byte `json:",omitempty"`
	Na []rune `json:",omitempty"`

	Oa []int    `json:",omitempty"`
	Pa []uint   `json:",omitempty"`
	Qa []string `json:",omitempty"`
	Ra []bool   `json:",omitempty"`

	Aap []*uint8  `json:",omitempty"`
	Bap []*uint16 `json:",omitempty"`
	Cap []*uint32 `json:",omitempty"`
	Dap []*uint64 `json:",omitempty"`

	Eap []*int8  `json:",omitempty"`
	Fap []*int16 `json:",omitempty"`
	Gap []*int32 `json:",omitempty"`
	Hap []*int64 `json:",omitempty"`

	Iap []*float32 `json:",omitempty"`
	Jap []*float64 `json:",omitempty"`

	Map []*byte `json:",omitempty"`
	Nap []*rune `json:",omitempty"`

	Oap []*int    `json:",omitempty"`
	Pap []*uint   `json:",omitempty"`
	Qap []*string `json:",omitempty"`
	Rap []*bool   `json:",omitempty"`
}

type FfFuzzOmitEmpty struct {
	A uint8  `json:",omitempty"`
	B uint16 `json:",omitempty"`
	C uint32 `json:",omitempty"`
	D uint64 `json:",omitempty"`

	E int8  `json:",omitempty"`
	F int16 `json:",omitempty"`
	G int32 `json:",omitempty"`
	H int64 `json:",omitempty"`

	I float32 `json:",omitempty"`
	J float64 `json:",omitempty"`

	M byte `json:",omitempty"`
	N rune `json:",omitempty"`

	O int       `json:",omitempty"`
	P uint      `json:",omitempty"`
	Q string    `json:",omitempty"`
	R bool      `json:",omitempty"`
	S time.Time `json:",omitempty"`

	Ap *uint8  `json:",omitempty"`
	Bp *uint16 `json:",omitempty"`
	Cp *uint32 `json:",omitempty"`
	Dp *uint64 `json:",omitempty"`

	Ep *int8  `json:",omitempty"`
	Fp *int16 `json:",omitempty"`
	Gp *int32 `json:",omitempty"`
	Hp *int64 `json:",omitempty"`

	Ip *float32 `json:",omitempty"`
	Jp *float64 `json:",omitempty"`

	Mp *byte `json:",omitempty"`
	Np *rune `json:",omitempty"`

	Op *int       `json:",omitempty"`
	Pp *uint      `json:",omitempty"`
	Qp *string    `json:",omitempty"`
	Rp *bool      `json:",omitempty"`
	Sp *time.Time `json:",omitempty"`

	Aa []uint8  `json:",omitempty"`
	Ba []uint16 `json:",omitempty"`
	Ca []uint32 `json:",omitempty"`
	Da []uint64 `json:",omitempty"`

	Ea []int8  `json:",omitempty"`
	Fa []int16 `json:",omitempty"`
	Ga []int32 `json:",omitempty"`
	Ha []int64 `json:",omitempty"`

	Ia []float32 `json:",omitempty"`
	Ja []float64 `json:",omitempty"`

	Ma []byte `json:",omitempty"`
	Na []rune `json:",omitempty"`

	Oa []int    `json:",omitempty"`
	Pa []uint   `json:",omitempty"`
	Qa []string `json:",omitempty"`
	Ra []bool   `json:",omitempty"`

	Aap []*uint8  `json:",omitempty"`
	Bap []*uint16 `json:",omitempty"`
	Cap []*uint32 `json:",omitempty"`
	Dap []*uint64 `json:",omitempty"`

	Eap []*int8  `json:",omitempty"`
	Fap []*int16 `json:",omitempty"`
	Gap []*int32 `json:",omitempty"`
	Hap []*int64 `json:",omitempty"`

	Iap []*float32 `json:",omitempty"`
	Jap []*float64 `json:",omitempty"`

	Map []*byte `json:",omitempty"`
	Nap []*rune `json:",omitempty"`

	Oap []*int    `json:",omitempty"`
	Pap []*uint   `json:",omitempty"`
	Qap []*string `json:",omitempty"`
	Rap []*bool   `json:",omitempty"`
}

// ffjson: skip
type FuzzString struct {
	A uint8  `json:",string"`
	B uint16 `json:",string"`
	C uint32 `json:",string"`
	D uint64 `json:",string"`

	E int8  `json:",string"`
	F int16 `json:",string"`
	G int32 `json:",string"`
	H int64 `json:",string"`

	I float32 `json:",string"`
	J float64 `json:",string"`

	M byte `json:",string"`
	N rune `json:",string"`

	O int  `json:",string"`
	P uint `json:",string"`

	Q string `json:",string"`

	R bool `json:",string"`
	// https://github.com/golang/go/issues/9812
	// S time.Time `json:",string"`

	Ap *uint8  `json:",string"`
	Bp *uint16 `json:",string"`
	Cp *uint32 `json:",string"`
	Dp *uint64 `json:",string"`

	Ep *int8  `json:",string"`
	Fp *int16 `json:",string"`
	Gp *int32 `json:",string"`
	Hp *int64 `json:",string"`

	Ip *float32 `json:",string"`
	Jp *float64 `json:",string"`

	Mp *byte `json:",string"`
	Np *rune `json:",string"`

	Op *int    `json:",string"`
	Pp *uint   `json:",string"`
	Qp *string `json:",string"`
	Rp *bool   `json:",string"`
	// https://github.com/golang/go/issues/9812
	// Sp *time.Time `json:",string"`
}

type FfFuzzString struct {
	A uint8  `json:",string"`
	B uint16 `json:",string"`
	C uint32 `json:",string"`
	D uint64 `json:",string"`

	E int8  `json:",string"`
	F int16 `json:",string"`
	G int32 `json:",string"`
	H int64 `json:",string"`

	I float32 `json:",string"`
	J float64 `json:",string"`

	M byte `json:",string"`
	N rune `json:",string"`

	O int  `json:",string"`
	P uint `json:",string"`

	Q string `json:",string"`

	R bool `json:",string"`
	// https://github.com/golang/go/issues/9812
	// S time.Time `json:",string"`

	Ap *uint8  `json:",string"`
	Bp *uint16 `json:",string"`
	Cp *uint32 `json:",string"`
	Dp *uint64 `json:",string"`

	Ep *int8  `json:",string"`
	Fp *int16 `json:",string"`
	Gp *int32 `json:",string"`
	Hp *int64 `json:",string"`

	Ip *float32 `json:",string"`
	Jp *float64 `json:",string"`

	Mp *byte `json:",string"`
	Np *rune `json:",string"`

	Op *int    `json:",string"`
	Pp *uint   `json:",string"`
	Qp *string `json:",string"`
	Rp *bool   `json:",string"`
	// https://github.com/golang/go/issues/9812
	// Sp *time.Time `json:",string"`
}

// ffjson: skip
type TTestMaps struct {
	Aa map[string]uint8
	Ba map[string]uint16
	Ca map[string]uint32
	Da map[string]uint64

	Ea map[string]int8
	Fa map[string]int16
	Ga map[string]int32
	Ha map[string]int64

	Ia map[string]float32
	Ja map[string]float64

	Ma map[string]byte
	Na map[string]rune

	Oa map[string]int
	Pa map[string]uint
	Qa map[string]string
	Ra map[string]bool

	AaP map[string]*uint8
	BaP map[string]*uint16
	CaP map[string]*uint32
	DaP map[string]*uint64

	EaP map[string]*int8
	FaP map[string]*int16
	GaP map[string]*int32
	HaP map[string]*int64

	IaP map[string]*float32
	JaP map[string]*float64

	MaP map[string]*byte
	NaP map[string]*rune

	OaP map[string]*int
	PaP map[string]*uint
	QaP map[string]*string
	RaP map[string]*bool
}

type XTestMaps struct {
	TTestMaps
}

// ffjson: noencoder
type NoEncoder struct {
	C string
	B int `json:"A"`
}

// ffjson: nodecoder
type NoDecoder struct {
	C string
	B int `json:"A"`
}

// ffjson: skip
type TEmbeddedStructures struct {
	X []interface{}
	Y struct {
		X int
	}
	Z []struct {
		X int
	}
	U map[string]struct {
		X int
	}
	V []map[string]struct {
		X int
	}
	W [5]map[string]struct {
		X int
	}
	Q [][]string
}

type XEmbeddedStructures struct {
	X []interface{}
	Y struct {
		X int
	}
	Z []struct {
		X int
	}
	U map[string]struct {
		X int
	}
	V []map[string]struct {
		X int
	}
	W [5]map[string]struct {
		X int
	}
	Q [][]string
}

// ffjson: skip
// Side-effect of this test is also to verify that Encoder/Decoder skipping works.
type TRenameTypes struct {
	X struct {
		X int
	} `json:"X-renamed"`
	Y NoEncoder  `json:"Y-renamed"`
	Z string     `json:"Z-renamed"`
	U *NoDecoder `json:"U-renamed"`
}

type XRenameTypes struct {
	X struct {
		X int
	} `json:"X-renamed"`
	Y NoEncoder  `json:"Y-renamed"`
	Z string     `json:"Z-renamed"`
	U *NoDecoder `json:"U-renamed"`
}

type ReTypedA uint8
type ReTypedB uint16
type ReTypedC uint32
type ReTypedD uint64

type ReTypedE int8
type ReTypedF int16
type ReTypedG int32
type ReTypedH int64

type ReTypedI float32
type ReTypedJ float64

type ReTypedM byte
type ReTypedN rune

type ReTypedO int
type ReTypedP uint
type ReTypedQ string
type ReTypedR bool
type ReTypedS time.Time

type ReTypedAp *uint8
type ReTypedBp *uint16
type ReTypedCp *uint32
type ReTypedDp *uint64

type ReTypedEp *int8
type ReTypedFp *int16
type ReTypedGp *int32
type ReTypedHp *int64

type ReTypedIp *float32
type ReTypedJp *float64

type ReTypedMp *byte
type ReTypedNp *rune

type ReTypedOp *int
type ReTypedPp *uint
type ReTypedQp *string
type ReTypedRp *bool
type ReTypedSp *time.Time

type ReTypedAa []uint8
type ReTypedBa []uint16
type ReTypedCa []uint32
type ReTypedDa []uint64

type ReTypedEa []int8
type ReTypedFa []int16
type ReTypedGa []int32
type ReTypedHa []int64

type ReTypedIa []float32
type ReTypedJa []float64

type ReTypedMa []byte
type ReTypedNa []rune

type ReTypedOa []int
type ReTypedPa []uint
type ReTypedQa []string
type ReTypedRa []bool

type ReTypedAap []*uint8
type ReTypedBap []*uint16
type ReTypedCap []*uint32
type ReTypedDap []*uint64

type ReTypedEap []*int8
type ReTypedFap []*int16
type ReTypedGap []*int32
type ReTypedHap []*int64

type ReTypedIap []*float32
type ReTypedJap []*float64

type ReTypedMap []*byte
type ReTypedNap []*rune

type ReTypedOap []*int
type ReTypedPap []*uint
type ReTypedQap []*string
type ReTypedRap []*bool
type ReTypedXa NoDecoder
type ReTypedXb NoEncoder
type ReTypedXc *NoDecoder
type ReTypedXd *NoEncoder

type ReReTypedA ReTypedA
type ReReTypedS ReTypedS
type ReReTypedAp ReTypedAp
type ReReTypedSp ReTypedSp
type ReReTypedAa ReTypedAa
type ReReTypedAap ReTypedAap
type ReReTypedXa ReTypedXa
type ReReTypedXb ReTypedXb
type ReReTypedXc ReTypedXc
type ReReTypedXd ReTypedXd

type RePReTypedA *ReTypedA
type ReSReTypedS []ReTypedS
type ReAReTypedAp [4]ReTypedAp

// ffjson: ignore
type TReTyped struct {
	A ReTypedA
	B ReTypedB
	C ReTypedC
	D ReTypedD

	E ReTypedE
	F ReTypedF
	G ReTypedG
	H ReTypedH

	I ReTypedI
	J ReTypedJ

	M ReTypedM
	N ReTypedN

	O ReTypedO
	P ReTypedP
	Q ReTypedQ
	R ReTypedR
	S ReTypedS

	Ap ReTypedAp
	Bp ReTypedBp
	Cp ReTypedCp
	Dp ReTypedDp

	Ep ReTypedEp
	Fp ReTypedFp
	Gp ReTypedGp
	Hp ReTypedHp

	Ip ReTypedIp
	Jp ReTypedJp

	Mp ReTypedMp
	Np ReTypedNp

	Op ReTypedOp
	Pp ReTypedPp
	Qp ReTypedQp
	Rp ReTypedRp
	// FIXME: https://github.com/pquerna/ffjson/issues/108
	//Sp ReTypedSp

	// Bug in encoding/json: Bug in encoding/json: json: cannot unmarshal string into Go value of type tff.ReTypedAa
	//Aa ReTypedAa
	Ba ReTypedBa
	Ca ReTypedCa
	Da ReTypedDa

	Ea ReTypedEa
	Fa ReTypedFa
	Ga ReTypedGa
	Ha ReTypedHa

	Ia ReTypedIa
	Ja ReTypedJa

	// Bug in encoding/json: json: cannot unmarshal string into Go value of type tff.ReTypedMa
	// Ma ReTypedMa
	Na ReTypedNa

	Oa ReTypedOa
	Pa ReTypedPa
	Qa ReTypedQa
	Ra ReTypedRa

	Aap ReTypedAap
	Bap ReTypedBap
	Cap ReTypedCap
	Dap ReTypedDap

	Eap ReTypedEap
	Fap ReTypedFap
	Gap ReTypedGap
	Hap ReTypedHap

	Iap ReTypedIap
	Jap ReTypedJap

	Map ReTypedMap
	Nap ReTypedNap

	Oap ReTypedOap
	Pap ReTypedPap
	Qap ReTypedQap
	Rap ReTypedRap

	Xa ReTypedXa
	Xb ReTypedXb

	Rra  ReReTypedA
	Rrs  ReReTypedS
	Rrap ReReTypedAp
	// FIXME: https://github.com/pquerna/ffjson/issues/108
	// Rrsp  ReReTypedSp
	// Rrxc  ReReTypedXc
	// Rrxd  ReReTypedXd

	// Bug in encoding/json: json: json: cannot unmarshal string into Go value of type tff.ReReTypedAa
	// Rraa  ReReTypedAa
	Rraap ReReTypedAap
	Rrxa  ReReTypedXa
	Rrxb  ReReTypedXb

	Rpra RePReTypedA
	Rsrs ReSReTypedS
}

type XReTyped struct {
	A ReTypedA
	B ReTypedB
	C ReTypedC
	D ReTypedD

	E ReTypedE
	F ReTypedF
	G ReTypedG
	H ReTypedH

	I ReTypedI
	J ReTypedJ

	M ReTypedM
	N ReTypedN

	O ReTypedO
	P ReTypedP
	Q ReTypedQ
	R ReTypedR
	S ReTypedS

	Ap ReTypedAp
	Bp ReTypedBp
	Cp ReTypedCp
	Dp ReTypedDp

	Ep ReTypedEp
	Fp ReTypedFp
	Gp ReTypedGp
	Hp ReTypedHp

	Ip ReTypedIp
	Jp ReTypedJp

	Mp ReTypedMp
	Np ReTypedNp

	Op ReTypedOp
	Pp ReTypedPp
	Qp ReTypedQp
	Rp ReTypedRp
	// FIXME: https://github.com/pquerna/ffjson/issues/108
	//Sp ReTypedSp

	// Bug in encoding/json: Bug in encoding/json: json: cannot unmarshal string into Go value of type tff.ReTypedAa
	// Aa ReTypedAa
	Ba ReTypedBa
	Ca ReTypedCa
	Da ReTypedDa

	Ea ReTypedEa
	Fa ReTypedFa
	Ga ReTypedGa
	Ha ReTypedHa

	Ia ReTypedIa
	Ja ReTypedJa

	// Bug in encoding/json: Bug in encoding/json: json: cannot unmarshal string into Go value of type tff.ReTypedMa
	//Ma ReTypedMa
	Na ReTypedNa

	Oa ReTypedOa
	Pa ReTypedPa
	Qa ReTypedQa
	Ra ReTypedRa

	Aap ReTypedAap
	Bap ReTypedBap
	Cap ReTypedCap
	Dap ReTypedDap

	Eap ReTypedEap
	Fap ReTypedFap
	Gap ReTypedGap
	Hap ReTypedHap

	Iap ReTypedIap
	Jap ReTypedJap

	Map ReTypedMap
	Nap ReTypedNap

	Oap ReTypedOap
	Pap ReTypedPap
	Qap ReTypedQap
	Rap ReTypedRap

	Xa ReTypedXa
	Xb ReTypedXb

	Rra  ReReTypedA
	Rrs  ReReTypedS
	Rrap ReReTypedAp
	// FIXME: https://github.com/pquerna/ffjson/issues/108
	// Rrsp  ReReTypedSp
	// Rrxc  ReReTypedXc
	// Rrxd  ReReTypedXd

	// Bug in encoding/json: json: json: cannot unmarshal string into Go value of type tff.ReReTypedAa
	// Rraa  ReReTypedAa
	Rraap ReReTypedAap
	Rrxa  ReReTypedXa
	Rrxb  ReReTypedXb

	Rpra RePReTypedA
	Rsrs ReSReTypedS
}

// ffjson: skip
type TInlineStructs struct {
	B struct {
		A uint8
		B uint16
		C uint32
		D uint64

		E int8
		F int16
		G int32
		H int64

		I float32
		J float64

		M byte
		N rune

		O int
		P uint
		Q string
		R bool
		S time.Time

		Ap *uint8
		Bp *uint16
		Cp *uint32
		Dp *uint64

		Ep *int8
		Fp *int16
		Gp *int32
		Hp *int64

		Ip *float32
		Jp *float64

		Mp *byte
		Np *rune

		Op *int
		Pp *uint
		Qp *string
		Rp *bool
		Sp *time.Time

		Aa []uint8
		Ba []uint16
		Ca []uint32
		Da []uint64

		Ea []int8
		Fa []int16
		Ga []int32
		Ha []int64

		Ia []float32
		Ja []float64

		Ma []byte
		Na []rune

		Oa []int
		Pa []uint
		Qa []string
		Ra []bool

		Aap []*uint8
		Bap []*uint16
		Cap []*uint32
		Dap []*uint64

		Eap []*int8
		Fap []*int16
		Gap []*int32
		Hap []*int64

		Iap []*float32
		Jap []*float64

		Map []*byte
		Nap []*rune

		Oap []*int
		Pap []*uint
		Qap []*string
		Rap []*bool
	}
	PtStr *struct {
		X int
	}
	InceptionStr struct {
		Y []struct {
			X *int
		}
	}
}

type XInlineStructs struct {
	B struct {
		A uint8
		B uint16
		C uint32
		D uint64

		E int8
		F int16
		G int32
		H int64

		I float32
		J float64

		M byte
		N rune

		O int
		P uint
		Q string
		R bool
		S time.Time

		Ap *uint8
		Bp *uint16
		Cp *uint32
		Dp *uint64

		Ep *int8
		Fp *int16
		Gp *int32
		Hp *int64

		Ip *float32
		Jp *float64

		Mp *byte
		Np *rune

		Op *int
		Pp *uint
		Qp *string
		Rp *bool
		Sp *time.Time

		Aa []uint8
		Ba []uint16
		Ca []uint32
		Da []uint64

		Ea []int8
		Fa []int16
		Ga []int32
		Ha []int64

		Ia []float32
		Ja []float64

		Ma []byte
		Na []rune

		Oa []int
		Pa []uint
		Qa []string
		Ra []bool

		Aap []*uint8
		Bap []*uint16
		Cap []*uint32
		Dap []*uint64

		Eap []*int8
		Fap []*int16
		Gap []*int32
		Hap []*int64

		Iap []*float32
		Jap []*float64

		Map []*byte
		Nap []*rune

		Oap []*int
		Pap []*uint
		Qap []*string
		Rap []*bool
	}
	PtStr *struct {
		X int
	}
	InceptionStr struct {
		Y []struct {
			X *int
		}
	}
}

// ffjson: skip
type TDominantField struct {
	X     *int `json:"Name,omitempty"`
	Y     *int `json:"Name,omitempty"`
	Other string
	Name  *int             `json",omitempty"`
	A     *struct{ X int } `json:"Name,omitempty"`
}

type XDominantField struct {
	X     *int `json:"Name,omitempty"`
	Y     *int `json:"Name,omitempty"`
	Other string
	Name  *int             `json",omitempty"`
	A     *struct{ X int } `json:"Name,omitempty"`
}
