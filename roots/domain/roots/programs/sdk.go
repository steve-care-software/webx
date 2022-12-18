package programs

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
)

// Builder represents a programs builder
type Builder interface {
	Create() Builder
	WithList(list []Program) Builder
	Now() (Programs, error)
}

// Programs represents programs
type Programs interface {
	Hash() hash.Hash
	List() []Program
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithName(name string) ProgramBuilder
	WithModules(modules modules.Modules) ProgramBuilder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Program, error)
}

// Program represents a program database
type Program interface {
	Hash() hash.Hash
	Name() string
	Modules() modules.Modules
	HasHistory() bool
	History() hashtrees.HashTree
}
