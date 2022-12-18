package commits

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type builder struct {
	hashAdapter hash.Adapter
	values      hashtrees.HashTree
	pCreatedOn  *time.Time
	parent      Commit
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		values:      nil,
		pCreatedOn:  nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
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

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createCommit(*pHash, app.values, *app.pCreatedOn), nil
}
