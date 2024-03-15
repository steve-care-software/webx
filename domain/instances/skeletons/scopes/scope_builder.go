package scopes

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type scopeBuilder struct {
	hashAdapter hash.Adapter
	prefix      []string
}

func createScopeBuilder(
	hashAdapter hash.Adapter,
) ScopeBuilder {
	out := scopeBuilder{
		hashAdapter: hashAdapter,
		prefix:      nil,
	}

	return &out
}

// Create intiializes the builder
func (app *scopeBuilder) Create() ScopeBuilder {
	return createScopeBuilder(
		app.hashAdapter,
	)
}

// WithPrefix adds a prefix to the builder
func (app *scopeBuilder) WithPrefix(prefix []string) ScopeBuilder {
	app.prefix = prefix
	return app
}

// Now builds a new Scope instance
func (app *scopeBuilder) Now() (Scope, error) {
	if app.prefix != nil && len(app.prefix) <= 0 {
		app.prefix = nil
	}

	if app.prefix == nil {
		return nil, errors.New("the prefix is mandatory in order to build a Scope instance")
	}

	data := [][]byte{}
	for _, oneName := range app.prefix {
		data = append(data, []byte(oneName))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createScope(*pHash, app.prefix), nil
}
