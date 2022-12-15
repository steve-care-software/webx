package references

import (
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type contentKey struct {
	hash      hash.Hash
	kind      uint
	content   Pointer
	trx       hash.Hash
	createdOn time.Time
}

func createContentKey(
	hash hash.Hash,
	kind uint,
	content Pointer,
	trx hash.Hash,
	createdOn time.Time,
) ContentKey {
	out := contentKey{
		hash:      hash,
		kind:      kind,
		content:   content,
		trx:       trx,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *contentKey) Hash() hash.Hash {
	return obj.hash
}

// Kind returns the kind
func (obj *contentKey) Kind() uint {
	return obj.kind
}

// Content returns the content
func (obj *contentKey) Content() Pointer {
	return obj.content
}

// Transaction returns the transaction
func (obj *contentKey) Transaction() hash.Hash {
	return obj.trx
}

// CreatedOn returns the creation time
func (obj *contentKey) CreatedOn() time.Time {
	return obj.createdOn
}
