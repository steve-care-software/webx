package executions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
)

type execution struct {
	hash     hash.Hash
	logic    links.Link
	database databases.Database
}

func createExecution(
	hash hash.Hash,
	logic links.Link,
	database databases.Database,
) Execution {
	out := execution{
		hash:     hash,
		logic:    logic,
		database: database,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Logic returns the logic
func (obj *execution) Logic() links.Link {
	return obj.logic
}

// Database returns the database
func (obj *execution) Database() databases.Database {
	return obj.database
}
