package entries

import "errors"

type relationsBuilder struct {
	list []Relation
}

func createRelationsBuilder() RelationsBuilder {
	out := relationsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *relationsBuilder) Create() RelationsBuilder {
	return createRelationsBuilder()
}

// WithList adds a list to the builder
func (app *relationsBuilder) WithList(list []Relation) RelationsBuilder {
	app.list = list
	return app
}

// Now builds a new Relations instance
func (app *relationsBuilder) Now() (Relations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Relation in order to build a Relations instance")
	}

	return createRelations(app.list), nil
}
