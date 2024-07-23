package inits

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type initStr struct {
	hash        hash.Hash
	path        string
	name        string
	description string
}

func createInit(
	hash hash.Hash,
	path string,
	name string,
	description string,
) Init {
	out := initStr{
		hash:        hash,
		path:        path,
		name:        name,
		description: description,
	}

	return &out
}

// Hash returns the hash
func (obj *initStr) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *initStr) Path() string {
	return obj.path
}

// Name returns the name
func (obj *initStr) Name() string {
	return obj.name
}

// Description returns the description
func (obj *initStr) Description() string {
	return obj.description
}
