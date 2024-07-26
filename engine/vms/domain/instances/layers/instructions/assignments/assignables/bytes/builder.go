package bytes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	join        []string
	compare     []string
	hashBytes   string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		join:        nil,
		compare:     nil,
		hashBytes:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithJoin adds a join to the builder
func (app *builder) WithJoin(join []string) Builder {
	app.join = join
	return app
}

// WithCompare adds a compare to the builder
func (app *builder) WithCompare(compare []string) Builder {
	app.compare = compare
	return app
}

// WithHashBytes adds an hashBytes to the builder
func (app *builder) WithHashBytes(hashBytes string) Builder {
	app.hashBytes = hashBytes
	return app
}

// Now builds a new Bytes instance
func (app *builder) Now() (Bytes, error) {
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
