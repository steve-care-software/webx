package databases

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/services"
)

type builder struct {
	hashAdapter hash.Adapter
	repository  repositories.Repository
	service     services.Service
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		repository:  nil,
		service:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithRepository adds a repository to the builder
func (app *builder) WithRepository(repository repositories.Repository) Builder {
	app.repository = repository
	return app
}

// WithService adds a service to the builder
func (app *builder) WithService(service services.Service) Builder {
	app.service = service
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	data := [][]byte{}
	if app.repository != nil {
		data = append(data, []byte("repository"))
		data = append(data, app.repository.Hash().Bytes())
	}

	if app.service != nil {
		data = append(data, []byte("service"))
		data = append(data, app.service.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Database is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.repository != nil {
		return createDatabaseWithRepository(*pHash, app.repository), nil
	}

	return createDatabaseWithService(*pHash, app.service), nil
}
