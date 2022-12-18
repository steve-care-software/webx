package references

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type contentKey struct {
	hash    hash.Hash
	kind    uint
	content Pointer
	commit  hash.Hash
}

func createContentKey(
	hash hash.Hash,
	kind uint,
	content Pointer,
	commit hash.Hash,
) ContentKey {
	out := contentKey{
		hash:    hash,
		kind:    kind,
		content: content,
		commit:  commit,
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

// Commit returns the commit
func (obj *contentKey) Commit() hash.Hash {
	return obj.commit
}
