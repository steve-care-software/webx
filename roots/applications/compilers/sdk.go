package compilers

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/compilers/compilers"
	"github.com/steve-care-software/webx/roots/domain/programs/programs"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
)

// Application represents a compiler application
type Application interface {
	Database
	Software
}

// Software represents the compiler software application
type Software interface {
	Execute(compiler compilers.Compiler, modules modules.Modules, script []byte) (programs.Program, error)
}

// Database represents the compiler database application
type Database interface {
	List(context uint) ([]hash.Hash, error)
	Retrieve(context uint, hash hash.Hash) (compilers.Compiler, error)
	Insert(context uint, compiler compilers.Compiler) error
	InsertAll(context uint, compilers []compilers.Compiler) error
}
