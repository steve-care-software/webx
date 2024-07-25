package applications

import (
	applications_layer "github.com/steve-care-software/webx/engine/vms/applications/layers"
	route_applications "github.com/steve-care-software/webx/engine/vms/applications/routes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

func NewApplication(
	routeApp route_applications.Application,
	layerApp applications_layer.Application,
	layerRepository layers.Repository,
	routeRepository routes.Repository,
	batchSize uint,
) Application {
	executionsBuilder := executions.NewBuilder()
	return createApplication(
		routeApp,
		layerApp,
		layerRepository,
		routeRepository,
		executionsBuilder,
		batchSize,
	)
}

// Application represents the application
type Application interface {
	Execute(input []byte) (executions.Executions, error)
}
