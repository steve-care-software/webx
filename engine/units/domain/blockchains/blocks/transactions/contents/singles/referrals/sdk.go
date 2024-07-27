package referrals

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

// Referral represents a referral
type Referral interface {
	Hash() hash.Hash
	Referree() hash.Hash
	Profile() hash.Hash
}
