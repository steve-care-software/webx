package deletes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/databases/pointers"
)

type deleteBuilder struct {
	keyname string
	pointer pointers.Pointer
}

func createDeleteBuilder() DeleteBuilder {
	out := deleteBuilder{
		keyname: "",
		pointer: nil,
	}

	return &out
}

// Create initializes the builder
func (app *deleteBuilder) Create() DeleteBuilder {
	return createDeleteBuilder()
}

// WithKeyname adds a keyname to the builder
func (app *deleteBuilder) WithKeyname(keyname string) DeleteBuilder {
	app.keyname = keyname
	return app
}

// WithPointer adds a pointer to the builder
func (app *deleteBuilder) WithPointer(pointer pointers.Pointer) DeleteBuilder {
	app.pointer = pointer
	return app
}

// Now creates a new Delete instance
func (app *deleteBuilder) Now() (Delete, error) {
	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build a Delete instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Delete instance")
	}

	return createDelete(
		app.keyname,
		app.pointer,
	), nil
}
