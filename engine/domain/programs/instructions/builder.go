package instructions

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Instruction
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
func (app *builder) WithList(list []Instruction) Builder {
	app.list = list
	return app
}

// Now builds a new Instructions instance
func (app *builder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build a Instructions instance")
	}

	mp := map[string]Instruction{}
	for _, oneToken := range app.list {
		keyname := oneToken.Block()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Instruction (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}

		mp[keyname] = oneToken
	}

	return createInstructions(app.list, mp), nil
}
