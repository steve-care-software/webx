package signs

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

type builder struct {
	hashAdapter hash.Adapter
	create      creates.Create
	validate    validates.Validate
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		create:      nil,
		validate:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCreate adds a create to the builder
func (app *builder) WithCreate(create creates.Create) Builder {
	app.create = create
	return app
}

// WithValidate adds a validate to the builder
func (app *builder) WithValidate(validate validates.Validate) Builder {
	app.validate = validate
	return app
}

// Now builds a new Sign instance
func (app *builder) Now() (Sign, error) {
	data := [][]byte{}
	if app.create != nil {
		data = append(data, []byte("create"))
		data = append(data, app.create.Hash().Bytes())
	}

	if app.validate != nil {
		data = append(data, []byte("validate"))
		data = append(data, app.validate.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Sign is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.create != nil {
		return createSignWithCreate(*pHash, app.create), nil
	}

	return createSignWithValidate(*pHash, app.validate), nil
}
