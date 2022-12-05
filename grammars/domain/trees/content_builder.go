package trees

import (
	"errors"
)

type contentBuilder struct {
	value Value
	tree  Tree
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		value: nil,
		tree:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithValue adds a value to the builder
func (app *contentBuilder) WithValue(value Value) ContentBuilder {
	app.value = value
	return app
}

// WithTree adds a tree to the builder
func (app *contentBuilder) WithTree(tree Tree) ContentBuilder {
	app.tree = tree
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.value != nil {
		return createContentWithValue(app.value), nil
	}

	if app.tree != nil {
		return createContentWithTree(app.tree), nil
	}

	return nil, errors.New("the Content is invalid")
}
