package asts

import "github.com/steve-care-software/webx/engine/domain/hash"

type nft struct {
	hash  hash.Hash
	bytes []byte
	nfts  []hash.Hash
}

func createNFTWithBytes(
	hash hash.Hash,
	bytes []byte,
) NFT {
	return createNFTInternally(hash, bytes, nil)
}

func createNFTWithNFTs(
	hash hash.Hash,
	nfts []hash.Hash,
) NFT {
	return createNFTInternally(hash, nil, nfts)
}

func createNFTInternally(
	hash hash.Hash,
	bytes []byte,
	nfts []hash.Hash,
) NFT {
	out := nft{
		hash:  hash,
		bytes: bytes,
		nfts:  nfts,
	}

	return &out
}

// Hash returns the hash
func (obj *nft) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *nft) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *nft) Bytes() []byte {
	return obj.bytes
}

// IsNFTs returns true if there is nfts, false otherwise
func (obj *nft) IsNFTs() bool {
	return obj.nfts != nil
}

// NFTs returns the nfts, if any
func (obj *nft) NFTs() []hash.Hash {
	return obj.nfts
}
