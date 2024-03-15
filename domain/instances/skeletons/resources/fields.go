package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type fields struct {
	hash hash.Hash
	list []Field
}

func createFields(
	hash hash.Hash,
	list []Field,
) Fields {
	out := fields{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *fields) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *fields) List() []Field {
	return obj.list
}
