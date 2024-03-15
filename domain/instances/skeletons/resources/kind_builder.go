package resources

import (
	"errors"
	"strings"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type kindBuilder struct {
	hashAdapter hash.Adapter
	native      Native
	reference   []string
	connection  string
}

func createKindBuilder(
	hashAdapter hash.Adapter,
) KindBuilder {
	out := kindBuilder{
		hashAdapter: hashAdapter,
		native:      nil,
		reference:   nil,
		connection:  "",
	}

	return &out
}

// Create initializes the builder
func (app *kindBuilder) Create() KindBuilder {
	return createKindBuilder(
		app.hashAdapter,
	)
}

// WithNative adds a native to the builder
func (app *kindBuilder) WithNative(native Native) KindBuilder {
	app.native = native
	return app
}

// WithReference adds a reference to the builder
func (app *kindBuilder) WithReference(reference []string) KindBuilder {
	app.reference = reference
	return app
}

// WithConnection adds a connection to the builder
func (app *kindBuilder) WithConnection(connection string) KindBuilder {
	app.connection = connection
	return app
}

// Now builds a new Kind isntance
func (app *kindBuilder) Now() (Kind, error) {
	if app.reference != nil && len(app.reference) <= 0 {
		app.reference = nil
	}

	data := [][]byte{}
	if app.native != nil {
		data = append(data, []byte("native"))
		data = append(data, app.native.Hash().Bytes())
	}

	if app.reference != nil {
		data = append(data, []byte("reference"))
		data = append(data, []byte(strings.Join(app.reference, "")))
	}

	if app.connection != "" {
		data = append(data, []byte("connection"))
		data = append(data, []byte(app.connection))
	}

	if len(data) != 2 {
		return nil, errors.New("the Kind is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.native != nil {
		return createKindWithNative(*pHash, app.native), nil
	}

	if app.reference != nil {
		return createKindWithReference(*pHash, app.reference), nil
	}

	if app.connection != "" {
		return createKindWithConnection(*pHash, app.connection), nil
	}

	return nil, errors.New("the Kind is invalid")
}
