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
	mine      Mine
	parent    Commit
}

func createCommit(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
) Commit {
	return createCommitInternally(hash, values, createdOn, nil, nil)
}

func createCommitWithMine(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	mine Mine,
) Commit {
	return createCommitInternally(hash, values, createdOn, mine, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	parent Commit,
) Commit {
	return createCommitInternally(hash, values, createdOn, nil, parent)
}

func createCommitWithMineAndParent(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	mine Mine,
	parent Commit,
) Commit {
	return createCommitInternally(hash, values, createdOn, mine, parent)
}

func createCommitInternally(
	hash hash.Hash,
	values hashtrees.HashTree,
	createdOn time.Time,
	mine Mine,
	parent Commit,
) Commit {
	out := commit{
		hash:      hash,
		values:    values,
		createdOn: createdOn,
		mine:      mine,
		parent:    parent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Height returns the height
func (obj *commit) Height() uint {
	amount := uint(1)
	if obj.HasParent() {
		amount += obj.parent.Height()
	}

	return amount
}

// Values returns the values
func (obj *commit) Values() hashtrees.HashTree {
	return obj.values
}

// CreatedOn returns the creation time
func (obj *commit) CreatedOn() time.Time {
	return obj.createdOn
}

// HasMine returns true if there is a mine, false otherwise
func (obj *commit) HasMine() bool {
	return obj.mine != nil
}

// Mine returns the mine, if any
func (obj *commit) Mine() Mine {
	return obj.mine
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() Commit {
	return obj.parent
}
