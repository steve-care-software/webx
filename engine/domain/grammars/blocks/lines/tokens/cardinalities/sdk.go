package cardinalities

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/cardinalities/uints"
	"github.com/steve-care-software/webx/engine/domain/nfts"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	cardinalityBuilder := NewBuilder()
	uintAdapter := uints.NewAdapter()
	nftsBuilder := nfts.NewBuilder()
	nftBuilder := nfts.NewNFTBuilder()
	return createAdapter(
		cardinalityBuilder,
		uintAdapter,
		nftsBuilder,
		nftBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the cardinality adapter
type Adapter interface {
	ToNFT(ins Cardinality) (nfts.NFT, error)
	ToInstance(nft nfts.NFT) (Cardinality, error)
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents a cardinality
type Cardinality interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
