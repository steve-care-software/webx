package singles

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents/singles/referrals"
	"github.com/steve-care-software/webx/engine/units/domain/identities/signers"
)

// Single represents a single
type Single interface {
	Hash() hash.Hash
	Referral() referrals.Referral
	Signature() signers.Signature
}
