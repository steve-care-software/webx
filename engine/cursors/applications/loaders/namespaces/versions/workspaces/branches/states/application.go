package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles"
)

type application struct {
	stateBuilder  states.Builder
	singleBuilder singles.Builder
}

func createApplication(
	stateBuilder states.Builder,
	singleBuilder singles.Builder,
) Application {
	out := application{
		stateBuilder:  stateBuilder,
		singleBuilder: singleBuilder,
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
