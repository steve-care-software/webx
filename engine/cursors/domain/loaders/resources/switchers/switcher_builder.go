package switchers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

type switcherBuilder struct {
	original singles.Single
	insert   inserts.Insert
	update   updates.Update
	delete   deletes.Delete
}

func createSwitcherBuilder() SwitcherBuilder {
	out := switcherBuilder{
		original: nil,
		insert:   nil,
		update:   nil,
		delete:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *switcherBuilder) Create() SwitcherBuilder {
	return createSwitcherBuilder()
}

// WithOriginal adds an original to the builder
func (app *switcherBuilder) WithOriginal(original singles.Single) SwitcherBuilder {
	app.original = original
	return app
}

// WithInsert adds an insert to the builder
func (app *switcherBuilder) WithInsert(insert inserts.Insert) SwitcherBuilder {
	app.insert = insert
	return app
}

// WithUpdate adds an update to the builder
func (app *switcherBuilder) WithUpdate(update updates.Update) SwitcherBuilder {
	app.update = update
	return app
}

// WithDelete adds a delete to the builder
func (app *switcherBuilder) WithDelete(delete deletes.Delete) SwitcherBuilder {
	app.delete = delete
	return app
}

// Now builds a new Switcher instance
func (app *switcherBuilder) Now() (Switcher, error) {
	if app.original != nil && app.insert != nil {
		return createSwitcherWithOriginalAndInsert(app.original, app.insert), nil
	}

	if app.original != nil && app.update != nil {
		return createSwitcherWithOriginalAndUpdate(app.original, app.update), nil
	}

	if app.original != nil && app.delete != nil {
		return createSwitcherWithOriginalAndDelete(app.original, app.delete), nil
	}

	if app.insert != nil {
		return createSwitcherWithInsert(app.insert), nil
	}

	if app.update != nil {
		return createSwitcherWithUpdate(app.update), nil
	}

	if app.delete != nil {
		return createSwitcherWithDelete(app.delete), nil
	}

	return nil, errors.New("the Switcher is invalid")
}
