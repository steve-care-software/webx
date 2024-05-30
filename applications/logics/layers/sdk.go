package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

// Application represents an application
type Application interface {
	Execute(input []byte, layer layers.Layer) (results.Result, error)
	ExecuteWithContext(input []byte, layer layers.Layer, context executions.Executions) (results.Result, error)
}
