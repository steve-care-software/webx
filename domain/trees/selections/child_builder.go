package selections

import "errors"

type childBuilder struct {
	selections Selections
	content    []byte
}

func createChildBuilder() ChildBuilder {
	out := childBuilder{
		selections: nil,
		content:    nil,
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

// WithContent adds content to the builder
func (app *childBuilder) WithContent(content []byte) ChildBuilder {
	app.content = content
	return app
}

// Now builds a new Child instance
func (app *childBuilder) Now() (Child, error) {
	if app.selections != nil {
		return createChildWithSelections(app.selections), nil
	}

	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content != nil {
		return createChildWithContent(app.content), nil
	}

	return nil, errors.New("the Child is invalid")
}
