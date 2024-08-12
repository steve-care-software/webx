package states

import (
	"errors"
	"fmt"

	storage_pointer_applications "github.com/steve-care-software/webx/engine/cursors/applications/sessions/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/updates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type application struct {
	pointerApp            storage_pointer_applications.Application
	statesBuilder         states.Builder
	stateBuilder          states.StateBuilder
	singleAdapter         singles.Adapter
	singleBuilder         singles.Builder
	updateBuilder         updates.Builder
	pointersBuilder       pointers.Builder
	pointerBuilder        pointers.PointerBuilder
	storagePointerBuilder storage_pointers.StorageBuilder
	delimiterBuilder      delimiters.DelimiterBuilder
	nextIndex             uint64
}

func createApplication(
	pointerApp storage_pointer_applications.Application,
	statesBuilder states.Builder,
	stateBuilder states.StateBuilder,
	singleAdapter singles.Adapter,
	singleBuilder singles.Builder,
	updateBuilder updates.Builder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	storagePointerBuilder storage_pointers.StorageBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
	nextIndex uint64,
) Application {
	out := application{
		pointerApp:            pointerApp,
		statesBuilder:         statesBuilder,
		stateBuilder:          stateBuilder,
		singleAdapter:         singleAdapter,
		singleBuilder:         singleBuilder,
		updateBuilder:         updateBuilder,
		pointersBuilder:       pointersBuilder,
		pointerBuilder:        pointerBuilder,
		storagePointerBuilder: storagePointerBuilder,
		delimiterBuilder:      delimiterBuilder,
		nextIndex:             nextIndex,
	}

	return &out
}

// List lists the message of the current states
func (app *application) List(input states.States) []string {
	return input.Messages()
}

// Delete deletes a state
func (app *application) Delete(input states.States, index uint64, message string) (states.States, error) {
	toDelState, err := input.Fetch(index)
	if err != nil {
		return nil, err
	}

	if toDelState.Current().IsDeleted() {
		str := fmt.Sprintf(stateAlreadyDeletedErrPattern, index)
		return nil, errors.New(str)
	}

	if !toDelState.HasOriginal() {
		str := fmt.Sprintf(cannotBeDeletedNeverCommitedBeforeErrPattern, index)
		return nil, errors.New(str)
	}

	toDelSingle := toDelState.Original()
	updatedSingleBuilder := app.singleBuilder.Create().
		IsDeleted().
		WithMessage(message)

	current := toDelState.Current()
	if current.HasPointers() {
		inPointers := toDelState.Current().Pointers()
		retPointers, err := app.pointerApp.DeleteData(inPointers, index)
		if err != nil {
			return nil, err
		}

		updatedSingleBuilder.WithPointers(retPointers)
	}

	updatedSingle, err := updatedSingleBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.updateSingleAtIndex(
		input,
		index,
		toDelSingle,
		updatedSingle,
	)
}

// Recover recovers a state
func (app *application) Recover(input states.States, index uint64, message string) (states.States, error) {
	toRecoverState, err := input.Fetch(index)
	if err != nil {
		return nil, err
	}

	if toRecoverState.Current().IsDeleted() {
		str := fmt.Sprintf(stateNotDeletedErrPattern, index)
		return nil, errors.New(str)
	}

	if !toRecoverState.HasOriginal() {
		str := fmt.Sprintf(cannotBeRecoveredNeverCommitedBeforeErrPattern, index)
		return nil, errors.New(str)
	}

	toRecoverSingle := toRecoverState.Original()
	updatedSingleBuilder := app.singleBuilder.Create().WithMessage(message)
	if toRecoverSingle.HasPointers() {
		pointers := toRecoverSingle.Pointers()
		updatedSingleBuilder.WithPointers(pointers)
	}

	updatedSingle, err := updatedSingleBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.updateSingleAtIndex(
		input,
		index,
		toRecoverSingle,
		updatedSingle,
	)
}

// InsertData inserts data
func (app *application) InsertData(input states.States, message string, data []byte) (states.States, error) {
	lastState, err := input.LastActive()
	if err != nil {
		return nil, err
	}

	current := lastState.Current()
	inPointers := current.Pointers()
	retPointers, err := app.pointerApp.InsertData(inPointers, data)
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuilder.Create().
		WithMessage(message).
		WithPointers(retPointers).
		Now()

	if err != nil {
		return nil, err
	}

	return app.addSingleToList(input, single)
}

// UpdateData updates data
func (app *application) UpdateData(input states.States, original delimiters.Delimiter, updated []byte) (states.States, error) {
	/*lastState, err := input.LastActive()
	if err != nil {
		return nil, err
	}

	current := lastState.Current()*/

	return nil, nil
}

// DeleteData deletes data
func (app *application) DeleteData(input states.States, delete delimiters.Delimiter) (states.States, error) {
	return nil, nil
}

// Commit commits data
func (app *application) Commit(input states.States, message string) (states.States, error) {
	return nil, nil
}

func (app *application) addSingleToList(input states.States, newSingle singles.Single) (states.States, error) {
	newState, err := app.stateBuilder.Create().
		WithOriginal(newSingle).
		Now()

	if err != nil {
		return nil, err
	}

	list := input.List()
	list = append(list, newState)
	return app.statesBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) updateSingleAtIndex(input states.States, index uint64, original singles.Single, updatedSingle singles.Single) (states.States, error) {
	bytes, err := app.singleAdapter.ToBytes(updatedSingle)
	if err != nil {
		return nil, err
	}

	update, err := app.updateBuilder.Create().
		WithSingle(updatedSingle).
		WithBytes(bytes).
		Now()

	if err != nil {
		return nil, err
	}

	updatedState, err := app.stateBuilder.Create().
		WithOriginal(original).
		WithUpdated(update).
		Now()

	if err != nil {
		return nil, err
	}

	list := input.List()
	list[index] = updatedState
	return app.statesBuilder.Create().
		WithList(list).
		Now()
}
