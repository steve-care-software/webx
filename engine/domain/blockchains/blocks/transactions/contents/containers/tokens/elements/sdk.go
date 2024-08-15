package elements

import "github.com/steve-care-software/webx/engine/domain/hash"

// Elements represents transactions content elements
type Elements interface {
	Hash() hash.Hash
	List() []Element
}

// Element represents a transaction content element
type Element interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsNft() bool
	Nft() hash.Hash
}
