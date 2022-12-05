package coverages

import "errors"

type executionsBuilder struct {
	list []Execution
}

func createExecutionsBuilder() ExecutionsBuilder {
	out := executionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionsBuilder) Create() ExecutionsBuilder {
	return createExecutionsBuilder()
}

// WithList adds a list to the builder
func (app *executionsBuilder) WithList(list []Execution) ExecutionsBuilder {
	app.list = list
	return app
}

// Now builds a new Executions instance
func (app *executionsBuilder) Now() (Executions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Execution in order to build a Executions instance")
	}

	return createExecutions(app.list), nil
}
