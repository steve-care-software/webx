package entities

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type referenceBuilder struct {
	identifier Identifier
	trx        hash.Hash
	block      hash.Hash
	chain      hash.Hash
}

func createReferenceBuilder() ReferenceBuilder {
	out := referenceBuilder{
		identifier: nil,
		trx:        nil,
		block:      nil,
		chain:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *referenceBuilder) Create() ReferenceBuilder {
	return createReferenceBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *referenceBuilder) WithIdentifier(identifier Identifier) ReferenceBuilder {
	app.identifier = identifier
	return app
}

// WithTransaction adds a trx hash to the builder
func (app *referenceBuilder) WithTransaction(trx hash.Hash) ReferenceBuilder {
	app.trx = trx
	return app
}

// WithBlock adds a block hash to the builder
func (app *referenceBuilder) WithBlock(block hash.Hash) ReferenceBuilder {
	app.block = block
	return app
}

// WithChain adds a chain hash to the builder
func (app *referenceBuilder) WithChain(chain hash.Hash) ReferenceBuilder {
	app.chain = chain
	return app
}

// Now builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Reference instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction hash is mandatory in order to build a Reference instance")
	}

	if app.block == nil {
		return nil, errors.New("the block hash is mandatory in order to build a Reference instance")
	}

	if app.chain == nil {
		return nil, errors.New("the chain hash is mandatory in order to build a Reference instance")
	}

	return createReference(app.identifier, app.trx, app.block, app.chain), nil
}
