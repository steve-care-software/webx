package hashtrees

import "errors"

type parentLeafBuilder struct {
	left  Leaf
	right Leaf
}

func createParentLeafBuilder() ParentLeafBuilder {
	out := parentLeafBuilder{
		left:  nil,
		right: nil,
	}

	return &out
}

// Create initializes the builder
func (app *parentLeafBuilder) Create() ParentLeafBuilder {
	return createParentLeafBuilder()
}

// WithLeft adds a left to the builder
func (app *parentLeafBuilder) WithLeft(left Leaf) ParentLeafBuilder {
	app.left = left
	return app
}

// WithRight adds a right to the builder
func (app *parentLeafBuilder) WithRight(right Leaf) ParentLeafBuilder {
	app.right = right
	return app
}

// Now builds a new ParentLeaf instance
func (app *parentLeafBuilder) Now() (ParentLeaf, error) {
	if app.left == nil {
		return nil, errors.New("the left is mandatory in order to build a ParentLeaf instance")
	}

	if app.right == nil {
		return nil, errors.New("the right is mandatory in order to build a ParentLeaf instance")
	}

	return createParentLeaf(app.left, app.right), nil
}
