package compilers

type valueBuilder struct {
	constant string
	//criteria criterias.Criteria
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		constant: "",
		//criteria: nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant string) ValueBuilder {
	app.constant = constant
	return app
}

// WithCriteria adds a criteria to the builder
/*func (app *valueBuilder) WithCriteria(criteria criterias.Criteria) ValueBuilder {
	app.criteria = criteria
	return app
}*/

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	/*if app.criteria != nil {
		return createValueWithCriteria(app.criteria), nil
	}*/

	if app.constant != "" {
		return createValueWithConstant(app.constant), nil
	}

	return createValueWithNil(), nil
}
