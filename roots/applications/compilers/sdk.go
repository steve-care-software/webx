package compilers

import (
	"github.com/steve-care-software/webx/compilers/domain/compilers"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// Application represents a compiler application
type Application interface {
	List(context uint) ([]hash.Hash, error)
	Retrieve(context uint, hash hash.Hash) (compilers.Compiler, error)
	Insert(context uint, compiler compilers.Compiler) error
	InsertAll(context uint, compilers []compilers.Compiler) error
}
