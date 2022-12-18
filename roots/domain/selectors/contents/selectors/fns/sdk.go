package fns

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a func builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithProgram(program hash.Hash) Builder
	WithParam(param uint) Builder
	IsSingle() Builder
	IsContent() Builder
	Now() (Fn, error)
}

// Fn represents a func
type Fn interface {
	Hash() hash.Hash
	IsSingle() bool
	IsContent() bool
	Program() hash.Hash
	Param() uint
}
