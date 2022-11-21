package databases

import (
	contents_application "github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/domain/databases/references"
)

type contentApplication struct {
	absPath string
}

func createContentApplication(
	absPath string,
) contents_application.Application {
	out := contentApplication{
		absPath: absPath,
	}

	return &out
}

// Reference returns the reference
func (app *contentApplication) Reference() (references.Reference, error) {
	// open the file:

	// read the first byte and convert it to a uint8:

	// read the length of the reference's content:

	// read the reference's content:

	// convert the content to a Reference instance:
	return nil, nil
}

// Retrieve retrieves a pointer's content
func (app *contentApplication) Retrieve(pointer references.Pointer) ([]byte, error) {
	return nil, nil
}

// List retrieves multiple pointers contents
func (app *contentApplication) List(pointers []references.Pointer) ([][]byte, error) {
	return nil, nil
}
