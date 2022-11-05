package routers

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a router builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithRoutes(routes []Route) Builder
	Now() (Router, error)
}

// Router represents a router
type Router interface {
	Entity() entities.Entity
	List() []Route
}

// RouteBuilder represents a route builder
type RouteBuilder interface {
	Create() RouteBuilder
	WithEntity(entity entities.Entity) RouteBuilder
	WithGrammar(grammar entities.Identifier) RouteBuilder
	WithProgram(program entities.Identifier) RouteBuilder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Program() entities.Identifier
}
