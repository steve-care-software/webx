package infrastructure

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	entities_applications "github.com/steve-care-software/webx/engine/databases/entities/applications"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/entities"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/stencils/domain/sessions"
	applications_vms "github.com/steve-care-software/webx/engine/vms/applications"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

type localApplication struct {
	entityApp      entities_applications.Application
	vmApp          applications_vms.Application
	entityAdapter  entities.Adapter
	sessionBuilder sessions.Builder
	hashAdapter    hash.Adapter
	contexts       map[string]uint
}

func createLocalApplication(
	entityApp entities_applications.Application,
	vmApp applications_vms.Application,
	entityAdapter entities.Adapter,
	sessionBuilder sessions.Builder,
	hashAdapter hash.Adapter,
) applications.Application {
	out := localApplication{
		entityApp:      entityApp,
		vmApp:          vmApp,
		entityAdapter:  entityAdapter,
		sessionBuilder: sessionBuilder,
		hashAdapter:    hashAdapter,
		contexts:       map[string]uint{},
	}

	return &out
}

// Begin begins a context
func (app *localApplication) Begin(keyname string) (hash.Hash, error) {
	pContext, err := app.entityApp.Begin(keyname)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
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

		prevSession, err := app.retrieveSession(context, identifier)
		if err != nil {
			return nil, err
		}

		executionsList := []executions.Executions{}
		if prevSession != nil {
			executionsList = prevSession.Executions()
		}

		executionsList = append(executionsList, executionsIns)
		updatedSession, err := app.sessionBuilder.Create().
			WithExecutions(executionsList).
			WithHash(identifier).
			Now()

		if err != nil {
			return nil, err
		}

		err = app.entityApp.Delete(context, identifier)
		if err != nil {
			return nil, err
		}

		err = app.entityApp.Insert(context, updatedSession)
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
		return app.entityApp.Commit(context)
	}

	return errors.New(contextIdentifierUndefinedPattern)
}

// Session returns the session
func (app *localApplication) Session(identifier hash.Hash) (sessions.Session, error) {
	keyname := identifier.String()
	if context, ok := app.contexts[keyname]; ok {
		retSession, err := app.entityApp.Retrieve(context, identifier)
		if err != nil {
			return nil, err
		}

		if casted, ok := retSession.(sessions.Session); ok {
			return casted, nil
		}

		str := fmt.Sprintf(notSessionErrPattern, identifier.String())
		return nil, errors.New(str)
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
		return app.entityApp.Close(context)
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

func (app *localApplication) retrieveSession(context uint, identifier hash.Hash) (sessions.Session, error) {
	instance, err := app.entityApp.Retrieve(context, identifier)
	if err != nil {
		return nil, nil
	}

	if casted, ok := instance.(sessions.Session); ok {
		return casted, nil
	}

	str := fmt.Sprintf(notSessionErrPattern, identifier.String())
	return nil, errors.New(str)
}
