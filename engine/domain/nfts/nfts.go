package nfts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nfts struct {
	hash hash.Hash
	list []NFT
	mp   map[string]NFT
}

func createNFTs(
	hash hash.Hash,
	list []NFT,
	mp map[string]NFT,
) NFTs {
	out := nfts{
		hash: hash,
		list: list,
		mp:   mp,
	}

	return &out
}

// Hash returns the nfts hash
func (obj *nfts) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of nfts
func (obj *nfts) List() []NFT {
	return obj.list
}

// Fetch fetches an nft by hash
func (obj *nfts) Fetch(hash hash.Hash) (NFT, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no NFT registered for the provided hash: %s", keyname)
	return nil, errors.New(str)
}

// Combine returns the unique list of NFTs presents in current and passed NFTs
func (obj *nfts) Combine(input NFTs) []NFT {
	combined := append(obj.list, input.List()...)
	mp := map[string]NFT{}
	for _, oneNFT := range combined {
		keyname := oneNFT.Hash().String()
		mp[keyname] = oneNFT
	}

	output := []NFT{}
	for _, oneNFT := range mp {
		output = append(output, oneNFT)
	}

	return output
}
