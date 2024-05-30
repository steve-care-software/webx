package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics"
)

type resource struct {
	hash     hash.Hash
	database heads.Head
	logics   logics.Logics
}

func createResource(
	hash hash.Hash,
	database heads.Head,
	logics logics.Logics,
) Resource {
	out := resource{
		hash:     hash,
		database: database,
		logics:   logics,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Database returns the database
func (obj *resource) Database() heads.Head {
	return obj.database
}

// Logics returns the logics
func (obj *resource) Logics() logics.Logics {
	return obj.logics
}
