package databases

import (
	"github.com/steve-care-software/webx/applications/databases"
	"github.com/steve-care-software/webx/applications/databases/contents"
	"github.com/steve-care-software/webx/applications/databases/transactions"
	"github.com/steve-care-software/webx/domain/databases/references"
)

// NewBuilder creates a new database application
func NewBuilder() Builder {
	contentAppBuilder := NewContentApplicationBuilder()
	trxAppBuilder := NewTransactionApplicationBuilder()
	referenceAdapter := references.NewAdapter()
	referenceFactory := references.NewFactory()
	return createBuilder(
		contentAppBuilder,
		trxAppBuilder,
		referenceAdapter,
		referenceFactory,
	)
}

// NewContentApplicationBuilder creates a new content application builder
func NewContentApplicationBuilder() contents.Builder {
	return createContentApplicationBuilder()
}

// NewTransactionApplicationBuilder creates a new transaction application builder
func NewTransactionApplicationBuilder() transactions.Builder {
	return createTransactionApplicationBuilder()
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithDirPath(dirPath string) Builder
	Now() (databases.Application, error)
}
