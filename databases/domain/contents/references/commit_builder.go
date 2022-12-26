package references

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type commitBuilder struct {
	hashAdapter hash.Adapter
	values      hashtrees.HashTree
	pParent     *hash.Hash
	pProof      *big.Int
	pCreatedOn  *time.Time
	miningValue byte
}

func createCommitBuilder(
	hashAdapter hash.Adapter,
	miningValue byte,
) CommitBuilder {
	out := commitBuilder{
		hashAdapter: hashAdapter,
		values:      nil,
		pParent:     nil,
		pProof:      nil,
		pCreatedOn:  nil,
		miningValue: miningValue,
	}

	return &out
}

// Create initializes the builder
func (app *commitBuilder) Create() CommitBuilder {
	return createCommitBuilder(
		app.hashAdapter,
		app.miningValue,
	)
}

// WithValues add values to the builder
func (app *commitBuilder) WithValues(values hashtrees.HashTree) CommitBuilder {
	app.values = values
	return app
}

// WithParent adds a parent to the builder
func (app *commitBuilder) WithParent(parent hash.Hash) CommitBuilder {
	app.pParent = &parent
	return app
}

// WithProof adds a proof to the builder
func (app *commitBuilder) WithProof(proof *big.Int) CommitBuilder {
	app.pProof = proof
	return app
}

// CreatedOn adds a creation time to the builder
func (app *commitBuilder) CreatedOn(createdOn time.Time) CommitBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Commit instance
func (app *commitBuilder) Now() (Commit, error) {
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

	if app.pParent != nil {
		data = append(data, app.pParent.Bytes())
	}

	if app.pProof != nil {
		data = append(data, app.pProof.Bytes())
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

	if app.pParent != nil && mine != nil {
		return createCommitWithParentAndMine(*pHash, app.values, *app.pCreatedOn, app.pParent, mine), nil
	}

	if app.pParent != nil {
		return createCommitWithParent(*pHash, app.values, *app.pCreatedOn, app.pParent), nil
	}

	if mine != nil {
		return createCommitWithMine(*pHash, app.values, *app.pCreatedOn, mine), nil
	}

	return createCommit(*pHash, app.values, *app.pCreatedOn), nil
}
