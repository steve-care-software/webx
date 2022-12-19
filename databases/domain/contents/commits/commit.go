package commits

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type commit struct {
	hash    hash.Hash
	values  hashtrees.HashTree
	pParent *hash.Hash
}

func createCommit(
	hash hash.Hash,
	values hashtrees.HashTree,
) Commit {
	return createCommitInternally(hash, values, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	values hashtrees.HashTree,
	pParent *hash.Hash,
) Commit {
	return createCommitInternally(hash, values, pParent)
}

func createCommitInternally(
	hash hash.Hash,
	values hashtrees.HashTree,
	pParent *hash.Hash,
) Commit {
	out := commit{
		hash:    hash,
		values:  values,
		pParent: pParent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Values returns the values
func (obj *commit) Values() hashtrees.HashTree {
	return obj.values
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() *hash.Hash {
	return obj.pParent
}
