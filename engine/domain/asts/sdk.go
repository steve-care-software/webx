package asts

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewNFTsBuilder creates a new nfts builder
func NewNFTsBuilder() NFTsBuilder {
	hashAdapter := hash.NewAdapter()
	return createNFTsBuilder(
		hashAdapter,
	)
}

// NewNFTBuilder creates a new nft builder
func NewNFTBuilder() NFTBuilder {
	hashAdapter := hash.NewAdapter()
	return createNFTBuilder(
		hashAdapter,
	)
}

// Builder represents an AST builder
type Builder interface {
	Create() Builder
	WithLibrary(library NFTs) Builder
	WithEntry(entry hash.Hash) Builder
	Now() (AST, error)
}

// AST represents an AST
type AST interface {
	Library() NFTs
	Entry() hash.Hash
	Complexity() map[string]uint
}

// NFTsBuilder represents an nfts builder
type NFTsBuilder interface {
	Create() NFTsBuilder
	WithList(list []NFT) NFTsBuilder
	Now() (NFTs, error)
}

// NFTs represents nfts
type NFTs interface {
	Hash() hash.Hash
	List() []NFT
	Complexity() map[string]uint
	Fetch(hash hash.Hash) (NFT, error)
}

// NFTBuilder represents an nft builder
type NFTBuilder interface {
	Create() NFTBuilder
	WithBytes(bytes []byte) NFTBuilder
	WithNFTs(nfts []hash.Hash) NFTBuilder
	Now() (NFT, error)
}

// NFT represents an nft
type NFT interface {
	Hash() hash.Hash
	Complexity() map[string]uint
	IsBytes() bool
	Bytes() []byte
	IsNFTs() bool
	NFTs() []hash.Hash
}
