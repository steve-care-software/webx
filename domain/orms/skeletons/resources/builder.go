package resources

import (
	"errors"
	"fmt"
)

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
		name := oneResource.Name()
		if idx, ok := mp[name]; ok {
			str := fmt.Sprintf("the Resource (index: %d, name: %s) already exists", idx, name)
			return nil, errors.New(str)
		}

		mp[name] = oneResource
	}

	return createResources(mp, app.list), nil
}
