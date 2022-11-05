package assignments

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

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
