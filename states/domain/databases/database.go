package databases

import (
	"github.com/steve-care-software/datastencil/states/domain/databases/commits"
	"github.com/steve-care-software/datastencil/states/domain/databases/metadatas"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type database struct {
	hash     hash.Hash
	head     commits.Commit
	metaData metadatas.MetaData
}

func createDatabase(
	hash hash.Hash,
	head commits.Commit,
	metaData metadatas.MetaData,
) Database {
	return &database{
		hash:     hash,
		head:     head,
		metaData: metaData,
	}
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head commit
func (obj *database) Head() commits.Commit {
	return obj.head
}

// MetaData returns the metadata
func (obj *database) MetaData() metadatas.MetaData {
	return obj.metaData
}
