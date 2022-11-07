package databases

import "errors"

type pointersBuilder struct {
	list []Pointer
}

func createPointersBuilder() PointersBuilder {
	out := pointersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointersBuilder) Create() PointersBuilder {
	return createPointersBuilder()
}

// WithList adds a list to the builder
func (app *pointersBuilder) WithList(list []Pointer) PointersBuilder {
	app.list = list
	return app
}

// Now builds a new Pointers instance
func (app *pointersBuilder) Now() (Pointers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Pointer in order to build an Pointers instance")
	}

	return createPointers(app.list), nil
}
