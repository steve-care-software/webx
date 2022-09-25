package outputs

import "errors"

type builder struct {
	list []Variable
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

// WithList adds a list of variables to the builder
func (app *builder) WithList(list []Variable) Builder {
	app.list = list
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Variable in order to build an Output instance")
	}

	mp := map[string]Variable{}
	for _, oneVariable := range app.list {
		name := oneVariable.Name()
		mp[name] = oneVariable
	}

	return createOutput(app.list, mp), nil
}
