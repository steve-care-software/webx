package logics

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	execution_links "github.com/steve-care-software/datastencil/domain/instances/executions/links"
	"github.com/steve-care-software/datastencil/domain/resources/logics"
)

// Application represents an application
type Application interface {
	Execute(input []byte, logic logics.Logic) (execution_links.Link, error)
	ExecuteWithContext(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error)
}
