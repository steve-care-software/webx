package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a single builder
type Builder interface {
	Create() Builder
	WithStorage(storage storages.Storage) Builder
	WithBytes(bytes []byte) Builder
	Now() (Single, error)
}

// Single represents a single
type Single interface {
	Storage() storages.Storage
	Bytes() []byte
}
