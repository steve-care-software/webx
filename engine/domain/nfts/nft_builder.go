package nfts

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nftBuilder struct {
	hashAdapter hash.Adapter
	pByte       *byte
	nfts        NFTs
	pRecursive  *uint
}

func createNFTBuilder(
	hashAdapter hash.Adapter,
) NFTBuilder {
	out := nftBuilder{
		hashAdapter: hashAdapter,
		pByte:       nil,
		nfts:        nil,
		pRecursive:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *nftBuilder) Create() NFTBuilder {
	return createNFTBuilder(
		app.hashAdapter,
	)
}

// WithByte adds a byte to the builder
func (app *nftBuilder) WithByte(byteValue byte) NFTBuilder {
	app.pByte = &byteValue
	return app
}

// WithNFTs add nfts to the builder
func (app *nftBuilder) WithNFTs(nfts NFTs) NFTBuilder {
	app.nfts = nfts
	return app
}

// WithRecursive adds a recursive level to the builder
func (app *nftBuilder) WithRecursive(recursive uint) NFTBuilder {
	app.pRecursive = &recursive
	return app
}

func (app *nftBuilder) hash() (hash.Hash, error) {
	if app.pByte != nil && app.nfts != nil {
		return nil, errors.New("the bytes and nfts cannot both be non-empty")
	}

	data := [][]byte{}
	if app.pByte != nil {
		byteValue := *app.pByte
		data = append(data, []byte{
			byteValue,
		})
	}

	if app.nfts != nil {
		data = append(data, app.nfts.Hash().Bytes())
	}

	if app.pRecursive != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.pRecursive))))
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

	if app.pByte != nil {
		return createNFTWithByte(pHash, app.pByte), nil
	}

	if app.nfts != nil {
		return createNFTWithNFTs(pHash, app.nfts), nil
	}

	if app.pRecursive != nil {
		return createNFTWithRecursive(pHash, app.pRecursive), nil
	}

	return nil, errors.New("the NFT is invalid")
}
