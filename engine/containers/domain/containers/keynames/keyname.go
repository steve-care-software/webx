package keynames

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type keyname struct {
	hash hash.Hash
	name string
}

func createKeyname(
	hash hash.Hash,
	name string,
) Keyname {
	out := keyname{
		hash: hash,
		name: name,
	}

	return &out
}

// Hash returns the hash
func (obj *keyname) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *keyname) Name() string {
	return obj.name
}
