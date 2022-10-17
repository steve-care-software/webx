package databases

import (
	"net/url"

	"github.com/steve-care-software/syntax/domain/blockchains/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/databases"
	"github.com/steve-care-software/syntax/domain/syntax/databases/blockchains/blocks"
	"github.com/steve-care-software/syntax/domain/syntax/identities"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	Now() (Application, error)
}

// Application represents a database application
type Application interface {
	List() ([]string, error)
	Retrieve(name string) (databases.Database, error)
	SearchBlock(database databases.Database, trx transactions.Transaction) (blocks.Block, error)
	Execute(database databases.Database, value []byte, signature signatures.RingSignature) ([]byte, error)
	Propose(database databases.Database, trx transactions.Transaction) error
	Connect(database databases.Database, url url.URL) error
	Migrate(migration databases.Migration) error
}
