package updates

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates/criterias"
)

type builder struct {
	hashAdapter hash.Adapter
	credentials string
	criteria    criterias.Criteria
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		credentials: "",
		criteria:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCredentials add credentials to the builder
func (app *builder) WithCredentials(credentials string) Builder {
	app.credentials = credentials
	return app
}

// WithCriteria add criteria to the builder
func (app *builder) WithCriteria(criteria criterias.Criteria) Builder {
	app.criteria = criteria
	return app
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	if app.credentials == "" {
		return nil, errors.New("the credentials is mandatory in order to build an Update instance")
	}

	if app.criteria == nil {
		return nil, errors.New("the criteria is mandatory in order to build an Update instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.credentials),
		app.criteria.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createUpdate(*pHash, app.credentials, app.criteria), nil
}
