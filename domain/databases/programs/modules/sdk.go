package modules

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents a module builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithName(name []byte) Builder
	Now() (Module, error)
}

// Module represents a module
type Module interface {
	Entity() entities.Entity
	Name() []byte
}
