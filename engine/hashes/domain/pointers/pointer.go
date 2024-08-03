package pointers

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type pointer struct {
	hash      hash.Hash
	delimiter delimiters.Delimiter
}

func createPointer(
	hash hash.Hash,
	delimiter delimiters.Delimiter,
) Pointer {
	out := pointer{
		hash:      hash,
		delimiter: delimiter,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Delimiter returns the pointer
func (obj *pointer) Delimiter() delimiters.Delimiter {
	return obj.delimiter
}
