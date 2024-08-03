package opens

import "github.com/steve-care-software/webx/engine/hashes/domain/hash"

type open struct {
	hash       hash.Hash
	path       string
	permission string
}

func createOpen(
	hash hash.Hash,
	path string,
	permission string,
) Open {
	out := open{
		hash:       hash,
		path:       path,
		permission: permission,
	}

	return &out
}

// Hash returns the hash
func (obj *open) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *open) Path() string {
	return obj.path
}

// Permission returns the permission
func (obj *open) Permission() string {
	return obj.permission
}
