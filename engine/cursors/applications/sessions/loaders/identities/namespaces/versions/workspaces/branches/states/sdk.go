package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

const cannotBeDeletedNeverCommitedBeforeErrPattern = "the state (index; %d) cannot be deleted because it has never been committed yet"
const stateAlreadyDeletedErrPattern = "the state (index: %d) has already been deleted"
const cannotBeRecoveredNeverCommitedBeforeErrPattern = "the state (index; %d) cannot be recovered because it has never been committed yet"
const stateNotDeletedErrPattern = "the state (index: %d) cannot be recovered because it has not been deleted"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithNextIndex(nextIndex uint64) Builder
	Now() (Application, error)
}

// Application represents the state loader application
type Application interface {
	List(input states.States) []string
	Delete(input states.States, index uint64, message string) (states.States, error)
	Recover(input states.States, index uint64, message string) (states.States, error)

	// data
	InsertData(input states.States, message string, data []byte) (states.States, error)
	UpdateData(input states.States, original delimiters.Delimiter, updated []byte) (states.States, error)
	DeleteData(input states.States, delete delimiters.Delimiter) (states.States, error)

	// commit:
	Commit(input states.States, message string) (states.States, error)
}
