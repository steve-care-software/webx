package votes

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	Transaction() hash.Hash
	IsApproved() bool
	IsDeclined() bool
	IsNeutral() bool
}
