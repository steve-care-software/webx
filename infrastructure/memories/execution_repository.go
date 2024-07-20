package memories

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/historydb/domain/hash"
)

type executionRepository struct {
}

func createExecutionRepository() executions.Repository {
	out := executionRepository{}
	return &out
}

// RetrieveAll retrieves executions
func (app *executionRepository) RetrieveAll(dbPath []string, hashes []hash.Hash) (executions.Executions, error) {
	return nil, nil
}
