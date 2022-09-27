package genesis

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	ticker      string
	description string
	supply      uint64
	owner       []hash.Hash
	pProof      *hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		ticker:      "",
		description: "",
		supply:      0,
		owner:       nil,
		pProof:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTicker adds a ticker to the builder
func (app *builder) WithTicker(ticker string) Builder {
	app.ticker = ticker
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithSupply adds a supply to the builder
func (app *builder) WithSupply(supply uint64) Builder {
	app.supply = supply
	return app
}

// WithOwner adds an owner to the builder
func (app *builder) WithOwner(owner []hash.Hash) Builder {
	app.owner = owner
	return app
}

// WithProof adds a proof to the builder
func (app *builder) WithProof(proof hash.Hash) Builder {
	app.pProof = &proof
	return app
}

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.supply <= 0 {
		return nil, errors.New("the supply must be greater than zero (0)")
	}

	if app.ticker == "" {
		return nil, errors.New("the ticker is mandatory in order to build a Genesis instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Genesis instance")
	}

	if app.owner != nil && len(app.owner) <= 0 {
		app.owner = nil
	}

	if app.owner == nil {
		return nil, errors.New("there must be at least 1 owner hash in order to build a Genesis instance")
	}

	data := [][]byte{
		[]byte(app.ticker),
		[]byte(app.description),
		[]byte(fmt.Sprintf("%d", app.supply)),
	}

	for _, oneHash := range app.owner {
		data = append(data, oneHash.Bytes())
	}

	if app.pProof != nil {
		data = append(data, app.pProof.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pProof != nil {
		return createGenesisWithProof(
			*hash,
			app.ticker,
			app.description,
			app.supply,
			app.owner,
			app.pProof,
		), nil
	}

	return createGenesis(
		*hash,
		app.ticker,
		app.description,
		app.supply,
		app.owner,
	), nil
}
