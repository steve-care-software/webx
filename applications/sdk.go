package applications

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources"
)

// Application represents the core application
type Application interface {
	Execute(input []byte, resources resources.Resources) (executions.Executions, error)
	ExecuteWithContext(input []byte, resources resources.Resources, context executions.Executions) (executions.Executions, error)
}
