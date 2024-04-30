package pointers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	entity      string
	field       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		entity:      "",
		field:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEntity adds a entity to the builder
func (app *builder) WithEntity(entity string) Builder {
	app.entity = entity
	return app
}

// WithField adds a field to the builder
func (app *builder) WithField(field string) Builder {
	app.field = field
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.entity == "" {
		return nil, errors.New("the entity is mandatory in order to build a Pointer instance")
	}

	if app.field == "" {
		return nil, errors.New("the field is mandatory in order to build a Pointer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.entity),
		[]byte(app.field),
	})

	if err != nil {
		return nil, err
	}

	return createPointer(*pHash, app.entity, app.field), nil
}
