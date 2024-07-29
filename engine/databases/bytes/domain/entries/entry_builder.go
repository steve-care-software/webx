package entries

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"
)

type entryBuilder struct {
	pointer pointers.Pointer
	bytes   []byte
}

func createEntryBuilder() EntryBuilder {
	out := entryBuilder{
		pointer: nil,
		bytes:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *entryBuilder) Create() EntryBuilder {
	return createEntryBuilder()
}

// WithPointer adds a pointer to the builder
func (app *entryBuilder) WithPointer(pointer pointers.Pointer) EntryBuilder {
	app.pointer = pointer
	return app
}

// WithBytes add bytes to the builder
func (app *entryBuilder) WithBytes(bytes []byte) EntryBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Entry instance
func (app *entryBuilder) Now() (Entry, error) {
	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build an Entry instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build an Entry instance")
	}

	return createEntry(
		app.pointer,
		app.bytes,
	), nil
}
