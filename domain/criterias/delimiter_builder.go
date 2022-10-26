package criterias

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type delimiterBuilder struct {
	hashAdapter hash.Adapter
	pIndex      *uint
	pAmount     *uint
}

func createDelimiterBuilder(
	hashAdapter hash.Adapter,
) DelimiterBuilder {
	out := delimiterBuilder{
		hashAdapter: hashAdapter,
		pIndex:      nil,
		pAmount:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *delimiterBuilder) Create() DelimiterBuilder {
	return createDelimiterBuilder(app.hashAdapter)
}

// WithIndex adds an index to the builder
func (app *delimiterBuilder) WithIndex(index uint) DelimiterBuilder {
	app.pIndex = &index
	return app
}

// WithAmount adds an amount to the builder
func (app *delimiterBuilder) WithAmount(amount uint) DelimiterBuilder {
	app.pAmount = &amount
	return app
}

// Now builds a new Delimiter instance
func (app *delimiterBuilder) Now() (Delimiter, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Delimiter instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pIndex)),
	}

	if app.pAmount != nil {
		data = append(data, []byte(fmt.Sprintf("%d", *app.pAmount)))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if app.pAmount != nil {
		return createDelimiterWithAmount(*pHash, *app.pIndex, app.pAmount), nil
	}

	return createDelimiter(*pHash, *app.pIndex), nil
}
