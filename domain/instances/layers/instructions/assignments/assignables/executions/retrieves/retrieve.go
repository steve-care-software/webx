package retrieves

import "github.com/steve-care-software/historydb/domain/hash"

type retrieve struct {
	hash    hash.Hash
	context string
	index   string
	ret     string
	length  string
}

func createRetrieve(
	hash hash.Hash,
	index string,
	context string,
	ret string,
) Retrieve {
	return createRetrieveInternally(hash, index, context, ret, "")
}

func createRetrieveWithLength(
	hash hash.Hash,
	index string,
	context string,
	ret string,
	length string,
) Retrieve {
	return createRetrieveInternally(hash, index, context, ret, length)
}

func createRetrieveInternally(
	hash hash.Hash,
	index string,
	context string,
	ret string,
	length string,
) Retrieve {
	out := retrieve{
		hash:    hash,
		context: context,
		index:   index,
		ret:     ret,
		length:  length,
	}

	return &out
}

// Hash returns the hash
func (obj *retrieve) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *retrieve) Context() string {
	return obj.context
}

// Index returns the index
func (obj *retrieve) Index() string {
	return obj.index
}

// Return returns the return
func (obj *retrieve) Return() string {
	return obj.ret
}

// HasLength returns true if there is a length, false otherwise
func (obj *retrieve) HasLength() bool {
	return obj.length != ""
}

// Length returns the length, if any
func (obj *retrieve) Length() string {
	return obj.length
}
