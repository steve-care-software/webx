package commits

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type builder struct {
	hashAdapter hash.Adapter
	values      hashtrees.HashTree
	pCreatedOn  *time.Time
	pProof      *big.Int
	parent      Commit
	miningValue byte
}

func createBuilder(
	hashAdapter hash.Adapter,
	miningValue byte,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		values:      nil,
		pCreatedOn:  nil,
		pProof:      nil,
		parent:      nil,
		miningValue: miningValue,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.miningValue,
	)
}

// WithValues add values to the builder
func (app *builder) WithValues(values hashtrees.HashTree) Builder {
	app.values = values
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent Commit) Builder {
	app.parent = parent
	return app
}

// WithProof adds a proof to the builder
func (app *builder) WithProof(proof *big.Int) Builder {
	app.pProof = proof
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.values == nil {
		return nil, errors.New("the values is mandatory in order to build a Commit instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		app.values.Head().Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.UnixNano())),
	}

	if app.pProof != nil {
		data = append(data, app.pProof.Bytes())
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	var mine Mine
	if app.pProof != nil {
		// make the result hash:
		pResult, err := app.hashAdapter.FromMultiBytes([][]byte{
			pHash.Bytes(),
			app.pProof.Bytes(),
		})

		if err != nil {
			return nil, err
		}

		score := uint(0)
		resultBytes := pResult.Bytes()
		for _, oneByte := range resultBytes {
			if oneByte == app.miningValue {
				score++
				continue
			}

			break
		}

		mine = createMine(*pResult, app.pProof, score)
	}

	if mine != nil && app.parent != nil {
		return createCommitWithMineAndParent(*pHash, app.values, *app.pCreatedOn, mine, app.parent), nil
	}

	if mine != nil {
		return createCommitWithMine(*pHash, app.values, *app.pCreatedOn, mine), nil
	}

	if app.parent != nil {
		return createCommitWithParent(*pHash, app.values, *app.pCreatedOn, app.parent), nil
	}

	return createCommit(*pHash, app.values, *app.pCreatedOn), nil
}
