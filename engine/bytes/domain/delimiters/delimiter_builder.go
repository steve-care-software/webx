package delimiters

import "errors"

type delimiterBuilder struct {
	pIndex *uint64
	length uint64
}

func createDelimiterBuilder() DelimiterBuilder {
	out := delimiterBuilder{
		pIndex: nil,
		length: 0,
	}

	return &out
}

// Create initializes the delimiterBuilder
func (app *delimiterBuilder) Create() DelimiterBuilder {
	return createDelimiterBuilder()
}

// WithIndex adds an index to the delimiterBuilder
func (app *delimiterBuilder) WithIndex(index uint64) DelimiterBuilder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the delimiterBuilder
func (app *delimiterBuilder) WithLength(length uint64) DelimiterBuilder {
	app.length = length
	return app
}

// Now builds a new Lister instance
func (app *delimiterBuilder) Now() (Delimiter, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Delimiter instance")
	}

	if app.length == 0 {
		return nil, errors.New("the length is mandatory in order to build a Delimiter instance")
	}

	return createDelimiter(
		*app.pIndex,
		app.length,
	), nil
}
