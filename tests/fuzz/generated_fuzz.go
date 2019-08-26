// +build gofuzz

package fuzz

// FuzzUnmarshal tests unmarshaling
func FuzzUnmarshal(fuzz []byte) int {
	data := &Data{}
	err := data.UnmarshalJSON(fuzz)
	if err != nil {
		return 0
	}
	return 1
}
