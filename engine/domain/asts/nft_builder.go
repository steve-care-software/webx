package asts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nftBuilder struct {
	hashAdapter hash.Adapter
	bytes       []byte
	nfts        []hash.Hash
}

func createNFTBuilder(
	hashAdapter hash.Adapter,
) NFTBuilder {
	out := nftBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		nfts:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *nftBuilder) Create() NFTBuilder {
	return createNFTBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *nftBuilder) WithBytes(bytes []byte) NFTBuilder {
	app.bytes = bytes
	return app
}

// WithNFTs add nfts to the builder
func (app *nftBuilder) WithNFTs(nfts []hash.Hash) NFTBuilder {
	app.nfts = nfts
	return app
}

func (app *nftBuilder) hash() (hash.Hash, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.nfts != nil && len(app.nfts) <= 0 {
		app.nfts = nil
	}

	if app.bytes != nil && app.nfts != nil {
		return nil, errors.New("the bytes and nfts cannot both be non-empty")
	}

	if app.nfts != nil && len(app.nfts) <= 1 {
		return app.nfts[0], nil
	}

	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, app.bytes)
	}

	if app.nfts != nil {
		for _, oneNFT := range app.nfts {
			data = append(data, oneNFT)
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return *pHash, nil
}

// Now builds a new NFT instance
func (app *nftBuilder) Now() (NFT, error) {
	pHash, err := app.hash()
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createNFTWithBytes(pHash, app.bytes), nil
	}

	if app.nfts != nil {
		// calculate the complexity:
		complexity := map[string]uint{}
		for _, oneHash := range app.nfts {
			keyname := oneHash.String()
			if _, ok := complexity[keyname]; ok {
				complexity[keyname]++
				continue
			}

			complexity[keyname] = 1
		}

		return createNFTWithNFTs(pHash, app.nfts, complexity), nil
	}

	return nil, errors.New("the NFT is invalid")
}
