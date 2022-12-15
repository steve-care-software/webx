package programs

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

// Adapter represents a programs adapter
type Adapter interface {
	ToProgram(content []byte) (Program, error)
	ToContent(ins Program) ([]byte, error)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithModules(modules []uint) Builder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Program, error)
}

// Program represents a program database
type Program interface {
	Hash() hash.Hash
	Name() string
	Modules() []uint
	HasHistory() bool
	History() hashtrees.HashTree
}
