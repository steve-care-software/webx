package selectors

import "errors"

type fetchersBuilder struct {
	list []Fetcher
}

func createFetchersBuilder() FetchersBuilder {
	out := fetchersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *fetchersBuilder) Create() FetchersBuilder {
	return createFetchersBuilder()
}

// WithList adds a list to the builder
func (app *fetchersBuilder) WithList(list []Fetcher) FetchersBuilder {
	app.list = list
	return app
}

// Now builds a new Fetchers instance
func (app *fetchersBuilder) Now() (Fetchers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Fetcher in order to build a Fetchers instance")
	}

	return createFetchers(app.list), nil
}
