package references

import (
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type commit struct {
	hash      hash.Hash
	pointer   Pointer
	createdOn time.Time
}

func createCommit(
	hash hash.Hash,
	pointer Pointer,
	createdOn time.Time,
) Commit {
	out := commit{
		hash:      hash,
		pointer:   pointer,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Pointer returns the pointer
func (obj *commit) Pointer() Pointer {
	return obj.pointer
}

// CreatedOn returns the creation time
func (obj *commit) CreatedOn() time.Time {
	return obj.createdOn
}
