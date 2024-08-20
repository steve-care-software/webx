package nfts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []NFT
	name        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
		name:        "",
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
func (app *builder) WithList(list []NFT) Builder {
	app.list = list
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new NFTs instance
func (app *builder) Now() (NFTs, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 NFT in order to build an NFTs instance")
	}

	data := [][]byte{}
	for _, oneNFT := range app.list {
		data = append(data, oneNFT.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.name != "" {
		return createNFTsWithName(
			*pHash,
			app.list,
			app.name,
		), nil
	}

	return createNFTs(
		*pHash,
		app.list,
	), nil

}
