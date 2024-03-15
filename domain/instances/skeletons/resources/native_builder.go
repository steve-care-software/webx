package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type nativeBuilder struct {
	hashAdapter hash.Adapter
	pSingle     *uint8
	list        List
}

func createNativeBuilder(
	hashAdapter hash.Adapter,
) NativeBuilder {
	out := nativeBuilder{
		hashAdapter: hashAdapter,
		pSingle:     nil,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *nativeBuilder) Create() NativeBuilder {
	return createNativeBuilder(
		app.hashAdapter,
	)
}

// WithSingle adds a single to the builder
func (app *nativeBuilder) WithSingle(single uint8) NativeBuilder {
	app.pSingle = &single
	return app
}

// WithList adds a list to the builder
func (app *nativeBuilder) WithList(list List) NativeBuilder {
	app.list = list
	return app
}

// Now builds a new Native instance
func (app *nativeBuilder) Now() (Native, error) {
	data := [][]byte{}
	if app.pSingle != nil {
		data = append(data, []byte("single"))
		data = append(data, []byte{
			*app.pSingle,
		})
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		data = append(data, app.list.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Native is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pSingle != nil {
		return createNativeWithSingle(*pHash, app.pSingle), nil
	}

	if app.list != nil {
		return createNativeWithList(*pHash, app.list), nil
	}

	return nil, errors.New("the Native is invalid")
}
