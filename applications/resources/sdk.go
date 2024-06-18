package resources

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/resources"
)

// Application represents a resources application
type Application interface {
	Execute(path []string) (resources.Resources, error)
	ExecuteWithContext(path []string, context executions.Executions) (resources.Resources, error)
}
