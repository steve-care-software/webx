package references

type builder struct {
	active    Keys
	pendings  Keys
	deleted   Keys
	links     Links
	relations Relations
}

func createBuilder() Builder {
	out := builder{
		active:    nil,
		pendings:  nil,
		deleted:   nil,
		links:     nil,
		relations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithActive add the active keys to the builder
func (app *builder) WithActive(active Keys) Builder {
	app.active = active
	return app
}

// WithPendings add pendings to the builder
func (app *builder) WithPendings(pendings Keys) Builder {
	app.pendings = pendings
	return app
}

// WithDeleted add the deleted keys to the builder
func (app *builder) WithDeleted(deleted Keys) Builder {
	app.deleted = deleted
	return app
}

// WithLinks add the links to the builder
func (app *builder) WithLinks(links Links) Builder {
	app.links = links
	return app
}

// WithRelations add the relations to the builder
func (app *builder) WithRelations(relations Relations) Builder {
	app.relations = relations
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	return createReference(
		app.active,
		app.pendings,
		app.deleted,
		app.links,
		app.relations,
	), nil
}
