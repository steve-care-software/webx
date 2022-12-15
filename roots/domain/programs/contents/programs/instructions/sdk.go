package instructions

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the instruction adapter
type Adapter interface {
	ToContent(ins Instruction) ([]byte, error)
	ToInstruction(content []byte) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithValue(value hash.Hash) Builder
	WithExecution(execution hash.Hash) Builder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	Content() Content
}

// Content represents an instruction's content
type Content interface {
	IsValue() bool
	Value() *hash.Hash
	IsExecution() bool
	Execution() *hash.Hash
}
