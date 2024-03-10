package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type resource struct {
	hash  hash.Hash
	field Pointer
	value interface{}
}

func createResourceWithField(
	hash hash.Hash,
	field Pointer,
) Resource {
	return createResourceInternally(hash, field, nil)
}

func createResourceWithValue(
	hash hash.Hash,
	value interface{},
) Resource {
	return createResourceInternally(hash, nil, value)
}

func createResourceInternally(
	hash hash.Hash,
	field Pointer,
	value interface{},
) Resource {
	out := resource{
		hash:  hash,
		field: field,
		value: value,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// IsField returns true if there is a field, false otherwise
func (obj *resource) IsField() bool {
	return obj.field != nil
}

// Field returns the field, if any
func (obj *resource) Field() Pointer {
	return obj.field
}

// IsValue returns true if there is a value, false otherwise
func (obj *resource) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *resource) Value() interface{} {
	return obj.value
}
