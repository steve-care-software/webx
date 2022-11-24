package blocks

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/hashtrees"
)

type builder struct {
	hashAdapter   hash.Adapter
	pHeight       *uint
	pNextScore    *big.Int
	pPendingScore *big.Int
	trx           hashtrees.HashTree
	pPrevious     *hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
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
	return createBuilder(
		app.hashAdapter,
	)
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
func (app *builder) WithTransactions(transactions hashtrees.HashTree) Builder {
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

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pHeight)),
		app.pNextScore.Bytes(),
		app.pPendingScore.Bytes(),
		app.trx.Head().Bytes(),
	}

	if app.pPrevious != nil {
		data = append(data, app.pPrevious.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pPrevious != nil {
		return createBlockWithPrevious(*pHash, *app.pHeight, *app.pNextScore, *app.pPendingScore, app.trx, app.pPrevious), nil
	}

	return createBlock(*pHash, *app.pHeight, *app.pNextScore, *app.pPendingScore, app.trx), nil
}
