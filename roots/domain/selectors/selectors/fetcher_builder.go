package selectors

import "errors"

type fetcherBuilder struct {
	recursive string
	selector  Selector
}

func createFetcherBuilder() FetcherBuilder {
	out := fetcherBuilder{
		recursive: "",
		selector:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *fetcherBuilder) Create() FetcherBuilder {
	return createFetcherBuilder()
}

// WithRecursive adds a recursive selector's token name to the builder
func (app *fetcherBuilder) WithRecursive(recursive string) FetcherBuilder {
	app.recursive = recursive
	return app
}

// WithSelector adds a selector to the builder
func (app *fetcherBuilder) WithSelector(selector Selector) FetcherBuilder {
	app.selector = selector
	return app
}

// Now builds a new Fetcher instance
func (app *fetcherBuilder) Now() (Fetcher, error) {
	if app.recursive != "" {
		return createFetcherWithRecursive(app.recursive), nil
	}

	if app.selector != nil {
		return createFetcherWithSelector(app.selector), nil
	}

	return nil, errors.New("the Fetcher is invalid")
}
