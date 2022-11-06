package suites

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a suite builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithContent(content []byte) Builder
	IsValid() Builder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Entity() entities.Entity
	IsValid() bool
	Content() []byte
}
