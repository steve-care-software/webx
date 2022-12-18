package programs

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the program adapter
type Adapter interface {
	ToContent(ins Program) ([]byte, error)
	ToProgram(content []byte) (Program, error)
}

// Builder represents a program
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithInstructions(instructions []hash.Hash) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Instructions() []hash.Hash
	HasOutputs() bool
	Outputs() []uint
}
