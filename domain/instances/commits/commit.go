package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents"
)

type commit struct {
	hash      hash.Hash
	content   contents.Content
	signature signers.Signature
}

func createCommit(
	hash hash.Hash,
	content contents.Content,
	signature signers.Signature,
) Commit {
	out := commit{
		hash:      hash,
		content:   content,
		signature: signature,
	}

	return &out
}

// Hash returns the hash
func (obj *commit) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *commit) Content() contents.Content {
	return obj.content
}

// Signature returns the signature
func (obj *commit) Signature() signers.Signature {
	return obj.signature
}

// Index returns the index
func (obj *commit) Index() uint {
	return obj.content.Previous().Index() + 1
}

// PublicKey returns the signature's publicKey
func (obj *commit) PublicKey() (signers.PublicKey, error) {
	msg := obj.Content().Hash().Bytes()
	return obj.signature.PublicKey(msg)
}
