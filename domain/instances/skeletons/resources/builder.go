package resources

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Resource
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
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Resource) Builder {
	app.list = list
	return app
}

// Now builds a new Resources instance
func (app *builder) Now() (Resources, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Resource in order to build a Resources instance")
	}

	mp := map[string]Resource{}
	for _, oneResource := range app.list {
		name := oneResource.Name()
		if idx, ok := mp[name]; ok {
			str := fmt.Sprintf("the Resource (index: %d, name: %s) already exists", idx, name)
			return nil, errors.New(str)
		}

		mp[name] = oneResource
	}

	data := [][]byte{}
	for _, oneResource := range app.list {
		data = append(data, oneResource.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createResources(*pHash, mp, app.list), nil
}
