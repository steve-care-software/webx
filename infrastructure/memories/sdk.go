package memories

import "github.com/steve-care-software/datastencil/domain/instances/executions"

// NewExecutionRepository creates a new execution repository
func NewExecutionRepository() executions.Repository {
	return createExecutionRepository()
}

// NewExecutionService creates a new execution service
func NewExecutionService() executions.Service {
	return createExecutionService()
}
