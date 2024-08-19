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
	name string
}

func createNFTs(
	hash hash.Hash,
	list []NFT,
	mp map[string]NFT,
) NFTs {
	return createNFTsInternally(hash, list, mp, "")
}

func createNFTsWithName(
	hash hash.Hash,
	list []NFT,
	mp map[string]NFT,
	name string,
) NFTs {
	return createNFTsInternally(hash, list, mp, name)
}

func createNFTsInternally(
	hash hash.Hash,
	list []NFT,
	mp map[string]NFT,
	name string,
) NFTs {
	out := nfts{
		hash: hash,
		list: list,
		mp:   mp,
		name: name,
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

// HasName returns true if there is a name, false otherwise
func (obj *nfts) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *nfts) Name() string {
	return obj.name
}
