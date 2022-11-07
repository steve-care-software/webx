package databases

import "errors"

type sectionsBuilder struct {
	list []Section
}

func createSectionsBuilder() SectionsBuilder {
	out := sectionsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *sectionsBuilder) Create() SectionsBuilder {
	return createSectionsBuilder()
}

// WithList adds a list to the builder
func (app *sectionsBuilder) WithList(list []Section) SectionsBuilder {
	app.list = list
	return app
}

// Now builds a new Sections instance
func (app *sectionsBuilder) Now() (Sections, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Section in order to build an Sections instance")
	}

	return createSections(app.list), nil
}
