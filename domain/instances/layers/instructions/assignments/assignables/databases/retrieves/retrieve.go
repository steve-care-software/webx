package retrieves

import "github.com/steve-care-software/datastencil/domain/hash"

type retrieve struct {
	hash     hash.Hash
	exists   string
	retrieve string
	isList   bool
}

func createRetrieveWithExists(
	hash hash.Hash,
	exists string,
) Retrieve {
	return createRetrieveInternally(hash, exists, "", false)
}

func createRetrieveWithRetrieve(
	hash hash.Hash,
	ret string,
) Retrieve {
	return createRetrieveInternally(hash, "", ret, false)
}

func createRetrieveWithList(
	hash hash.Hash,
) Retrieve {
	return createRetrieveInternally(hash, "", "", true)
}

func createRetrieveInternally(
	hash hash.Hash,
	exists string,
	ret string,
	isList bool,
) Retrieve {
	out := retrieve{
		hash:     hash,
		exists:   exists,
		retrieve: ret,
		isList:   isList,
	}

	return &out
}

// Hash returns the hash
func (obj *retrieve) Hash() hash.Hash {
	return obj.hash
}

// IsList returns true if there is a list, false otherwise
func (obj *retrieve) IsList() bool {
	return obj.isList
}

// IsExists returns true if there is an exists, false otherwise
func (obj *retrieve) IsExists() bool {
	return obj.exists != ""
}

// Exists returns the exists, if any
func (obj *retrieve) Exists() string {
	return obj.exists
}

// IsRetrieve returns true if there is a retrieve, false otherwise
func (obj *retrieve) IsRetrieve() bool {
	return obj.retrieve != ""
}

// Retrieve returns the retrieve, if any
func (obj *retrieve) Retrieve() string {
	return obj.retrieve
}
