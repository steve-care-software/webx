package trees

import "errors"

type contentsBuilder struct {
	list []Content
}

func createContentsBuilder() ContentsBuilder {
	out := contentsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentsBuilder) Create() ContentsBuilder {
	return createContentsBuilder()
}

// WithList adds a list to the builder
func (app *contentsBuilder) WithList(list []Content) ContentsBuilder {
	app.list = list
	return app
}

// Now builds a new Contents instance
func (app *contentsBuilder) Now() (Contents, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Content in order to build a Contents instance")
	}

	return createContents(app.list), nil
}
