package units

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
)

type unit struct {
	hash       hash.Hash
	content    Content
	signatures []signatures.RingSignature
}

func createUnit(
	hash hash.Hash,
	content Content,
	signatures []signatures.RingSignature,
) Unit {
	out := unit{
		hash:       hash,
		content:    content,
		signatures: signatures,
	}

	return &out
}

// Hash returns the hash
func (obj *unit) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *unit) Content() Content {
	return obj.content
}

// Signatures returns the signatures
func (obj *unit) Signatures() []signatures.RingSignature {
	return obj.signatures
}
