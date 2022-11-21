package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents a content transaction application
type Application interface {
	Begin() (*uint, error)
	Insert(context uint, content []byte) error
	Delete(context uint, hash hash.Hash) error
	Approve(context uint, hash hash.Hash, proof big.Int) error
	Push(context uint) error
	Cancel(context uint) error
}
