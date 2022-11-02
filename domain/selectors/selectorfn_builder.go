package selectors

import "errors"

type selectorFnBuilder struct {
	single SingleSelectorFn
	multi  MultiSelectorFn
}

func createSelectorFnBuilder() SelectorFnBuilder {
	out := selectorFnBuilder{
		single: nil,
		multi:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *selectorFnBuilder) Create() SelectorFnBuilder {
	return createSelectorFnBuilder()
}

// WithSingle adds a single func to the builder
func (app *selectorFnBuilder) WithSingle(single SingleSelectorFn) SelectorFnBuilder {
	app.single = single
	return app
}

// WithMulti adds a multi func to the builder
func (app *selectorFnBuilder) WithMulti(multi MultiSelectorFn) SelectorFnBuilder {
	app.multi = multi
	return app
}

// Now builds a new Selector func
func (app *selectorFnBuilder) Now() (SelectorFn, error) {
	if app.single != nil {
		return createSelectorFnWithSingle(app.single), nil
	}

	if app.multi != nil {
		return createSelectorFnWithMulti(app.multi), nil
	}

	return nil, errors.New("the SelectorFn is invalid")
}
