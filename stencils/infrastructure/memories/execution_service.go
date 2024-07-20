package memories

import "github.com/steve-care-software/datastencil/stencils/domain/instances/executions"

type executionService struct {
}

func createExecutionService() executions.Service {
	out := executionService{}
	return &out
}

// Save saves an execution
func (app *executionService) Save(dbPath []string, ins executions.Execution) error {
	return nil
}
