package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

type update struct {
	hash    hash.Hash
	content Content
	vote    signers.Vote
}

func createUpdate(
	hash hash.Hash,
	content Content,
	vote signers.Vote,
) Update {
	out := update{
		hash:    hash,
		content: content,
		vote:    vote,
	}

	return &out
}

// Hash returns the hash
func (obj *update) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *update) Content() Content {
	return obj.content
}

// Vote returns the vote
func (obj *update) Vote() signers.Vote {
	return obj.vote
}
