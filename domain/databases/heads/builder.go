package heads

type builder struct {
	active            Keys
	deleted           Keys
	links             Links
	relations         Relations
	weightedRelations WeightedRelations
}

func createBuilder() Builder {
	out := builder{
		active:            nil,
		deleted:           nil,
		links:             nil,
		relations:         nil,
		weightedRelations: nil,
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

// WithWeightedRelations add the weighted relations to the builder
func (app *builder) WithWeightedRelations(weightedRelations WeightedRelations) Builder {
	app.weightedRelations = weightedRelations
	return app
}

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	return createHead(app.active, app.deleted, app.links, app.relations, app.weightedRelations), nil
}
