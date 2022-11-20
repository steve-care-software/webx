package entries

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type relation struct {
	new      Entries
	existing []hash.Hash
}

func createRelationWithNew(
	new Entries,
) Relation {
	return createRelationInternally(new, nil)
}

func createRelationWithExisting(
	existing []hash.Hash,
) Relation {
	return createRelationInternally(nil, existing)
}

func createRelationInternally(
	new Entries,
	existing []hash.Hash,
) Relation {
	out := relation{
		new:      new,
		existing: existing,
	}

	return &out
}

// IsNew returns true if new, false otherwise
func (obj *relation) IsNew() bool {
	return obj.new != nil
}

// New returns new entries, if any
func (obj *relation) New() Entries {
	return obj.new
}

// IsExisting returns true if existing, false otherwise
func (obj *relation) IsExisting() bool {
	return obj.existing != nil
}

// Existing returns existing hashes, if any
func (obj *relation) Existing() []hash.Hash {
	return obj.existing
}
