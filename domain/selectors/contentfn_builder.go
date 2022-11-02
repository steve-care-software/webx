package selectors

import "errors"

type contentFnBuilder struct {
	single SingleContentFn
	multi  MultiContentFn
}

func createContentFnBuilder() ContentFnBuilder {
	out := contentFnBuilder{
		single: nil,
		multi:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentFnBuilder) Create() ContentFnBuilder {
	return createContentFnBuilder()
}

// WithSingle adds a single func to the builder
func (app *contentFnBuilder) WithSingle(single SingleContentFn) ContentFnBuilder {
	app.single = single
	return app
}

// WithMulti adds a multi func to the builder
func (app *contentFnBuilder) WithMulti(multi MultiContentFn) ContentFnBuilder {
	app.multi = multi
	return app
}

// Now builds a new Content func
func (app *contentFnBuilder) Now() (ContentFn, error) {
	if app.single != nil {
		return createContentFnWithSingle(app.single), nil
	}

	if app.multi != nil {
		return createContentFnWithMulti(app.multi), nil
	}

	return nil, errors.New("the ContentFn is invalid")
}
