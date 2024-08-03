package pointers

import (
	bytes_pointers "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type pointer struct {
	hash    hash.Hash
	pointer bytes_pointers.Pointer
}

func createPointer(
	hash hash.Hash,
	ptr bytes_pointers.Pointer,
) Pointer {
	out := pointer{
		hash:    hash,
		pointer: ptr,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Pointer returns the pointer
func (obj *pointer) Pointer() bytes_pointers.Pointer {
	return obj.pointer
}
