package fuzz

import (
	"time"
)

// Data provides an unmarshaling target
type Data struct {
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

	IP *float32
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
