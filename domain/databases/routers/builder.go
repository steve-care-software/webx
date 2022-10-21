package routers

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/routers/routes"
)

type builder struct {
	hashAdapter hash.Adapter
	routes      []routes.Route
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		routes:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithRoutes add routes to the builder
func (app *builder) WithRoutes(list []routes.Route) Builder {
	app.routes = list
	return app
}

// Now builds a new Router instance
func (app *builder) Now() (Router, error) {
	if app.routes != nil && len(app.routes) <= 0 {
		app.routes = nil
	}

	if app.routes == nil {
		return nil, errors.New("there must be at least 1 Route in order to build a Router instance")
	}

	data := [][]byte{}
	for _, oneRoute := range app.routes {
		data = append(data, oneRoute.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createRouter(*pHash, app.routes), nil
}
