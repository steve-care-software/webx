package asts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nftsBuilder struct {
	hashAdapter hash.Adapter
	list        []NFT
}

func createNFTsBuilder(
	hashAdapter hash.Adapter,
) NFTsBuilder {
	out := nftsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *nftsBuilder) Create() NFTsBuilder {
	return createNFTsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *nftsBuilder) WithList(list []NFT) NFTsBuilder {
	app.list = list
	return app
}

// Now builds a new NFTs instance
func (app *nftsBuilder) Now() (NFTs, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 NFT in order to build an NFTs instance")
	}

	mp := map[string]NFT{}
	data := [][]byte{}
	for _, oneNFT := range app.list {
		data = append(data, oneNFT.Bytes())
		keyname := oneNFT.Hash().String()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the NFT (hash: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}

		mp[keyname] = oneNFT
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	// calculate the complexity:
	complexity := map[string]uint{}
	for _, oneNFT := range app.list {
		if oneNFT.IsBytes() {
			continue
		}

		subComplexity := oneNFT.Complexity()
		for subHashStr, score := range subComplexity {
			if _, ok := complexity[subHashStr]; ok {
				complexity[subHashStr] = score + 1
				continue
			}

			complexity[subHashStr] = 1
		}
	}

	return createNFTs(
		*pHash,
		app.list,
		mp,
		complexity,
	), nil

}
