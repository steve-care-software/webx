package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers"
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/hash"
)

type execution struct {
	hash     hash.Hash
	layer    layers.Layer
	database databases.Database
}

func createExecution(
	hash hash.Hash,
	layer layers.Layer,
	database databases.Database,
) Execution {
	out := execution{
		hash:     hash,
		layer:    layer,
		database: database,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *execution) Layer() layers.Layer {
	return obj.layer
}

// Database returns the database
func (obj *execution) Database() databases.Database {
	return obj.database
}
