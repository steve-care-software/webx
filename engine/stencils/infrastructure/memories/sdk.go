package memories

import "github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"

// NewExecutionRepositoryAndService creates a new execution repository and service
func NewExecutionRepositoryAndService() (executions.Repository, executions.Service) {
	memory := map[string]map[string]executions.Execution{}
	builder := executions.NewBuilder()
	return createExecutionRepository(memory, builder), createExecutionService(memory)
}
