package asts

import "github.com/steve-care-software/webx/engine/domain/hash"

type nft struct {
	hash  hash.Hash
	pByte *byte
	nfts  []hash.Hash
}

func createNFTWithByte(
	hash hash.Hash,
	pByte *byte,
) NFT {
	return createNFTInternally(hash, pByte, nil)
}

func createNFTWithNFTs(
	hash hash.Hash,
	nfts []hash.Hash,
) NFT {
	return createNFTInternally(hash, nil, nfts)
}

func createNFTInternally(
	hash hash.Hash,
	pByte *byte,
	nfts []hash.Hash,
) NFT {
	out := nft{
		hash:  hash,
		pByte: pByte,
		nfts:  nfts,
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
func (obj *nft) NFTs() []hash.Hash {
	return obj.nfts
}
