package units

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/units/genesis"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	amount      uint64
	owner       []hash.Hash
	units       Units
	genesis     genesis.Genesis
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		amount:      0,
		owner:       nil,
		units:       nil,
		genesis:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
	)
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount uint64) ContentBuilder {
	app.amount = amount
	return app
}

// WithOwner adds an owner to the builder
func (app *contentBuilder) WithOwner(owner []hash.Hash) ContentBuilder {
	app.owner = owner
	return app
}

// WithUnits adds a units to the builder
func (app *contentBuilder) WithUnits(units Units) ContentBuilder {
	app.units = units
	return app
}

// WithGenesis adds a genesis to the builder
func (app *contentBuilder) WithGenesis(genesis genesis.Genesis) ContentBuilder {
	app.genesis = genesis
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Content instance")
	}

	if app.owner != nil && len(app.owner) <= 0 {
		app.owner = nil
	}

	if app.owner == nil {
		return nil, errors.New("there must be at least 1 owner hash in order to build a Content instance")
	}

	if app.units != nil {
		prevAmount := app.units.Amount()
		if app.amount > prevAmount {
			str := fmt.Sprintf("the amount (%d) cannot be greater than the previous amount (%d)", app.amount, prevAmount)
			return nil, errors.New(str)
		}

		previous := createPreviousWithUnits(app.units)
		data := [][]byte{
			previous.Hash().Bytes(),
			[]byte(fmt.Sprintf("%d", app.amount)),
		}

		for _, oneHash := range app.owner {
			data = append(data, oneHash.Bytes())
		}

		hash, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		return createContent(*hash, app.amount, app.owner, previous), nil
	}

	if app.genesis != nil {
		prevAmount := app.genesis.Supply()
		if app.amount > prevAmount {
			str := fmt.Sprintf("the amount (%d) cannot be greater than the previous amount (%d)", app.amount, prevAmount)
			return nil, errors.New(str)
		}

		previous := createPreviousWithGenesis(app.genesis)
		data := [][]byte{
			previous.Hash().Bytes(),
			[]byte(fmt.Sprintf("%d", app.amount)),
		}

		for _, oneHash := range app.owner {
			data = append(data, oneHash.Bytes())
		}

		hash, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		return createContent(*hash, app.amount, app.owner, previous), nil
	}

	return nil, errors.New("the Unit is invalid")
}
