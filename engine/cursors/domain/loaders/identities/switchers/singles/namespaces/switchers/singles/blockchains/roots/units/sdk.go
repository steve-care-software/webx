package units

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots/units/purses"
)

// Builder represents a unit builder
type Builder interface {
	Create() Builder
	WithSymbol(symbol string) Builder
	WithPurses(purses purses.Purses) Builder
	WithLock(lock []byte) Builder
	Now() (Unit, error)
}

// Unit represents unit description
type Unit interface {
	Hash() hash.Hash
	Symbol() string
	Purses() purses.Purses
	Lock() []byte
}
