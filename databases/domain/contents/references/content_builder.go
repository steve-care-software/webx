package references

type contentBuilder struct {
	active   ContentKeys
	pendings ContentKeys
	deleted  ContentKeys
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		active:   nil,
		pendings: nil,
		deleted:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithActive add the active keys to the builder
func (app *contentBuilder) WithActive(active ContentKeys) ContentBuilder {
	app.active = active
	return app
}

// WithPendings add pendings to the builder
func (app *contentBuilder) WithPendings(pendings ContentKeys) ContentBuilder {
	app.pendings = pendings
	return app
}

// WithDeleted add the deleted keys to the builder
func (app *contentBuilder) WithDeleted(deleted ContentKeys) ContentBuilder {
	app.deleted = deleted
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	return createContent(
		app.active,
		app.pendings,
		app.deleted,
	), nil
}
