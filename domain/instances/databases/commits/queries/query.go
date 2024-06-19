package queries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries/chunks"
)

type query struct {
	hash  hash.Hash
	bytes []byte
	chunk chunks.Chunk
}

func createQueryWithBytes(
	hash hash.Hash,
	bytes []byte,
) Query {
	return createQueryInternally(hash, bytes, nil)
}

func createQueryWithChunk(
	hash hash.Hash,
	chunk chunks.Chunk,
) Query {
	return createQueryInternally(hash, nil, chunk)
}

func createQueryInternally(
	hash hash.Hash,
	bytes []byte,
	chunk chunks.Chunk,
) Query {
	out := query{
		hash:  hash,
		bytes: bytes,
		chunk: chunk,
	}

	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *query) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *query) Bytes() []byte {
	return obj.bytes
}

// IsChunk returns true if there is chunk, false otherwise
func (obj *query) IsChunk() bool {
	return obj.chunk != nil
}

// Chunk returns the bytes, if any
func (obj *query) Chunk() chunks.Chunk {
	return obj.chunk
}
