package singles

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains"

// Adapter represents a single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents a single builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithBlockchain(blockchain blockchains.Blockchain) Builder
	Now() (Single, error)
}

// Single represents a single namespace
type Single interface {
	Name() string
	Description() string
	Blockchain() blockchains.Blockchain
}
