package modifications

import (
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states/containers"
)

// Adapter represents a modification adapter
type Adapter interface {
	ToUpdatedState(modification Modification) (states.State, error)
}

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithOriginal(original states.State) Builder
	WithInsertions(insertions entries.Entries) Builder
	WithDeletions(deletions containers.Container) Builder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Original() states.State
	HasInsertions() bool
	Insertions() entries.Entries
	HasDeletions()
	Deletions() containers.Container
}
