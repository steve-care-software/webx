package executions

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
)

type execution struct {
	hash  hash.Hash
	bytes []byte
	chunk chunks.Chunk
}

func createExecutionWithBytes(
	hash hash.Hash,
	bytes []byte,
) Execution {
	return &execution{
		hash:  hash,
		bytes: bytes,
	}
}

func createExecutionWithChunk(
	hash hash.Hash,
	chunk chunks.Chunk,
) Execution {
	return &execution{
		hash:  hash,
		chunk: chunk,
	}
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *execution) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes if any
func (obj *execution) Bytes() []byte {
	return obj.bytes
}

// IsChunk returns true if there is a chunk
func (obj *execution) IsChunk() bool {
	return obj.chunk != nil
}

// Chunk returns the chunk if any
func (obj *execution) Chunk() chunks.Chunk {
	return obj.chunk
}
