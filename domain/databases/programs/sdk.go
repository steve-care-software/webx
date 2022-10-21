package programs

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithEngine(engine string) Builder
	WithCompiler(compiler []byte) Builder
	WithScript(script []byte) Builder
	Now() (Program, error)
}

// Program represents  program
type Program interface {
	Hash() hash.Hash
	Engine() string
	Compiler() []byte
	Script() []byte
}
