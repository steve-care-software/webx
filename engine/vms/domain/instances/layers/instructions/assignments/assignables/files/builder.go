package files

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/files/opens"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/files/reads"
)

type builder struct {
	hashAdapter hash.Adapter
	open        opens.Open
	read        reads.Read
	exists      string
	length      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		open:        nil,
		read:        nil,
		exists:      "",
		length:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOpen adds an open to the builder
func (app *builder) WithOpen(open opens.Open) Builder {
	app.open = open
	return app
}

// WithRead adds a read to the builder
func (app *builder) WithRead(read reads.Read) Builder {
	app.read = read
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length string) Builder {
	app.length = length
	return app
}

// WithExists adds an exists to the builder
func (app *builder) WithExists(exists string) Builder {
	app.exists = exists
	return app
}

// Now builds a new File instance
func (app *builder) Now() (File, error) {
	data := [][]byte{}
	if app.open != nil {
		data = append(data, []byte("open"))
		data = append(data, app.open.Hash().Bytes())
	}

	if app.read != nil {
		data = append(data, []byte("read"))
		data = append(data, app.read.Hash().Bytes())
	}

	if app.length != "" {
		data = append(data, []byte("length"))
		data = append(data, []byte(app.length))
	}

	if app.exists != "" {
		data = append(data, []byte("exists"))
		data = append(data, []byte(app.exists))
	}

	if len(data) != 2 {
		return nil, errors.New("the File is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.open != nil {
		return createFileWithOpen(*pHash, app.open), nil
	}

	if app.read != nil {
		return createFileWithRead(*pHash, app.read), nil
	}

	if app.length != "" {
		return createFileWithLength(*pHash, app.length), nil
	}

	return createFileWithExists(*pHash, app.exists), nil

}
