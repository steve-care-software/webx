package deletes

import "github.com/steve-care-software/datastencil/domain/hash"

type delete struct {
	hash       hash.Hash
	context    string
	path       string
	identifier string
}

func createDelete(
	hash hash.Hash,
	context string,
	path string,
	identifier string,
) Delete {
	out := delete{
		hash:       hash,
		context:    context,
		path:       path,
		identifier: identifier,
	}

	return &out
}

// Hash returns the hash
func (obj *delete) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *delete) Context() string {
	return obj.context
}

// Path returns the path
func (obj *delete) Path() string {
	return obj.path
}

// Identifier returns the identifier
func (obj *delete) Identifier() string {
	return obj.identifier
}
