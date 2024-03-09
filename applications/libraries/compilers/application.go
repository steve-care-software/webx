package compilers

import "github.com/steve-care-software/datastencil/domain/libraries"

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Compile compiles the application
func (app *application) Compile(input []byte) (libraries.Library, error) {
	return nil, nil
}

// Decompile decompiles the application
func (app *application) Decompile(library libraries.Library) ([]byte, error) {
	return nil, nil
}
