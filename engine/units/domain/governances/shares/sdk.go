package shares

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// Share represents the governance share
type Share interface {
	Hash() hash.Hash
	Developer() uint8
	Affiliate() uint8
	Miner() uint8
}
