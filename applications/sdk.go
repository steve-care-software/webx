package logics

import (
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/receipts"
	"github.com/steve-care-software/identity/domain/credentials"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials credentials.Credentials) Builder
	Now() (Application, error)
}

// Application represents the logic application
type Application interface {
	Execute(
		input []byte,
		layer layers.Layer,
		library libraries.Library,
		context receipts.Receipt,
	) (receipts.Receipt, error)
}
