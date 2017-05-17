package v1

import "fmt"

// ErrUnknowFields error is return when some unknow fields are found
type ErrUnknowFields struct {
	Fields []string
}

// NewErrUnknowFields create an instance UnknowFields error
func NewErrUnknowFields(fields []string) *ErrUnknowFields {
	return &ErrUnknowFields{
		Fields: fields,
	}
}
func (e *ErrUnknowFields) Error() string {
	return fmt.Sprintf("unknow fields %v", e.Fields)
}
