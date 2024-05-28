package outputs

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new output builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the output adapter
type Adapter interface {
	ToBytes(ins Output) ([]byte, error)
	ToInstance(bytes []byte) (Output, error)
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithInput(input []byte) Builder
	WithExecute(execute []byte) Builder
	Now() (Output, error)
}

// Output represents an output
type Output interface {
	Hash() hash.Hash
	Value() []byte
	Input() []byte
	HasExecute() bool
	Execute() []byte
}
