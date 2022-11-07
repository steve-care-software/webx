package transactions

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	pAsset *hash.Hash
	pProof *big.Int
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		pAsset: nil,
		pProof: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntity adds an entity to the builder
func (app *builder) WithEntity(entity entities.Entity) Builder {
	app.entity = entity
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
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Transaction instance")
	}

	if app.pAsset == nil {
		return nil, errors.New("the asset is mandatory in order to build a Transaction instance")
	}

	if app.pProof == nil {
		return nil, errors.New("the proof is mandatory in order to build a Transaction instance")
	}

	return createTransaction(app.entity, *app.pAsset, *app.pProof), nil
}
