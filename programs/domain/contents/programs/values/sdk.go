package values

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the value adapter
type Adapter interface {
	ToContent(ins Value) ([]byte, error)
	ToValue(content []byte) (Value, error)
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithInput(input uint) Builder
	WithValue(value hash.Hash) Builder
	WithConstant(constant []byte) Builder
	WithExecution(execution hash.Hash) Builder
	WithProgram(program hash.Hash) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	Content() Content
}

// Content represents a value content
type Content interface {
	IsInput() bool
	Input() *uint
	IsValue() bool
	Value() *hash.Hash
	IsConstant() bool
	Constant() []byte
	IsExecution() bool
	Execution() *hash.Hash
	IsProgram() bool
	Program() *hash.Hash
}
