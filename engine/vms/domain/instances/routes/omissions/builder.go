package omissions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
)

type builder struct {
	hashAdapter hash.Adapter
	prefix      elements.Element
	suffix      elements.Element
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		prefix:      nil,
		suffix:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPrefix adds a prefix to the builder
func (app *builder) WithPrefix(prefix elements.Element) Builder {
	app.prefix = prefix
	return app
}

// WithSuffix adds a suffix to the builder
func (app *builder) WithSuffix(suffix elements.Element) Builder {
	app.suffix = suffix
	return app
}

// Now builds a new Omission instance
func (app *builder) Now() (Omission, error) {
	data := [][]byte{}
	if app.prefix != nil {
		data = append(data, []byte("prefix"))
		data = append(data, app.prefix.Bytes())
	}

	if app.suffix != nil {
		data = append(data, []byte("suffix"))
		data = append(data, app.suffix.Bytes())
	}

	length := len(data)
	if length != 2 && length != 4 {
		return nil, errors.New("the Omission is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.prefix != nil && app.suffix != nil {
		return createOmissionWithPrefixAndSuffix(*pHash, app.prefix, app.suffix), nil
	}

	if app.prefix != nil {
		return createOmissionWithPrefix(*pHash, app.prefix), nil
	}

	return createOmissionWithSuffix(*pHash, app.suffix), nil
}
