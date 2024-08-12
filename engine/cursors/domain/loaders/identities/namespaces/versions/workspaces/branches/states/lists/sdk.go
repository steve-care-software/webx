package lists

import "github.com/steve-care-software/webx/engine/cursors/domain/strings"

// Adapter represents a list adapter
type Adapter interface {
	ToBytes(ins List) ([]byte, error)
	ToInstance(data []byte) (List, error)
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithResources(resources []string) Builder
	IsUnique() Builder
	Now() (List, error)
}

// List represents a list
type List interface {
	IsUnique() bool
	HasResources() bool
	Resources() strings.Strings
}
