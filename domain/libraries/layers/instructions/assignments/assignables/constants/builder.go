package constants

import "github.com/steve-care-software/datastencil/domain/hash"

type builder struct {
	hashAdapter hash.Adapter
	boolValue   *bool
	bytes       []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		boolValue:   nil,
		bytes:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBool adds a bool to the builder
func (app *builder) WithBool(boolValue bool) Builder {
	app.boolValue = &boolValue
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Constant instance
func (app *builder) Now() (Constant, error) {
	data := [][]byte{}
	if app.boolValue != nil {
		bytes := [][]byte{
			[]byte("false"),
		}

		val := *app.boolValue
		if val {
			bytes = [][]byte{
				[]byte("true"),
			}
		}

		bytes = append(bytes, []byte("bool"))
		data = append(data, bytes...)
	}

	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.boolValue != nil {
		return createConstantWithBool(*pHash, app.boolValue), nil
	}

	return createConstantWithData(*pHash, app.bytes), nil
}
