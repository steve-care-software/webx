package layers

import (
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

// Application represents a layer application
type Application interface {
	Execute(layer layers.Layer) (execution_layers.Execution, error)
	ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Execution, error)
}
