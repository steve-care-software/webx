package logics

import (
	"github.com/steve-care-software/datastencil/applications/logics/layers"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	execution_links "github.com/steve-care-software/datastencil/domain/instances/executions/links"
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics"
	bridged_layers "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions"
)

type application struct {
	layerApplication       layers.Application
	executionLayersBuilder execution_layers.Builder
	executionLayerBuilder  execution_layers.LayerBuilder
	executionLinkBuilder   execution_links.Builder
}

func createApplication(
	layerApplication layers.Application,
	executionLayersBuilder execution_layers.Builder,
	executionLayerBuilder execution_layers.LayerBuilder,
	executionLinkBuilder execution_links.Builder,
) Application {
	out := application{
		layerApplication:       layerApplication,
		executionLayersBuilder: executionLayersBuilder,
		executionLayerBuilder:  executionLayerBuilder,
		executionLinkBuilder:   executionLinkBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte, logic logics.Logic) (execution_links.Link, error) {
	return app.execute(input, logic, nil)
}

// Execute executes the application with context
func (app *application) ExecuteWithContext(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error) {
	return app.execute(input, logic, context)
}

func (app *application) execute(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error) {
	bridges := logic.Bridges()
	executedLayersList := []execution_layers.Layer{}
	link := logic.Link()
	elementsList := link.Elements().List()
	for _, oneElement := range elementsList {
		layerPath := oneElement.Layer()
		bridge, err := bridges.Fetch(layerPath)
		if err != nil {
			return nil, err
		}

		layer := bridge.Layer()
		retResult, err := app.executeLayer(input, layer, context)
		if err != nil {
			return nil, err
		}

		executedLayer, err := app.executionLayerBuilder.Create().WithInput(input).WithSource(layer).WithResult(retResult).Now()
		if err != nil {
			return nil, err
		}

		executedLayersList = append(executedLayersList, executedLayer)
		if retResult.IsSuccess() {
			continue
		}

		interruption := retResult.Interruption()
		if interruption.IsStop() {
			break
		}

		failure := interruption.Failure()
		code := failure.Code()
		isRaisedInLayer := failure.IsRaisedInLayer()
		condition := oneElement.Condition()
		bContinue, err := app.respectCondition(condition, code, isRaisedInLayer)
		if err != nil {
			return nil, err
		}

		if !bContinue {
			break
		}
	}

	executedLayers, err := app.executionLayersBuilder.Create().
		WithList(executedLayersList).
		Now()

	if err != nil {
		return nil, err
	}

	return app.executionLinkBuilder.Create().
		WithInput(input).
		WithLayers(executedLayers).
		WithSource(link).
		Now()
}

func (app *application) respectCondition(exepectedCondition conditions.Condition, code uint, isRaisedInLayer bool) (bool, error) {
	return true, nil
}

func (app *application) executeLayer(input []byte, layer bridged_layers.Layer, context executions.Executions) (results.Result, error) {
	if context != nil {
		return app.layerApplication.ExecuteWithContext(input, layer, context)
	}

	return app.layerApplication.Execute(input, layer)
}
