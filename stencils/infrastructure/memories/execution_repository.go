package memories

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions"
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
