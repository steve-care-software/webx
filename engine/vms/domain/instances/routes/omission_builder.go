package routes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type omissionBuilder struct {
	hashAdapter hash.Adapter
	prefix      Element
	suffix      Element
}

func createOmissionBuilder(
	hashAdapter hash.Adapter,
) OmissionBuilder {
	out := omissionBuilder{
		hashAdapter: hashAdapter,
		prefix:      nil,
		suffix:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *omissionBuilder) Create() OmissionBuilder {
	return createOmissionBuilder(
		app.hashAdapter,
	)
}

// WithPrefix adds a prefix to the builder
func (app *omissionBuilder) WithPrefix(prefix Element) OmissionBuilder {
	app.prefix = prefix
	return app
}

// WithSuffix adds a suffix to the builder
func (app *omissionBuilder) WithSuffix(suffix Element) OmissionBuilder {
	app.suffix = suffix
	return app
}

// Now builds a new Omission instance
func (app *omissionBuilder) Now() (Omission, error) {
	data := [][]byte{}
	if app.prefix != nil {
		data = append(data, []byte("prefix"))
		data = append(data, app.prefix.Bytes())
	}

	if app.suffix != nil {
		data = append(data, []byte("suffix"))
		data = append(data, app.suffix.Bytes())
	}

	if len(data) != 2 {
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
