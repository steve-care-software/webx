package uints

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/nfts"
)

type adapter struct {
	nftsBuilder nfts.Builder
	nftBuilder  nfts.NFTBuilder
}

func createAdapter(
	nftsBuilder nfts.Builder,
	nftBuilder nfts.NFTBuilder,
) Adapter {
	out := adapter{
		nftsBuilder: nftsBuilder,
		nftBuilder:  nftBuilder,
	}

	return &out
}

// ToNFT converts a value to an nft
func (app *adapter) ToNFT(value uint64) (nfts.NFT, error) {
	data := uintToBytes(value)
	nftsList := []nfts.NFT{}
	for _, oneByte := range data {
		ins, err := app.nftBuilder.Create().WithByte(oneByte).Now()
		if err != nil {
			return nil, err
		}

		nftsList = append(nftsList, ins)
	}

	nfts, err := app.nftsBuilder.Create().WithList(nftsList).Now()
	if err != nil {
		return nil, err
	}

	return app.nftBuilder.Create().
		WithNFTs(nfts).
		Now()
}

// ToValue converts an NFT to a value
func (app *adapter) ToValue(root nfts.NFT) (*uint64, error) {
	if root.IsByte() {
		return nil, errors.New("the root AST was expected to contain NFT's, not a byte")
	}

	nftsList := root.NFTs().List()
	length := len(nftsList)
	if length != AmountOfBytesIntUint64 {
		str := fmt.Sprintf("the root NFT was expected to contain %d sub-NFT's, %d provided", AmountOfBytesIntUint64, length)
		return nil, errors.New(str)
	}

	data := []byte{}
	for idx, oneNFT := range nftsList {
		if !oneNFT.IsByte() {
			str := fmt.Sprintf("the NFT (index: %d) was expected to contain a byte", idx)
			return nil, errors.New(str)
		}

		data = append(data, *oneNFT.Byte())
	}

	value := bytesToUInt(data)
	return &value, nil
}
