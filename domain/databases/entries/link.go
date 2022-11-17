package entries

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type link struct {
	new       Entry
	pExisting *hash.Hash
}

func createLinkWithNew(
	new Entry,
) Link {
	return createLinkInternally(new, nil)
}

func createLinkWithExisting(
	pExisting *hash.Hash,
) Link {
	return createLinkInternally(nil, pExisting)
}

func createLinkInternally(
	new Entry,
	pExisting *hash.Hash,
) Link {
	out := link{
		new:       new,
		pExisting: pExisting,
	}

	return &out
}

// IsNew returns true if new, false otherwise
func (obj *link) IsNew() bool {
	return obj.new != nil
}

// New returns the new entry, if any
func (obj *link) New() Entry {
	return obj.new
}

// IsExisting returns true if existing, false otherwise
func (obj *link) IsExisting() bool {
	return obj.pExisting != nil
}

// Existing returns the existing hash, if any
func (obj *link) Existing() *hash.Hash {
	return obj.pExisting
}
