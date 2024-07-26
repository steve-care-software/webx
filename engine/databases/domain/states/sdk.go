package states

import (
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers"
)

// Adapter represents a state adapter
type Adapter interface {
	ToUpdatedHeader(state State) (headers.Header, error)
}

// Builder represents a state builder
type Builder interface {
	Create() Builder
	WithOriginal(original headers.Header) Builder
	WithInsertions(insertions entries.Entries) Builder
	WithDeletions(deletions containers.Container) Builder
	Now() (State, error)
}

// State represents a state
type State interface {
	Original() headers.Header
	HasInsertions() bool
	Insertions() entries.Entries
	HasDeletions()
	Deletions() containers.Container
}
