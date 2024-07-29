package retrievals

import "errors"

type retrievalRetrievalBuilder struct {
	pIndex *uint64
	length uint64
}

func createRetrievalBuilder() RetrievalBuilder {
	out := retrievalRetrievalBuilder{
		pIndex: nil,
		length: 0,
	}

	return &out
}

// Create initializes the retrievalRetrievalBuilder
func (app *retrievalRetrievalBuilder) Create() RetrievalBuilder {
	return createRetrievalBuilder()
}

// WithIndex adds an index to the retrievalRetrievalBuilder
func (app *retrievalRetrievalBuilder) WithIndex(index uint64) RetrievalBuilder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the retrievalRetrievalBuilder
func (app *retrievalRetrievalBuilder) WithLength(length uint64) RetrievalBuilder {
	app.length = length
	return app
}

// Now builds a new Lister instance
func (app *retrievalRetrievalBuilder) Now() (Retrieval, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Lister instance")
	}

	if app.length == 0 {
		return nil, errors.New("the length is mandatory in order to build a Lister instance")
	}

	return createRetrieval(
		*app.pIndex,
		app.length,
	), nil
}
