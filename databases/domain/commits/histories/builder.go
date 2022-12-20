package histories

import "errors"

type builder struct {
	list []History
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

// WithList add histories to the builder
func (app *builder) WithList(list []History) Builder {
	app.list = list
	return app
}

// Now builds Histories instance
func (app *builder) Now() (Histories, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 History in order to build a Histories instance")
	}

	mp := map[string]History{}
	for _, oneHistory := range app.list {
		keyname := oneHistory.Commit().String()
		mp[keyname] = oneHistory
	}

	return createHistories(mp, app.list), nil
}
