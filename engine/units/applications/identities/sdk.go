package identities

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/identities"
)

// Application represents the units database
type Application interface {
	Insert(context uint, identity identities.Identity, password []byte) error
	Update(context uint, original hash.Hash, updated identities.Identity, oeiginalPassword []byte, newPassword []byte) error
	Delete(context uint, hash hash.Hash) error
}
