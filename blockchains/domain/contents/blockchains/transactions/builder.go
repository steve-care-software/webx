package transactions

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash  *hash.Hash
	pAsset *hash.Hash
	pProof *big.Int
}

func createBuilder() Builder {
	out := builder{
		pHash:  nil,
		pAsset: nil,
		pProof: nil,
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

// WithAsset adds an asset to the builder
func (app *builder) WithAsset(asset hash.Hash) Builder {
	app.pAsset = &asset
	return app
}

// WithProof adds a proof to the builder
func (app *builder) WithProof(proof big.Int) Builder {
	app.pProof = &proof
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Transaction instance")
	}

	if app.pAsset == nil {
		return nil, errors.New("the asset is mandatory in order to build a Transaction instance")
	}

	if app.pProof == nil {
		return nil, errors.New("the proof is mandatory in order to build a Transaction instance")
	}

	return createTransaction(*app.pHash, *app.pAsset, app.pProof), nil
}
