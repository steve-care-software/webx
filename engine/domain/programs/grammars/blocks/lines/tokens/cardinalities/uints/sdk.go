package uints

import "github.com/steve-care-software/webx/engine/domain/nfts"

// AmountOfBytesIntUint64 represents the amount of bytes an uint64 contains
const AmountOfBytesIntUint64 = 8

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	nftsBuilder := nfts.NewBuilder()
	nftBuilder := nfts.NewNFTBuilder()
	return createAdapter(
		nftsBuilder,
		nftBuilder,
	)
}

// Adapter represents an uint adapter
type Adapter interface {
	ToNFT(value uint64) (nfts.NFT, error)
	ToValue(nft nfts.NFT) (*uint64, error)
}
