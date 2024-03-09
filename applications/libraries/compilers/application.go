package compilers

import "github.com/steve-care-software/datastencil/domain/libraries"

type application struct {
	adapter libraries.Adapter
}

func createApplication(
	adapter libraries.Adapter,
) Application {
	out := application{
		adapter: adapter,
	}

	return &out
}

// Compile compiles the application
func (app *application) Compile(input []byte) (libraries.Library, error) {
	return app.adapter.ToInstance(input)
}

// Decompile decompiles the application
func (app *application) Decompile(library libraries.Library) ([]byte, error) {
	return app.adapter.ToData(library)
}
