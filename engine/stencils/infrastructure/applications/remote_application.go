package applications

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
)

type remoteApplication struct {
}

func createRemoteApplication() applications.Application {
	out := remoteApplication{}
	return &out
}

// List lists the database paths
func (app *remoteApplication) List() ([][]string, error) {
	return nil, nil
}

// Init initializes a new database and begins a context on it
func (app *remoteApplication) Init(dbPath []string, name string, description string) (*uint, error) {
	return nil, nil
}

// Begin begins a context
func (app *remoteApplication) Begin(dbPath []string) (*uint, error) {
	return nil, nil
}

// Execute executes data on a context
func (app *remoteApplication) Execute(contextIdentifier uint, input []byte) ([]byte, error) {
	return nil, nil
}

// ExecuteWithPath reads the path and exectes the data on the context
func (app *remoteApplication) ExecuteWithPath(context uint, inputPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayer executes data with a layer on a context
func (app *remoteApplication) ExecuteLayer(context uint, input []byte, layerPath []string) ([]byte, error) {
	return nil, nil
}

// ExecuteLayerWithPath reads the path and exectes the data on the context using a layer path
func (app *remoteApplication) ExecuteLayerWithPath(context uint, inputPath []string, layerPath []string) ([]byte, error) {
	return nil, nil
}

// RetrieveAll retrieves all the executions of a context, between the index and length, if any
func (app *remoteApplication) RetrieveAll(contextIdentifier uint, index uint, length uint) (executions.Executions, error) {
	return nil, nil
}

// RetrieveAt retrieves an execution of the current context at index, if any
func (app *remoteApplication) RetrieveAt(context uint, index uint) (executions.Execution, error) {
	return nil, nil
}

// Amount returns the amount of executions in the current context
func (app *remoteApplication) Amount(context uint) (*uint, error) {
	return nil, nil
}

// Head returns the database head hash of the current context
func (app *remoteApplication) Head(context uint) (hash.Hash, error) {
	return nil, nil
}

// Commit commits executions to a context
func (app *remoteApplication) Commit(contextIdentifier uint) error {
	return nil
}

// Rollback rollsback a context
func (app *remoteApplication) Rollback(context uint) error {
	return nil
}

// Cancel cancels a context
func (app *remoteApplication) Cancel(context uint) error {
	return nil
}

// Merge merges the context on the base context
func (app *remoteApplication) Merge(baseContext uint, content uint) error {
	return nil
}
