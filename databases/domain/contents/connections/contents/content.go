package contents

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type content struct {
	hash hash.Hash
	data []byte
	kind uint
}

func createContent(
	hash hash.Hash,
	data []byte,
	kind uint,
) Content {
	out := content{
		hash: hash,
		data: data,
		kind: kind,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Data returns the data
func (obj *content) Data() []byte {
	return obj.data
}

// Kind returns the kind
func (obj *content) Kind() uint {
	return obj.kind
}
