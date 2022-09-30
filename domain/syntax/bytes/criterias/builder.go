package criterias

import "errors"

type builder struct {
	name            string
	pIndex          *uint
	includeChannels bool
	child           Criteria
	match           []byte
}

func createBuilder() Builder {
	out := builder{
		name:            "",
		pIndex:          nil,
		includeChannels: false,
		child:           nil,
		match:           nil,
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

// IncludeChannels flags the builder as include channels
func (app *builder) IncludeChannels() Builder {
	app.includeChannels = true
	return app
}

// WithChild adds a child to the builder
func (app *builder) WithChild(child Criteria) Builder {
	app.child = child
	return app
}

// WithMatch adds a match to the builder
func (app *builder) WithMatch(match []byte) Builder {
	app.match = match
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

	if app.match != nil {
		content := createContentWithMatch(app.match)
		return createCriteriaWithContent(app.name, *app.pIndex, app.includeChannels, content), nil
	}

	if app.child != nil {
		content := createContentWithChild(app.child)
		return createCriteriaWithContent(app.name, *app.pIndex, app.includeChannels, content), nil
	}

	return createCriteria(app.name, *app.pIndex, app.includeChannels), nil
}
