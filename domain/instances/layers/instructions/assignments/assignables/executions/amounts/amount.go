package amounts

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

type amount struct {
	hash    hash.Hash
	context string
	ret     string
}

func createAmount(
	hash hash.Hash,
	context string,
	ret string,
) Amount {
	out := amount{
		hash:    hash,
		context: context,
		ret:     ret,
	}

	return &out
}

// Hash returns the hash
func (obj *amount) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *amount) Context() string {
	return obj.context
}

// Return returns the retunr
func (obj *amount) Return() string {
	return obj.ret
}
