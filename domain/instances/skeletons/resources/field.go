package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type field struct {
	hash     hash.Hash
	name     string
	kind     Kind
	canBeNil bool
}

func createField(
	hash hash.Hash,
	name string,
	kind Kind,
	canBeNil bool,
) Field {
	return createFieldInternally(hash, name, kind, canBeNil)
}

func createFieldInternally(
	hash hash.Hash,
	name string,
	kind Kind,
	canBeNil bool,
) Field {
	out := field{
		hash:     hash,
		name:     name,
		kind:     kind,
		canBeNil: canBeNil,
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

// Kind returns the kind
func (obj *field) Kind() Kind {
	return obj.kind
}

// CanBeNil returns true if canBeNil, false otherwise
func (obj *field) CanBeNil() bool {
	return obj.canBeNil
}
