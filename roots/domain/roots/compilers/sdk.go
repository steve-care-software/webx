package compilers

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
	"github.com/steve-care-software/webx/roots/domain/roots/programs"
	"github.com/steve-care-software/webx/roots/domain/roots/selectors"
)

// Builder represents a compiler builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithModules(modules modules.Modules) Builder
	WithSelectors(selectors selectors.Selectors) Builder
	WithPrograms(programs programs.Programs) Builder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Compiler, error)
}

// Compiler represents the compiler database application
type Compiler interface {
	Hash() hash.Hash
	Name() string
	Modules() modules.Modules
	Selectors() selectors.Selectors
	Programs() programs.Programs
	HasHistory() bool
	History() hashtrees.HashTree
}
