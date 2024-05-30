package lists

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/lists/fetches"
)

type list struct {
	hash   hash.Hash
	fetch  fetches.Fetch
	length string
	create string
}

func createListWithFetch(
	hash hash.Hash,
	fetch fetches.Fetch,
) List {
	return createListInternally(hash, fetch, "", "")
}

func createListWithLength(
	hash hash.Hash,
	length string,
) List {
	return createListInternally(hash, nil, length, "")
}

func createListWithCreate(
	hash hash.Hash,
	create string,
) List {
	return createListInternally(hash, nil, "", create)
}

func createListInternally(
	hash hash.Hash,
	fetch fetches.Fetch,
	length string,
	create string,
) List {
	out := list{
		hash:   hash,
		fetch:  fetch,
		length: length,
		create: create,
	}

	return &out
}

// Hash returns the hash
func (obj *list) Hash() hash.Hash {
	return obj.hash
}

// IsFetch returns true if there is a fetch, false otherwise
func (obj *list) IsFetch() bool {
	return obj.fetch != nil
}

// Fetch returns the fetch, if any
func (obj *list) Fetch() fetches.Fetch {
	return obj.fetch
}

// IsLength returns true if there is a length, false otherwise
func (obj *list) IsLength() bool {
	return obj.length != ""
}

// Length returns the length, if any
func (obj *list) Length() string {
	return obj.length
}

// IsCreate returns true if there is a create, false otherwise
func (obj *list) IsCreate() bool {
	return obj.create != ""
}

// Create returns the create, if any
func (obj *list) Create() string {
	return obj.create
}
