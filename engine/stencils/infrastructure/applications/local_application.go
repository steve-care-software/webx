package applications

import (
	"errors"
	"fmt"

	db_applications "github.com/steve-care-software/webx/engine/states/applications"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
	applications_layers "github.com/steve-care-software/webx/engine/vms/applications"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
)

type localApplication struct {
	dbApp                db_applications.Application
	layerApp             applications_layers.Application
	layerAdapter         layers.Adapter
	executionsRepository executions.Repository
	executionsService    executions.Service
	executionsAdapter    executions.Adapter
	executionsBuilder    executions.Builder
	contextBuilder       contexts.Builder
	contextRepository    contexts.Repository
	contextService       contexts.Service
	executions           map[uint]context
}

func createLocalApplication(
	dbApp db_applications.Application,
	layerApp applications_layers.Application,
	layerAdapter layers.Adapter,
	executionsRepository executions.Repository,
	executionsService executions.Service,
	executionsAdapter executions.Adapter,
	executionsBuilder executions.Builder,
	contextBuilder contexts.Builder,
	contextRepository contexts.Repository,
	contextService contexts.Service,
) applications.Application {
	out := localApplication{
		dbApp:                dbApp,
		layerApp:             layerApp,
		layerAdapter:         layerAdapter,
		executionsRepository: executionsRepository,
		executionsService:    executionsService,
		executionsAdapter:    executionsAdapter,
		executionsBuilder:    executionsBuilder,
		contextBuilder:       contextBuilder,
		contextRepository:    contextRepository,
		contextService:       contextService,
		executions:           map[uint]context{},
	}

	return &out
}

// List lists the database paths
func (app *localApplication) List() ([][]string, error) {
	return nil, nil
}

// Init initializes a new database and begins a context on it
func (app *localApplication) Init(dbPath []string, name string, description string) (*uint, error) {
	pContext, err := app.dbApp.BeginWithInit(dbPath, name, description)
	if err != nil {
		return nil, err
	}

	app.executions[*pContext] = context{
		dbPath:     dbPath,
		executions: []hash.Hash{},
	}

	return pContext, nil
}

// Begin begins a context
func (app *localApplication) Begin(dbPath []string) (*uint, error) {
	pContext, err := app.dbApp.Begin(dbPath)
	if err != nil {
		return nil, err
	}

	// read the context:
	contextIns, err := app.contextRepository.Retrieve(dbPath)
	if err != nil {
		return nil, err
	}

	app.executions[*pContext] = context{
		dbPath:     dbPath,
		executions: contextIns.Executions(),
	}

	return pContext, nil
}

// Execute executes data on a context
func (app *localApplication) Execute(contextIdentifier uint, input []byte) ([]byte, error) {
	layer, err := app.layerAdapter.ToInstance(input)
	if err != nil {
		return nil, err
	}

	if currentContext, ok := app.executions[contextIdentifier]; ok {
		layerExecution, err := app.layerApp.Execute(layer)
		if err != nil {
			return nil, err
		}

		output, err := app.executionsAdapter.InstanceToBytes(layerExecution)
		if err != nil {
			return nil, err
		}

		// Execute executes in the database
		err = app.dbApp.Execute(contextIdentifier, input)
		if err != nil {
			return nil, err
		}

		// save the execution:
		err = app.executionsService.Save(currentContext.dbPath, layerExecution)
		if err != nil {
			return nil, err
		}

		app.executions[contextIdentifier] = context{
			dbPath:     currentContext.dbPath,
			executions: append(currentContext.executions, layerExecution.Hash()),
		}

		return output, nil
	}

	str := fmt.Sprintf(invalidPatternErr, contextIdentifier)
	return nil, errors.New(str)
}

// ExecuteWithPath reads the path and exectes the data on the context
func (app *localApplication) ExecuteWithPath(context uint, inputPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayer executes data with a layer on a context
func (app *localApplication) ExecuteLayer(context uint, input []byte, layerPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayerWithPath reads the path and exectes the data on the context using a layer path
func (app *localApplication) ExecuteLayerWithPath(context uint, inputPath []string, layerPath []string) ([]byte, error) {
	return nil, nil
}

// RetrieveAll retrieves all the executions of a context, between the index and length, if any
func (app *localApplication) RetrieveAll(contextIdentifier uint, index uint, length uint) (executions.Executions, error) {
	if currentContext, ok := app.executions[contextIdentifier]; ok {
		return app.executionsRepository.RetrieveAll(
			currentContext.dbPath,
			currentContext.executions,
		)
	}

	str := fmt.Sprintf(invalidPatternErr, contextIdentifier)
	return nil, errors.New(str)
}

// RetrieveAt retrieves an execution of the current context at index, if any
func (app *localApplication) RetrieveAt(context uint, index uint) (executions.Execution, error) {
	return nil, nil
}

// Amount returns the amount of executions in the current context
func (app *localApplication) Amount(context uint) (*uint, error) {
	return nil, nil
}

// Head returns the database head hash of the current context
func (app *localApplication) Head(context uint) (hash.Hash, error) {
	return nil, nil
}

// Commit commits executions to a context
func (app *localApplication) Commit(contextIdentifier uint) error {
	if currentContext, ok := app.executions[contextIdentifier]; ok {
		// commit to the database:
		err := app.dbApp.Commit(contextIdentifier)
		if err != nil {
			return err
		}

		// push to the database
		err = app.dbApp.Push(contextIdentifier)
		if err != nil {
			return err
		}

		builder := app.contextBuilder.Create().
			WithIdentifier(contextIdentifier).
			WithExecutions(currentContext.executions)

		// read the database:
		dbIns, err := app.dbApp.Retrieve(currentContext.dbPath)
		if err == nil {
			head := dbIns.Head().Hash()
			builder.WithHead(head)
		}

		contextIns, err := builder.Now()
		if err != nil {
			return err
		}

		return app.contextService.Save(currentContext.dbPath, contextIns)
	}

	str := fmt.Sprintf(invalidPatternErr, contextIdentifier)
	return errors.New(str)
}

// Rollback rollsback a context
func (app *localApplication) Rollback(context uint) error {
	return nil
}

// Cancel cancels a context
func (app *localApplication) Cancel(context uint) error {
	return nil
}

// Merge merges the context on the base context
func (app *localApplication) Merge(baseContext uint, content uint) error {
	return nil
}
