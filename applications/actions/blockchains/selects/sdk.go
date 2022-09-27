package selects

import (
	"net/url"

	"github.com/steve-care-software/syntax/domain/blockchains"
	"github.com/steve-care-software/syntax/domain/blockchains/blocks"
	"github.com/steve-care-software/syntax/domain/blockchains/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	repository blockchains.Repository,
	service blockchains.Service,
	blockRepository blocks.Repository,
) Builder {
	builder := blockchains.NewBuilder()
	transactionsBuilder := transactions.NewBuilder()
	return createBuilder(
		builder,
		repository,
		service,
		blockRepository,
		transactionsBuilder,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithReference(ref hash.Hash) Builder
	Now() (Application, error)
}

// Application represents a selected blockchain application
type Application interface {
	Retrieve() (blockchains.Blockchain, error)
	Transact(trx transactions.Transaction) error
	BlockByPreviousHash(prev hash.Hash) (blocks.Block, error)
	BlockByHash(hash hash.Hash) (blocks.Block, error)
	BlockByHeight(height uint) (blocks.Block, error)
	Search(trx hash.Hash) (blocks.Block, transactions.Transactions, error)
	ConnectList(conns []*url.URL) error
	Connect(conn *url.URL) error
}
