package compilers

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

// Adapter represents a compiler adapter
type Adapter interface {
	ToCompiler(content []byte) (Compiler, error)
	ToContent(ins Compiler) ([]byte, error)
}

// Builder represents a compiler builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithModules(modules []uint) Builder
	WithSelectors(selectors []hash.Hash) Builder
	WithPrograms(programs []hash.Hash) Builder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Compiler, error)
}

// Compiler represents the compiler database application
type Compiler interface {
	Hash() hash.Hash
	Name() string
	Modules() []uint
	Selectors() []hash.Hash
	Programs() []hash.Hash
	HasHistory() bool
	History() hashtrees.HashTree
}
