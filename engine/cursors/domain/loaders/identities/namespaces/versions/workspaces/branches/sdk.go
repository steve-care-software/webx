package branches

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/states"

// Adapter represents a branch adapter
type Adapter interface {
	ToBytes(ins Branch) ([]byte, error)
	ToInstance(data []byte) (Branch, error)
}

// Builder represents a branch builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithStates(states states.States) Builder
	Now() (Branch, error)
}

// Branch represents a branch
type Branch interface {
	Description() string
	HasStates() bool
	States() states.States
}
