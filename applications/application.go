package applications

import (
	"errors"
	"fmt"

	applications_layers "github.com/steve-care-software/datastencil/applications/layers"
	"github.com/steve-care-software/datastencil/domain/contexts"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/historydb/applications"
	"github.com/steve-care-software/historydb/domain/hash"
)

type application struct {
	dbApp                applications.Application
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

func createApplication(
	dbApp applications.Application,
	layerApp applications_layers.Application,
	layerAdapter layers.Adapter,
	executionsRepository executions.Repository,
	executionsService executions.Service,
	executionsAdapter executions.Adapter,
	executionsBuilder executions.Builder,
	contextBuilder contexts.Builder,
	contextRepository contexts.Repository,
	contextService contexts.Service,
) Application {
	out := application{
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
func (app *application) List() ([][]string, error) {
	return nil, nil
}

// Init initializes a new database and begins a context on it
func (app *application) Init(dbPath []string, name string, description string) (*uint, error) {
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
func (app *application) Begin(dbPath []string) (*uint, error) {
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
func (app *application) Execute(contextIdentifier uint, input []byte) ([]byte, error) {
	layer, err := app.layerAdapter.ToInstance(input)
	if err != nil {
		return nil, err
	}

	if currentContext, ok := app.executions[contextIdentifier]; ok {
		layerExecution, err := app.layerApp.Execute(layer)
		if err != nil {
			return nil, err
		}

		output, err := app.executionsAdapter.ToBytes(layerExecution)
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
func (app *application) ExecuteWithPath(context uint, inputPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayer executes data with a layer on a context
func (app *application) ExecuteLayer(context uint, input []byte, layerPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayerWithPath reads the path and exectes the data on the context using a layer path
func (app *application) ExecuteLayerWithPath(context uint, inputPath []string, layerPath []string) ([]byte, error) {
	return nil, nil
}

// RetrieveAll retrieves all the executions of a context, between the index and length, if any
func (app *application) RetrieveAll(contextIdentifier uint, index uint, length uint) (executions.Executions, error) {
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
func (app *application) RetrieveAt(context uint, index uint) (executions.Execution, error) {
	return nil, nil
}

// Amount returns the amount of executions in the current context
func (app *application) Amount(context uint) (*uint, error) {
	return nil, nil
}

// Head returns the database head hash of the current context
func (app *application) Head(context uint) (hash.Hash, error) {
	return nil, nil
}

// Commit commits executions to a context
func (app *application) Commit(contextIdentifier uint) error {
	if currentContext, ok := app.executions[contextIdentifier]; ok {
		// read the database:
		dbIns, err := app.dbApp.Retrieve(currentContext.dbPath)
		if err != nil {
			return err
		}

		// commit to the database:
		err = app.dbApp.Commit(contextIdentifier)
		if err != nil {
			return err
		}

		head := dbIns.Head().Hash()
		contextIns, err := app.contextBuilder.Create().
			WithIdentifier(contextIdentifier).
			WithHead(head).
			WithExecutions(currentContext.executions).
			Now()

		if err != nil {
			return err
		}

		return app.contextService.Save(contextIns)
	}

	str := fmt.Sprintf(invalidPatternErr, contextIdentifier)
	return errors.New(str)
}

// Rollback rollsback a context
func (app *application) Rollback(context uint) error {
	return nil
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	return nil
}

// Merge merges the context on the base context
func (app *application) Merge(baseContext uint, content uint) error {
	return nil
}
