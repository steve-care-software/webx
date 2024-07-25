package applications

import (
	applications_layer "github.com/steve-care-software/webx/engine/vms/applications/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

func NewApplication(
	layerApp applications_layer.Application,
	layerRepository layers.Repository,
	routeRepository routes.Repository,
	batchSize uint,
) Application {
	executionsBuilder := executions.NewBuilder()
	return createApplication(
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
