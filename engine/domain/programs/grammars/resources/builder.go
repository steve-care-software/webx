package resources

import "errors"

type builder struct {
	list []Resource
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
func (app *builder) WithList(list []Resource) Builder {
	app.list = list
	return app
}

// Now builds a new Resources instance
func (app *builder) Now() (Resources, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Resource in order to build a Resources instance")
	}

	mp := map[string]Resource{}
	for _, oneResource := range app.list {
		keyname := oneResource.Name()
		mp[keyname] = oneResource
	}

	return createResources(
		app.list,
		mp,
	), nil
}
