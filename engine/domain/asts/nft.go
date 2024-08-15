package asts

import "github.com/steve-care-software/webx/engine/domain/hash"

type nft struct {
	hash       hash.Hash
	bytes      []byte
	nfts       []hash.Hash
	complexity map[string]uint
}

func createNFTWithBytes(
	hash hash.Hash,
	bytes []byte,
) NFT {
	return createNFTInternally(hash, bytes, nil, map[string]uint{})
}

func createNFTWithNFTs(
	hash hash.Hash,
	nfts []hash.Hash,
	complexity map[string]uint,
) NFT {
	return createNFTInternally(hash, nil, nfts, complexity)
}

func createNFTInternally(
	hash hash.Hash,
	bytes []byte,
	nfts []hash.Hash,
	complexity map[string]uint,
) NFT {
	out := nft{
		hash:       hash,
		bytes:      bytes,
		nfts:       nfts,
		complexity: complexity,
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

// Complexity returns the complexity
func (obj *nft) Complexity() map[string]uint {
	return obj.complexity
}
