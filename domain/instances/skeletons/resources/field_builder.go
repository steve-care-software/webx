package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type fieldBuilder struct {
	hashAdapter hash.Adapter
	name        string
	kind        Kind
	canBeNil    bool
}

func createFieldBuilder(
	hashAdapter hash.Adapter,
) FieldBuilder {
	out := fieldBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		kind:        nil,
		canBeNil:    false,
	}

	return &out
}

// Create initializes the builder
func (app *fieldBuilder) Create() FieldBuilder {
	return createFieldBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *fieldBuilder) WithName(name string) FieldBuilder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *fieldBuilder) WithKind(kind Kind) FieldBuilder {
	app.kind = kind
	return app
}

// CanBeNil flags the builder as canBeNil
func (app *fieldBuilder) CanBeNil() FieldBuilder {
	app.canBeNil = true
	return app
}

// Now builds a new Field instance
func (app *fieldBuilder) Now() (Field, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Field instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Field instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createField(
		*pHash,
		app.name,
		app.kind,
		app.canBeNil,
	), nil
}
