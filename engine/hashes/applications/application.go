package applications

import (
	"errors"
	"fmt"

	bytes_applications "github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

type application struct {
	bytesApp         bytes_applications.Application
	pointerAdapter   pointers.Adapter
	pointersBuilder  pointers.Builder
	pointerBuilder   pointers.PointerBuilder
	delimiterBuilder delimiters.DelimiterBuilder
	contexts         map[uint]*context
}

func createApplication(
	bytesApp bytes_applications.Application,
	pointerAdapter pointers.Adapter,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
) Application {
	out := application{
		bytesApp:         bytesApp,
		pointerAdapter:   pointerAdapter,
		pointersBuilder:  pointersBuilder,
		pointerBuilder:   pointerBuilder,
		delimiterBuilder: delimiterBuilder,
		contexts:         map[uint]*context{},
	}

	return &out
}

// Begin begins a context
func (app *application) Begin(name string) (*uint, error) {
	pContext, err := app.bytesApp.Begin(name)
	if err != nil {
		return nil, err
	}

	states, err := app.bytesApp.States(*pContext)
	if err != nil {
		return nil, err
	}

	var currentPointers pointers.Pointers
	if states.HasRoot() {
		retBytes, err := app.bytesApp.Retrieve(*pContext, states.Root())
		if err != nil {
			return nil, err
		}

		currentPointers, _, err = app.pointerAdapter.BytesToInstances(retBytes)
		if err != nil {
			return nil, err
		}
	}

	app.contexts[*pContext] = &context{
		current: currentPointers,
		inserts: []pointers.Pointer{},
		deletes: []hash.Hash{},
	}

	return pContext, nil
}

// Retrieve retrieves data from an hash
func (app *application) Retrieve(identifier uint, hash hash.Hash) ([]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		if pContext.current == nil {
			return nil, errors.New("there is zero entry in our database")
		}

		pointer, err := pContext.current.Retrieve(hash)
		if err != nil {
			return nil, err
		}

		delimiter := pointer.Delimiter()
		return app.bytesApp.Retrieve(identifier, delimiter)
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// Insert inserts data
func (app *application) Insert(identifier uint, hash hash.Hash, data []byte) error {
	if pContext, ok := app.contexts[identifier]; ok {
		delimiter, err := app.bytesApp.Insert(identifier, data)
		if err != nil {
			return err
		}

		pointer, err := app.pointerBuilder.Create().WithDelimiter(delimiter).WithHash(hash).Now()
		if err != nil {
			return err
		}

		inserts := append(pContext.inserts, pointer)
		app.contexts[identifier] = &context{
			current: pContext.current,
			inserts: inserts,
			deletes: pContext.deletes,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Delete deletes data
func (app *application) Delete(identifier uint, hash hash.Hash) error {
	if pContext, ok := app.contexts[identifier]; ok {
		pointer, err := pContext.current.Retrieve(hash)
		if err != nil {
			return err
		}

		delimiter := pointer.Delimiter()
		err = app.bytesApp.Delete(identifier, delimiter)
		if err != nil {
			return err
		}

		deletes := append(pContext.deletes, hash)
		app.contexts[identifier] = &context{
			current: pContext.current,
			inserts: pContext.inserts,
			deletes: deletes,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Commit commits a context
func (app *application) Commit(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		pointersList := []pointers.Pointer{}
		if pContext.current != nil && len(pContext.deletes) > 0 {
			list := pContext.current.List()
			for _, onePointer := range list {
				isDeleted := false
				for _, oneDelHash := range pContext.deletes {
					if onePointer.Hash().Compare(oneDelHash) {
						isDeleted = true
						break
					}
				}

				if !isDeleted {
					pointersList = append(pointersList, onePointer)
				}
			}
		}

		pointersIns, err := app.pointersBuilder.Create().WithList(append(pointersList, pContext.inserts...)).Now()
		if err != nil {
			return err
		}

		pointerData, err := app.pointerAdapter.InstancesToBytes(pointersIns)
		if err != nil {
			return err
		}

		rootDelimiter, err := app.bytesApp.Insert(identifier, pointerData)
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			current: pointersIns,
			inserts: []pointers.Pointer{},
			deletes: []hash.Hash{},
		}

		return app.bytesApp.CommitWithRoot(identifier, rootDelimiter)
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeleteState deletes a state
func (app *application) DeleteState(context uint, stateIndex uint) error {
	return app.bytesApp.DeleteState(context, stateIndex)
}

// RecoverState recovers a state
func (app *application) RecoverState(context uint, stateIndex uint) error {
	return app.bytesApp.RecoverState(context, stateIndex)
}

// StatesAmount returns the amount of states
func (app *application) StatesAmount(context uint) (*uint, error) {
	return app.bytesApp.StatesAmount(context)
}

// DeletedStateIndexes returns the deleted state indexes
func (app *application) DeletedStateIndexes(context uint) ([]uint, error) {
	return app.bytesApp.DeletedStateIndexes(context)
}

// Close closes the context
func (app *application) Close(identifier uint) error {
	if _, ok := app.contexts[identifier]; ok {
		delete(app.contexts, identifier)
		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Purge purges the context
func (app *application) Purge(context uint) error {
	return nil
}
