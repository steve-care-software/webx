package connections

import (
	"errors"
	"fmt"
)

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

// WithList adds a list to the builder
func (app *builder) WithList(list []Connection) Builder {
	app.list = list
	return app
}

// Now builds a new Connections instance
func (app *builder) Now() (Connections, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Connection in order to build a Connections instance")
	}

	mp := map[string]Connection{}
	for _, oneConnection := range app.list {
		name := oneConnection.Name()
		if idx, ok := mp[name]; ok {
			str := fmt.Sprintf("the Connection (index: %d, name: %s) already exists", idx, name)
			return nil, errors.New(str)
		}

		mp[name] = oneConnection
	}

	return createConnections(mp, app.list), nil
}
