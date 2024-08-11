package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
)

// Adapter represents a single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithPointers(pointers pointers.Pointers) Builder
	IsDeleted() Builder
	Now() (Single, error)
}

// Single represents the single state
type Single interface {
	Message() string
	IsDeleted() bool
	HasPointers() bool
	Pointers() pointers.Pointers
}
