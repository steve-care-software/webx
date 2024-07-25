package applications

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	execution_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

// Application represents the application
type Application interface {
	Init() error // load the initial routes
	Execute(input []byte) (execution_layers.Execution, error)
	Insert(route routes.Route) error
	Delete(hash hash.Hash) error
}
