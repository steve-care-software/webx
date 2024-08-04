package applications

import (
	"reflect"

	"github.com/steve-care-software/webx/engine/entities/domain/entities"
	hash_applications "github.com/steve-care-software/webx/engine/hashes/applications"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type application struct {
	hashApp       hash_applications.Application
	entityAdapter entities.Adapter
}

func createApplication(
	hashApp hash_applications.Application,
	entityAdapter entities.Adapter,
) Application {
	out := application{
		hashApp:       hashApp,
		entityAdapter: entityAdapter,
	}

	return &out
}

// Begin begins the application
func (app *application) Begin(name string) (*uint, error) {
	return app.hashApp.Begin(name)
}

// Established returns true if the context exists, false otherwise
func (app *application) Established(identifier uint) bool {
	return app.hashApp.Established(identifier)
}

// Retrieve retrieves an entity by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (entities.Entity, error) {
	bytes, err := app.hashApp.Retrieve(context, hash)
	if err != nil {
		return nil, err
	}

	return app.entityAdapter.ToInstance(bytes)
}

// Insert inserts an entity
func (app *application) Insert(context uint, entity entities.Entity) error {
	refIns := reflect.TypeOf(entity)
	amount := refIns.NumMethod()
	for i := 0; i < amount; i++ {
		methodIns := refIns.Method(i)
		methodType := methodIns.Type
		amountParams := methodType.NumIn() - 1 // excluding the receiver
		if amountParams > 0 {
			continue
		}

		output := methodIns.Func.Call([]reflect.Value{})
		if len(output) != 1 {
			continue
		}

		value := output[0].Interface()
		if casted, ok := value.(entities.Entity); ok {
			return app.Insert(context, casted)
		}
	}

	return nil
}

// Delete deletes an entity
func (app *application) Delete(context uint, hash hash.Hash) error {
	return app.hashApp.Delete(context, hash)
}

// Commit executes a commit
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

// StatesAmount returns the amount of states
func (app *application) StatesAmount(context uint) (*uint, error) {
	return app.hashApp.StatesAmount(context)
}

// DeletedStateIndexes returns the deleted state indexes
func (app *application) DeletedStateIndexes(context uint) ([]uint, error) {
	return app.hashApp.DeletedStateIndexes(context)
}

// Close closes the context
func (app *application) Close(context uint) error {
	return app.hashApp.Close(context)
}

// Purge purges the context
func (app *application) Purge(context uint) error {
	return app.hashApp.Purge(context)
}
