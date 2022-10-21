package routers

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/routers/routes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents the router builder
type Builder interface {
	Create() Builder
	WithRoutes(list []routes.Route) Builder
	Now() (Router, error)
}

// Router represents a router
type Router interface {
	Hash() hash.Hash
	Routes() []routes.Route
}
