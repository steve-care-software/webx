package applications

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers"
	route_applications "github.com/steve-care-software/webx/engine/vms/applications/routes"
	execution_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

func NewApplication(
	layerApp layers.Application,
	routeApp route_applications.Application,
	routeRepository routes.Repository,
	batchSize uint,
) Application {
	return createApplication(
		layerApp,
		routeApp,
		routeRepository,
		batchSize,
	)
}

// Application represents the application
type Application interface {
	Execute(input []byte) (execution_layers.Execution, error)
}
