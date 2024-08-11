package switchers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
)

type switcherBuilder struct {
	original singles.Single
	updated  singles.Single
}

func createSwitcherBuilder() SwitcherBuilder {
	out := switcherBuilder{
		original: nil,
		updated:  nil,
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

// WithUpdated adds an updated to the builder
func (app *switcherBuilder) WithUpdated(updated singles.Single) SwitcherBuilder {
	app.updated = updated
	return app
}

// Now builds a new Switcher instance
func (app *switcherBuilder) Now() (Switcher, error) {
	if app.original != nil && app.updated != nil {
		return createSwitcherWithOriginalAndUpdated(app.original, app.updated), nil
	}

	if app.original != nil {
		return createSwitcherWithOriginal(app.original), nil
	}

	if app.updated != nil {
		return createSwitcherWithUpdated(app.updated), nil
	}

	return nil, errors.New("the Switcher is invalid")
}
