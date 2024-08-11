package states

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
	storage_states "github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
)

type application struct {
	stateBuilder       states.Builder
	singleBuilder      singles.Builder
	storeStatesBuilder storage_states.Builder
	storeStateBuilder  storage_states.StateBuilder
}

func createApplication(
	stateBuilder states.Builder,
	singleBuilder singles.Builder,
	storeStatesBuilder storage_states.Builder,
	storeStateBuilder storage_states.StateBuilder,
) Application {
	out := application{
		stateBuilder:       stateBuilder,
		singleBuilder:      singleBuilder,
		storeStatesBuilder: storeStatesBuilder,
		storeStateBuilder:  storeStateBuilder,
	}

	return &out
}

// Set sets the state by name
func (app *application) Set(state states.State, name string) (states.State, error) {
	all := state.All()
	stateIns, err := all.FetchByName(name)
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuilder.Create().
		WithState(stateIns).
		Now()

	if err != nil {
		return nil, err
	}

	return app.stateBuilder.Create().
		WithStates(all).
		WithSingle(single).
		Now()
}

// Climb climb from the current state
func (app *application) Climb(state states.State) (states.State, error) {
	if !state.HasSingle() {
		return nil, errors.New("there is no current state to climb from")
	}
	all := state.All()
	return app.stateBuilder.Create().
		WithStates(all).
		Now()
}

// Insert inserts a state
func (app *application) Insert(state states.State, original originals.Original) (states.State, error) {
	stateIns, err := app.storeStateBuilder.Create().WithOriginal(original).Now()
	if err != nil {
		return nil, err
	}

	list := state.All().List()
	list = append(list, stateIns)
	storageStates, err := app.storeStatesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, err
	}

	builder := app.stateBuilder.Create().WithStates(storageStates)
	if state.HasSingle() {
		single := state.Single()
		builder.WithSingle(single)
	}

	return builder.Now()
}

// Update updates a state
func (app *application) Update(state states.State, original string, updated originals.Original) (states.State, error) {
	return nil, nil
}

// Delete eletes a state
func (app *application) Delete(state states.State, name string) (states.State, error) {
	return nil, nil
}

// Recover recovers a state
func (app *application) Recover(state states.State, name string) (states.State, error) {
	return nil, nil
}

// Purge purges a state
func (app *application) Purge(state states.State, name string) (states.State, error) {
	return nil, nil
}

// PurgeAll purges all states
func (app *application) PurgeAll(state states.State) (states.State, error) {
	return nil, nil
}

// InsertData inserts data in the current state
func (app *application) InsertData(state states.State, delimiter delimiters.Delimiter) (states.State, error) {
	return nil, nil
}

// UpdateData updates data in the current state
func (app *application) UpdateData(state states.State, original delimiters.Delimiter, updated []byte) (states.State, error) {
	return nil, nil
}

// DeleteData deletes data from the current state
func (app *application) DeleteData(state states.State, delete delimiters.Delimiter) (states.State, error) {
	return nil, nil
}
