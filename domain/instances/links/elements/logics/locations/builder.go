package locations

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	single      []byte
	list        [][]byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		single:      nil,
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

// WithSingle adds a sigle command to the builder
func (app *builder) WithSingle(single []byte) Builder {
	app.single = single
	return app
}

// WithSiWithListngle adds a command list to the builder
func (app *builder) WithList(list [][]byte) Builder {
	app.list = list
	return app
}

// Now builds a new Location instance
func (app *builder) Now() (Location, error) {
	if app.single != nil && len(app.single) <= 0 {
		app.single = nil
	}

	if app.list != nil {
		cmdList := [][]byte{}
		for _, oneCommand := range app.list {
			if oneCommand != nil && len(oneCommand) <= 0 {
				continue
			}

			cmdList = append(cmdList, oneCommand)
		}

		app.list = cmdList
		if len(app.list) <= 0 {
			app.list = nil
		}
	}

	data := [][]byte{}
	if app.single != nil {
		data = append(data, []byte("single"))
		data = append(data, []byte(app.single))
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		for _, oneCommand := range app.list {
			data = append(data, oneCommand)
		}
	}

	if len(data) < 2 {
		return nil, errors.New("the Location is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.single != nil {
		return createLocationWithSingle(*pHash, app.single), nil
	}

	return createLocationWithList(*pHash, app.list), nil

}
