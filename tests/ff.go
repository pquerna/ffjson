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
