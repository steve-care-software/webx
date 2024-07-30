package entries

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"
)

type entryBuilder struct {
	keyname   string
	delimiter delimiters.Delimiter
	bytes     []byte
}

func createEntryBuilder() EntryBuilder {
	out := entryBuilder{
		keyname:   "",
		delimiter: nil,
		bytes:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *entryBuilder) Create() EntryBuilder {
	return createEntryBuilder()
}

// WithKeyname adds a keyname to the builder
func (app *entryBuilder) WithKeyname(keyname string) EntryBuilder {
	app.keyname = keyname
	return app
}

// WithDelimiter adds a delimiter to the builder
func (app *entryBuilder) WithDelimiter(delimiter delimiters.Delimiter) EntryBuilder {
	app.delimiter = delimiter
	return app
}

// WithBytes add bytes to the builder
func (app *entryBuilder) WithBytes(bytes []byte) EntryBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Entry instance
func (app *entryBuilder) Now() (Entry, error) {
	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build an Entry instance")
	}

	if app.delimiter == nil {
		return nil, errors.New("the delimiter is mandatory in order to build an Entry instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build an Entry instance")
	}

	return createEntry(
		app.keyname,
		app.delimiter,
		app.bytes,
	), nil
}
