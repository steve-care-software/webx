package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	compile     string
	decompile   string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		compile:     "",
		decompile:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithCompile adds a compile to the builder
func (app *builder) WithCompile(compile string) Builder {
	app.compile = compile
	return app
}

// WithDecompile adds a decompile to the builder
func (app *builder) WithDecompile(decompile string) Builder {
	app.decompile = decompile
	return app
}

// Now builds a new Compiler instance
func (app *builder) Now() (Compiler, error) {
	data := [][]byte{}
	if app.compile != "" {
		data = append(data, []byte("compile"))
		data = append(data, []byte(app.compile))
	}

	if app.decompile != "" {
		data = append(data, []byte("decompile"))
		data = append(data, []byte(app.decompile))
	}

	if len(data) != 2 {
		return nil, errors.New("the Compiler is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.compile != "" {
		return createCompilerWithCompile(*pHash, app.compile), nil
	}

	return createCompilerWithDecompile(*pHash, app.decompile), nil
}
