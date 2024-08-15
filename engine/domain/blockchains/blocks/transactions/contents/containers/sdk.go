package containers

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents/containers/tokens"
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents/containers/transfers"
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// Containers represents containers
type Containers interface {
	Hash() hash.Hash
	List() []Container
}

// Container represents a container
type Container interface {
	Hash() hash.Hash
	IsToken() bool
	Token() tokens.Token
	IsTransfer() bool
	Transfer() transfers.Transfer
}
