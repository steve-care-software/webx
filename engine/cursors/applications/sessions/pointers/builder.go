package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/resources/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

type builder struct {
	dbApp                databases.Application
	storagePointerBulder storage_pointers.StorageBuilder
	pointersBuilder      pointers.Builder
	pointerBuilder       pointers.PointerBuilder
	delimiterBuilder     delimiters.DelimiterBuilder
	pNextIndex           *uint64
}

func createBuilder(
	dbApp databases.Application,
	storagePointerBulder storage_pointers.StorageBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
) Builder {
	out := builder{
		dbApp:                dbApp,
		storagePointerBulder: storagePointerBulder,
		pointersBuilder:      pointersBuilder,
		pointerBuilder:       pointerBuilder,
		delimiterBuilder:     delimiterBuilder,
		pNextIndex:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.dbApp,
		app.storagePointerBulder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.delimiterBuilder,
	)
}

// WithNextIndex adds a nextIndex to the builder
func (app *builder) WithNextIndex(nextIndex uint64) Builder {
	app.pNextIndex = &nextIndex
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.pNextIndex == nil {
		return nil, errors.New("the next index is mandatory in order to build the Application")
	}

	return createApplication(
		app.dbApp,
		app.storagePointerBulder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.delimiterBuilder,
		*app.pNextIndex,
	), nil
}
