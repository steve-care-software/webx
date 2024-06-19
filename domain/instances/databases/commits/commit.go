package commits

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries"
)

type commit struct {
	hash    hash.Hash
	queries queries.Queries
	parent  hash.Hash
}

func createCommit(
	hash hash.Hash,
	queries queries.Queries,
) Commit {
	return createCommitInternally(hash, queries, nil)
}

func createCommitWithParent(
	hash hash.Hash,
	queries queries.Queries,
	parent hash.Hash,
) Commit {
	return createCommitInternally(hash, queries, parent)
}

func createCommitInternally(
	hash hash.Hash,
	queries queries.Queries,
	parent hash.Hash,
) Commit {
	out := commit{
		hash:    hash,
		queries: queries,
		parent:  parent,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Queries returns the queries
func (obj *commit) Queries() queries.Queries {
	return obj.queries
}

// HasParent returns true if there is parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *commit) Parent() hash.Hash {
	return obj.parent
}
