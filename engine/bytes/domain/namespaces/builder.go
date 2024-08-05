package namespaces

import (
	"errors"
	"fmt"
)

type namespacesIns struct {
	list []Namespace
}

func createBuilder() Builder {
	out := namespacesIns{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *namespacesIns) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *namespacesIns) WithList(list []Namespace) Builder {
	app.list = list
	return app
}

// Now builds a new NAmespaces instance
func (app *namespacesIns) Now() (Namespaces, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	names := []string{}
	mp := map[string]Namespace{}
	for idx, oneIns := range app.list {
		name := oneIns.Name()
		if _, ok := mp[name]; ok {
			str := fmt.Sprintf("thre namespace (%s) at index (%d) is a duplicate", name, idx)
			return nil, errors.New(str)
		}

		mp[name] = oneIns
		names = append(names, name)
	}

	return createNamespaces(
		mp,
		app.list,
		names,
	), nil
}
