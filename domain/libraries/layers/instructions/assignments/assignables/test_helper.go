package assignables

import "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"

// NewAssignableWithBytesForTests creates a new assignable with bytes for tests
func NewAssignableWithBytesForTests(input bytes.Bytes) Assignable {
	ins, err := NewBuilder().Create().WithBytes(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
