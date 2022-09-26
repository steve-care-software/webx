package connections

import "errors"

type builder struct {
	connections []Connection
}

func createBuilder() Builder {
	out := builder{
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add connections to the builder
func (app *builder) WithList(connections []Connection) Builder {
	app.connections = connections
	return app
}

// Now builds a new Connections instance
func (app *builder) Now() (Connections, error) {
	if app.connections != nil && len(app.connections) <= 0 {
		app.connections = nil
	}

	if app.connections == nil {
		return nil, errors.New("there must be at least 1 Connection in order to build an Connections instance")
	}

	return createConnections(app.connections), nil
}
