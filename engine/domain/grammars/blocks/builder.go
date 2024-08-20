package blocks

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Block
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
func (app *builder) WithList(list []Block) Builder {
	app.list = list
	return app
}

// Now builds a new Blocks instance
func (app *builder) Now() (Blocks, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Block in order to build a Blocks instance")
	}

	// reverse the list:
	for i, j := 0, len(app.list)-1; i < j; i, j = i+1, j-1 {
		app.list[i], app.list[j] = app.list[j], app.list[i]
	}

	mp := map[string]Block{}
	for _, oneBlock := range app.list {
		keyname := oneBlock.Name()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Block (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}
		mp[keyname] = oneBlock
	}

	return createBlocks(app.list, mp), nil
}
