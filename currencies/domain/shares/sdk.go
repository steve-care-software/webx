package shares

import (
	"math/big"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// Shares represents shares
type Shares interface {
	Hash() hash.Hash
	List() []Share
}

// Share represents a share
type Share interface {
	Hash() hash.Hash
	Name() string
	Ring() []hash.Hash
	Units() big.Int
}
