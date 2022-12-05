package references

import "errors"

type blockchainBuilder struct {
	chain  BlockchainKey
	blocks BlockchainKeys
	trx    BlockchainKeys
}

func createBlockchainBuilder() BlockchainBuilder {
	out := blockchainBuilder{
		chain:  nil,
		blocks: nil,
		trx:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockchainBuilder) Create() BlockchainBuilder {
	return createBlockchainBuilder()
}

// WithChain adds a chain to the builder
func (app *blockchainBuilder) WithChain(chain BlockchainKey) BlockchainBuilder {
	app.chain = chain
	return app
}

// WithBlocks adds a blocks to the builder
func (app *blockchainBuilder) WithBlocks(blocks BlockchainKeys) BlockchainBuilder {
	app.blocks = blocks
	return app
}

// WithTransactions add transactions to the builder
func (app *blockchainBuilder) WithTransactions(trx BlockchainKeys) BlockchainBuilder {
	app.trx = trx
	return app
}

// Now builds a new Blockchain instance
func (app *blockchainBuilder) Now() (Blockchain, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Blockchain instance")
	}

	if app.blocks == nil {
		return nil, errors.New("the blocks is mandatory in order to build a Blockchain instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Blockchain instance")
	}

	return createBlockchain(app.chain, app.blocks, app.trx), nil
}
