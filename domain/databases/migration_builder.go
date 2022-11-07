package databases

import "errors"

type migrationBuilder struct {
	previous    Head
	pheight     *uint
	description string
}

func createMigrationBuilder() MigrationBuilder {
	out := migrationBuilder{
		previous:    nil,
		pheight:     nil,
		description: "",
	}

	return &out
}

// Create initializes the builder
func (app *migrationBuilder) Create() MigrationBuilder {
	return createMigrationBuilder()
}

// WithPrevious adds previous head to the builder
func (app *migrationBuilder) WithPrevious(previous Head) MigrationBuilder {
	app.previous = previous
	return app
}

// WithHeight adds an height to the builder
func (app *migrationBuilder) WithHeight(height uint) MigrationBuilder {
	app.pheight = &height
	return app
}

// WithDescrition adds a description to the builder
func (app *migrationBuilder) WithDescrition(description string) MigrationBuilder {
	app.description = description
	return app
}

// Now builds a new Migration instance
func (app *migrationBuilder) Now() (Migration, error) {
	if app.previous == nil {
		return nil, errors.New("the previous content is mandatory in order to build a Migration instance")
	}

	if app.pheight == nil {
		return nil, errors.New("the height is mandatory in order to build a Migration instance")
	}

	return createMigration(app.previous, *app.pheight, app.description), nil
}
