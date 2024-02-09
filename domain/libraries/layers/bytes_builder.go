package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type bytesBuilder struct {
	hashAdapter hash.Adapter
	join        []string
	compare     []string
	hashBytes   string
}

func createBytesBuilder(
	hashAdapter hash.Adapter,
) BytesBuilder {
	out := bytesBuilder{
		hashAdapter: hashAdapter,
		join:        nil,
		compare:     nil,
		hashBytes:   "",
	}

	return &out
}

// Create initializes the builder
func (app *bytesBuilder) Create() BytesBuilder {
	return createBytesBuilder(
		app.hashAdapter,
	)
}

// WithJoin adds a join to the builder
func (app *bytesBuilder) WithJoin(join []string) BytesBuilder {
	app.join = join
	return app
}

// WithCompare adds a compare to the builder
func (app *bytesBuilder) WithCompare(compare []string) BytesBuilder {
	app.compare = compare
	return app
}

// WithHashBytes adds an hashBytes to the builder
func (app *bytesBuilder) WithHashBytes(hashBytes string) BytesBuilder {
	app.hashBytes = hashBytes
	return app
}

// Now builds a new Bytes instance
func (app *bytesBuilder) Now() (Bytes, error) {
	data := [][]byte{}
	if app.join != nil && len(app.join) <= 0 {
		app.join = nil
	}

	if app.join != nil {
		data = append(data, []byte("join"))
		for _, oneVariable := range app.join {
			data = append(data, []byte(oneVariable))
		}
	}

	if app.compare != nil && len(app.compare) <= 0 {
		app.compare = nil
	}

	if app.compare != nil {
		data = append(data, []byte("compare"))
		for _, oneVariable := range app.compare {
			data = append(data, []byte(oneVariable))
		}
	}

	if app.hashBytes != "" {
		data = append(data, []byte("hash"))
		data = append(data, []byte(app.hashBytes))
	}

	if len(data) <= 0 {
		return nil, errors.New("the Bytes is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.join != nil {
		return createBytesWithJoin(*pHash, app.join), nil
	}

	if app.hashBytes != "" {
		return createBytesWithHashBytes(*pHash, app.hashBytes), nil
	}

	return createBytesWithCompare(*pHash, app.compare), nil
}
