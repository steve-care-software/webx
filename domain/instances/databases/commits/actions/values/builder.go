package values

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values/transforms"
)

type builder struct {
	hashAdapter hash.Adapter
	instance    instances.Instance
	transform   transforms.Transform
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		instance:    nil,
		transform:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithInstance adds an instance to the builder
func (app *builder) WithInstance(instance instances.Instance) Builder {
	app.instance = instance
	return app
}

// WithTransform adds a transform to the builder
func (app *builder) WithTransform(transform transforms.Transform) Builder {
	app.transform = transform
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	data := [][]byte{}
	if app.instance != nil {
		data = append(data, []byte("instance"))
		data = append(data, app.instance.Hash().Bytes())
	}

	if app.transform != nil {
		data = append(data, []byte("transform"))
		data = append(data, app.transform.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Value is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.instance != nil {
		return createValueWithInstance(*pHash, app.instance), nil
	}

	return createValueWithTransform(*pHash, app.transform), nil
}
