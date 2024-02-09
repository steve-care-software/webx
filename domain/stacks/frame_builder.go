package stacks

type frameBuilder struct {
	assignments Assignments
}

func createFrameBuilder() FrameBuilder {
	out := frameBuilder{
		assignments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *frameBuilder) Create() FrameBuilder {
	return createFrameBuilder()
}

// WithAssignments add assignments to the builder
func (app *frameBuilder) WithAssignments(assignments Assignments) FrameBuilder {
	app.assignments = assignments
	return app
}

// Now builds a new Frame instance
func (app *frameBuilder) Now() (Frame, error) {
	if app.assignments != nil {
		return createFrameWithAssignments(app.assignments), nil
	}

	return createFrame(), nil
}
