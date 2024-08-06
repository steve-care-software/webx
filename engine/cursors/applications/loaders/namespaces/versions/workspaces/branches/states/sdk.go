package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Application represents the state loader application
type Application interface {
	Set(state states.State, name string) (states.State, error)
	Climb(state states.State) (states.State, error)
	Insert(state states.State, original originals.Original) (states.State, error)
	Update(state states.State, original string, updated originals.Original) (states.State, error)
	Delete(state states.State, name string) (states.State, error)
	Recover(state states.State, name string) (states.State, error)
	Purge(state states.State, name string) (states.State, error)
	PurgeAll(state states.State) (states.State, error)

	// data:
	InsertData(state states.State, delimiter delimiters.Delimiter) (states.State, error)
	UpdateData(state states.State, original delimiters.Delimiter, updated []byte) (states.State, error)
	DeleteData(state states.State, delete delimiters.Delimiter) (states.State, error)
}
