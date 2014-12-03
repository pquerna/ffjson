package tff

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
