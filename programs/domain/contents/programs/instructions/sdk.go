package instructions

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/contents/programs/assignments"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithAssignment(assignment assignments.Assignment) Builder
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
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsExecution() bool
	Execution() hash.Hash
}
