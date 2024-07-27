package identities

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/identities/profiles"
	"github.com/steve-care-software/webx/engine/units/domain/identities/signers"
)

// Identity represents an identity
type Identity interface {
	Hash() hash.Hash
	Profile() profiles.Profile
	Signer() signers.Signer
	HasReferree() bool
	Referree() Identity
}
