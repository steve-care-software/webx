package inserts

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

type insert struct {
	hash      hash.Hash
	name      string
	bytes     []byte
	whitelist []hash.Hash
}

func createInsert(
	hash hash.Hash,
	name string,
	bytes []byte,
	whitelist []hash.Hash,
) Insert {
	out := insert{
		hash:      hash,
		name:      name,
		bytes:     bytes,
		whitelist: whitelist,
	}

	return &out
}

// Hash returns the hash
func (obj *insert) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *insert) Name() string {
	return obj.name
}

// Bytes returns the bytes
func (obj *insert) Bytes() []byte {
	return obj.bytes
}

// Whitelist returns the whitelist
func (obj *insert) Whitelist() []hash.Hash {
	return obj.whitelist
}
