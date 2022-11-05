package fetchers

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a fetcher builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithRecursive(recursive entities.Identifier) Builder
	WithSelector(selector entities.Identifier) Builder
	Now() (Fetcher, error)
}

// Fetcher represents a fetcher
type Fetcher interface {
	Entity() entities.Entity
	Content() Content
}

// Content represents a fetcher's content
type Content interface {
	IsRecursive() bool
	Recursive() entities.Identifier
	IsSelector() bool
	Selector() entities.Identifier
}
