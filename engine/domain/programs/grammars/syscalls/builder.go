package syscalls

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Syscall
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
func (app *builder) WithList(list []Syscall) Builder {
	app.list = list
	return app
}

// Now builds a new Syscalls instance
func (app *builder) Now() (Syscalls, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Syscall in order to build a Syscalls instance")
	}

	mp := map[string]Syscall{}
	for _, oneToken := range app.list {
		keyname := oneToken.Name()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Syscall (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}

		mp[keyname] = oneToken
	}

	return createSyscalls(app.list, mp), nil
}
