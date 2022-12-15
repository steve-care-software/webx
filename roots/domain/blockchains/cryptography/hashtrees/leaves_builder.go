package hashtrees

import "errors"

type leavesBuilder struct {
	list []Leaf
}

func createLeavesBuilder() LeavesBuilder {
	out := leavesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *leavesBuilder) Create() LeavesBuilder {
	return createLeavesBuilder()
}

// WithList add list to the builder
func (app *leavesBuilder) WithList(list []Leaf) LeavesBuilder {
	app.list = list
	return app
}

// Now builds a new Leaves instance
func (app *leavesBuilder) Now() (Leaves, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Leaf in order to build a Leaves instance")
	}

	return createLeaves(app.list), nil
}
