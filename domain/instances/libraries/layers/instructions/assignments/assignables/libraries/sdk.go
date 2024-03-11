package libraries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a library builder
type Builder interface {
	Create() Builder
	WithCompiler(compiler compilers.Compiler) Builder
	WithDatabase(database databases.Database) Builder
	Now() (Library, error)
}

// Library represents a library assignable
type Library interface {
	Hash() hash.Hash
	IsCompiler() bool
	Compiler() compilers.Compiler
	IsDatabase() bool
	Database() databases.Database
}
