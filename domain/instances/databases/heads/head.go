package heads

import "github.com/steve-care-software/datastencil/domain/hash"

type head struct {
	hash        hash.Hash
	path        []string
	description string
	isActive    bool
}

func createHead(
	hash hash.Hash,
	path []string,
	description string,
	isActive bool,
) Head {
	out := head{
		hash:        hash,
		path:        path,
		description: description,
		isActive:    isActive,
	}

	return &out
}

// Hash returns the hash
func (obj *head) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *head) Path() []string {
	return obj.path
}

// Description returns the description
func (obj *head) Description() string {
	return obj.description
}

// IsActive returns true if active, false otherwise
func (obj *head) IsActive() bool {
	return obj.isActive
}
