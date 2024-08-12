package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

// Adapter represents a single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithDelimiter(delimiter delimiters.Delimiter) Builder
	IsDeleted() Builder
	Now() (Single, error)
}

// Single represents the single state
type Single interface {
	Delimiter() delimiters.Delimiter
	IsDeleted() bool
}
