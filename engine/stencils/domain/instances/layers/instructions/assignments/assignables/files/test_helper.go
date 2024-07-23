package files

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files/opens"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files/reads"
)

// NewFileWithOpenForTests creates a new file with open for tests
func NewFileWithOpenForTests(open opens.Open) File {
	ins, err := NewBuilder().Create().WithOpen(open).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFileWithReadForTests creates a new file with read for tests
func NewFileWithReadForTests(read reads.Read) File {
	ins, err := NewBuilder().Create().WithRead(read).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFileWithExistsForTests creates a new file with exists for tests
func NewFileWithExistsForTests(exists string) File {
	ins, err := NewBuilder().Create().WithExists(exists).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewFileWithLengthForTests creates a new file with length for tests
func NewFileWithLengthForTests(length string) File {
	ins, err := NewBuilder().Create().WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
