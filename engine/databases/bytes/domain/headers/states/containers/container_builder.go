package containers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"
)

type containerBuilder struct {
	keyname  string
	pointers pointers.Pointers
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		keyname:  "",
		pointers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithKeyname adds a keyname to the builder
func (app *containerBuilder) WithKeyname(keyname string) ContainerBuilder {
	app.keyname = keyname
	return app
}

// WithPointers add pointers to the builder
func (app *containerBuilder) WithPointers(pointers pointers.Pointers) ContainerBuilder {
	app.pointers = pointers
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build a Container instance")
	}

	if app.pointers == nil {
		return nil, errors.New("the pointers is mandatory in order to build a Container instance")
	}

	return createContainer(
		app.keyname,
		app.pointers,
	), nil
}
