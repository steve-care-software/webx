package results

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
)

type success struct {
	hash  hash.Hash
	bytes []byte
	kind  layers.Kind
}

func createSuccess(
	hash hash.Hash,
	bytes []byte,
	kind layers.Kind,
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
func (obj *success) Kind() layers.Kind {
	return obj.kind
}
