package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/containers/domain/containers"
	"github.com/steve-care-software/webx/engine/containers/domain/containers/keynames"
	hash_applications "github.com/steve-care-software/webx/engine/hashes/applications"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type application struct {
	hashApp          hash_applications.Application
	containerAdapter containers.Adapter
	containerBuilder containers.Builder
	keynameBuilder   keynames.Builder
}

func createApplication(
	hashApp hash_applications.Application,
	containerAdapter containers.Adapter,
	containerBuilder containers.Builder,
	keynameBuilder keynames.Builder,
) Application {
	out := application{
		hashApp:          hashApp,
		containerAdapter: containerAdapter,
		containerBuilder: containerBuilder,
		keynameBuilder:   keynameBuilder,
	}

	return &out
}

// Begin begins the context
func (app *application) Begin(name string) (*uint, error) {
	pContext, err := app.hashApp.Begin(name)
	if err != nil {
		return nil, err
	}

	return pContext, nil
}

// Established returns true if the context exists, false otherwise
func (app *application) Established(context uint) bool {
	return app.hashApp.Established(context)
}

// Retrieve retrieves a container
func (app *application) Retrieve(context uint, keyname string) (containers.Container, error) {
	// build the keyname:
	keynameIns, err := app.keynameBuilder.Create().WithName(keyname).Now()
	if err != nil {
		return nil, err
	}

	// retrieve the container bytes:
	retBytes, err := app.hashApp.Retrieve(context, keynameIns.Hash())
	if err != nil {
		return nil, err
	}

	// convert the container:
	return app.containerAdapter.ToInstance(retBytes)
}

// Amount returns the amount of elements a container contains
func (app *application) Amount(context uint, keyname string) (*uint, error) {
	// retrieve the container:
	retContainer, err := app.Retrieve(context, keyname)
	if err != nil {
		return nil, err
	}

	amount := uint(len(retContainer.Elements()))
	return &amount, nil
}

// List returns the hashes related to the container
func (app *application) List(context uint, keyname string, index uint64, length uint64) ([]hash.Hash, error) {
	// retrieve the container:
	retContainer, err := app.Retrieve(context, keyname)
	if err != nil {
		return nil, err
	}

	retList, err := retContainer.Subset(index, length)
	if err != nil {
		return nil, err
	}

	return retList, nil
}

// Insert inserts values in a container
func (app *application) Insert(context uint, keyname string, values []hash.Hash) error {
	// retrieve the container:
	retContainer, err := app.Retrieve(context, keyname)
	if err != nil {
		return err
	}

	// update the container:
	elements := retContainer.Elements()
	elements = append(elements, values...)
	return app.update(context, keyname, elements)
}

// Delete delete values from a container
func (app *application) Delete(context uint, keyname string, index uint64, length uint64) error {
	// retrieve the container:
	retContainer, err := app.Retrieve(context, keyname)
	if err != nil {
		return err
	}

	elements := retContainer.Elements()
	amount := uint64(len(elements))
	if amount <= index {
		str := fmt.Sprintf("the comtainer (%s) contains %d elements, the index (%d) is therefore invalid", keyname, amount, index)
		return errors.New(str)
	}

	total := index + length
	if total > amount {
		str := fmt.Sprintf("the comtainer (%s) contains %d elements, therefore index + length (%d) is invalid", keyname, amount, total)
		return errors.New(str)
	}

	subset := elements[index:length]
	return app.update(context, keyname, subset)
}

// Remove removes a container
func (app *application) Remove(context uint, keyname string) error {
	// build the keyname:
	keynameIns, err := app.keynameBuilder.Create().WithName(keyname).Now()
	if err != nil {
		return err
	}

	return app.hashApp.Delete(context, keynameIns.Hash())
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return app.hashApp.Commit(context)
}

// DeleteState deletes a state
func (app *application) DeleteState(context uint, stateIndex uint) error {
	return app.hashApp.DeleteState(context, stateIndex)
}

// RecoverState recovers a state
func (app *application) RecoverState(context uint, stateIndex uint) error {
	return app.hashApp.RecoverState(context, stateIndex)
}

// StatesAmount returns the state amount
func (app *application) StatesAmount(context uint) (*uint, error) {
	return app.hashApp.StatesAmount(context)
}

// DeletedStateIndexes deletes state indexes
func (app *application) DeletedStateIndexes(context uint) ([]uint, error) {
	return app.hashApp.DeletedStateIndexes(context)
}

// Close closes a context
func (app *application) Close(context uint) error {
	return app.hashApp.Close(context)
}

// Purge purges a context
func (app *application) Purge(context uint) error {
	return app.hashApp.Purge(context)
}

func (app *application) update(context uint, keyname string, elements []hash.Hash) error {
	// update the container:
	updated, err := app.containerBuilder.Create().WithElements(elements).WithKeyname(keyname).Now()
	if err != nil {
		return err
	}

	// delete the container:
	keynameHash := updated.Keyname().Hash()
	err = app.hashApp.Delete(context, keynameHash)
	if err != nil {
		return err
	}

	// convert the container to bytes:
	retBytes, err := app.containerAdapter.ToBytes(updated)
	if err != nil {
		return err
	}

	// insert the updated container:
	return app.hashApp.Insert(context, keynameHash, retBytes)
}
