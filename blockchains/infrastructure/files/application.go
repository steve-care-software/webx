package files

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/contents/references"
)

type application struct {
	dirPath string
}

func createApplication(
	dirPath string,
) applications.Application {
	out := application{
		dirPath: dirPath,
	}

	return &out
}

// Open opens a context
func (app *application) Open(name string) (*uint, error) {
	return nil, nil
}

// Read reads a pointer on a context
func (app *application) Read(context uint, pointer references.Pointer) ([]byte, error) {
	return nil, nil
}

// ReadAll read pointers on a context
func (app *application) ReadAll(context uint, pointers []references.Pointer) ([][]byte, error) {
	return nil, nil
}

// Write writes data to a context
func (app *application) Write(data []byte) error {
	return nil
}

// WriteAll writes a list of data to a context
func (app *application) WriteAll(data [][]byte) error {
	return nil
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	return nil
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return nil
}

// Push pushes a context
func (app *application) Push(context uint) error {
	return nil
}

// Close closes a context
func (app *application) Close(context uint) error {
	return nil
}
