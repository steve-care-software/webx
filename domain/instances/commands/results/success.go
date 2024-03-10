package results

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
)

type success struct {
	hash  hash.Hash
	bytes []byte
	kind  kinds.Kind
}

func createSuccess(
	hash hash.Hash,
	bytes []byte,
	kind kinds.Kind,
) Success {
	out := success{
		hash:  hash,
		bytes: bytes,
		kind:  kind,
	}

	return &out
}

// Hash returns the hash
func (obj *success) Hash() hash.Hash {
	return obj.hash
}

// Bytes returns the bytes
func (obj *success) Bytes() []byte {
	return obj.bytes
}

// Kind returns the kind
func (obj *success) Kind() kinds.Kind {
	return obj.kind
}
