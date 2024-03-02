package resources

import "errors"

type nativeBuilder struct {
	pSingle *uint8
	list    List
}

func createNativeBuilder() NativeBuilder {
	out := nativeBuilder{
		pSingle: nil,
		list:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *nativeBuilder) Create() NativeBuilder {
	return createNativeBuilder()
}

// WithSingle adds a single to the builder
func (app *nativeBuilder) WithSingle(single uint8) NativeBuilder {
	app.pSingle = &single
	return app
}

// WithList adds a list to the builder
func (app *nativeBuilder) WithList(list List) NativeBuilder {
	app.list = list
	return app
}

// Now builds a new Native instance
func (app *nativeBuilder) Now() (Native, error) {
	if app.pSingle != nil {
		return createNativeWithSingle(app.pSingle), nil
	}

	if app.list != nil {
		return createNativeWithList(app.list), nil
	}

	return nil, errors.New("the Native is invalid")
}
