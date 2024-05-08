package updates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

type update struct {
	hash   hash.Hash
	delete deletes.Delete
	bytes  []byte
}

func createUpdate(
	hash hash.Hash,
	delete deletes.Delete,
	bytes []byte,
) Update {
	out := update{
		hash:   hash,
		delete: delete,
		bytes:  bytes,
	}

	return &out
}

// Hash returns the hash
func (obj *update) Hash() hash.Hash {
	return obj.hash
}

// Delete returns the delete
func (obj *update) Delete() deletes.Delete {
	return obj.delete
}

// Bytes returns the bytes
func (obj *update) Bytes() []byte {
	return obj.bytes
}
