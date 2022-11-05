package insides

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents an inside builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithFn(fn entities.Identifier) Builder
	WithFetchers(fetchers []entities.Identifier) Builder
	Now() (Inside, error)
}

// Inside represents an inside
type Inside interface {
	Entity() entities.Entity
	Content() Content
}

// Content represents an inside content
type Content interface {
	IsFn() bool
	Fn() entities.Identifier
	IsFetchers() bool
	Fetchers() []entities.Identifier
}
