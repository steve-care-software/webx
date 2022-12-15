package transactions

import (
	"errors"
	"math/big"
	"sort"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Transaction
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list of transactions to the builder
func (app *builder) WithList(list []Transaction) Builder {
	app.list = list
	return app
}

// Now builds a new Transactions instance
func (app *builder) Now() (Transactions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Transaction in order to build a Transactions instance")
	}

	scoresList := []*big.Int{}
	scoresTrx := map[string]Transaction{}
	for _, oneTrx := range app.list {
		pScore := oneTrx.Score()
		scoresTrx[pScore.String()] = oneTrx
		scoresList = append(scoresList, pScore)
	}

	sort.Slice(scoresList, func(i, j int) bool {
		return scoresList[i].Cmp(scoresList[j]) < 0
	})

	mp := map[string]Transaction{}
	data := [][]byte{}
	total := big.NewInt(0)
	trxList := []Transaction{}
	for _, oneScore := range scoresList {
		scoreKeyname := oneScore.String()
		trx := scoresTrx[scoreKeyname]

		hashKeyname := trx.Hash().String()
		mp[hashKeyname] = trx

		trxList = append(trxList, trx)
		data = append(data, trx.Hash().Bytes())
		total = total.Add(total, oneScore)
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTransactions(*hash, trxList, mp, total), nil
}
