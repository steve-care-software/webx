package compilers

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the compiler adapter
type Adapter interface {
	ToBytes(ins Compiler) ([]byte, error)
	ToInstance(bytes []byte) (Compiler, error)
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
