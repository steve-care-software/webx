package blocks

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	pHeight       *uint
	pPrevious     *hash.Hash
	trx           transactions.Transactions
	pNextScore    *big.Int
	pPendingScore *big.Int
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		pHeight:       nil,
		pPrevious:     nil,
		trx:           nil,
		pNextScore:    nil,
		pPendingScore: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height uint) Builder {
	app.pHeight = &height
	return app
}

// WithPrevious adds a previous hash to the builder
func (app *builder) WithPrevious(previous hash.Hash) Builder {
	app.pPrevious = &previous
	return app
}

// WithTransactions adds a transactions to the builder
func (app *builder) WithTransactions(trx transactions.Transactions) Builder {
	app.trx = trx
	return app
}

// WithNextScore adds a nextScore to the builder
func (app *builder) WithNextScore(nextScore big.Int) Builder {
	app.pNextScore = &nextScore
	return app
}

// WithPendingScore adds a pendingScore to the builder
func (app *builder) WithPendingScore(pendingScore big.Int) Builder {
	app.pPendingScore = &pendingScore
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.pHeight == nil {
		return nil, errors.New("the height is mandatory in order to build a Block instance")
	}

	if app.pPrevious == nil && (*app.pHeight > 0) {
		str := fmt.Sprintf("the height (%d) must be zero when there is no previous block hash, in order to build a Block instance", *app.pHeight)
		return nil, errors.New(str)
	}

	if app.trx == nil {
		return nil, errors.New("the transactions is mandatory in order to build a Block instance")
	}

	if app.pNextScore == nil {
		return nil, errors.New("the nextScore is mandatory in order to build a Block instance")
	}

	if app.pPendingScore == nil {
		return nil, errors.New("the pendingScore is mandatory in order to build a Block instance")
	}

	trxScore := app.trx.Score()
	pTotal := trxScore.Add(trxScore, app.pNextScore)
	pTotal = pTotal.Add(pTotal, app.pPendingScore)
	score := createScore(app.pNextScore, app.pPendingScore, pTotal)

	length := big.NewInt(int64(len(app.trx.List())))
	pMaxTrxScore := transactions.FetchMaxScore()
	pMaxBlockScore := pMaxTrxScore.Mul(pMaxTrxScore, length)

	cmpValue := pTotal.Cmp(pMaxBlockScore)
	isStuck := (cmpValue >= 0)

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pHeight)),
		app.trx.Hash().Bytes(),
	}

	if app.pPrevious != nil {
		data = append(data, app.pPrevious.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pPrevious != nil {
		return createBlockWithPrevious(*hash, *app.pHeight, score, isStuck, app.trx, app.pPrevious), nil
	}

	return createBlock(*hash, *app.pHeight, score, isStuck, app.trx), nil
}
