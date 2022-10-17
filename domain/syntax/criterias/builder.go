package criterias

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

type builder struct {
	hashAdapter     hash.Adapter
	name            string
	pIndex          *uint
	includeChannels bool
	child           Criteria
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:     hashAdapter,
		name:            "",
		pIndex:          nil,
		includeChannels: false,
		child:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// IncludeChannels flags the builder as include channels
func (app *builder) IncludeChannels() Builder {
	app.includeChannels = true
	return app
}

// WithChild adds a child to the builder
func (app *builder) WithChild(child Criteria) Builder {
	app.child = child
	return app
}

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Criteria instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Criteria instance")
	}

	includeChannelsStr := "false"
	if app.includeChannels {
		includeChannelsStr = "true"
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(*app.pIndex))),
		[]byte(includeChannelsStr),
	}

	if app.child != nil {
		data = append(data, app.child.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.child != nil {
		return createCriteriaWithChild(*pHash, app.name, *app.pIndex, app.includeChannels, app.child), nil
	}

	return createCriteria(*pHash, app.name, *app.pIndex, app.includeChannels), nil
}
