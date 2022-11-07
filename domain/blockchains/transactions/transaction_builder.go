package transactions

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	pAsset      *hash.Hash
	pProof      *big.Int
}

func createTransactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		pAsset:      nil,
		pProof:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(app.hashAdapter)
}

// WithAssset adds an asset to the builder
func (app *transactionBuilder) WithAssset(asset hash.Hash) TransactionBuilder {
	app.pAsset = &asset
	return app
}

// WithProof adds a proof to the builder
func (app *transactionBuilder) WithProof(proof big.Int) TransactionBuilder {
	app.pProof = &proof
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.pAsset == nil {
		return nil, errors.New("the asset is mandatory in order to build a Transaction instance")
	}

	if app.pProof == nil {
		return nil, errors.New("the proof is mndatory in order to build a Transaction instance")
	}

	mine, err := ExecuteMiner(app.hashAdapter, *app.pAsset, *app.pProof)
	if err != nil {
		return nil, err
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.pAsset.Bytes(),
		app.pProof.Bytes(),
		mine.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	amount := FetchMinedAmount(*mine)
	pScore := CalculateScore(amount)
	return createTransaction(*hash, *app.pAsset, *app.pProof, *mine, pScore), nil
}
