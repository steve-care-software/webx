package blockchains

import (
	"errors"

	"github.com/steve-care-software/webx/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type builder struct {
	pReference *hash.Hash
	head       blocks.Block
}

func createBuilder() Builder {
	out := builder{
		pReference: nil,
		head:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReference adds a reference to the builder
func (app *builder) WithReference(reference hash.Hash) Builder {
	app.pReference = &reference
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head blocks.Block) Builder {
	app.head = head
	return app
}

// Now builds a new Blockchain instance
func (app *builder) Now() (Blockchain, error) {
	if app.pReference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Blockchain instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Blockchain instance")
	}

	return createBlockchain(*app.pReference, app.head), nil
}
