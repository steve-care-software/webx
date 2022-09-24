package criterias

import "errors"

type builder struct {
	name        string
	pIndex      *uint
	requirement []uint
	child       Criteria
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		pIndex:      nil,
		requirement: nil,
		child:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithRequirement adds a requirement to the builder
func (app *builder) WithRequirement(requirement []uint) Builder {
	app.requirement = requirement
	return app
}

// WithChild adds a child to the builder
func (app *builder) WithChild(child Criteria) Builder {
	app.child = child
	return app
}

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Criteria instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Criteria instance")
	}

	if app.requirement != nil && len(app.requirement) <= 0 {
		app.requirement = nil
	}

	if app.requirement != nil {
		content := createContentWithRequirement(app.requirement)
		return createCriteria(app.name, *app.pIndex, content), nil
	}

	if app.child != nil {
		content := createContentWithChild(app.child)
		return createCriteria(app.name, *app.pIndex, content), nil
	}

	return nil, errors.New("the Criteria is invalid")
}
