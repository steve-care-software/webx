package pointers

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
)

type pointer struct {
	hash     hash.Hash
	head     hash.Hash
	metaData metadatas.MetaData
}

func createPointer(
	hash hash.Hash,
	head hash.Hash,
	metaData metadatas.MetaData,
) Pointer {
	out := pointer{
		hash:     hash,
		head:     head,
		metaData: metaData,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *pointer) Head() hash.Hash {
	return obj.head
}

// MetaData returns the metaData
func (obj *pointer) MetaData() metadatas.MetaData {
	return obj.metaData
}
