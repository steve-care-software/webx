package applications

import (
	"errors"

	"github.com/steve-care-software/webx/engine/vms/applications/layers"
	route_applications "github.com/steve-care-software/webx/engine/vms/applications/routes"
	execution_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

type application struct {
	layerApp        layers.Application
	routeApp        route_applications.Application
	routeRepository routes.Repository
	batchSize       uint
}

func createApplication(
	layerApp layers.Application,
	routeApp route_applications.Application,
	routeRepository routes.Repository,
	batchSize uint,
) Application {
	out := application{
		layerApp:        layerApp,
		routeApp:        routeApp,
		routeRepository: routeRepository,
		batchSize:       batchSize,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte) (execution_layers.Execution, error) {
	pTotal, err := app.routeRepository.Amount()
	if err != nil {
		return nil, err
	}

	amountLoops := int(float64(*pTotal / app.batchSize))
	if amountLoops*int(app.batchSize) < int(*pTotal) {
		amountLoops++
	}

	for i := 0; i < amountLoops; i++ {
		index := uint(i) * app.batchSize
		amount := app.batchSize

		list, err := app.routeRepository.List(index, amount)
		if err != nil {
			return nil, err
		}

		for _, oneHash := range list {
			route, err := app.routeRepository.Retrieve(oneHash)
			if err != nil {
				return nil, err
			}

			retLayer, err := app.routeApp.Execute(input, route)
			if err != nil {
				continue
			}

			return app.layerApp.ExecuteWithInput(retLayer, input)
		}

	}

	return nil, errors.New("there is no Route that matches the provided input data")
}
