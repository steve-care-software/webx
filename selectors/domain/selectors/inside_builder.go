package selectors

import "errors"

type insideBuilder struct {
	fn       ContentFn
	fetchers Fetchers
}

func createInsideBuilder() InsideBuilder {
	out := insideBuilder{
		fn:       nil,
		fetchers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *insideBuilder) Create() InsideBuilder {
	return createInsideBuilder()
}

// WithFn adds a func to the Builder
func (app *insideBuilder) WithFn(fn ContentFn) InsideBuilder {
	app.fn = fn
	return app
}

// WithFetchers add fetchers to the Builder
func (app *insideBuilder) WithFetchers(fetchers Fetchers) InsideBuilder {
	app.fetchers = fetchers
	return app
}

// Now builds a new Inside instance
func (app *insideBuilder) Now() (Inside, error) {
	if app.fn != nil {
		return createInsideWithFunc(app.fn), nil
	}

	if app.fetchers != nil {
		return createInsideWithFetchers(app.fetchers), nil
	}

	return nil, errors.New("the Inside is invalid")
}
