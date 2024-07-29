package listers

import "errors"

type builder struct {
	keyname string
	pIndex  *uint64
	length  uint64
}

func createBuilder() Builder {
	out := builder{
		keyname: "",
		pIndex:  nil,
		length:  0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithKeyname adds a keyname to the builder
func (app *builder) WithKeyname(keyname string) Builder {
	app.keyname = keyname
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint64) Builder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length uint64) Builder {
	app.length = length
	return app
}

// Now builds a new Lister instance
func (app *builder) Now() (Lister, error) {
	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build a Lister instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Lister instance")
	}

	if app.length == 0 {
		return nil, errors.New("the length is mandatory in order to build a Lister instance")
	}

	return createLister(
		app.keyname,
		*app.pIndex,
		app.length,
	), nil
}
