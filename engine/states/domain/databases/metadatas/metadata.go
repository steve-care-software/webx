package metadatas

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

type metadata struct {
	hash        hash.Hash
	path        []string
	name        string
	description string
}

func createMetadata(
	hash hash.Hash,
	path []string,
	name string,
	description string,
) MetaData {
	return &metadata{
		hash:        hash,
		path:        path,
		name:        name,
		description: description,
	}
}

// Hash returns the hash
func (obj *metadata) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *metadata) Path() []string {
	return obj.path
}

// Name returns the name
func (obj *metadata) Name() string {
	return obj.name
}

// Description returns the description
func (obj *metadata) Description() string {
	return obj.description
}
