package transforms

import "github.com/steve-care-software/datastencil/domain/hash"

type transform struct {
	hash  hash.Hash
	query []byte
	bytes []byte
}

func createTransform(
	hash hash.Hash,
	query []byte,
	bytes []byte,
) Transform {
	out := transform{
		hash:  hash,
		query: query,
		bytes: bytes,
	}

	return &out
}

// Hash returns the hash
func (obj *transform) Hash() hash.Hash {
	return obj.hash
}

// Query returns the query
func (obj *transform) Query() []byte {
	return obj.query
}

// Bytes returns the bytes
func (obj *transform) Bytes() []byte {
	return obj.bytes
}
