package scopes

import "github.com/steve-care-software/datastencil/domain/hash"

type builder struct {
	hashAdapter hash.Adapter
	list        []Scope
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Scope) Builder {
	app.list = list
	return app
}

// Now builds a new Scope instance
func (app *builder) Now() (Scopes, error) {
	data := [][]byte{}
	for _, oneScope := range app.list {
		data = append(data, oneScope.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createScopes(*pHash, app.list), nil
}
