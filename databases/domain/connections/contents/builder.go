package contents

import "errors"

type builder struct {
	list []Content
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add contents to the builder
func (app *builder) WithList(list []Content) Builder {
	app.list = list
	return app
}

// Now builds Contents instance
func (app *builder) Now() (Contents, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Content in order to build a Contents instance")
	}

	mp := map[string]Content{}
	for _, oneContent := range app.list {
		name := oneContent.Hash().String()
		mp[name] = oneContent
	}

	return createContents(mp, app.list), nil
}
