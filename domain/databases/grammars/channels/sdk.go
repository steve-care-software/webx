package channels

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a channel builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithToken(token entities.Identifier) Builder
	WithPrevious(previous entities.Identifier) Builder
	WithNext(next entities.Identifier) Builder
	Now() (Channel, error)
}

// Channel represents a chanel
type Channel interface {
	Entity() entities.Entity
	Token() entities.Identifier
	HasPrevious() bool
	Previous() entities.Identifier
	HasNext() bool
	Next() entities.Identifier
}
