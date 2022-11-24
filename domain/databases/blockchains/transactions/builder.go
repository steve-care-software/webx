package transactions

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pAsset      *hash.Hash
	pProof      *big.Int
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pAsset:      nil,
		pProof:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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
	if app.pAsset == nil {
		return nil, errors.New("the asset is mandatory in order to build a Transaction instance")
	}

	if app.pProof == nil {
		return nil, errors.New("the proof is mandatory in order to build a Transaction instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.pAsset.Bytes(),
		app.pProof.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*pHash, *app.pAsset, *app.pProof), nil
}
