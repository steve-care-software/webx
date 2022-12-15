package references

import (
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type blockchainKey struct {
	hash      hash.Hash
	content   Pointer
	createdOn time.Time
}

func createBlockchainKey(
	hash hash.Hash,
	content Pointer,
	createdOn time.Time,
) BlockchainKey {
	out := blockchainKey{
		hash:      hash,
		content:   content,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *blockchainKey) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *blockchainKey) Content() Pointer {
	return obj.content
}

// CreatedOn returns the creation time
func (obj *blockchainKey) CreatedOn() time.Time {
	return obj.createdOn
}
