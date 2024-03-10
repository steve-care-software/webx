package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type pointer struct {
	hash   hash.Hash
	entity string
	field  string
}

func createPointer(
	hash hash.Hash,
	entity string,
	field string,
) Pointer {
	out := pointer{
		hash:   hash,
		entity: entity,
		field:  field,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Entity returns the entity
func (obj *pointer) Entity() string {
	return obj.entity
}

// Field returns the field
func (obj *pointer) Field() string {
	return obj.field
}
