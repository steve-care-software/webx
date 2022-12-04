package elements

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	pHash       *hash.Hash
	cardinality Cardinality
	pValue      *uint8
	pExternal   *hash.Hash
	pToken      *hash.Hash
	pEverything *hash.Hash
	pRecursive  *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:       nil,
		cardinality: nil,
		pValue:      nil,
		pExternal:   nil,
		pToken:      nil,
		pEverything: nil,
		pRecursive:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *builder) WithCardinality(cardinality Cardinality) Builder {
	app.cardinality = cardinality
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value uint8) Builder {
	app.pValue = &value
	return app
}

// WithExternal adds an pExternal to the builder
func (app *builder) WithExternal(external hash.Hash) Builder {
	app.pExternal = &external
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token hash.Hash) Builder {
	app.pToken = &token
	return app
}

// WithEverything adds an everything to the builder
func (app *builder) WithEverything(everything hash.Hash) Builder {
	app.pEverything = &everything
	return app
}

// WithRecursive adds a recursive to the builder
func (app *builder) WithRecursive(recursive hash.Hash) Builder {
	app.pRecursive = &recursive
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Element instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build an Element instance")
	}

	if app.pValue != nil {
		content := createContentWithValue(app.pValue)
		return createElement(*app.pHash, app.cardinality, content), nil
	}

	if app.pExternal != nil {
		content := createContentWithExternal(app.pExternal)
		return createElement(*app.pHash, app.cardinality, content), nil
	}

	if app.pToken != nil {
		content := createContentWithToken(app.pToken)
		return createElement(*app.pHash, app.cardinality, content), nil
	}

	if app.pEverything != nil {
		content := createContentWithEverything(app.pEverything)
		return createElement(*app.pHash, app.cardinality, content), nil
	}

	if app.pRecursive != nil {
		content := createContentWithRecursive(app.pRecursive)
		return createElement(*app.pHash, app.cardinality, content), nil
	}

	return nil, errors.New("the Element is invalid")
}
