package applications

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	containers_applications "github.com/steve-care-software/webx/engine/containers/applications"
	entities_applications "github.com/steve-care-software/webx/engine/entities/applications"
	"github.com/steve-care-software/webx/engine/entities/domain/entities"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	applications_vms "github.com/steve-care-software/webx/engine/vms/applications"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

type localApplication struct {
	entityApp     entities_applications.Application
	containerApp  containers_applications.Application
	vmApp         applications_vms.Application
	entityAdapter entities.Adapter
	hashAdapter   hash.Adapter
	contexts      map[string]uint
}

func createLocalApplication(
	entityApp entities_applications.Application,
	containerApp containers_applications.Application,
	vmApp applications_vms.Application,
	entityAdapter entities.Adapter,
	hashAdapter hash.Adapter,
) applications.Application {
	out := localApplication{
		entityApp:     entityApp,
		containerApp:  containerApp,
		vmApp:         vmApp,
		entityAdapter: entityAdapter,
		hashAdapter:   hashAdapter,
		contexts:      map[string]uint{},
	}

	return &out
}

// Begin begins a context
func (app *localApplication) Begin(keyname string) (hash.Hash, error) {
	pContext, err := app.entityApp.Begin(keyname)
	if err != nil {
		return nil, err
	}

	if !app.containerApp.Established(*pContext) {
		str := fmt.Sprintf("the container application was expected to be already established for context: %d", *pContext)
		return nil, errors.New(str)
	}

	number := rand.Int()
	pHash, err := app.hashAdapter.FromBytes([]byte(strconv.Itoa(number)))
	if err != nil {
		return nil, err
	}

	app.contexts[pHash.String()] = *pContext
	return *pHash, nil
}

// Execute executes the application
func (app *localApplication) Execute(identifier hash.Hash, input []byte) ([]byte, error) {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		executionsIns, err := app.vmApp.Execute(input)
		if err != nil {
			return nil, err
		}

		hashes := []hash.Hash{}
		executionsList := executionsIns.List()
		for _, oneExecution := range executionsList {
			hashes = append(hashes, oneExecution.Hash())
			err = app.entityApp.Insert(context, oneExecution)
			if err != nil {
				return nil, err
			}
		}

		// insert the elements in the container:
		err = app.containerApp.Insert(context, identifier.String(), hashes)
		if err != nil {
			return nil, err
		}

		return app.entityAdapter.ToBytes(executionsIns)
	}

	return nil, errors.New(contextIdentifierUndefinedPattern)
}

// Commit commits in the database
func (app *localApplication) Commit(identifier hash.Hash) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		err := app.entityApp.Commit(context)
		if err != nil {
			return err
		}

		if app.containerApp.Established(context) {
			str := fmt.Sprintf("the container application was expected to NOT be established for context: %d", context)
			return errors.New(str)
		}

		return nil
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// Executions retrieves the past executions of the identifier
func (app *localApplication) Executions(identifier hash.Hash) ([]executions.Executions, error) {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		// retrieve the container:
		retContainer, err := app.containerApp.Retrieve(context, identifier.String())
		if err != nil {
			return nil, err
		}

		executionsList := []executions.Executions{}
		elements := retContainer.Elements()
		for _, oneHash := range elements {
			retExecutions, err := app.entityApp.Retrieve(context, oneHash)
			if err != nil {
				return nil, err
			}

			if casted, ok := retExecutions.(executions.Executions); ok {
				executionsList = append(executionsList, casted)
				continue
			}

			str := fmt.Sprintf("the context (%d) is trying to retrieving a container (hash: %s) that was expected to contain an Executions instance (hash: %s), but the data could not be catsed properly", context, identifier.String(), oneHash.String())
			return nil, errors.New(str)
		}

		return executionsList, nil
	}

	return nil, errors.New(contextIdentifierUndefinedPattern)
}

// Sessions returns the list of sessions
func (app *localApplication) Sessions() ([]hash.Hash, error) {
	list := []hash.Hash{}
	for hashStr := range app.contexts {
		pHash, err := app.hashAdapter.FromString(hashStr)
		if err != nil {
			return nil, err
		}

		list = append(list, *pHash)
	}

	return list, nil
}

// Delete deletes the executions of a session
func (app *localApplication) Delete(identifier hash.Hash) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		err := app.entityApp.Delete(context, identifier)
		if err != nil {
			return err
		}

		delete(app.contexts, keyname)
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// DeleteState deletes a state
func (app *localApplication) DeleteState(identifier hash.Hash, stateIndex uint) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		return app.entityApp.DeleteState(context, stateIndex)
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// RecoverState recovers the state
func (app *localApplication) RecoverState(identifier hash.Hash, stateIndex uint) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		return app.entityApp.RecoverState(context, stateIndex)
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// StatesAmount returns the states amount
func (app *localApplication) StatesAmount(identifier hash.Hash) (*uint, error) {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		return app.entityApp.StatesAmount(context)
	}

	return nil, errors.New(contextIdentifierUndefinedPattern)
}

// DeletedStateIndexes returns the deletes state indexes
func (app *localApplication) DeletedStateIndexes(identifier hash.Hash) ([]uint, error) {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		return app.entityApp.DeletedStateIndexes(context)
	}

	return nil, errors.New(contextIdentifierUndefinedPattern)
}

// Close closes the context
func (app *localApplication) Close(identifier hash.Hash) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		err := app.entityApp.Close(context)
		if err != nil {
			return err
		}

		if app.containerApp.Established(context) {
			str := fmt.Sprintf("the container application was expected to NOT be established for context: %d", context)
			return errors.New(str)
		}

		return nil
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// Purge purges the context
func (app *localApplication) Purge(identifier hash.Hash) error {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		return app.entityApp.Purge(context)
	}

	return errors.New(contextIdentifierUndefinedPattern)
}
