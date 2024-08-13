package switchers

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Switcher
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
func (app *builder) WithList(list []Switcher) Builder {
	app.list = list
	return app
}

// Now builds a new Switchers instance
func (app *builder) Now() (Switchers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Switcher in order to build a Switchers instance")
	}

	mp := map[string]Switcher{}
	for _, oneSwitcher := range app.list {
		name := oneSwitcher.Name()
		if idx, ok := mp[name]; ok {
			str := fmt.Sprintf("the switcher (name: %s, index: %d) is duplicate", name, idx)
			return nil, errors.New(str)
		}

		mp[name] = oneSwitcher
	}

	return createSwitchers(app.list, mp), nil
}
