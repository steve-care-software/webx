package connections

import "github.com/steve-care-software/datastencil/domain/hash"

type field struct {
	hash hash.Hash
	name string
	path []string
}

func createField(
	hash hash.Hash,
	name string,
	path []string,
) Field {
	out := field{
		hash: hash,
		name: name,
		path: path,
	}

	return &out
}

// Hash returns the hash
func (obj *field) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Path returns the path
func (obj *field) Path() []string {
	return obj.path
}
