package heads

import "errors"

type relationBuilder struct {
	pFrom *uint
	to    []uint
}

func createRelationBuilder() RelationBuilder {
	out := relationBuilder{
		pFrom: nil,
		to:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *relationBuilder) Create() RelationBuilder {
	return createRelationBuilder()
}

// From adds a from index to the builder
func (app *relationBuilder) From(from uint) RelationBuilder {
	app.pFrom = &from
	return app
}

// To adds a to index to the builder
func (app *relationBuilder) To(to []uint) RelationBuilder {
	app.to = to
	return app
}

// Now builds a new Relation instance
func (app *relationBuilder) Now() (Relation, error) {
	if app.pFrom == nil {
		return nil, errors.New("the from index is mandatory in order to buiold a Relation instance")
	}

	if app.to != nil && len(app.to) <= 0 {
		app.to = nil
	}

	if app.to == nil {
		return nil, errors.New("the to indexes is mandatory in order to buiold a Relation instance")
	}

	return createRelation(*app.pFrom, app.to), nil
}
