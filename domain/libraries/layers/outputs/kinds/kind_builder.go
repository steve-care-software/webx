package kinds

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type kindBuilder struct {
	hashAdapter hash.Adapter
	isPrompt    bool
	isContinue  bool
}

func createKindBuilder(
	hashAdapter hash.Adapter,
) KindBuilder {
	out := kindBuilder{
		hashAdapter: hashAdapter,
		isPrompt:    false,
		isContinue:  false,
	}

	return &out
}

// Create initializes the builder
func (app *kindBuilder) Create() KindBuilder {
	return createKindBuilder(
		app.hashAdapter,
	)
}

// IsPrompt flags the builder as a prompt
func (app *kindBuilder) IsPrompt() KindBuilder {
	app.isPrompt = true
	return app
}

// IsContinue flags the builder as a continue
func (app *kindBuilder) IsContinue() KindBuilder {
	app.isContinue = true
	return app
}

// Now builds a new Kind instance
func (app *kindBuilder) Now() (Kind, error) {
	data := [][]byte{}
	if app.isPrompt {
		data = append(data, []byte("isPrompt"))
	}

	if app.isContinue {
		data = append(data, []byte("isContinue"))
	}

	if len(data) <= 0 {
		return nil, errors.New("the Kind is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isPrompt {
		return createKindWithPrompt(*pHash), nil
	}

	return createKindWithContinue(*pHash), nil
}
