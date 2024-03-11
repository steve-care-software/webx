package inserts

import "github.com/steve-care-software/datastencil/domain/hash"

type insert struct {
	hash     hash.Hash
	context  string
	instance string
	path     string
}

func createInsert(
	hash hash.Hash,
	context string,
	instance string,
	path string,
) Insert {
	out := insert{
		hash:     hash,
		context:  context,
		instance: instance,
		path:     path,
	}

	return &out
}

// Hash returns the hash
func (obj *insert) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *insert) Context() string {
	return obj.context
}

// Instance returns the instance
func (obj *insert) Instance() string {
	return obj.instance
}

// Path returns the path
func (obj *insert) Path() string {
	return obj.path
}
