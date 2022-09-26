package genesis

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	supply      uint64
	owner       []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		supply:      0,
		owner:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
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

// Now builds a new Genesis instance
func (app *builder) Now() (Genesis, error) {
	if app.supply <= 0 {
		return nil, errors.New("the supply must be greater than zero (0)")
	}

	if app.owner != nil && len(app.owner) <= 0 {
		app.owner = nil
	}

	if app.owner == nil {
		return nil, errors.New("there must be at least 1 owner hash in order to build a Genesis instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", app.supply)),
	}

	for _, oneHash := range app.owner {
		data = append(data, oneHash.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createGenesis(*hash, app.supply, app.owner), nil
}
