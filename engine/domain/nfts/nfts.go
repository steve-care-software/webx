package nfts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nfts struct {
	hash hash.Hash
	list []NFT
	name string
}

func createNFTs(
	hash hash.Hash,
	list []NFT,
) NFTs {
	return createNFTsInternally(hash, list, "")
}

func createNFTsWithName(
	hash hash.Hash,
	list []NFT,
	name string,
) NFTs {
	return createNFTsInternally(hash, list, name)
}

func createNFTsInternally(
	hash hash.Hash,
	list []NFT,
	name string,
) NFTs {
	out := nfts{
		hash: hash,
		list: list,
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

// name fetches an nfts by name
func (obj *nfts) Fetch(name string) (NFT, error) {
	for _, oneNFT := range obj.list {
		retNFT, err := oneNFT.Fetch(name)
		if err != nil {
			continue
		}

		return retNFT, nil
	}

	str := fmt.Sprintf("the NFT (name: %s) could not be found", name)
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
