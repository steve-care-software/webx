package asts

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
