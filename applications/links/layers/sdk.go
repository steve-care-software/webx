package layers

import (
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

// Application represents a layer application
type Application interface {
	ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Layer, error)
	Execute(layer layers.Layer) (execution_layers.Layer, error)
}
