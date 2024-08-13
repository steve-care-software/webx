package deletes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

type delete struct {
	hash hash.Hash
	name string
	vote signers.Vote
}

func createDelete(
	hash hash.Hash,
	name string,
	vote signers.Vote,
) Delete {
	out := delete{
		hash: hash,
		name: name,
		vote: vote,
	}

	return &out
}

// Hash returns the hash
func (obj *delete) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *delete) Name() string {
	return obj.name
}

// Vote returns the vote
func (obj *delete) Vote() signers.Vote {
	return obj.vote
}
