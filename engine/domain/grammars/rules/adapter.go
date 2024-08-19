package rules

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/nfts"
)

type adapter struct {
	rulesBuilder Builder
	ruleBuilder  RuleBuilder
	nftsBuilder  nfts.Builder
	nftBuilder   nfts.NFTBuilder
}

func createAdapter(
	rulesBuilder Builder,
	ruleBuilder RuleBuilder,
	nftsBuilder nfts.Builder,
	nftBuilder nfts.NFTBuilder,
) Adapter {
	out := adapter{
		rulesBuilder: rulesBuilder,
		ruleBuilder:  ruleBuilder,
		nftsBuilder:  nftsBuilder,
		nftBuilder:   nftBuilder,
	}

	return &out
}

// RuleToNFT comverts a rule to an NFT
func (app *adapter) RuleToNFT(ins Rule) (nfts.NFT, error) {
	list := []nfts.NFT{}
	data := ins.Bytes()
	for _, oneByte := range data {
		nft, err := app.nftBuilder.Create().
			WithByte(oneByte).
			Now()

		if err != nil {
			return nil, err
		}

		list = append(list, nft)
	}

	name := ins.Name()
	nfts, err := app.nftsBuilder.Create().
		WithName(name).
		WithList(list).
		Now()

	if err != nil {
		return nil, err
	}

	return app.nftBuilder.Create().
		WithNFTs(nfts).
		Now()
}

// NFTToInstance comverts an NFT to a rule
func (app *adapter) NFTToInstance(root nfts.NFT) (Rule, error) {
	if !root.IsNFTs() {
		return nil, errors.New("the provided NFT was expected to contain sub NFT's")
	}

	nfts := root.NFTs()
	if !nfts.HasName() {
		return nil, errors.New("the provided NFT was expected to contain a sub NFT's with a name")
	}

	values := []byte{}
	name := nfts.Name()
	nftsList := nfts.List()
	for idx, oneNFT := range nftsList {
		if !oneNFT.IsByte() {
			str := fmt.Sprintf("the NFT's (name: %s) was expected to contain NFT's with byte values at index: %d", name, idx)
			return nil, errors.New(str)
		}

		values = append(values, *oneNFT.Byte())
	}

	return app.ruleBuilder.Create().
		WithName(name).
		WithBytes(values).
		Now()
}
