package modifications

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a modification adapter
type Adapter interface {
	ToUpdatedState(modification Modification) (states.State, error)
}

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithOriginal(original states.State) Builder
	WithInsertions(insertions entries.Entries) Builder
	WithDeletions(deletions deletes.Deletes) Builder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Original() states.State
	HasInsertions() bool
	Insertions() entries.Entries
	HasDeletions() bool
	Deletions() deletes.Deletes
}
