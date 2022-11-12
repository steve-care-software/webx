package references

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type key struct {
	hash      hash.Hash
	index     uint
	kind      uint8
	content   Pointer
	isEntity  bool
	createdOn time.Time
	pTrx      *uint
}

func createKey(
	hash hash.Hash,
	index uint,
	kind uint8,
	content Pointer,
	isEntity bool,
	createdOn time.Time,
) Key {
	return createKeyInternally(hash, index, kind, content, isEntity, createdOn, nil)
}

func createKeyWithTransaction(
	hash hash.Hash,
	index uint,
	kind uint8,
	content Pointer,
	isEntity bool,
	createdOn time.Time,
	pTrx *uint,
) Key {
	return createKeyInternally(hash, index, kind, content, isEntity, createdOn, pTrx)
}

func createKeyInternally(
	hash hash.Hash,
	index uint,
	kind uint8,
	content Pointer,
	isEntity bool,
	createdOn time.Time,
	pTrx *uint,
) Key {
	out := key{
		hash:      hash,
		index:     index,
		kind:      kind,
		content:   content,
		isEntity:  isEntity,
		createdOn: createdOn,
		pTrx:      pTrx,
	}

	return &out
}

// Hash returns the hash
func (obj *key) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *key) Index() uint {
	return obj.index
}

// Kind returns the kind
func (obj *key) Kind() uint8 {
	return obj.kind
}

// Content returns the content
func (obj *key) Content() Pointer {
	return obj.content
}

// IsEntity returns true if there is an entity, false otherwise
func (obj *key) IsEntity() bool {
	return obj.isEntity
}

// CreatedOn returns the creation time
func (obj *key) CreatedOn() time.Time {
	return obj.createdOn
}

// HasTransaction returns true if there is a transaction, false otherwise
func (obj *key) HasTransaction() bool {
	return obj.pTrx != nil
}

// Transaction returns the transaction, if any
func (obj *key) Transaction() *uint {
	return obj.pTrx
}
