package applications

import (
	"errors"

	layer_applications "github.com/steve-care-software/webx/engine/vms/applications/layers"
	route_applications "github.com/steve-care-software/webx/engine/vms/applications/routes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

type application struct {
	routeApp          route_applications.Application
	layerApp          layer_applications.Application
	layerRepository   layers.Repository
	routeRepository   routes.Repository
	executionsBuilder executions.Builder
	batchSize         uint
}

func createApplication(
	routeApp route_applications.Application,
	layerApp layer_applications.Application,
	layerRepository layers.Repository,
	routeRepository routes.Repository,
	executionsBuilder executions.Builder,
	batchSize uint,
) Application {
	out := application{
		routeApp:          routeApp,
		layerApp:          layerApp,
		layerRepository:   layerRepository,
		routeRepository:   routeRepository,
		executionsBuilder: executionsBuilder,
		batchSize:         batchSize,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte) (executions.Executions, error) {
	lastRemaining := []byte{}
	executionsList := []executions.Execution{}
	for {
		retExecution, retRemaining, err := app.executeOnce(input)
		if err != nil {
			return nil, err
		}

		executionsList = append(executionsList, retExecution)
		if len(retRemaining) < len(lastRemaining) {
			continue
		}

		break
	}

	return app.executionsBuilder.Create().
		WithList(executionsList).
		Now()
}

func (app *application) executeOnce(input []byte) (executions.Execution, []byte, error) {
	pTotal, err := app.routeRepository.Amount()
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}

		for _, oneHash := range list {
			route, err := app.routeRepository.Retrieve(oneHash)
			if err != nil {
				return nil, nil, err
			}

			// verify if the route has a match against the data
			isMatch, remaining, err := app.routeApp.Execute(input, route)
			if err != nil {
				return nil, nil, err
			}

			// if there is no match, continue
			if !isMatch {
				continue
			}

			hash := route.Layer()
			retLayer, err := app.layerRepository.Retrieve(hash)
			if err != nil {
				continue
			}

			retExecution, err := app.layerApp.ExecuteWithInput(retLayer, input)
			if err != nil {
				return nil, nil, err
			}

			return retExecution, remaining, nil
		}

	}

	return nil, nil, errors.New("there is no Route that matches the provided input data")
}
