package asts

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewASTBuilder creates a new ast builder
func NewASTBuilder() AstBuilder {
	hashAdapter := hash.NewAdapter()
	return createAstBuilder(
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

// Builder represents an ast list
type Builder interface {
	Create() Builder
	WithList(list []AST) Builder
	Now() (ASTs, error)
}

// ASTs represents asts
type ASTs interface {
	List() []AST
}

// AstBuilder represents an AST builder
type AstBuilder interface {
	Create() AstBuilder
	WithLibrary(library NFTs) AstBuilder
	WithEntry(entry hash.Hash) AstBuilder
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
