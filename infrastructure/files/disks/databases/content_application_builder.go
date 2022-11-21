package databases

import (
	"errors"

	"github.com/steve-care-software/webx/applications/databases/contents"
)

type contentApplicationBuilder struct {
	name string
}

func createContentApplicationBuilder() contents.Builder {
	out := contentApplicationBuilder{}
	return &out
}

// Create initializes the builder
func (app *contentApplicationBuilder) Create() contents.Builder {
	return createContentApplicationBuilder()
}

// WithName adds a name to the builder
func (app *contentApplicationBuilder) WithName(name string) contents.Builder {
	app.name = name
	return app
}

// Now builds a new content application
func (app *contentApplicationBuilder) Now() (contents.Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a content application")
	}

	return createContentApplication(app.name), nil
}
