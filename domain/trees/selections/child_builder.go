package selections

import "errors"

type childBuilder struct {
	selections Selections
	bytes      []byte
}

func createChildBuilder() ChildBuilder {
	out := childBuilder{
		selections: nil,
		bytes:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *childBuilder) Create() ChildBuilder {
	return createChildBuilder()
}

// WithSelections adds a selections to the builder
func (app *childBuilder) WithSelections(selections Selections) ChildBuilder {
	app.selections = selections
	return app
}

// WithBytes adds bytes to the builder
func (app *childBuilder) WithBytes(bytes []byte) ChildBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Child instance
func (app *childBuilder) Now() (Child, error) {
	if app.selections != nil {
		return createChildWithSelections(app.selections), nil
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createChildWithBytes(app.bytes), nil
	}

	return nil, errors.New("the Child is invalid")
}
