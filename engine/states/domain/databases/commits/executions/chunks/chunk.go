package chunks

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type chunk struct {
	hash        hash.Hash
	path        []string
	fingerPrint hash.Hash
}

func createChunk(
	hash hash.Hash,
	path []string,
	fingerPrint hash.Hash,
) Chunk {
	return &chunk{
		hash:        hash,
		path:        path,
		fingerPrint: fingerPrint,
	}
}

// Hash returns the hash
func (obj *chunk) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *chunk) Path() []string {
	return obj.path
}

// FingerPrint returns the fingerprint
func (obj *chunk) FingerPrint() hash.Hash {
	return obj.fingerPrint
}
