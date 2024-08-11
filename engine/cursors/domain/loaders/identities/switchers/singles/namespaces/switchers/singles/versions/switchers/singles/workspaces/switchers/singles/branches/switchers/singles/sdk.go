package singles

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states"

// Adapter represents a single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents a single builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithStates(states states.States) Builder
	Now() (Single, error)
}

// Single represents a single branch
type Single interface {
	Name() string
	Description() string
	HasStates() bool
	States() states.States
}
