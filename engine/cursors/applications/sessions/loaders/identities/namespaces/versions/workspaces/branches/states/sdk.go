package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/states/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
)

// Application represents the state loader application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error) // creates a new state
	/*List(input states.States) []string
	Delete(input states.States, index uint64, message string) (states.States, error)
	Recover(input states.States, index uint64, message string) (states.States, error)

	// data
	InsertData(input states.States, message string, data []byte) (states.States, error)
	UpdateData(input states.States, original delimiters.Delimiter, updated []byte) (states.States, error)
	DeleteData(input states.States, delete delimiters.Delimiter) (states.States, error)

	// commit:
	Commit(input states.States, message string) (states.States, error)*/
}
