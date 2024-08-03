package authorizations

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/identities/signers"
)

// Authorization represents an authorization
type Authorization interface {
	Hash() hash.Hash
	IsSignature() bool
	Signature() signers.Signature
	IsVote() bool
	Vote() signers.Vote
}
