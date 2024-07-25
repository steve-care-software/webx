package routes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	layer       hash.Hash
	bytes       []byte
	str         string
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		layer:       nil,
		bytes:       nil,
		str:         "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *elementBuilder) WithLayer(layer hash.Hash) ElementBuilder {
	app.layer = layer
	return app
}

// WithBytes add bytes to the builder
func (app *elementBuilder) WithBytes(bytes []byte) ElementBuilder {
	app.bytes = bytes
	return app
}

// WithString add string to the builder
func (app *elementBuilder) WithString(str string) ElementBuilder {
	app.str = str
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.layer != nil && len(app.layer) <= 0 {
		app.layer = nil
	}

	data := [][]byte{}
	if app.layer != nil {
		data = append(data, []byte("layer"))
		data = append(data, app.layer.Bytes())
	}

	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes)
	}

	if app.str != "" {
		data = append(data, []byte("string"))
		data = append(data, []byte(app.str))
	}

	if len(data) != 2 {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.layer != nil {
		return createElementWithLayer(*pHash, app.layer), nil
	}

	if app.bytes != nil {
		return createElementWithBytes(*pHash, app.bytes), nil
	}

	return createElementWithString(*pHash, app.str), nil
}
