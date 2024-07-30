package delimiters

import "errors"

type delimiterDelimiterBuilder struct {
	pIndex *uint64
	length uint64
}

func createDelimiterBuilder() DelimiterBuilder {
	out := delimiterDelimiterBuilder{
		pIndex: nil,
		length: 0,
	}

	return &out
}

// Create initializes the delimiterDelimiterBuilder
func (app *delimiterDelimiterBuilder) Create() DelimiterBuilder {
	return createDelimiterBuilder()
}

// WithIndex adds an index to the delimiterDelimiterBuilder
func (app *delimiterDelimiterBuilder) WithIndex(index uint64) DelimiterBuilder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the delimiterDelimiterBuilder
func (app *delimiterDelimiterBuilder) WithLength(length uint64) DelimiterBuilder {
	app.length = length
	return app
}

// Now builds a new Lister instance
func (app *delimiterDelimiterBuilder) Now() (Delimiter, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Lister instance")
	}

	if app.length == 0 {
		return nil, errors.New("the length is mandatory in order to build a Lister instance")
	}

	return createDelimiter(
		*app.pIndex,
		app.length,
	), nil
}
