package chunks

import "github.com/steve-care-software/datastencil/domain/hash"

type chunk struct {
	hash        hash.Hash
	path        []string
	fingerprint hash.Hash
}

func createChunk(
	hash hash.Hash,
	path []string,
	fingerprint hash.Hash,
) Chunk {
	out := chunk{
		hash:        hash,
		path:        path,
		fingerprint: fingerprint,
	}

	return &out
}

// Hash returns the hash
func (obj *chunk) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *chunk) Path() []string {
	return obj.path
}

// Fingerprint returns the fingerprint
func (obj *chunk) Fingerprint() hash.Hash {
	return obj.fingerprint
}
