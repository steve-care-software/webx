package routers

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes"
)

type router struct {
	hash   hash.Hash
	routes []routes.Route
}

func createRouter(
	hash hash.Hash,
	routes []routes.Route,
) Router {
	out := router{
		hash:   hash,
		routes: routes,
	}

	return &out
}

// Hash returns the hash
func (obj *router) Hash() hash.Hash {
	return obj.hash
}

// Routes returns the routes
func (obj *router) Routes() []routes.Route {
	return obj.routes
}
