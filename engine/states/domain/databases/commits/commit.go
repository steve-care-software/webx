package commits

import (
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type commit struct {
	hash       hash.Hash
	executions executions.Executions
	parent     hash.Hash
}

func createCommit(
	hash hash.Hash,
	executions executions.Executions,
) Commit {
	return &commit{
		hash:       hash,
		executions: executions,
		parent:     nil,
	}
}

func createCommitWithParent(
	hash hash.Hash,
	executions executions.Executions,
	parent hash.Hash,
) Commit {
	return &commit{
		hash:       hash,
		executions: executions,
		parent:     parent,
	}
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Executions returns the executions
func (obj *commit) Executions() executions.Executions {
	return obj.executions
}

// HasParent returns true if there is a parent, false otherwise
func (obj *commit) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent
func (obj *commit) Parent() hash.Hash {
	return obj.parent
}
