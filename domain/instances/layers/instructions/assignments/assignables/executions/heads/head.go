package heads

import "github.com/steve-care-software/historydb/domain/hash"

type head struct {
	hash    hash.Hash
	context string
	ret     string
}

func createHead(
	hash hash.Hash,
	context string,
	ret string,
) Head {
	out := head{
		hash:    hash,
		context: context,
		ret:     ret,
	}

	return &out
}

// Hash returns the hash
func (obj *head) Hash() hash.Hash {
	return obj.hash
}

// Context returns the context
func (obj *head) Context() string {
	return obj.context
}

// ReturnHash returns the return hash
func (obj *head) ReturnHash() string {
	return obj.ret
}
