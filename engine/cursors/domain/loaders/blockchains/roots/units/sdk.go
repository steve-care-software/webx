package units

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/roots/units/purses"
)

// Unit represents unit description
type Unit interface {
	Hash() hash.Hash
	Symbol() string
	Name() string
	Description() string
	Purses() purses.Purses
	Lock() hash.Hash
}
