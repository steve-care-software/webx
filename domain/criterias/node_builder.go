package criterias

import "errors"

type nodeBuilder struct {
	next Criteria
	tail Tail
}

func createNodeBuilder() NodeBuilder {
	out := nodeBuilder{
		next: nil,
		tail: nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder()
}

// WithNext adds a next criteria to the builder
func (app *nodeBuilder) WithNext(next Criteria) NodeBuilder {
	app.next = next
	return app
}

// WithTail adds a tail to the builder
func (app *nodeBuilder) WithTail(tail Tail) NodeBuilder {
	app.tail = tail
	return app
}

//  Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.next != nil {
		return createNodeWithNext(app.next), nil
	}

	if app.tail != nil {
		return createNodeWithTail(app.tail), nil
	}

	return nil, errors.New("the Node is invalid")
}
