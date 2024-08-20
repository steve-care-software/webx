package nfts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type nft struct {
	hash       hash.Hash
	pByte      *byte
	nfts       NFTs
	pRecursive *uint
}

func createNFTWithByte(
	hash hash.Hash,
	pByte *byte,
) NFT {
	return createNFTInternally(hash, pByte, nil, nil)
}

func createNFTWithNFTs(
	hash hash.Hash,
	nfts NFTs,
) NFT {
	return createNFTInternally(hash, nil, nfts, nil)
}

func createNFTWithRecursive(
	hash hash.Hash,
	pRecursive *uint,
) NFT {
	return createNFTInternally(hash, nil, nil, pRecursive)
}

func createNFTInternally(
	hash hash.Hash,
	pByte *byte,
	nfts NFTs,
	pRecursive *uint,
) NFT {
	out := nft{
		hash:       hash,
		pByte:      pByte,
		nfts:       nfts,
		pRecursive: pRecursive,
	}

	return &out
}

// Hash returns the hash
func (obj *nft) Hash() hash.Hash {
	return obj.hash
}

// IsByte returns true if there is a byte, false otherwise
func (obj *nft) IsByte() bool {
	return obj.pByte != nil
}

// Byte returns the byte, if any
func (obj *nft) Byte() *byte {
	return obj.pByte
}

// IsNFTs returns true if there is nfts, false otherwise
func (obj *nft) IsNFTs() bool {
	return obj.nfts != nil
}

// NFTs returns the nfts, if any
func (obj *nft) NFTs() NFTs {
	return obj.nfts
}

// IsRecursive returns true if recursive, false otherwise
func (obj *nft) IsRecursive() bool {
	return obj.pRecursive != nil
}

// Recursive returns the recursive level, if any
func (obj *nft) Recursive() *uint {
	return obj.pRecursive
}

// Fetch fetches an nft by name
func (obj *nft) Fetch(name string) (NFT, error) {
	if !obj.IsNFTs() {
		str := fmt.Sprintf("the NFT (name: %s) could not be found", name)
		return nil, errors.New(str)
	}

	if obj.nfts.HasName() && obj.nfts.Name() == name {
		return obj, nil
	}

	return obj.nfts.Fetch(name)
}
