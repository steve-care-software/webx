package layers

import (
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	Now() (Application, error)
}

// Application represents a layer application
type Application interface {
	Execute(layer layers.Layer) (execution_layers.Execution, error)
	ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Execution, error)
}
