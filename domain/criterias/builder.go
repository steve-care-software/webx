package criterias

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	current     Tail
	next        Node
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		current:     nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithCurrent adds a current tail to the builder
func (app *builder) WithCurrent(current Tail) Builder {
	app.current = current
	return app
}

// WithNext adds a next node to the builder
func (app *builder) WithNext(next Node) Builder {
	app.next = next
	return app
}

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	if app.current == nil {
		return nil, errors.New("the current tail is mandatory in order to build a Criteria instance")
	}

	data := [][]byte{
		app.current.Hash().Bytes(),
	}

	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createCriteriaWithNext(*pHash, app.current, app.next), nil
	}

	return createCriteria(*pHash, app.current), nil
}
