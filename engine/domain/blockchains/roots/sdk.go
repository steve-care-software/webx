package roots

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
)

// Root represents a root block
type Root interface {
	Hash() hash.Hash
	Package() string
	MiningValue() uint8
	BaseDifficulty() uint8
	IncreaseDiffPerBytes() float64
}
