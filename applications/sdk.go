package applications

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
)

// Application represents the core application
type Application interface {
	Execute(input []byte) (executions.Executions, error)
	ExecuteWithContext(input []byte, context executions.Executions) (executions.Executions, error)
}
