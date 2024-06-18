package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/resources/logics"
)

type resourceBuilder struct {
	hashAdapter hash.Adapter
	database    heads.Head
	logics      logics.Logics
}

func createResourceBuilder(
	hashAdapter hash.Adapter,
) ResourceBuilder {
	out := resourceBuilder{
		hashAdapter: hashAdapter,
		database:    nil,
		logics:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder(
		app.hashAdapter,
	)
}

// WithDatabase adds a database to the builder
func (app *resourceBuilder) WithDatabase(database heads.Head) ResourceBuilder {
	app.database = database
	return app
}

// WithLogics add logics to the builder
func (app *resourceBuilder) WithLogics(logics logics.Logics) ResourceBuilder {
	app.logics = logics
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.database == nil {
		return nil, errors.New("the database is mandatory in order to build a Resource instance")
	}

	if app.logics == nil {
		return nil, errors.New("the logics is mandatory in order to build a Resource instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.database.Hash().Bytes(),
		app.logics.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.database, app.logics), nil
}
