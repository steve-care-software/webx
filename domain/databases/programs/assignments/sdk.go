package assignments

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an assignment builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithValue(value entities.Identifier) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Index() uint
	Value() entities.Identifier
}
