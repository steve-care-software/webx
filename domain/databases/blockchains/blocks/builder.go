package blocks

import (
	"errors"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity        entities.Entity
	pHeight       *uint
	pNextScore    *big.Int
	pPendingScore *big.Int
	trx           entities.Identifiers
	pPrevious     *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		entity:        nil,
		pHeight:       nil,
		pNextScore:    nil,
		pPendingScore: nil,
		trx:           nil,
		pPrevious:     nil,
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

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.pHeight = &height
	return app
}

// WithNextScore adds a nextScore to the builder
func (app *builder) WithNextScore(nextScore big.Int) Builder {
	app.pNextScore = &nextScore
	return app
}

// WithPendingScope adds a pendingScore to the builder
func (app *builder) WithPendingScope(pendingScore big.Int) Builder {
	app.pPendingScore = &pendingScore
	return app
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(transactions entities.Identifiers) Builder {
	app.trx = transactions
	return app
}

// WithPrevious adds a previous hash to the builder
func (app *builder) WithPrevious(previous hash.Hash) Builder {
	app.pPrevious = &previous
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Block instance")
	}

	if app.pHeight == nil {
		return nil, errors.New("the height is mandatory in order to build a Block instance")
	}

	if app.pNextScore == nil {
		return nil, errors.New("the nextScore is mandatory in order to build a Block instance")
	}

	if app.pPendingScore == nil {
		return nil, errors.New("the pendingScore is mandatory in order to build a Block instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Block instance")
	}

	if app.pPrevious != nil {
		return createBlockWithPrevious(app.entity, *app.pHeight, *app.pNextScore, *app.pPendingScore, app.trx, app.pPrevious), nil
	}

	return createBlock(app.entity, *app.pHeight, *app.pNextScore, *app.pPendingScore, app.trx), nil
}
