package links

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
)

// Application represents a link application
type Application interface {
	Execute(input []byte, elements elements.Elements, references map[string]instances.Instance) (executions.Executions, error)
	ExecuteWithContext(input []byte, elements elements.Elements, references map[string]instances.Instance, context executions.Executions) (executions.Executions, error)
}
