package programs

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	engine      string
	compiler    []byte
	script      []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		engine:      "",
		compiler:    nil,
		script:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithEngine adds an engine to the builder
func (app *builder) WithEngine(engine string) Builder {
	app.engine = engine
	return app
}

// WithCompiler adds a compiler to the builder
func (app *builder) WithCompiler(compiler []byte) Builder {
	app.compiler = compiler
	return app
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script []byte) Builder {
	app.script = script
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.engine == "" {
		return nil, errors.New("the engine is mandatory in order to build a Program instance")
	}

	if app.compiler != nil && len(app.compiler) <= 0 {
		app.compiler = nil
	}

	if app.compiler == nil {
		return nil, errors.New("the compiler is mandatory in order to build a Program instance")
	}

	if app.script != nil && len(app.script) <= 0 {
		app.script = nil
	}

	if app.script == nil {
		return nil, errors.New("the script is mandatory in order to build a Program instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.engine),
		app.compiler,
		app.script,
	})

	if err != nil {
		return nil, err
	}

	return createProgram(*hash, app.engine, app.compiler, app.script), nil
}
