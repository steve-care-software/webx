package fns

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash     *hash.Hash
	isSingle  bool
	isContent bool
	program   hash.Hash
	pParam    *uint
}

func createBuilder() Builder {
	out := builder{
		pHash:     nil,
		isSingle:  false,
		isContent: false,
		program:   nil,
		pParam:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program hash.Hash) Builder {
	app.program = program
	return app
}

// WithParam adds a param to the builder
func (app *builder) WithParam(param uint) Builder {
	app.pParam = &param
	return app
}

// IsSingle flags the builder as single
func (app *builder) IsSingle() Builder {
	app.isSingle = true
	return app
}

// IsContent flags the builder as content
func (app *builder) IsContent() Builder {
	app.isContent = true
	return app
}

// Now builds a new Fn isntance
func (app *builder) Now() (Fn, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Fn instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Fn instance")
	}

	if app.pParam == nil {
		return nil, errors.New("the param is mandatory in order to build a Fn instance")
	}

	return createFn(*app.pHash, app.isSingle, app.isContent, app.program, *app.pParam), nil
}
