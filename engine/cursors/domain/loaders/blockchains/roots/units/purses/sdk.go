package purses

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// Purses represents purses
type Purses interface {
	Hash() hash.Hash
	List() []Purse
}

// Purse represents a purse
type Purse interface {
	Hash() hash.Hash
	Index() uint64
	Amount() uint64
}
