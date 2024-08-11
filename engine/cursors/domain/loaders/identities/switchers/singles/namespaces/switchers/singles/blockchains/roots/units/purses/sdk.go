package purses

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// Builder represents a purses builder
type Builder interface {
	Create() Builder
	WithList(list []Purse) Builder
	Now() (Purses, error)
}

// Purses represents purses
type Purses interface {
	Hash() hash.Hash
	List() []Purse
}

// PurseBuilder represents a purse builder
type PurseBuilder interface {
	Create() PurseBuilder
	WithIndex(index uint64) PurseBuilder
	WithAmount(amount uint64) PurseBuilder
	Now() (Purse, error)
}

// Purse represents a purse
type Purse interface {
	Hash() hash.Hash
	Index() uint64
	Amount() uint64
}
