package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

type database struct {
	hash   hash.Hash
	commit commits.Commit
	head   heads.Head
}

func createDatabase(
	hash hash.Hash,
	commit commits.Commit,
	head heads.Head,
) Database {
	out := database{
		hash:   hash,
		commit: commit,
		head:   head,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// Commit returns the commit
func (obj *database) Commit() commits.Commit {
	return obj.commit
}

// Head returns the head
func (obj *database) Head() heads.Head {
	return obj.head
}
