package databases

import "errors"

type sectionBuilder struct {
	pIndex   *uint
	pKind    *uint8
	pointers Pointers
}

func createSectionBuilder() SectionBuilder {
	out := sectionBuilder{
		pIndex:   nil,
		pKind:    nil,
		pointers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *sectionBuilder) Create() SectionBuilder {
	return createSectionBuilder()
}

// WithIndex adds an index to the builder
func (app *sectionBuilder) WithIndex(index uint) SectionBuilder {
	app.pIndex = &index
	return app
}

// WithKind adds a kind to the builder
func (app *sectionBuilder) WithKind(kind uint8) SectionBuilder {
	app.pKind = &kind
	return app
}

// WithPointers add pointers to the builder
func (app *sectionBuilder) WithPointers(pointers Pointers) SectionBuilder {
	app.pointers = pointers
	return app
}

// Now builds a new Section instance
func (app *sectionBuilder) Now() (Section, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Section instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Section instance")
	}

	if app.pointers != nil {
		return createSectionWithPointers(*app.pIndex, *app.pKind, app.pointers), nil
	}

	return createSection(*app.pIndex, *app.pKind), nil
}
