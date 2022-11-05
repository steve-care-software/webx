package everythings

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents an everything builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithException(exception entities.Identifier) Builder
	WithEscape(escape entities.Identifier) Builder
	Now() (Everything, error)
}

// Everything represents an everything
type Everything interface {
	Entity() entities.Entity
	Exception() entities.Identifier
	HasEscape() bool
	Escape() entities.Identifier
}
