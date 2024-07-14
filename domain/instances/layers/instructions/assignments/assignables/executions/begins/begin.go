package begins

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

type begin struct {
	hash    hash.Hash
	path    string
	context string
}

func createBegin(
	hash hash.Hash,
	path string,
	context string,
) Begin {
	out := begin{
		hash:    hash,
		path:    path,
		context: context,
	}

	return &out
}

// Hash returns the hash
func (obj *begin) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *begin) Path() string {
	return obj.path
}

// Context returns the context
func (obj *begin) Context() string {
	return obj.context
}
