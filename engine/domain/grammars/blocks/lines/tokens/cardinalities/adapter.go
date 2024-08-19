package cardinalities

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/cardinalities/uints"
	"github.com/steve-care-software/webx/engine/domain/nfts"
)

type adapter struct {
	cardinalityBuilder Builder
	uintAdapter        uints.Adapter
	nftsBuilder        nfts.Builder
	nftBuilder         nfts.NFTBuilder
}

func createAdapter(
	cardinalityBuilder Builder,
	uintAdapter uints.Adapter,
	nftsBuilder nfts.Builder,
	nftBuilder nfts.NFTBuilder,
) Adapter {
	out := adapter{
		cardinalityBuilder: cardinalityBuilder,
		uintAdapter:        uintAdapter,
		nftsBuilder:        nftsBuilder,
		nftBuilder:         nftBuilder,
	}

	return &out
}

// ToNFT converts an instance to an NFT
func (app *adapter) ToNFT(ins Cardinality) (nfts.NFT, error) {
	min := ins.Min()
	minNFT, err := app.uintAdapter.ToNFT(uint64(min))
	if err != nil {
		return nil, err
	}

	nftList := []nfts.NFT{
		minNFT,
	}

	if ins.HasMax() {
		pMax := ins.Max()
		maxAST, err := app.uintAdapter.ToNFT(uint64(*pMax))
		if err != nil {
			return nil, err
		}

		nftList = append(nftList, maxAST)
	}

	nfts, err := app.nftsBuilder.Create().
		WithList(nftList).
		Now()

	if err != nil {
		return nil, err
	}

	return app.nftBuilder.Create().
		WithNFTs(nfts).
		Now()
}

// ToInstance converts an NFT to an instance
func (app *adapter) ToInstance(root nfts.NFT) (Cardinality, error) {
	if root.IsByte() {
		return nil, errors.New("the root AST was expected to contain NFT's, not a byte")
	}

	nftsList := root.NFTs().List()
	length := len(nftsList)
	if length != 1 && length != 2 {
		str := fmt.Sprintf("the cardinality was expected to contain 1 (min only) or 2 (min and max) NFT's, %d provided", length)
		return nil, errors.New(str)
	}

	pMinValue, err := app.uintAdapter.ToValue(nftsList[0])
	if err != nil {
		return nil, err
	}

	builder := app.cardinalityBuilder.Create().
		WithMin(uint(*pMinValue))

	if length == 2 {
		pMaxValue, err := app.uintAdapter.ToValue(nftsList[1])
		if err != nil {
			return nil, err
		}

		builder.WithMax(uint(*pMaxValue))
	}

	return builder.Now()
}
