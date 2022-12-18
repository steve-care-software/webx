package commits

import (
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type commit struct {
	hash      hash.Hash
	values    hashtrees.HashTree
	createdOn time.Time
	pParent   *hash.Hash
}

func createCommit(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
) Commit {
	return createCommitInternally(hash, values, createdOn, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	pParent *hash.Hash,
) Commit {
	return createCommitInternally(hash, values, createdOn, pParent)
}

func createCommitInternally(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	pParent *hash.Hash,
) Commit {
	out := commit{
		hash:      hash,
		values:    values,
		createdOn: createdOn,
		pParent:   pParent,
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

// CreatedOn returns the creation time
func (obj *commit) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.pParent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() *hash.Hash {
	return obj.pParent
}
