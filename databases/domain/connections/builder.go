package connections

import "errors"

type builder struct {
	list []Connection
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add connections to the builder
func (app *builder) WithList(list []Connection) Builder {
	app.list = list
	return app
}

// Now builds Connections instance
func (app *builder) Now() (Connections, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Connection in order to build a Connections instance")
	}

	mp := map[uint]Connection{}
	for _, oneConnection := range app.list {
		identifier := oneConnection.Identifier()
		mp[identifier] = oneConnection
	}

	return createConnections(mp, app.list), nil
}
