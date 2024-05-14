package databases

import "github.com/steve-care-software/datastencil/domain/hash"

type database struct {
	hash        hash.Hash
	path        string
	description string
	head        string
	isActive    string
}

func createDatabase(
	hash hash.Hash,
	path string,
	description string,
	head string,
	isActive string,
) Database {
	out := database{
		hash:        hash,
		path:        path,
		description: description,
		head:        head,
		isActive:    isActive,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *database) Path() string {
	return obj.path
}

// Description returns the description
func (obj *database) Description() string {
	return obj.description
}

// Head returns the head
func (obj *database) Head() string {
	return obj.head
}

// IsActive returns the isActive
func (obj *database) IsActive() string {
	return obj.isActive
}
