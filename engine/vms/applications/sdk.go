package applications

import (
	execution_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

// Application represents the application
type Application interface {
	Init() error // load the initial routes
	Execute(input []byte) (execution_layers.Execution, error)
}
