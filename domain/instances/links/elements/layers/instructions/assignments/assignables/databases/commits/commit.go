package commits

import "github.com/steve-care-software/datastencil/domain/hash"

type commit struct {
	hash        hash.Hash
	description string
	actions     string
	parent      string
}

func createCommit(
	hash hash.Hash,
	description string,
	actions string,
) Commit {
	return createCommitInternally(hash, description, actions, "")
}

func createCommitWithParent(
	hash hash.Hash,
	description string,
	actions string,
	parent string,
) Commit {
	return createCommitInternally(hash, description, actions, parent)
}

func createCommitInternally(
	hash hash.Hash,
	description string,
	actions string,
	parent string,
) Commit {
	out := commit{
		hash:        hash,
		description: description,
		actions:     actions,
		parent:      parent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Description returns the description
func (obj *commit) Description() string {
	return obj.description
}

// Actions returns the actions
func (obj *commit) Actions() string {
	return obj.actions
}

// HashParent returns true if there is a parent, false otherwise
func (obj *commit) HashParent() bool {
	return obj.parent != ""
}

// Parent returns the parent, if any
func (obj *commit) Parent() string {
	return obj.parent
}
