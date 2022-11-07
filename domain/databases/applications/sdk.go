package applications

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithInit(init entities.Identifiers) Builder
	WithStop(stop entities.Identifiers) Builder
	WithStart(start entities.Identifiers) Builder
	WithDaemon(daemon entities.Identifiers) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Entity() entities.Entity
	Content() Content
}

// Content represents content
type Content interface {
	HasInit() bool
	Init() entities.Identifiers
	HasStop() bool
	Stop() entities.Identifiers
	HasStart() bool
	Start() entities.Identifiers
	HasDaemon() bool
	Daemon() entities.Identifiers
}
