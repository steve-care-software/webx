package grammars

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	cardinality cardinalities.Cardinality
	value       values.Value
	external    External
	instance    Instance
	recursive   string
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		cardinality: nil,
		value:       nil,
		external:    nil,
		instance:    nil,
		recursive:   "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithCardinality adds a cardinality to the builder
func (app *elementBuilder) WithCardinality(cardinality cardinalities.Cardinality) ElementBuilder {
	app.cardinality = cardinality
	return app
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value values.Value) ElementBuilder {
	app.value = value
	return app
}

// WithExternal adds an external grammar to the builder
func (app *elementBuilder) WithExternal(external External) ElementBuilder {
	app.external = external
	return app
}

// WithInstance adds an instance to the builder
func (app *elementBuilder) WithInstance(instance Instance) ElementBuilder {
	app.instance = instance
	return app
}

// WithRecursive adds a recursive to the builder
func (app *elementBuilder) WithRecursive(recursive string) ElementBuilder {
	app.recursive = recursive
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build an Element instance")
	}

	data := [][]byte{
		app.cardinality.Hash().Bytes(),
	}

	var content ElementContent
	if app.value != nil {
		hash := app.value.Hash()
		content = createElementContentWithValue(hash, app.value)
		data = append(data, content.Hash().Bytes())
	}

	if app.external != nil {
		hash := app.external.Hash()
		content = createElementContentWithExternalToken(hash, app.external)
		data = append(data, content.Hash().Bytes())
	}

	if app.instance != nil {
		hash := app.instance.Hash()
		content = createElementContentWithInstance(hash, app.instance)
		data = append(data, content.Hash().Bytes())
	}

	if app.recursive != "" {
		pHash, err := app.hashAdapter.FromBytes([]byte(app.recursive))
		if err != nil {
			return nil, err
		}

		content = createElementContentWithRecursive(*pHash, app.recursive)
		data = append(data, content.Hash().Bytes())
	}

	if content == nil {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createElement(*pHash, content, app.cardinality), nil
}
