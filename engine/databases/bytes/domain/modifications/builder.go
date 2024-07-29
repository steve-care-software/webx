package modifications

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
)

type builder struct {
	original   states.State
	insertions entries.Entries
	deletions  deletes.Deletes
}

func createBuilder() Builder {
	out := builder{
		original:   nil,
		insertions: nil,
		deletions:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOriginal adds an original to the builder
func (app *builder) WithOriginal(original states.State) Builder {
	app.original = original
	return app
}

// WithInsertions add insertions to the builder
func (app *builder) WithInsertions(insertions entries.Entries) Builder {
	app.insertions = insertions
	return app
}

// WithDeletions add deletions to the builder
func (app *builder) WithDeletions(deletions deletes.Deletes) Builder {
	app.deletions = deletions
	return app
}

// Now builds a new Modification instance
func (app *builder) Now() (Modification, error) {
	if app.original == nil {
		return nil, errors.New("the original is mandatory in order to build a Modification instance")
	}

	if app.insertions != nil && app.deletions != nil {
		return createModificationWithInsertionsAndDeletions(app.original, app.insertions, app.deletions), nil
	}

	if app.insertions != nil {
		return createModificationWithInsertions(app.original, app.insertions), nil
	}

	if app.deletions != nil {
		return createModificationWithDeletions(app.original, app.deletions), nil
	}

	return nil, errors.New("the Modification is invalid")
}
