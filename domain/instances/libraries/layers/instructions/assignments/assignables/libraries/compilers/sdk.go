package compilers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the compiler application
type Builder interface {
	Create() Builder
	WithCompile(compile string) Builder
	WithDecompile(decompile string) Builder
	Now() (Compiler, error)
}

// Compiler represents a compiler
type Compiler interface {
	Hash() hash.Hash
	IsCompile() bool
	Compile() string
	IsDecompile() bool
	Decompile() string
}
